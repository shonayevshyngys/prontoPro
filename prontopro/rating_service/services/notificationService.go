package services

import (
	"encoding/json"
	"github.com/shonayevshyngys/prontopro/rating_service/database"
	"github.com/shonayevshyngys/prontopro/rating_service/models"
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
