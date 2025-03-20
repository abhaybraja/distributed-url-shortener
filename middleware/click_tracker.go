package middleware

import (
	"fmt"
	"time"
	"url-shortener/repositories"
)

// TrackClick increments the visit count for a given short code.
func TrackClick(shortCode string) error {

	// Get current analytics data
	analytics, err := repositories.GetAnalyticsAndInit(shortCode)
	if err != nil {
		fmt.Println("Failed to get analytics data:", err)
		return err
	}

	// Increment the click count
	analytics.Clicks++
	analytics.UpdatedAt = time.Now()

	// Save updated analytics
	err = repositories.SaveAnalytics(shortCode, analytics)
	return err
}
