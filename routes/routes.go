package routes

import (
	"url-shortener/internal/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, urlHandler handlers.URLHandler, analyticsHandler handlers.AnalyticsHandler) {
	api := app.Group("/api")
	api.Post("/shorten", urlHandler.ShortenURL)
	api.Get("/:shortcode", urlHandler.RedirectURL)
	api.Get("/analytics/:shortcode", analyticsHandler.GetAnalytics)
}
