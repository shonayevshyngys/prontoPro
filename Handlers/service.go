package Handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shonayevshyngys/database"
	"shonayevshyngys/models"
	"shonayevshyngys/util"
)

func CreateUser() gin.HandlerFunc {
	fn := func(c *gin.Context) {
		var userBody models.User
		err := c.ShouldBindJSON(&userBody)
		if err != nil {
			errMsg := util.ErrorMessage{Code: 400, Message: "Wrong body"}
			c.JSON(http.StatusBadRequest, errMsg)
			return
		}
		database.Instance.Create(&userBody)
		c.JSON(http.StatusCreated, userBody)
	}
	return fn
}
