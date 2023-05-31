package main

import (
	"github.com/gin-gonic/gin"
	"github.com/shonayevshyngys/prontopro/pkg/models"
	"github.com/shonayevshyngys/prontopro/pkg/util"
	"net/http"
	"strconv"
)

func GetNotificationHandler() NotificationHandler {
	var service = GetNotificationService()
	var handler = NotificationHandler{
		service: &service,
	}
	return handler
}

type NotificationHandler struct {
	service notificationServiceInterface
	NotificationHandlerInterface
}

type NotificationHandlerInterface interface {
	createNotification() gin.HandlerFunc
	subscribeUserToProvider() gin.HandlerFunc
	getProviderNotification() gin.HandlerFunc
	getUserNotifications() gin.HandlerFunc
}

// @Summary Creates a notification object
// @ID createNotification
// @Tags Internal
// @Produce json
// @Param Notification body models.Notification true "To create notification you need to pass providerId and notification text".
// @Success 201 {object} models.Notification
// @Failure 400 {object} util.ErrorMessage
// @Router /notification [post]
func (h *NotificationHandler) createNotification() gin.HandlerFunc {
	fn := func(context *gin.Context) {
		var notificationBody models.Notification
		err := context.ShouldBindJSON(&notificationBody)
		if err != nil {
			errMsg := util.ErrorMessage{Code: 400, Message: wrongBodyErrorText}
			context.JSON(http.StatusBadRequest, errMsg)
			return
		}
		h.service.saveNotification(&notificationBody)
		context.JSON(http.StatusCreated, notificationBody)
	}
	return fn
}

// @Summary Subscribes user to a provider
// @ID Subscribe
// @Tags Notification
// @Produce json
// @Param SubscriptionBody body util.SubscriptionBody true "Subscribes user to provider to get notifications".
// @Success 200 {object} util.SuccessMessage
// @Failure 400 {object} util.ErrorMessage
// @Router /notification/subscribe [post]
func (h *NotificationHandler) subscribeUserToProvider() gin.HandlerFunc {
	fn := func(context *gin.Context) {
		var subscriptionBody util.SubscriptionBody
		err := context.ShouldBindJSON(&subscriptionBody)
		if err != nil {
			errMsg := util.ErrorMessage{Code: 400, Message: wrongBodyErrorText}
			context.JSON(http.StatusBadRequest, errMsg)
			return
		}
		err = h.service.subscribeUserToProvider(subscriptionBody.ProviderId, subscriptionBody.UserId, &subscriptionBody)
		if err != nil {
			context.JSON(http.StatusBadRequest, err)
			return
		}
		successMsg := util.SuccessMessage{Code: 200, Message: "subscribed"}
		context.JSON(http.StatusOK, successMsg)
	}
	return fn
}

// @Summary Get provider's notification
// @ID getProviderNotifications
// @Tags Notification
// @Produce json
// @Param provider_id path int true "id of a provider"
// @Success 200 {object} util.SuccessMessage
// @Failure 400 {object} util.ErrorMessage
// @Router /notification/provider/{provider_id} [get]
func (h *NotificationHandler) getProviderNotification() gin.HandlerFunc {
	fn := func(context *gin.Context) {
		id, err := strconv.Atoi(context.Param("id"))
		if err != nil || id < 1 {
			errMsg := util.ErrorMessage{Code: 400, Message: util.BadIdText}
			context.JSON(http.StatusBadRequest, errMsg)
			return
		}
		notifications, err := h.service.getProviderNotifications(id)
		if err != nil || len(notifications) == 0 {
			errMsg := util.ErrorMessage{Code: 400, Message: "Not found entries"}
			context.JSON(http.StatusNotFound, errMsg)
			return
		}
		context.JSON(http.StatusOK, notifications)
	}
	return fn
}

// @Summary Get user's subbed notifications
// @ID getUserNotifications
// @Tags Notification
// @Produce json
// @Param user_id path int true "id of a user"
// @Success 200 {object} util.SuccessMessage
// @Failure 400 {object} util.ErrorMessage
// @Router /notification/user/{user_id} [get]
func (h *NotificationHandler) getUserNotifications() gin.HandlerFunc {
	fn := func(context *gin.Context) {
		id, err := strconv.Atoi(context.Param("id"))
		if err != nil || id < 1 {
			errMsg := util.ErrorMessage{Code: 400, Message: util.BadIdText}
			context.JSON(http.StatusBadRequest, errMsg)
			return
		}
		notifications, err := h.service.getUserNotifications(id)
		if err != nil || len(notifications) == 0 {
			errMsg := util.ErrorMessage{Code: 404, Message: "Not found entries"}
			context.JSON(http.StatusNotFound, errMsg)
			return
		}
		context.JSON(http.StatusOK, notifications)
	}
	return fn
}
