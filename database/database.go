package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DatabaseInstance *gorm.DB

func ConnectToDatabase() {

	dsn := "host=localhost user=postgres password=postgres dbname=rating_service port=5432"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed connect to database")
	}
	DatabaseInstance = db // newer version of golang forces this
}
