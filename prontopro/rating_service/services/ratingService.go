package services

import (
	"errors"
	"github.com/shonayevshyngys/prontopro/rating_service/database"
	"github.com/shonayevshyngys/prontopro/rating_service/models"
	"github.com/shonayevshyngys/prontopro/rating_service/util"
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
	return nil
}
