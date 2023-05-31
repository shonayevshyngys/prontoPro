package main

import (
	_ "github.com/shonayevshyngys/prontopro/docs"
	"github.com/shonayevshyngys/prontopro/pkg/database"
	"github.com/shonayevshyngys/prontopro/pkg/models"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func setup() {
	//connection
	database.ConnectToDatabase()

	//This one is created only for local testing, for persistent db should be deleted
	var err error
	err = database.GormInstance.AutoMigrate(&models.Provider{})
	if err != nil {
		log.Fatal("Provider table wasn't created")
	}
	err = database.GormInstance.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("User table wasn't created")
	}
	err = database.GormInstance.AutoMigrate(&models.Review{})
	if err != nil {
		log.Fatal("Review table wasn't created")
	}
}

// @title ProntoPro
// @version 1.0
// @description This is a take home assignment for pronto pro

// @host localhost:80
// @BasePath /
// @query.collection.format multi

func main() {
	setup()
	log.Println(os.Getenv("DATASOURCE"))
	log.Println(os.Getenv("PORT"))
	r := gin.Default()

	controller := GetRatingController()
	controller.checkRoutes(r)
	controller.providerRoutes(r)
	controller.reviewRoutes(r)
	controller.userRoutes(r)
	err := r.Run()
	if err != nil {
		log.Fatal(err)
		return
	}
}
