package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/shonayevshyngys/prontopro/pkg/database"
	"github.com/shonayevshyngys/prontopro/pkg/models"
	"github.com/shonayevshyngys/prontopro/pkg/util"
	"log"
	"net/http"
	"strconv"
)

const dbErrorText = "Something went wrong during saving to DB"
const wrongBodyErrorText = "Wrong body"
const baseUrl = "/rating"

func UserRoutes(route *gin.Engine) {
	user := route.Group(baseUrl)
	user.POST("/user", createUser())
}

// @Summary Creates a user
// @ID createUser
// @Param User body models.User.Username true "Binding required only for username, id will be adjusted by DB".
// @Success 201 {object} models.User
// @Failure 400 {object} util.ErrorMessage
// @Router /rating/user [post]
func createUser() gin.HandlerFunc {
	fn := func(context *gin.Context) {
		var userBody models.User
		err := context.ShouldBindJSON(&userBody)
		userBody.ID = 0
		if err != nil {
			errMsg := util.ErrorMessage{Code: 400, Message: wrongBodyErrorText}
			context.JSON(http.StatusBadRequest, errMsg)
			return
		}
		err = CreateUser(&userBody)
		if err != nil {
			errMsg := util.ErrorMessage{Code: 400, Message: dbErrorText}
			context.JSON(http.StatusBadRequest, errMsg)
		}
		context.JSON(http.StatusCreated, userBody)
	}
	return fn
}

func ProviderRoutes(route *gin.Engine) {
	provider := route.Group(baseUrl)
	provider.POST("/provider", createProvider())
	provider.GET("/provider/:id", getProvider())
}

// @Summary Creates a provider
// @ID createProvider
// @Param Provider body models.Provider true "Binding required only for name and description, id will be adjusted by DB".
// @Success 201 {object} models.Provider
// @Failure 400 {object} util.ErrorMessage
// @Router /rating/provider [post]
func createProvider() gin.HandlerFunc {
	fn := func(context *gin.Context) {

		var providerBody models.Provider
		err := context.ShouldBindJSON(&providerBody)
		if err != nil {
			errMsg := util.ErrorMessage{Code: 400, Message: wrongBodyErrorText}
			context.JSON(http.StatusBadRequest, errMsg)
			return
		}
		err = CreateProvider(&providerBody)
		if err != nil {
			errMsg := util.ErrorMessage{Code: 400, Message: dbErrorText}
			context.JSON(http.StatusBadRequest, errMsg)
			return
		}
		context.JSON(http.StatusCreated, providerBody)
	}
	return fn
}

// @Summary Gets a provider with average rating from review
// @ID getProvider
// @Tags rating
// @Param provider_id path int true "id of a provider"
// @Success 200 {object} models.Provider
// @Failure 400 {object} util.ErrorMessage
// @Failure 404 {object} util.ErrorMessage
// @Router /rating/{provider_id} [get]
func getProvider() gin.HandlerFunc {
	fn := func(context *gin.Context) {
		var provider models.Provider
		id, err := strconv.Atoi(context.Param("id"))
		if err != nil || id < 1 {
			errMsg := util.ErrorMessage{Code: 400, Message: util.BadIdText}
			context.JSON(http.StatusBadRequest, errMsg)
			return
		}

		err = GetProvider(&provider, id)
		if err != nil {
			errMsg := util.ErrorMessage{Code: 404, Message: "Not found"}
			context.JSON(http.StatusNotFound, errMsg)
			return
		}

		context.JSON(http.StatusOK, provider)
	}
	return fn
}

func ReviewRoutes(route *gin.Engine) {
	rating := route.Group(baseUrl)

	rating.POST("/review", func(context *gin.Context) {
		var reviewBody util.CreateReviewDTO
		err := context.ShouldBindJSON(&reviewBody)
		if err != nil || reviewBody.UserId < 1 || reviewBody.ProviderId < 1 {
			errMsg := util.ErrorMessage{Code: 400, Message: wrongBodyErrorText}
			context.JSON(http.StatusBadRequest, errMsg)
			return
		}
		if reviewBody.Rating < 0 || reviewBody.Rating > 5 {
			errMsg := util.ErrorMessage{Code: 400, Message: "Rating can be in range of 0 to 5"}
			context.JSON(http.StatusBadRequest, errMsg)
			return
		}
		review, err := CreateReview(&reviewBody)
		if err != nil {
			errMsg := util.ErrorMessage{Code: 400, Message: dbErrorText}
			context.JSON(http.StatusBadRequest, errMsg)
			return
		}
		context.JSON(http.StatusCreated, review)

		go func() {

			notification := models.Notification{
				ProviderID:   review.ProviderID,
				Notification: fmt.Sprintf("New rating %d submitted by %s", review.Rating, review.User.Username),
			}

			errNotif := util.SaveNotification(&notification)
			if errNotif != nil {
				log.Println("Something went wrong during saving notification ", err)
			}
		}()
	})
}

func CheckRoutes(route *gin.Engine) {
	check := route.Group(baseUrl)
	check.GET("/check/:providerID/:userID", func(context *gin.Context) {
		providerId, err := strconv.Atoi(context.Param("providerID"))
		if err != nil || providerId < 1 {
			errMsg := util.ErrorMessage{Code: 400, Message: "Bad format for id"}
			context.JSON(http.StatusBadRequest, errMsg)
			return
		}
		userId, err := strconv.Atoi(context.Param("userID"))
		if err != nil || userId < 1 {
			errMsg := util.ErrorMessage{Code: 400, Message: "Bad format for id"}
			context.JSON(http.StatusBadRequest, errMsg)
			return
		}
		if database.ProviderExists(uint(providerId)) && database.UserExists(uint(userId)) {
			context.JSON(http.StatusOK, "ok")
			return
		} else {
			context.JSON(http.StatusBadRequest, "not ok")
		}
	})
}
