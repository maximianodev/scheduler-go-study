package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/maximianodev/scheduler/service"
)

func GetSchedulesController(context *gin.Context) {
	service.GetSchedulesService(context)
}
