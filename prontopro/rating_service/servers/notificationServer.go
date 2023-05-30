package servers

import (
	"github.com/gin-gonic/gin"
	"github.com/shonayevshyngys/prontopro/rating_service/controllers"
	"github.com/shonayevshyngys/prontopro/rating_service/database"
	"github.com/shonayevshyngys/prontopro/rating_service/models"
	"log"
)

func InitRatingServer() {
	database.ConnectToDatabase()
	database.ConnectToRedis()
	var err error
	err = database.Instance.AutoMigrate(&models.Notification{})
	if err != nil {
		log.Fatal("Notification table wasn't created")
	}
}

func RunRatingService() {
	r := gin.Default()
	controllers.NotificationRoutes(r)
	err := r.Run()
	if err != nil {
		log.Fatal(err)
		return
	}
}
