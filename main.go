package main

import (
	"shonayevshyngys/database"
	"shonayevshyngys/models"

	"github.com/gin-gonic/gin"
)

func init() {
	//connection
	database.ConnectToDatabase()

	//This one is created only for local testing, for persistent db should be deleted
	database.DatabaseInstance.AutoMigrate()
	database.DatabaseInstance.AutoMigrate(&models.Provider{})
	database.DatabaseInstance.AutoMigrate(&models.User{})
}

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
