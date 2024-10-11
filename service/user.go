package service

import (
	"github.com/gin-gonic/gin"
	"github.com/maximianodev/scheduler/clients"
	"github.com/maximianodev/scheduler/models"
)

func GetAllUsersService(context *gin.Context) *[]models.User {
	return clients.GetAllUsers()
}

func GetUserByIdService(userId string) *models.User {
	return clients.GetUserById(userId)
}
