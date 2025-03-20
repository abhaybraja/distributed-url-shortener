package handlers

import (
	"fmt"
	"url-shortener/middleware"
	"url-shortener/services"

	"github.com/gofiber/fiber/v2"
)

type URLHandler struct {
	URLService services.URLService
}

type ShortenRequest struct {
	OriginalURL string `json:"original_url"`
}

// RedirectURL handles redirection and click tracking.
func (h *URLHandler) RedirectURL(c *fiber.Ctx) error {
	shortCode := c.Params("shortcode")

	// Retrieve URL from Redis
	url, err := h.URLService.GetURL(shortCode)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Short URL not found"})
	}

	// Track click in Redis
	go middleware.TrackClick(shortCode)

	// Redirect to original URL
	fmt.Println(url.OriginalURL)
	return c.Redirect(url.OriginalURL, fiber.StatusMovedPermanently)
}

func (h *URLHandler) ShortenURL(c *fiber.Ctx) error {
	var req ShortenRequest

	// Parse JSON body into struct
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	if req.OriginalURL == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "original_url is required"})
	}

	// Retrieve URL from Redis
	url, err := h.URLService.GenerateShortURL(req.OriginalURL)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Shorten URL service error!"})
	}

	return c.JSON(fiber.Map{"url": url})
}
