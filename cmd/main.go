package main

import (
	"log"
	"url-shortener/config"
	"url-shortener/database"
	"url-shortener/internal/handlers"
	"url-shortener/repositories"
	"url-shortener/routes"
	"url-shortener/services"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	cfg := config.LoadConfig()

	// Initialize Database
	database.InitDB(cfg.RedisHost, cfg.RedisPassword, cfg.RedisDB)
	// Initialize Repositories
	urlRepo := repositories.URLRepository{Cfg: cfg}

	// Initialize Services
	urlService := services.URLService{URLRepo: urlRepo}
	analyticsService := services.AnalyticsService{}

	// Initialize Handlers
	urlHandler := handlers.URLHandler{URLService: urlService}
	analyticsHandler := handlers.AnalyticsHandler{AnalyticsService: analyticsService}

	app := fiber.New()

	// Middleware
	app.Use(logger.New()) // Logging requests
	app.Use(cors.New())   // Enable CORS

	routes.SetupRoutes(app, urlHandler, analyticsHandler)

	log.Fatal(app.Listen(":" + cfg.Port))
}
