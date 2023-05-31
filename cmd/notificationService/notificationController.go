package main

import (
	"github.com/gin-gonic/gin"
)

const wrongBodyErrorText = "Wrong body"

func GetNotificationController() NotificationController {
	var handlers = GetNotificationHandler()
	var controller = NotificationController{handlers: &handlers}
	return controller
}

type NotificationController struct {
	handlers NotificationHandlerInterface
	NotificationControllerInterface
}

type NotificationControllerInterface interface {
	NotificationRoutes(route *gin.Engine)
}

func (controller *NotificationController) NotificationRoutes(route *gin.Engine) {
	notification := route.Group("/notification")

	notification.POST("/", controller.handlers.createNotification())

	notification.GET("/provider/:id", controller.handlers.getProviderNotification())

	notification.GET("/user/:id", controller.handlers.getUserNotifications())

	notification.POST("/subscribe", controller.handlers.subscribeUserToProvider())
}
