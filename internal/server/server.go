package server

import (
	"database/sql"

	"github.com/Andrew-Nzioki/kenyan-yelp/internal/business"
	"github.com/Andrew-Nzioki/kenyan-yelp/internal/business/repository"
	"github.com/Andrew-Nzioki/kenyan-yelp/internal/config"

	"github.com/gin-gonic/gin"
)

// internal/server/router.go
func NewGinRouter(cfg *config.Config, db *sql.DB) *gin.Engine {
    router := gin.Default()

    // Initialize your dependencies
    businessRepo := repository.NewPostgresRepository(db)
    businessService := business.NewService(businessRepo)
    businessHandler := business.NewHandler(businessService)

    // API routes
    v1 := router.Group("/api/v1")
    {
        businesses := v1.Group("/businesses")
        {
            businesses.POST("", businessHandler.CreateBusiness)
            // Add other routes...
        }
    }

    return router
}