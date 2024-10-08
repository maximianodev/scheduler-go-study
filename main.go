package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/maximianodev/scheduler/models"
)

func getUsers(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, models.Users)
}

func getSchedules(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, models.Schedules)
}

func main() {
	router := gin.Default()

	router.GET("/users", getUsers)
	router.GET("/schedules", getSchedules)

	router.Run("localhost:8080")
}
