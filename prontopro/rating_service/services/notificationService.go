package services

import (
	"encoding/json"
	"errors"
	"github.com/shonayevshyngys/prontopro/rating_service/database"
	"github.com/shonayevshyngys/prontopro/rating_service/models"
	"github.com/shonayevshyngys/prontopro/rating_service/util"
	"log"
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
}

func GetProviderNotifications(id int) ([]models.Notification, error) {
	key := "ProviderID_" + strconv.Itoa(id)
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

func SubscribeUserToProvider(provider int, user int) error {
	exists, err := util.CheckIfUserAndProviderExists(provider, user)
	log.Println(exists, err)
	if err != nil {
		log.Println("Couldn't verify if user/provider exists")
		return err
	}
	if !exists {
		var res string
		res = "user and/or provider doesn't exist"
		log.Println(res)
		return errors.New(res)
	}
	return nil
}

func Subscribe(subscriptionBody *util.SubscriptionBody) {
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
