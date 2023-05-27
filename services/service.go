package services

import (
	"errors"
	"shonayevshyngys/database"
	"shonayevshyngys/models"
)

const objectNotCreatedErrorText = "object wasn't created"

func CreateUser(user *models.User) error {
	database.Instance.Create(&user)
	if user.ID == 0 {
		return errors.New(objectNotCreatedErrorText)
	}
	return nil
}

func CreateProvider(provider *models.Provider) error {
	database.Instance.Create(&provider)
	if provider.ID == 0 {
		return errors.New(objectNotCreatedErrorText)
	}
	return nil
}

func CreateReview(review *models.Review) error {
	database.Instance.Create(&review)
	if review.ID == 0 {
		return errors.New(objectNotCreatedErrorText)
	}
	return nil
}
