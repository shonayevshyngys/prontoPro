package main

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
}

type RatingControllerInterface interface {
	userRoutes(route *gin.Engine)
	providerRoutes(route *gin.Engine)
	reviewRoutes(route *gin.Engine)
	checkRoutes(route *gin.Engine)
}

func (r *RatingController) userRoutes(route *gin.Engine) {
	user := route.Group(baseUrl)
	user.POST("/user", r.handlers.createUser())
}

func (r *RatingController) providerRoutes(route *gin.Engine) {
	provider := route.Group(baseUrl)
	provider.POST("/provider", r.handlers.createProvider())
	provider.GET("/provider/:id", r.handlers.getProvider())
}

func (r *RatingController) reviewRoutes(route *gin.Engine) {
	rating := route.Group(baseUrl)
	rating.POST("/review", r.handlers.createReview())
}

func (r *RatingController) checkRoutes(route *gin.Engine) {
	check := route.Group(baseUrl)
	check.GET("/check/:providerID/:userID", r.handlers.checkIfExists())
}
