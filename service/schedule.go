package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/maximianodev/scheduler/clients"
)

func GetSchedulesService(context *gin.Context) {
	data := clients.GetAllSchedules()

	context.IndentedJSON(http.StatusOK, data)
}
