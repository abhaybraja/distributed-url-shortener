package test

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"

	"url-shortener/config"
	"url-shortener/database"
	"url-shortener/internal/handlers"
	"url-shortener/repositories"
	"url-shortener/routes"
	"url-shortener/services"
)

// 🛠️ Setup Test App
func setupTestApp() *fiber.App {
	cfg := config.LoadConfig()

	// Initialize Mock Redis (Clear DB Before Running Tests)
	database.InitDB(cfg.RedisHost, cfg.RedisPassword, cfg.RedisDB)
	ctx := context.Background()
	database.RedisClient.FlushAll(ctx)

	// Initialize Dependencies
	urlRepo := repositories.URLRepository{Cfg: cfg}
	urlService := services.URLService{URLRepo: urlRepo}
	analyticsService := services.AnalyticsService{}

	urlHandler := handlers.URLHandler{URLService: urlService}
	analyticsHandler := handlers.AnalyticsHandler{AnalyticsService: analyticsService}

	app := fiber.New()

	// Setup Routes
	routes.SetupRoutes(app, urlHandler, analyticsHandler)

	return app
}

// 🧪 Test URL Shortening
func TestShortenURL(t *testing.T) {
	app := setupTestApp()

	requestBody := []byte(`{"original_url":"https://google.com"}`)
	req := httptest.NewRequest(http.MethodPost, "/api/shorten", bytes.NewBuffer(requestBody))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req, -1)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)

	// Decode Response
	var response map[string]string
	err = json.NewDecoder(resp.Body).Decode(&response)
	assert.NoError(t, err)
	assert.Contains(t, response, "short_code")

	log.Println("✅ Shorten URL Test Passed:", response)
}

// 🧪 Test URL Redirection
func TestRedirectURL(t *testing.T) {
	app := setupTestApp()

	// Manually store URL in Redis
	shortCode := "test123"
	originalURL := "https://example.com"
	database.RedisClient.Set(context.Background(), shortCode, originalURL, 7*24*time.Hour)

	req := httptest.NewRequest(http.MethodGet, "/"+shortCode, nil)
	resp, err := app.Test(req, -1)

	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusMovedPermanently, resp.StatusCode)
	assert.Equal(t, originalURL, resp.Header.Get("Location"))

	log.Println("✅ Redirect URL Test Passed")
}
