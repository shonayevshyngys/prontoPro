package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/shonayevshyngys/prontopro/rating_service/models"
	"github.com/shonayevshyngys/prontopro/rating_service/services"
	"github.com/shonayevshyngys/prontopro/rating_service/util"
	"net/http"
	"strconv"
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
	provider.GET("/:id", func(context *gin.Context) {
		var provider models.Provider
		id, err := strconv.Atoi(context.Param("id"))
		if err != nil || id < 1 {
			errMsg := util.ErrorMessage{Code: 400, Message: "Bad format for id"}
			context.JSON(http.StatusBadRequest, errMsg)
			return
		}

		err = services.GetProvider(&provider, id)
		if err != nil {
			errMsg := util.ErrorMessage{Code: 404, Message: "Not found"}
			context.JSON(http.StatusNotFound, errMsg)
			return
		}

		context.JSON(http.StatusOK, provider)
	})
}

func ReviewRoutes(route *gin.Engine) {
	rating := route.Group("/review")
	rating.POST("", func(context *gin.Context) {
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
		review, err := services.CreateReview(&reviewBody)
		if err != nil {
			errMsg := util.ErrorMessage{Code: 400, Message: dbErrorText}
			context.JSON(http.StatusBadRequest, errMsg)
			return
		}
		context.JSON(http.StatusCreated, review)

	})
}
