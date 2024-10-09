package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/maximianodev/scheduler/service"
)

func GetUsersController(context *gin.Context) {
	service.GetUsersService(context)
}
