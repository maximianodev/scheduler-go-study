package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/maximianodev/scheduler/service"
)

func GetAllSchedulesController(context *gin.Context) {
	data := service.GetAllSchedulesService(context)

	context.IndentedJSON(http.StatusOK, data)
}

func GetScheduleByEmailController(context *gin.Context) {
	userEmail := context.Param("email")

	if userEmail == "" {
		return
	}

	schedule := service.GetScheduleByEmailService(userEmail)

	if len(*schedule) == 0 {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Schedule not found"})

		return
	}

	context.IndentedJSON(http.StatusOK, schedule)
}
