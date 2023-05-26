package controllers

import (
	"github.com/gin-gonic/gin"
	"shonayevshyngys/Handlers"
)

func UserRoutes(route *gin.Engine) {
	user := route.Group("/user")
	user.POST("", Handlers.CreateUser())
}

func ProviderRoutes(route *gin.Engine) {
	provider := route.Group("/provider")
	provider.POST("")
}

func RatingRoutes(route *gin.Engine) {
	user := route.Group("/rating")
	user.GET("")
	user.POST("")
}
