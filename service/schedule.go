package service

import (
	"github.com/gin-gonic/gin"
	"github.com/maximianodev/scheduler/clients"
	"github.com/maximianodev/scheduler/models"
)

func GetAllSchedulesService(context *gin.Context) *[]models.Schedule {
	return clients.ListAllSchedules()
}

func GetScheduleByEmailService(userEmail string) *[]models.Schedule {
	return clients.ListScheduleByEmail(userEmail)
}
