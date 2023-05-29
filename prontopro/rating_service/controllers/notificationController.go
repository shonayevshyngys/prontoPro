package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/shonayevshyngys/prontopro/rating_service/models"
	"github.com/shonayevshyngys/prontopro/rating_service/services"
	"github.com/shonayevshyngys/prontopro/rating_service/util"
	"net/http"
)

func NotificationRoutes(route *gin.Engine) {
	notification := route.Group("/notification")

	notification.POST("", func(context *gin.Context) {
		var notificationBody models.Notification
		err := context.ShouldBindJSON(&notificationBody)
		if err != nil {
			errMsg := util.ErrorMessage{Code: 400, Message: wrongBodyErrorText}
			context.JSON(http.StatusBadRequest, errMsg)
			return
		}
		services.SaveNotification(&notificationBody)
		context.JSON(http.StatusCreated, notificationBody)
	})

	notification.POST("/subscribeAsUser", func(context *gin.Context) {

		type subscriberBody struct {
			userId     uint
			providerId uint
		}

		var subscriber subscriberBody
		err := context.ShouldBindJSON(&subscriber)
		if err != nil {
			errMsg := util.ErrorMessage{Code: 400, Message: wrongBodyErrorText}
			context.JSON(http.StatusBadRequest, errMsg)
			return
		}

	})

}
