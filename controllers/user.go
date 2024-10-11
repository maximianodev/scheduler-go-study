package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/maximianodev/scheduler/service"
)

func GetUsersController(context *gin.Context) {
	users := service.GetAllUsersService(context)

	context.IndentedJSON(http.StatusOK, users)
}

func GetUserByIdController(context *gin.Context) {
	userId := context.Param("id")

	user := service.GetUserByIdService(userId)

	if user == nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "User not found"})

		return
	}

	context.IndentedJSON(http.StatusOK, user)
}
