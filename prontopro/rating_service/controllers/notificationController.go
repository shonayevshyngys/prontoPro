package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/shonayevshyngys/prontopro/rating_service/database"
	"github.com/shonayevshyngys/prontopro/rating_service/models"
	"github.com/shonayevshyngys/prontopro/rating_service/util"
	"log"
	"net/http"
)

func NotificationRoutes(route *gin.Engine) {
	notification := route.Group("/notification")

	notification.POST("", func(context *gin.Context) {
		log.Println("Saving new notification")
		var notificationBody models.Notification
		err := context.ShouldBindJSON(&notificationBody)
		if err != nil {
			errMsg := util.ErrorMessage{Code: 400, Message: wrongBodyErrorText}
			context.JSON(http.StatusBadRequest, errMsg)
			return
		}
		database.Instance.Create(&notificationBody)
		context.JSON(http.StatusCreated, notificationBody)
	})

}
