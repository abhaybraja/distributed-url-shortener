package handlers

import (
	"url-shortener/services"

	"github.com/gofiber/fiber/v2"
)

type AnalyticsHandler struct {
	AnalyticsService services.AnalyticsService
}

func (h *AnalyticsHandler) GetAnalytics(c *fiber.Ctx) error {
	shortCode := c.Params("shortcode")

	analytics, err := h.AnalyticsService.FetchAnalyticsData(shortCode)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "No analytics found for this URL"})
	}

	// Return analytics data
	return c.JSON(fiber.Map{
		"short_code":   shortCode,
		"clicks":       analytics.Clicks,
		"last_clicked": analytics.UpdatedAt.Format("2006-01-02 15:04:05"),
	})
}
