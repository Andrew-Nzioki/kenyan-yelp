package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	docs "github.com/Andrew-Nzioki/kenyan-yelp/docs"
	"github.com/Andrew-Nzioki/kenyan-yelp/internal/config"
	"github.com/Andrew-Nzioki/kenyan-yelp/internal/database"
	"github.com/Andrew-Nzioki/kenyan-yelp/internal/server"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Kenyan Yelp API
// @version         1.0
// @description     A Yelp-like service for Kenyan businesses
// @host            localhost:8080
// @BasePath        /api/v1
// @schemes         http
func main() {
    if os.Getenv("GIN_MODE") != "debug" {
        gin.SetMode(gin.ReleaseMode)
    }

    cfg, err := config.Load()
    if err != nil {
        log.Fatalf("Failed to load configuration: %v", err)
    }

    db, err := database.Connect(cfg.DatabaseURL)
    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }
    defer db.Close()

    router := server.NewGinRouter(cfg, db)
    
    // Swagger docs
    docs.SwaggerInfo.BasePath = "/api/v1"
    router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

    srv := &http.Server{
        Addr:         cfg.ServerAddr,
        Handler:      router,
        ReadTimeout:  cfg.ReadTimeout,
        WriteTimeout: cfg.WriteTimeout,
        IdleTimeout:  time.Minute,
    }

    go func() {
        log.Printf("Starting server on %s", cfg.ServerAddr)
        if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
            log.Fatalf("Server failed: %v", err)
        }
    }()

    quit := make(chan os.Signal, 1)
    signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
    <-quit
    log.Println("Shutting down server...")

    ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
    defer cancel()

    if err := srv.Shutdown(ctx); err != nil {
        log.Fatalf("Server forced to shutdown: %v", err)
    }

    log.Println("Server exited gracefully")
}