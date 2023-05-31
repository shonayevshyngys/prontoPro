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
	err = database.DataBase.Instance.AutoMigrate(&models.Notification{})
	if err != nil {
		log.Fatal("Notification table wasn't created")
	}
}

// @title ProntoPro
// @version 1.0
// @description This is a take home assignment for pronto pro

// @host localhost:80
// @BasePath /
// @query.collection.format multi
func main() {
	r := gin.Default()
	var controller = GetNotificationController()
	controller.NotificationRoutes(r)
	err := r.Run()
	if err != nil {
		log.Fatal(err)
		return
	}
}
