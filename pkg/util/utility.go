package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/shonayevshyngys/prontopro/pkg/models"
	"log"
	"net/http"
)

type ErrorMessage struct {
	Code    uint
	Message string
}
type SuccessMessage struct {
	Code    uint
	Message string
}

type Subscribers struct {
	Ids map[int]string
}

const BadIdText = "Bad format for id"

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

func CheckIfUserAndProviderExists(provider int, user int) (bool, error) {

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
