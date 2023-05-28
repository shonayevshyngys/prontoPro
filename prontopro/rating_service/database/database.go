package database

import (
	"fmt"
	"github.com/shonayevshyngys/prontopro/rating_service/models"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Instance *gorm.DB

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
