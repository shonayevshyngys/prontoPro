package database

import "github.com/shonayevshyngys/prontopro/pkg/models"

// DB implementation
type DB struct{}

type DBInterface interface {
	CreateUser(user *models.User)
	CreateProvider(provider *models.Provider)
	CreateReview(review *models.Review)
	CreateNotification(notification *models.Notification)
	GetUser(user *models.User, id int)
	GetProvider(provider *models.Provider, id int)
	GetReview(review *models.Review, id uint)
	UpdateProvider(provider *models.Provider)
	FetchAverageRating(id uint, rating *float32)
	UserExists(id uint) bool
	ProviderExists(id uint) bool
}

func (d *DB) CreateNotification(notification *models.Notification) {
	GormInstance.Create(&notification)
}

func (d *DB) CreateUser(user *models.User) {
	GormInstance.Create(&user)
}

func (d *DB) CreateProvider(provider *models.Provider) {
	GormInstance.Create(&provider)
}

func (d *DB) CreateReview(review *models.Review) {
	GormInstance.Create(&review)
}

func (d *DB) GetUser(user *models.User, id int) {
	GormInstance.Find(&user, id)
}

func (d *DB) GetProvider(provider *models.Provider, id int) {
	GormInstance.Find(&provider, id)
}

func (d *DB) GetReview(review *models.Review, id uint) {
	GormInstance.Preload("User").Preload("Provider").Find(&review, review.ID)
}

func (d *DB) UpdateProvider(provider *models.Provider) {
	GormInstance.Save(&provider)
}

func (d *DB) FetchAverageRating(id uint, rating *float32) {
	GormInstance.Raw("SELECT AVG(Rating) FROM reviews WHERE provider_id = ?", id).Scan(&rating)
}

func (d *DB) UserExists(id uint) bool {
	var user models.User
	GormInstance.First(&user, id)
	if user.ID == 0 {
		return false
	}
	return true
}

func (d *DB) ProviderExists(id uint) bool {
	var provider models.Provider
	GormInstance.First(&provider, id)
	if provider.ID == 0 {
		return false
	}
	return true
}
