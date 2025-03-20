package models

import "time"

// ClickAnalytics tracks the number of visits for a short URL.
type ClickAnalytics struct {
	ShortCode string    `json:"short_code" redis:"short_code"`
	Clicks    int       `json:"clicks" redis:"clicks"`
	UpdatedAt time.Time `json:"updated_at" redis:"updated_at"`
}
