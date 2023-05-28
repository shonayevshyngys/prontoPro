package main

import (
	"github.com/gin-gonic/gin"
	"github.com/shonayevshyngys/prontopro/rating_service/controllers"
	"github.com/shonayevshyngys/prontopro/rating_service/database"
	"github.com/shonayevshyngys/prontopro/rating_service/models"
	"log"
	"os"
)

func init() {
	database.ConnectToDatabase()
	var err error
	err = database.Instance.AutoMigrate(&models.Notification{})
	if err != nil {
		log.Fatal("Notification table wasn't created")
	}
}

func main() {
	log.Println(os.Getenv("DATASOURCE"))
	log.Println(os.Getenv("PORT"))
	r := gin.Default()
	controllers.NotificationRoutes(r)
	err := r.Run()
	if err != nil {
		log.Fatal(err)
		return
	}
}
