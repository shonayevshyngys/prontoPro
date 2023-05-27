package main

import (
	"log"
	"shonayevshyngys/controllers"
	"shonayevshyngys/database"
	"shonayevshyngys/models"

	"github.com/gin-gonic/gin"
)

func init() {
	//connection
	database.ConnectToDatabase()

	//This one is created only for local testing, for persistent db should be deleted
	var err error
	err = database.Instance.AutoMigrate(&models.Provider{})
	if err != nil {
		log.Fatal("Provider table wasn't created")
	}
	err = database.Instance.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("User table wasn't created")
	}
	err = database.Instance.AutoMigrate(&models.Review{})
	if err != nil {
		log.Fatal("Review table wasn't created")
	}
}

func main() {
	r := gin.Default()
	controllers.UserRoutes(r)
	controllers.ProviderRoutes(r)
	controllers.ReviewRoutes(r)
	err := r.Run()
	if err != nil {
		return
	}
}
