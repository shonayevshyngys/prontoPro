package test

import (
	"github.com/shonayevshyngys/prontopro/pkg/database"
	"github.com/shonayevshyngys/prontopro/pkg/models"
)

type DBMock struct {
	users               map[int]models.User
	providers           map[int]models.Provider
	reviews             map[int]models.Review
	notifications       map[int]models.Notification
	usersCounter        int
	providerCounter     int
	reviewCounter       int
	notificationCounter int
	DBInterface         database.DBInterface
}

func GetNewMockDB() DBMock {
	return DBMock{
		users:         map[int]models.User{},
		providers:     map[int]models.Provider{},
		reviews:       map[int]models.Review{},
		notifications: map[int]models.Notification{},
	}
}

func (D *DBMock) CreateUser(user *models.User) {
	D.usersCounter++
	user.ID = uint(D.usersCounter)
	D.users[D.usersCounter] = *user
}

func (D *DBMock) CreateProvider(provider *models.Provider) {
	D.providerCounter = D.providerCounter + 1
	provider.ID = uint(D.providerCounter)
	D.providers[D.providerCounter] = *provider
}

func (D *DBMock) CreateReview(review *models.Review) {
	D.reviewCounter++
	review.ID = uint(D.reviewCounter)
	D.reviews[D.reviewCounter] = *review
}

func (D *DBMock) CreateNotification(notification *models.Notification) {
	D.notificationCounter++
	notification.ID = uint(D.notificationCounter)
	D.notifications[D.notificationCounter] = *notification
}

func (D *DBMock) GetUser(user *models.User, id int) {
	newUser, ok := D.users[id]
	if ok {
		user.Username = newUser.Username
		user.ID = newUser.ID
	}
}

func (D *DBMock) GetProvider(provider *models.Provider, id int) {
	newProvider, ok := D.providers[id]
	if ok {
		provider.Name = newProvider.Name
		provider.Description = newProvider.Description
		provider.Rating = newProvider.Rating
		provider.ID = newProvider.ID
	}
}

func (D *DBMock) GetReview(review *models.Review, id uint) {
	newReview, ok := D.reviews[int(id)]
	if ok {
		newProvider, okP := D.providers[int(review.ProviderID)]
		newUser, okU := D.users[int(review.UserID)]
		if okU && okP {
			review.ReviewText = newReview.ReviewText
			review.Rating = newReview.Rating
			review.ID = newReview.ID
			review.User = newUser
			review.Provider = newProvider
		}
	}

}

func (D *DBMock) UpdateProvider(provider *models.Provider) {
	D.providers[int(provider.ID)] = *provider
}

func (D *DBMock) FetchAverageRating(id uint, rating *float32) {
	amount := 0
	sum := 0
	for _, v := range D.reviews {
		if v.ProviderID == id {
			amount++
			sum = int(v.Rating) + sum
		}
	}
	floatSum := float32(sum)
	floatAmount := float32(amount)
	temp := floatSum / floatAmount
	*rating = temp
}

func (D *DBMock) UserExists(id uint) bool {
	_, ok := D.users[int(id)]
	return ok
}

func (D *DBMock) ProviderExists(id uint) bool {
	_, ok := D.providers[int(id)]
	return ok
}
