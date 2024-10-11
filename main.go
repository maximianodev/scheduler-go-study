package main

import (
	"github.com/gin-gonic/gin"
	"github.com/maximianodev/scheduler/controllers"
)

func main() {
	router := gin.Default()

	router.GET("/users", controllers.GetUsersController)
	router.GET("/users/:id", controllers.GetUserByIdController)

	router.GET("/schedules", controllers.GetSchedulesController)

	router.Run("localhost:8080")
}
