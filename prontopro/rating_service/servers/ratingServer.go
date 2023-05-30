package servers

import (
	"github.com/shonayevshyngys/prontopro/rating_service/controllers"
	"github.com/shonayevshyngys/prontopro/rating_service/database"
	"github.com/shonayevshyngys/prontopro/rating_service/models"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func InitNotificationService() {
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

func RunNotificationService() {
	log.Println(os.Getenv("DATASOURCE"))
	log.Println(os.Getenv("PORT"))
	r := gin.Default()
	controllers.UserRoutes(r)
	controllers.ProviderRoutes(r)
	controllers.ReviewRoutes(r)
	controllers.CheckRoutes(r)
	err := r.Run()
	if err != nil {
		log.Fatal(err)
		return
	}
}
