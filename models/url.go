package models

import "time"

// URL represents the shortened URL entity.
type URL struct {
	ShortCode   string    `json:"short_code" redis:"short_code"`
	OriginalURL string    `json:"original_url" redis:"original_url"`
	CreatedAt   time.Time `json:"created_at" redis:"created_at"`
	ExpiresAt   time.Time `json:"expires_at" redis:"expires_at"`
}
