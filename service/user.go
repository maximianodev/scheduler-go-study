package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/maximianodev/scheduler/clients"
)

func GetUsersService(context *gin.Context) {
	data := clients.GetUsers()

	context.IndentedJSON(http.StatusOK, data)
}
