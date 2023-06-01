package ratingService

import (
	"github.com/gin-gonic/gin"
	"github.com/shonayevshyngys/prontopro/pkg/database"
	"github.com/shonayevshyngys/prontopro/pkg/models"
	"github.com/shonayevshyngys/prontopro/pkg/util"
	"net/http"
	"strconv"
)

func GetRatingHandler() RatingHandler {
	var service = GetRatingService()
	return RatingHandler{service: &service}
}

type RatingHandler struct {
	service RatingServiceInterface
	ratingHandlerInterface
}

type ratingHandlerInterface interface {
	createUser() gin.HandlerFunc
	createProvider() gin.HandlerFunc
	getProvider() gin.HandlerFunc
	createReview() gin.HandlerFunc
	checkIfExists() gin.HandlerFunc
}

// @Summary Creates a user
// @ID createUser
// @Tags Rating
// @Produce json
// @Param User body models.User true "Binding required only for username, id will be adjusted by DB".
// @Success 201 {object} models.User
// @Failure 400 {object} util.ErrorMessage
// @Router /rating/user [post]
func (r *RatingHandler) createUser() gin.HandlerFunc {
	fn := func(context *gin.Context) {
		var userBody models.User
		err := context.ShouldBindJSON(&userBody)
		userBody.ID = 0
		if err != nil {
			errMsg := util.ErrorMessage{Code: 400, Message: wrongBodyErrorText}
			context.JSON(http.StatusBadRequest, errMsg)
			return
		}

		err = r.service.CreateUser(&userBody)
		if err != nil {
			errMsg := util.ErrorMessage{Code: 400, Message: dbErrorText}
			context.JSON(http.StatusBadRequest, errMsg)
		}
		context.JSON(http.StatusCreated, userBody)
	}
	return fn
}

// @Summary Creates a provider
// @ID createProvider
// @Tags Rating
// @Produce json
// @Param Provider body models.Provider true "Binding required only for name and description, id will be adjusted by DB".
// @Success 201 {object} models.Provider
// @Failure 400 {object} util.ErrorMessage
// @Router /rating/provider [post]
func (r *RatingHandler) createProvider() gin.HandlerFunc {
	fn := func(context *gin.Context) {

		var providerBody models.Provider
		err := context.ShouldBindJSON(&providerBody)
		if err != nil {
			errMsg := util.ErrorMessage{Code: 400, Message: wrongBodyErrorText}
			context.JSON(http.StatusBadRequest, errMsg)
			return
		}
		err = r.service.CreateProvider(&providerBody)
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
// @Tags Rating
// @Produce json
// @Param provider_id path int true "id of a provider"
// @Success 200 {object} models.Provider
// @Failure 400 {object} util.ErrorMessage
// @Failure 404 {object} util.ErrorMessage
// @Router /rating/provider/{provider_id} [get]
func (r *RatingHandler) getProvider() gin.HandlerFunc {
	fn := func(context *gin.Context) {
		var provider models.Provider
		id, err := strconv.Atoi(context.Param("id"))
		if err != nil || id < 1 {
			errMsg := util.ErrorMessage{Code: 400, Message: util.BadIdText}
			context.JSON(http.StatusBadRequest, errMsg)
			return
		}

		err = r.service.GetProvider(&provider, id)
		if err != nil {
			errMsg := util.ErrorMessage{Code: 404, Message: "Not found"}
			context.JSON(http.StatusNotFound, errMsg)
			return
		}

		context.JSON(http.StatusOK, provider)
	}
	return fn
}

// @Summary Creates a review binded to user and provider
// @ID createReview
// @Tags Rating
// @Produce json
// @Param Review body util.CreateReviewDTO true "To create review you need to pass userId, providerId, text and rating".
// @Success 201 {object} models.Review
// @Failure 400 {object} util.ErrorMessage
// @Router /rating/review [post]
func (r *RatingHandler) createReview() gin.HandlerFunc {
	fn := func(context *gin.Context) {
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
		review, err := r.service.CreateReview(&reviewBody)
		if err != nil {
			errMsg := util.ErrorMessage{Code: 400, Message: dbErrorText}
			context.JSON(http.StatusBadRequest, errMsg)
			return
		}
		context.JSON(http.StatusCreated, review)

		go r.service.sendNotification(&review)
	}
	return fn
}

// @Summary check if user and provider exists. It's needed for validation on notification service. Only for internal usage
// @ID check
// @Tags Internal
// @Param provider_id path int true "id of a provider"
// @Param user_id path int true "id of a user"
// @Success 200 {object} util.SuccessMessage
// @Failure 400 {object} util.ErrorMessage
// @Failure 404 {object} util.ErrorMessage
// @Router /rating/check/{provider_id}/{user_id} [get]
func (r *RatingHandler) checkIfExists() gin.HandlerFunc {
	fn := func(context *gin.Context) {
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
		if database.DataBase.DBInterface.ProviderExists(uint(providerId)) && database.DataBase.DBInterface.UserExists(uint(userId)) {
			successMsg := util.SuccessMessage{Code: 200, Message: "Both exist in db"}
			context.JSON(http.StatusOK, successMsg)
			return
		} else {
			errMsg := util.ErrorMessage{Code: 404, Message: "Not found"}
			context.JSON(http.StatusBadRequest, errMsg)
		}
	}
	return fn
}
