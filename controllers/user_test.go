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

func TestGetUserByIdController(t *testing.T) {
	expected := gin.H{"message": "User not found"}
	expectedJSON, _ := json.Marshal(expected)

	router := gin.Default()
	router.GET("/users/:id", controllers.GetUserByIdController)

	recorder := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/users/7766778", nil)
	router.ServeHTTP(recorder, req)

	assert.JSONEq(t, string(expectedJSON), recorder.Body.String())
}
