package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/shonayevshyngys/prontopro/rating_service/models"
	"github.com/shonayevshyngys/prontopro/rating_service/services"
	"github.com/shonayevshyngys/prontopro/rating_service/util"
	"net/http"
	"strconv"
)

func NotificationRoutes(route *gin.Engine) {
	notification := route.Group("")

	notification.POST("/notification", func(context *gin.Context) {
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

	notification.GET("/provider/:id", func(context *gin.Context) {

		id, err := strconv.Atoi(context.Param("id"))
		if err != nil || id < 1 {
			errMsg := util.ErrorMessage{Code: 400, Message: "Bad format for id"}
			context.JSON(http.StatusBadRequest, errMsg)
			return
		}
		notifications, err := services.GetProviderNotifications(id)
		if err != nil || len(notifications) == 0 {
			errMsg := util.ErrorMessage{Code: 400, Message: "Not found entries"}
			context.JSON(http.StatusNotFound, errMsg)
			return
		}
		context.JSON(http.StatusOK, notifications)
	})

	notification.GET("/user/:id", func(context *gin.Context) {
		id, err := strconv.Atoi(context.Param("id"))
		if err != nil || id < 1 {
			errMsg := util.ErrorMessage{Code: 400, Message: "Bad format for id"}
			context.JSON(http.StatusBadRequest, errMsg)
			return
		}
		notifications, err := services.GetUserNotifications(id)
		if err != nil || len(notifications) == 0 {
			errMsg := util.ErrorMessage{Code: 400, Message: "Not found entries"}
			context.JSON(http.StatusNotFound, errMsg)
			return
		}
		context.JSON(http.StatusOK, notifications)
	})

	notification.POST("/subscribe", func(context *gin.Context) {
		var subscriptionBody util.SubscriptionBody
		err := context.ShouldBindJSON(&subscriptionBody)
		if err != nil {
			errMsg := util.ErrorMessage{Code: 400, Message: wrongBodyErrorText}
			context.JSON(http.StatusBadRequest, errMsg)
			return
		}
		err = services.SubscribeUserToProvider(subscriptionBody.ProviderId, subscriptionBody.UserId, &subscriptionBody)
		if err != nil {
			context.JSON(http.StatusBadRequest, err)
			return
		}
		context.JSON(http.StatusOK, "subscribed")
	})
}
