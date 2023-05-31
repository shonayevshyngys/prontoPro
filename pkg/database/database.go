package database

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/shonayevshyngys/prontopro/pkg/models"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DataBase DB
var RedisInstance *redis.Client
var RedisContext context.Context

type DB struct {
	Instance *gorm.DB
}

type DBInterface interface {
	CreateUser(user *models.User)
	CreateProvider(provider *models.Provider)
	CreateReview(review *models.Review)
	CreateNotification(notification models.Notification)
	GetUser(user *models.User, id int)
	GetProvider(provider *models.Provider, id int)
	GetReview(review *models.Review, id uint)
	UpdateProvider(provider *models.Provider)
	FetchAverageRating(id uint, rating *float32)
	UserExists(id uint) bool
	ProviderExists(id uint) bool
}

func (d *DB) CreateNotification(notification *models.Notification) {
	d.Instance.Create(&notification)
}

func (d *DB) CreateUser(user *models.User) {
	d.Instance.Create(&user)
}

func (d *DB) CreateProvider(provider *models.Provider) {
	d.Instance.Create(&provider)
}

func (d *DB) CreateReview(review *models.Review) {
	d.Instance.Create(&review)
}

func (d *DB) GetUser(user *models.User, id int) {
	d.Instance.Find(&user, id)
}

func (d *DB) GetProvider(provider *models.Provider, id int) {
	d.Instance.Find(&provider, id)
}

func (d *DB) GetReview(review *models.Review, id uint) {
	d.Instance.Preload("User").Preload("Provider").Find(&review, review.ID)
}

func (d *DB) UpdateProvider(provider *models.Provider) {
	d.Instance.Save(&provider)
}

func (d *DB) FetchAverageRating(id uint, rating *float32) {
	d.Instance.Raw("SELECT AVG(Rating) FROM reviews WHERE provider_id = ?", id).Scan(&rating)
}

func ConnectToDatabase() {

	fmt.Println(os.Getenv("DATASOURCE"))
	dsn := os.Getenv("DATASOURCE")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed connect to database ", err)
	}
	db.Set("gorm:auto_preload", true)
	DataBase = DB{
		Instance: db,
	}
}

func ConnectToRedis() {
	client := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_URL"),
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	RedisInstance = client
	RedisContext = context.Background()
}

func (d *DB) UserExists(id uint) bool {
	var user models.User
	d.Instance.First(&user, id)
	if user.ID == 0 {
		return false
	}
	return true
}

func (d *DB) ProviderExists(id uint) bool {
	var provider models.Provider
	d.Instance.First(&provider, id)
	if provider.ID == 0 {
		return false
	}
	return true
}
