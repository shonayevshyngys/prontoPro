package services

import (
	"errors"
	"shonayevshyngys/database"
	"shonayevshyngys/models"
	"shonayevshyngys/util"
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
	err := database.Instance.Raw("SELECT AVG(Rating) FROM reviews WHERE provider_id = ?", provider.ID).Scan(&rating).Error
	if err != nil {
		return err
	}
	provider.Rating = rating
	return nil
}
