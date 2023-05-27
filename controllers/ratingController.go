package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shonayevshyngys/models"
	"shonayevshyngys/services"
	"shonayevshyngys/util"
)

const dbErrorText = "Something went wrong during saving to DB"
const wrongBodyErrorText = "Wrong body"

func UserRoutes(route *gin.Engine) {
	user := route.Group("/user")

	user.POST("", func(context *gin.Context) {
		var userBody models.User
		err := context.ShouldBindJSON(&userBody)
		userBody.ID = 0
		if err != nil {
			errMsg := util.ErrorMessage{Code: 400, Message: wrongBodyErrorText}
			context.JSON(http.StatusBadRequest, errMsg)
			return
		}
		err = services.CreateUser(&userBody)
		if err != nil {
			errMsg := util.ErrorMessage{Code: 400, Message: dbErrorText}
			context.JSON(http.StatusBadRequest, errMsg)
		}
		context.JSON(http.StatusCreated, userBody)
	})
}

func ProviderRoutes(route *gin.Engine) {
	provider := route.Group("/provider")

	provider.POST("", func(context *gin.Context) {
		var providerBody models.Provider
		err := context.ShouldBindJSON(&providerBody)
		providerBody.ID = 0
		if err != nil {
			errMsg := util.ErrorMessage{Code: 400, Message: wrongBodyErrorText}
			context.JSON(http.StatusBadRequest, errMsg)
			return
		}
		err = services.CreateProvider(&providerBody)
		if err != nil {
			errMsg := util.ErrorMessage{Code: 400, Message: dbErrorText}
			context.JSON(http.StatusBadRequest, errMsg)
			return
		}
		context.JSON(http.StatusCreated, providerBody)
	})
	provider.GET("")
}

func ReviewRoutes(route *gin.Engine) {
	rating := route.Group("/review")
	rating.POST("", func(context *gin.Context) {
		var reviewBody models.Review
		err := context.ShouldBindJSON(&reviewBody)
		reviewBody.ID = 0
		if err != nil {
			errMsg := util.ErrorMessage{Code: 400, Message: wrongBodyErrorText}
			context.JSON(http.StatusBadRequest, errMsg)
			return
		}
		err = services.CreateReview(&reviewBody)
		if err != nil {
			errMsg := util.ErrorMessage{Code: 400, Message: dbErrorText}
			context.JSON(http.StatusBadRequest, errMsg)
			return
		}
		context.JSON(http.StatusCreated, reviewBody)

	})
}
