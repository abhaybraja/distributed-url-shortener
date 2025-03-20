package services

import (
	"math/rand"
	"url-shortener/models"
	"url-shortener/repositories"
)

type URLService struct {
	URLRepo repositories.URLRepository
}

func (s *URLService) GenerateShortURL(originalURL string) (string, error) {
	shortCode := generateRandomCode(6)
	url := models.URL{ShortCode: shortCode, OriginalURL: originalURL}
	err := s.URLRepo.SaveURL(url)
	if err != nil {
		return "", err
	}
	return shortCode, nil
}

func (s *URLService) GetURL(shortCode string) (*models.URL, error) {
	url, err := s.URLRepo.GetURL(shortCode)
	return url, err
}

func generateRandomCode(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	code := make([]byte, length)
	for i := range code {
		code[i] = charset[rand.Intn(len(charset))]
	}
	return string(code)
}
