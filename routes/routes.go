package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/maximianodev/scheduler/controllers"
)

func Routes() *gin.Engine {
	router := gin.Default()

	router.GET("/users", controllers.GetUsersController)
	router.GET("/schedules", controllers.GetSchedulesController)

	return router
}
