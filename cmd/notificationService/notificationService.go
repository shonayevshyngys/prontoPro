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

func SaveNotification(notification *models.Notification) {
	database.Instance.Create(&notification)

	jsonBody, err := json.Marshal(notification)
	if err != nil {
		log.Println("body wasn't persisted to redis", err)
		return
	}
	go database.RedisInstance.RPush(database.RedisContext,
		"ProviderID_"+strconv.Itoa(int(notification.ProviderID)),
		jsonBody,
	)
	go sendNotificationToSubbedUsers(notification.ProviderID, jsonBody)
}

func sendNotificationToSubbedUsers(id uint, jsonBody []byte) {
	res, err := database.RedisInstance.Get(database.RedisContext,
		"ProviderID_"+strconv.Itoa(int(id))+"_subs").Result()
	if err != nil {
		log.Printf("Provider %d doesnt have subscirbers\n", id)
	}
	var subscribers util.Subscribers
	err = json.Unmarshal([]byte(res), &subscribers)
	if err != nil {
		log.Println("Something went wrong on parsing")
	}
	for key, _ := range subscribers.Ids {
		var redisKey string
		redisKey = "UserID_" + strconv.Itoa(key)
		go func() {
			_, err := database.RedisInstance.RPush(database.RedisContext, redisKey, jsonBody).Result()
			if err != nil {
				log.Println("failed to persist to redis", err)
			}
		}()
	}

}

func GetProviderNotifications(id int) ([]models.Notification, error) {
	key := "ProviderID_" + strconv.Itoa(id)
	notificaitons, err := getNotifications(id, key)
	return notificaitons, err
}

func GetUserNotifications(id int) ([]models.Notification, error) {
	key := "UserID_" + strconv.Itoa(id)
	notificaitons, err := getNotifications(id, key)
	return notificaitons, err
}

func getNotifications(id int, key string) ([]models.Notification, error) {
	stringNotifications, err := database.RedisInstance.LRange(database.RedisContext, key, 0, -1).Result()
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
	go database.RedisInstance.Del(database.RedisContext, key)
	return notificaitons, err
}

func SubscribeUserToProvider(provider int, user int, subBody *util.SubscriptionBody) error {
	exists, err := checkIfUserAndProviderExists(provider, user)
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
	res, err := database.RedisInstance.Get(database.RedisContext, key).Result()
	if err != nil {
		ids := make(map[int]string)
		ids[subscriptionBody.UserId] = "subbed" //it's Golang way for set, the values are not needed.
		jsonIds, _ := json.Marshal(&util.Subscribers{Ids: ids})
		database.RedisInstance.Set(database.RedisContext, key, jsonIds, 0)
	} else {
		var subs util.Subscribers
		errJson := json.Unmarshal([]byte(res), &subs)
		if errJson != nil {
			return
		}
		subs.Ids[subscriptionBody.UserId] = "Subbed"
		jsonIds, _ := json.Marshal(&subs)
		database.RedisInstance.Set(database.RedisContext, key, jsonIds, 0)
	}
}

func checkIfUserAndProviderExists(provider int, user int) (bool, error) {

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
