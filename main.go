package main

import (
	"github.com/gin-gonic/gin"
	"github.com/maximianodev/scheduler/controllers"
)

func main() {
	router := gin.Default()

	router.GET("/users", controllers.GetAllUsersController)
	router.GET("/users/:id", controllers.GetUserByIdController)

	router.GET("/schedules", controllers.GetAllSchedulesController)
	router.GET("/schedules/:email", controllers.GetScheduleByEmailController)

	router.Run("localhost:8080")
}
