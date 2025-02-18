package server

import (
	"github.com/Andrew-Nzioki/kenyan-yelp/internal/config"
	"github.com/Andrew-Nzioki/kenyan-yelp/internal/database"
	"github.com/Andrew-Nzioki/kenyan-yelp/internal/handlers"
	"github.com/gin-gonic/gin"
)

func NewGinRouter(cfg *config.Config, db *database.DB) *gin.Engine {
    router := gin.New()

    // Middleware
    router.Use(gin.Logger())
    router.Use(gin.Recovery())

    // Health check
   
    // API routes
    api := router.Group("/api/v1")
    {
        api.GET("/test", handlers.TestHandler)
		api.GET("/health", handlers.HealthCheck)

    }

    return router
}