package repositories

import (
	"context"
	"encoding/json"
	"time"
	"url-shortener/config"
	"url-shortener/database"
	"url-shortener/models"
)

type URLRepository struct {
	Cfg *config.Config
}

// SaveURL stores a shortened URL in Redis with an expiration time.
func (r *URLRepository) SaveURL(url models.URL) error {
	ctx := context.Background()
	url.CreatedAt = time.Now()
	url.ExpiresAt = time.Now().Add(r.Cfg.URLExpTime)

	data, err := json.Marshal(url)
	if err != nil {
		return err
	}

	err = database.RedisClient.Set(ctx, url.ShortCode, data, r.Cfg.URLExpTime).Err()
	if err != nil {
		return err
	}

	return nil
}

// GetURL retrieves a URL from Redis by its short code.
func (r *URLRepository) GetURL(shortCode string) (*models.URL, error) {
	ctx := context.Background()
	data, err := database.RedisClient.Get(ctx, shortCode).Result()
	if err != nil {
		return nil, err
	}

	var url models.URL
	err = json.Unmarshal([]byte(data), &url)
	if err != nil {
		return nil, err
	}

	return &url, nil
}
