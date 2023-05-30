package main

import (
	"github.com/gin-gonic/gin"
	"github.com/shonayevshyngys/prontopro/pkg/database"
	"github.com/shonayevshyngys/prontopro/pkg/models"
	"log"
)

func init() {
	database.ConnectToDatabase()
	database.ConnectToRedis()
	var err error
	err = database.Instance.AutoMigrate(&models.Notification{})
	if err != nil {
		log.Fatal("Notification table wasn't created")
	}
}

func main() {
	r := gin.Default()
	NotificationRoutes(r)
	err := r.Run()
	if err != nil {
		log.Fatal(err)
		return
	}
}
