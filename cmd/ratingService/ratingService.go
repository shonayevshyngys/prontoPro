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

func GetRatingService() RatingService {
	return RatingService{}
}

type RatingService struct{}

type ratingServiceInterface interface {
	createUser(user *models.User) error
	createProvider(provider *models.Provider) error
	createReview(reviewDTO *util.CreateReviewDTO) (models.Review, error)
	getProvider(provider *models.Provider, id int) error
	sendNotification(review *models.Review)
}

func (r *RatingService) createUser(user *models.User) error {
	user.ID = 0
	database.DataBase.CreateUser(user)
	if user.ID == 0 {
		return errors.New(objectNotCreatedErrorText)
	}
	return nil
}

func (r *RatingService) createProvider(provider *models.Provider) error {
	provider.ID = 0
	provider.Rating = 0
	if provider.Name == "" || provider.Description == "" {
		return errors.New(objectNotCreatedErrorText)
	}

	database.DataBase.CreateProvider(provider)
	if provider.ID == 0 {
		return errors.New(objectNotCreatedErrorText)
	}
	return nil
}

func (r *RatingService) createReview(reviewDTO *util.CreateReviewDTO) (models.Review, error) {
	review := models.Review{
		ReviewText: reviewDTO.ReviewText,
		Rating:     reviewDTO.Rating,
		UserID:     reviewDTO.UserId,
		ProviderID: reviewDTO.ProviderId,
	}
	if !database.DataBase.UserExists(review.UserID) || !database.DataBase.ProviderExists(review.ProviderID) {
		return review, errors.New("user or provider doesn't exist")
	}
	database.DataBase.CreateReview(&review)
	if review.ID == 0 {
		return review, errors.New(objectNotCreatedErrorText)
	}
	database.DataBase.GetReview(&review, review.ID)

	return review, nil
}

func (r *RatingService) getProvider(provider *models.Provider, id int) error {
	database.DataBase.GetProvider(provider, id)
	if provider.ID == 0 {
		return errors.New(objectNotCreatedErrorText)
	}
	var rating float32
	database.DataBase.FetchAverageRating(provider.ID, &rating)
	provider.Rating = rating
	if provider.Rating != rating {
		go database.DataBase.UpdateProvider(provider)
	}
	return nil
}

func (r *RatingService) sendNotification(review *models.Review) {

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
