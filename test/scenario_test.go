package test

import (
	"fmt"
	"github.com/shonayevshyngys/prontopro/pkg/database"
	"github.com/shonayevshyngys/prontopro/pkg/models"
	"github.com/shonayevshyngys/prontopro/pkg/notificationService"
	"github.com/shonayevshyngys/prontopro/pkg/ratingService"
	"github.com/shonayevshyngys/prontopro/pkg/util"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
	"time"
)

func Test_RatingService(t *testing.T) {
	rService := ratingService.GetRatingService()
	nService := notificationService.GetNotificationService()
	mockDb := GetNewMockDB()
	mockRedis := GetNewMiniRedis()
	database.DataBase = database.DBWrapper{
		DBInterface: &mockDb,
	}
	database.RedisBase = database.RedisWrapper{
		RedisInterface: &mockRedis,
	}

	//create users
	names := []string{"Paul", "Allen", "Mary", "Macbeth"}
	for i, name := range names {
		user := models.User{Username: name}
		err := rService.CreateUser(&user)
		if err != nil {
			t.Fatal("User wasn't created")
		}
		assert.EqualValues(t, user.ID, i+1)
		assert.EqualValues(t, user.Username, name)
	}
	log.Println("We created four users")

	//create providers
	providers := []string{"Cleaning", "Building", "Curing", "Wielding"}
	for i, name := range providers {
		provider := models.Provider{
			Name:        name,
			Description: "Doing some stuff",
		}
		err := rService.CreateProvider(&provider)
		if err != nil {
			t.Fatal("Provider wasn't created")
		}
		assert.EqualValues(t, provider.ID, i+1)
		assert.EqualValues(t, provider.Name, name)
	}
	log.Println("Then created four providers")
	log.Println("And subscribe each user to each provider")
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
			subBody := util.SubscriptionBody{
				UserId:     int(reviewDTO.UserId),
				ProviderId: int(reviewDTO.ProviderId),
			}
			//before we create review let's subscribe all users to proiders
			nService.Subscribe(&subBody)
			review, err := rService.CreateReview(&reviewDTO)
			if err != nil {
				log.Println(err)
			}
			// mocking creating of notificaitons as they are happening between services
			notifcation := models.Notification{
				ProviderID:   review.ProviderID,
				Notification: fmt.Sprintf("New rating %d submitted by %s", review.Rating, review.User.Username),
			}
			nService.SaveNotification(&notifcation)
			assert.EqualValues(t, review.ID, rId)
			rId++
		}
	}
	log.Println("Each user left one review for each provider and we sent notification both to users and providers")
	log.Println("Notifications are sent in goroutines so test can be unstable")
	//Get providers
	for i := 1; i <= 4; i++ {
		provider := models.Provider{}
		err := rService.GetProvider(&provider, i)
		if err != nil {
			t.Fatal("Provider didn't appear")
		}
		assert.EqualValues(t, provider.ID, i)
		assert.EqualValues(t, provider.Rating, 2.5)
	}
	time.Sleep(time.Second * 2) // let's wait a little bit for notifications
	log.Println("Each provider should have average rating of 2.5")
	log.Println("Let's try to get user notifications")
	for i := 1; i <= 4; i++ {
		notifications, err := nService.GetUserNotifications(i)
		if err != nil {
			t.Fatal("Notifications didn't appear")
		}
		assert.NotEqualf(t, 0, len(notifications), "User has zero notifications")
	}
	time.Sleep(time.Second * 2) // let's wait a little bit for deletion of notifcations
	log.Println("So we fetched all users' notifcations, let's retrieve them once again")
	for i := 1; i <= 4; i++ {
		notifations, err := nService.GetUserNotifications(i)
		if err != nil {
			assert.EqualError(t, err, "ERR no such key")
		} else {
			assert.EqualValues(t, 0, len(notifations))
		}
	}
	log.Println("And we can see that they are empty. Now let's do it with providers")
	for i := 1; i <= 4; i++ {
		notifications, err := nService.GetProviderNotifications(i)
		if err != nil {
			t.Fatal("Notifications didn't appear")
		}
		assert.EqualValues(t, 4, len(notifications), "User has zero notifications")
	}
	log.Println("So every provider got 4 notifications, which make sense as each user sent one review to each provider")
	log.Println("Let's try it again")
	time.Sleep(time.Second * 2) // let's wait a little bit for deletion of notifcations
	for i := 1; i <= 4; i++ {
		notifations, err := nService.GetProviderNotifications(i)
		if err != nil {
			assert.EqualError(t, err, "ERR no such key")
		} else {
			assert.EqualValues(t, 0, len(notifations))
		}
	}
	log.Println("And nothing appeared!")
	log.Println("PASS")

}
