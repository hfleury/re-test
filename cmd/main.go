package main

import (
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/hfleury/re-test/config"
	"github.com/hfleury/re-test/internal/handlers"
	"github.com/hfleury/re-test/internal/services"
)

func main() {
	AppConfig, err := config.NewConfig("config/configuration.yaml")
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Initlialize services
	packSizesService := services.NewPackSizeService(AppConfig.PackSize)
	// Initialize handlers
	packSizeHandler := handlers.NewPackSizeHandler(packSizesService)

	router := gin.Default()

	// Set up CORS
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Initialize routes
	packSizeHandler.RegisterRoutes(router)

	log.Println("Starting server on port 8081")
	if err := router.Run(":8081"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
