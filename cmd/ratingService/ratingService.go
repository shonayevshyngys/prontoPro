package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/shonayevshyngys/prontopro/pkg/util"
	"log"
	"net/http"

	"github.com/shonayevshyngys/prontopro/pkg/database"
	"github.com/shonayevshyngys/prontopro/pkg/models"
)

const objectNotCreatedErrorText = "object wasn't created"

func CreateUser(user *models.User) error {
	user.ID = 0
	database.Instance.Create(&user)
	if user.ID == 0 {
		return errors.New(objectNotCreatedErrorText)
	}
	return nil
}

func CreateProvider(provider *models.Provider) error {
	provider.ID = 0
	provider.Rating = 0
	if provider.Name == "" || provider.Description == "" {
		return errors.New(objectNotCreatedErrorText)
	}

	database.Instance.Create(&provider)
	if provider.ID == 0 {
		return errors.New(objectNotCreatedErrorText)
	}
	return nil
}

func CreateReview(reviewDTO *util.CreateReviewDTO) (models.Review, error) {
	review := models.Review{
		ReviewText: reviewDTO.ReviewText,
		Rating:     reviewDTO.Rating,
		UserID:     reviewDTO.UserId,
		ProviderID: reviewDTO.ProviderId,
	}
	if !database.UserExists(review.UserID) || !database.ProviderExists(review.ProviderID) {
		return review, errors.New("user or provider doesn't exist")
	}
	database.Instance.Create(&review)
	if review.ID == 0 {
		return review, errors.New(objectNotCreatedErrorText)
	}
	database.Instance.Preload("User").Preload("Provider").Find(&review, review.ID)

	return review, nil
}

func GetProvider(provider *models.Provider, id int) error {
	database.Instance.Find(&provider, id)
	if provider.ID == 0 {
		return errors.New(objectNotCreatedErrorText)
	}
	var rating float32
	database.Instance.Raw("SELECT AVG(Rating) FROM reviews WHERE provider_id = ?", provider.ID).Scan(&rating)
	provider.Rating = rating
	if provider.Rating != rating {
		go database.Instance.Save(&provider)
	}
	return nil
}

func SendNotification(review *models.Review) {

	notification := models.Notification{
		ProviderID:   review.ProviderID,
		Notification: fmt.Sprintf("New rating %d submitted by %s", review.Rating, review.User.Username),
	}

	something, err := json.Marshal(notification)
	if err != nil {
		log.Println("marshalling of object failed ", err)
		return
	}
	log.Println("Sending request to notification service")
	//change later
	_, err = http.Post("http://notification-service:7001/notification", "application/json", bytes.NewBuffer(something))
	if err != nil {
		log.Println("Error on send to notification service ", err)
		return
	}
}
