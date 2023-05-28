package util

import (
	"bytes"
	"encoding/json"
	"github.com/shonayevshyngys/prontopro/rating_service/models"
	"log"
	"net/http"
)

type ErrorMessage struct {
	Code    uint
	Message string
}

func SaveNotification(notification *models.Notification) error {
	something, err := json.Marshal(notification)
	if err != nil {
		log.Println("marshalling of object failed ", err)
		return err
	}
	log.Println("Sending request to notification service")
	//change later
	_, err = http.Post("http://notification-service:7001/notification", "application/json", bytes.NewBuffer(something))
	if err != nil {
		return err
	}
	return nil
}
