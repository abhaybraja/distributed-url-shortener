package test

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"

	"url-shortener/database"
)

// 🧪 Test Click Analytics
func TestGetAnalytics(t *testing.T) {
	app := setupTestApp()

	// Store Analytics Data in Redis
	shortCode := "click123"
	clickData := map[string]interface{}{
		"short_code": shortCode,
		"clicks":     5,
		"updated_at": time.Now().Format(time.RFC3339),
	}

	// Convert to JSON
	clickDataJSON, _ := json.Marshal(clickData)
	database.RedisClient.Set(context.Background(), "analytics:"+shortCode, clickDataJSON, 30*24*time.Hour)

	req := httptest.NewRequest(http.MethodGet, "/api/analytics/"+shortCode, nil)
	resp, err := app.Test(req, -1)

	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)

	// Decode Response
	var response map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&response)
	assert.NoError(t, err)
	assert.Equal(t, shortCode, response["short_code"])
	assert.GreaterOrEqual(t, int(response["clicks"].(float64)), 5)

	log.Println("✅ Analytics Test Passed:", response)
}
