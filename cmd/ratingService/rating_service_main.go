package main

import (
	_ "github.com/shonayevshyngys/prontopro/docs"
	"github.com/shonayevshyngys/prontopro/pkg/database"
	"github.com/shonayevshyngys/prontopro/pkg/models"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func init() {
	//connection
	database.ConnectToDatabase()

	//This one is created only for local testing, for persistent db should be deleted
	var err error
	err = database.DataBase.Instance.AutoMigrate(&models.Provider{})
	if err != nil {
		log.Fatal("Provider table wasn't created")
	}
	err = database.DataBase.Instance.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("User table wasn't created")
	}
	err = database.DataBase.Instance.AutoMigrate(&models.Review{})
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
