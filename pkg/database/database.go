package database

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/postgres"
	"log"
	"os"

	"gorm.io/gorm"
)

var DataBase DBWrapper
var RedisBase RedisWrapper

var RedisInstance *redis.Client
var RedisContext context.Context
var GormInstance *gorm.DB

type DBWrapper struct {
	DBInterface DBInterface
}

type RedisWrapper struct {
	RedisInterface RedisInterface
}

func ConnectToDatabase() {

	fmt.Println(os.Getenv("DATASOURCE"))
	dsn := os.Getenv("DATASOURCE")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed connect to database ", err)
	}
	db.Set("gorm:auto_preload", true)
	GormInstance = db
	DataBase = DBWrapper{
		DBInterface: &DB{},
	}
}

func ConnectToRedis() {
	client := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_URL"),
		Password: "", // no password Set
		DB:       0,  // use default DB
	})
	RedisInstance = client
	RedisContext = context.Background()
	RedisBase = RedisWrapper{
		RedisInterface: &Redis{},
	}
}
