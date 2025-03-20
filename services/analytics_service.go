package services

import (
	"url-shortener/models"
	"url-shortener/repositories"
)

type AnalyticsService struct {
	// AnalyticsRepo repositories.AnalyticsRepository
}

func (s *AnalyticsService) FetchAnalyticsData(shortCode string) (*models.ClickAnalytics, error) {
	analytics, err := repositories.GetAnalytics(shortCode)
	if err != nil {
		return nil, err
	}
	return analytics, nil
}
