package repositories

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
	"url-shortener/database"
	"url-shortener/models"
)

// GetAnalyticsData retrieves data from Redis by its short code.
func GetAnalyticsAndInit(shortCode string) (*models.ClickAnalytics, error) {
	ctx := context.Background()
	data, err := database.RedisClient.Get(ctx, "analytics:"+shortCode).Result()
	var analytics models.ClickAnalytics

	if err == nil {
		// Unmarshal existing data
		err = json.Unmarshal([]byte(data), &analytics)
		if err != nil {
			return nil, err
		}
	} else {
		// Initialize if not found
		analytics = models.ClickAnalytics{
			ShortCode: shortCode,
			Clicks:    0,
			UpdatedAt: time.Now(),
		}
	}

	return &analytics, nil
}

func GetAnalytics(shortCode string) (*models.ClickAnalytics, error) {
	ctx := context.Background()
	data, err := database.RedisClient.Get(ctx, "analytics:"+shortCode).Result()

	// Parse JSON data
	var analytics models.ClickAnalytics
	err = json.Unmarshal([]byte(data), &analytics)
	return &analytics, err
}

func SaveAnalytics(shortCode string, analytics *models.ClickAnalytics) error {
	ctx := context.Background()
	updatedData, _ := json.Marshal(analytics)
	err := database.RedisClient.Set(ctx, "analytics:"+shortCode, updatedData, 30*24*time.Hour).Err() // 30 days retention

	if err != nil {
		fmt.Println("❌ Failed to track click:", err)
		return err
	}

	fmt.Printf("✅ Click tracked for %s: %d clicks\n", shortCode, analytics.Clicks)
	return nil
}
