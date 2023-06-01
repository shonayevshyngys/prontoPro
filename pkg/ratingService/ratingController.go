package ratingService

import (
	"github.com/gin-gonic/gin"
)

const dbErrorText = "Something went wrong during saving to DB"
const wrongBodyErrorText = "Wrong body"
const baseUrl = "/rating"

func GetRatingController() RatingController {
	var handler = GetRatingHandler()
	return RatingController{handlers: &handler}
}

type RatingController struct {
	handlers ratingHandlerInterface
	RatingControllerInterface
}

type RatingControllerInterface interface {
	userRoutes(route *gin.Engine)
	providerRoutes(route *gin.Engine)
	reviewRoutes(route *gin.Engine)
	checkRoutes(route *gin.Engine)
}

func (r *RatingController) UserRoutes(route *gin.Engine) {
	user := route.Group(baseUrl)
	user.POST("/user", r.handlers.createUser())
}

func (r *RatingController) ProviderRoutes(route *gin.Engine) {
	provider := route.Group(baseUrl)
	provider.POST("/provider", r.handlers.createProvider())
	provider.GET("/provider/:id", r.handlers.getProvider())
}

func (r *RatingController) ReviewRoutes(route *gin.Engine) {
	rating := route.Group(baseUrl)
	rating.POST("/review", r.handlers.createReview())
}

func (r *RatingController) CheckRoutes(route *gin.Engine) {
	check := route.Group(baseUrl)
	check.GET("/check/:providerID/:userID", r.handlers.checkIfExists())
}
