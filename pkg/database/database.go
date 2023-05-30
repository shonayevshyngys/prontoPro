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

var Instance *gorm.DB
var RedisInstance *redis.Client
var RedisContext context.Context

func ConnectToDatabase() {

	fmt.Println(os.Getenv("DATASOURCE"))
	dsn := os.Getenv("DATASOURCE")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed connect to database ", err)
	}
	db.Set("gorm:auto_preload", true)
	Instance = db // newer version of golang forces this
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

func UserExists(id uint) bool {
	var user models.User
	Instance.First(&user, id)
	if user.ID == 0 {
		return false
	}
	return true
}

func ProviderExists(id uint) bool {
	var provider models.Provider
	Instance.First(&provider, id)
	if provider.ID == 0 {
		return false
	}
	return true
}
