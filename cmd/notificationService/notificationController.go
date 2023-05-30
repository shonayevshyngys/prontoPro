package main

import (
	"github.com/gin-gonic/gin"
	"github.com/shonayevshyngys/prontopro/pkg/models"
	"github.com/shonayevshyngys/prontopro/pkg/util"
	"net/http"
	"strconv"
)

const wrongBodyErrorText = "Wrong body"

func NotificationRoutes(route *gin.Engine) {
	notification := route.Group("/notification")

	notification.POST("/", func(context *gin.Context) {
		var notificationBody models.Notification
		err := context.ShouldBindJSON(&notificationBody)
		if err != nil {
			errMsg := util.ErrorMessage{Code: 400, Message: wrongBodyErrorText}
			context.JSON(http.StatusBadRequest, errMsg)
			return
		}
		SaveNotification(&notificationBody)
		context.JSON(http.StatusCreated, notificationBody)
	})

	notification.GET("/provider/:id", func(context *gin.Context) {

		id, err := strconv.Atoi(context.Param("id"))
		if err != nil || id < 1 {
			errMsg := util.ErrorMessage{Code: 400, Message: util.BadIdText}
			context.JSON(http.StatusBadRequest, errMsg)
			return
		}
		notifications, err := GetProviderNotifications(id)
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
			errMsg := util.ErrorMessage{Code: 400, Message: util.BadIdText}
			context.JSON(http.StatusBadRequest, errMsg)
			return
		}
		notifications, err := GetUserNotifications(id)
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
		err = SubscribeUserToProvider(subscriptionBody.ProviderId, subscriptionBody.UserId, &subscriptionBody)
		if err != nil {
			context.JSON(http.StatusBadRequest, err)
			return
		}
		successMsg := util.SuccessMessage{Code: 200, Message: "subscribed"}
		context.JSON(http.StatusOK, successMsg)
	})
}
