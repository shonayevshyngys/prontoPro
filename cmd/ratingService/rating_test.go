package main

import (
	"fmt"
	"github.com/shonayevshyngys/prontopro/pkg/database"
	"github.com/shonayevshyngys/prontopro/pkg/models"
	"github.com/shonayevshyngys/prontopro/pkg/util"
	"github.com/shonayevshyngys/prontopro/test"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_RatingService(t *testing.T) {
	ratingService := GetRatingService()

	mockDb := test.GetNewMockDB()
	database.DataBase = database.DBWrapper{
		DBInterface: &mockDb,
	}
	//setup db

	//create users
	names := []string{"Paul", "Allen", "Mary", "Macbeth"}
	for i, name := range names {
		user := models.User{Username: name}
		ratingService.createUser(&user)
		assert.EqualValues(t, user.ID, i+1)
		assert.EqualValues(t, user.Username, name)
	}
	fmt.Println("We created four users")

	//create providers
	providers := []string{"Cleaning", "Building", "Curing", "Wielding"}
	for i, name := range providers {
		provider := models.Provider{
			Name:        name,
			Description: "Doing some stuff",
		}
		ratingService.createProvider(&provider)
		assert.EqualValues(t, provider.ID, i+1)
		assert.EqualValues(t, provider.Name, name)
	}
	fmt.Println("Then created four providers")
	//create notifications
	rId := 1
	for i := 1; i <= 4; i++ {
		for j := 1; j <= 4; j++ {
			reviewDTO := util.CreateReviewDTO{
				UserId:     uint(i),
				ProviderId: uint(j),
				ReviewText: "test rating",
				Rating:     uint8(i),
			}
			review, err := ratingService.createReview(&reviewDTO)
			if err != nil {
				fmt.Println(err)
			}
			assert.EqualValues(t, review.ID, rId)
			rId++
		}
	}
	fmt.Println("Each user left one review for each provider")

	//Get providers
	for i := 1; i <= 4; i++ {
		provider := models.Provider{}
		ratingService.getProvider(&provider, i)
		assert.EqualValues(t, provider.ID, i)
		assert.EqualValues(t, provider.Rating, 2.5)
	}
	fmt.Println("Each provider should have average rating of 2.5")

}
