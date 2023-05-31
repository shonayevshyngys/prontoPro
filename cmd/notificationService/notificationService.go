package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/shonayevshyngys/prontopro/pkg/database"
	"github.com/shonayevshyngys/prontopro/pkg/models"
	"github.com/shonayevshyngys/prontopro/pkg/util"
	"log"
	"net/http"
	"strconv"
)

func GetNotificationService() NotificationService {
	return NotificationService{}
}

type NotificationService struct {
	notificationServiceInterface
}

type notificationServiceInterface interface {
	saveNotification(notification *models.Notification)
	sendNotificationToSubbedUsers(id uint, jsonBody []byte)
	getProviderNotifications(id int) ([]models.Notification, error)
	getUserNotifications(id int) ([]models.Notification, error)
	getNotifications(id int, key string) ([]models.Notification, error)
	subscribeUserToProvider(provider int, user int, subBody *util.SubscriptionBody) error
	checkIfUserAndProviderExists(provider int, user int) (bool, error)
}

func (service *NotificationService) saveNotification(notification *models.Notification) {
	database.DataBase.DBInterface.CreateNotification(notification)

	jsonBody, err := json.Marshal(notification)
	if err != nil {
		log.Println("body wasn't persisted to redis", err)
		return
	}
	go database.RedisBase.RedisInterface.RPush(
		"ProviderID_"+strconv.Itoa(int(notification.ProviderID)),
		jsonBody,
	)
	go service.sendNotificationToSubbedUsers(notification.ProviderID, jsonBody)
}

func (service *NotificationService) sendNotificationToSubbedUsers(id uint, jsonBody []byte) {
	res, err := database.RedisBase.RedisInterface.Get("ProviderID_" + strconv.Itoa(int(id)) + "_subs")
	if err != nil {
		log.Printf("Provider %d doesnt have subscirbers\n", id)
		return
	}
	var subscribers util.Subscribers
	err = json.Unmarshal([]byte(res), &subscribers)
	if err != nil {
		log.Println("Something went wrong on parsing")
		return
	}
	for key, _ := range subscribers.Ids {
		var redisKey string
		redisKey = "UserID_" + strconv.Itoa(key)
		go func() {
			_, err := database.RedisBase.RedisInterface.RPush(redisKey, jsonBody)
			if err != nil {
				log.Println("failed to persist to redis", err)
			}
		}()
	}

}

func (service *NotificationService) getProviderNotifications(id int) ([]models.Notification, error) {
	key := "ProviderID_" + strconv.Itoa(id)
	notificaitons, err := service.getNotifications(id, key)
	return notificaitons, err
}

func (service *NotificationService) getUserNotifications(id int) ([]models.Notification, error) {
	key := "UserID_" + strconv.Itoa(id)
	notificaitons, err := service.getNotifications(id, key)
	return notificaitons, err
}

func (service *NotificationService) getNotifications(id int, key string) ([]models.Notification, error) {
	stringNotifications, err := database.RedisBase.RedisInterface.LRange(key)
	if err != nil || len(stringNotifications) == 0 {
		return nil, err
	}
	notificaitons := make([]models.Notification, 0)
	for _, notification := range stringNotifications {
		temp := models.Notification{}
		err = json.Unmarshal([]byte(notification), &temp)
		if err != nil {
			continue
		}
		notificaitons = append(notificaitons, temp)
	}
	go database.RedisBase.RedisInterface.Del(key)
	return notificaitons, err
}

func (service *NotificationService) subscribeUserToProvider(provider int, user int, subBody *util.SubscriptionBody) error {
	exists, err := service.checkIfUserAndProviderExists(provider, user)
	if err != nil {
		log.Println("Couldn't verify if user/provider exists")
		return err
	}
	if !exists {
		var res string
		res = "user and/or provider doesn't exist"
		return errors.New(res)
	}
	go subscribe(subBody)
	return nil
}

func subscribe(subscriptionBody *util.SubscriptionBody) {
	key := "ProviderID_" + strconv.Itoa(subscriptionBody.ProviderId) + "_subs"
	res, err := database.RedisBase.RedisInterface.Get(key)
	if err != nil {
		ids := make(map[int]string)
		ids[subscriptionBody.UserId] = "subbed" //it's Golang way for set, the values are not needed.
		jsonIds, _ := json.Marshal(&util.Subscribers{Ids: ids})
		database.RedisBase.RedisInterface.Set(key, jsonIds)
	} else {
		var subs util.Subscribers
		errJson := json.Unmarshal([]byte(res), &subs)
		if errJson != nil {
			return
		}
		subs.Ids[subscriptionBody.UserId] = "Subbed"
		jsonIds, _ := json.Marshal(&subs)
		database.RedisBase.RedisInterface.Set(key, jsonIds)
	}
}

func (service *NotificationService) checkIfUserAndProviderExists(provider int, user int) (bool, error) {

	log.Println("Checking if user or provider exists")
	//change later
	url := fmt.Sprintf("http://rating-service:7000/rating/check/%d/%d", provider, user)
	resp, err := http.Get(url)
	if err != nil {
		return false, err
	}
	if resp.StatusCode == http.StatusOK {
		return true, nil
	} else {
		return false, nil
	}

}
