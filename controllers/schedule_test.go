package controllers_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/maximianodev/scheduler/controllers"
	"github.com/stretchr/testify/assert"
)

func GetScheduleByIdController(t *testing.T) {
	expected := gin.H{"message": "Schedule not found"}
	expectedJSON, _ := json.Marshal(expected)

	router := gin.Default()
	router.GET("/users/:email", controllers.GetScheduleByEmailController)

	recorder := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/users/asdasdas", nil)
	router.ServeHTTP(recorder, req)

	assert.JSONEq(t, string(expectedJSON), recorder.Body.String())
}
