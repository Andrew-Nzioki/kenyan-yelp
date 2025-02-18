package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// TestResponse represents the test endpoint response
type TestResponse struct {
    Message string    `json:"message" example:"Hello from Kenyan Yelp API"`
    Time    time.Time `json:"time" example:"2024-02-18T12:00:00Z"`
}

// TestHandler godoc
// @Summary     Test endpoint
// @Description Get a test response from the API
// @Tags        test
// @Accept      json
// @Produce     json
// @Success     200 {object} TestResponse
// @Router      /test [get]
func TestHandler(c *gin.Context) {
    c.JSON(http.StatusOK, TestResponse{
        Message: "Live !!!!!",
        Time:    time.Now(),
    })
}

// HealthCheck godoc
// @Summary     Health check
// @Description Check if the API is running
// @Tags        health
// @Produce     plain
// @Success     200 {string} string "OK"
// @Router      /health [get]
func HealthCheck(c *gin.Context) {
    c.String(http.StatusOK, "OK")
}