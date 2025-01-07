package shortener

import (
	"fmt"

	"github.com/hdurham99/skinny-url/internal/storage"
)

type Service struct {
	DB     storage.DBStorage
	Cache  storage.CacheStorage
	Domain string
}

func NewShortenerService(db storage.DBStorage, cache storage.CacheStorage, domain string) *Service {
	return &Service{
		DB:     db,
		Cache:  cache,
		Domain: domain,
	}
}

func (s *Service) ShortenURL(original_url string) (string, error) {
	const retryLimit = 5

	for i := 0; i < retryLimit; i++ {
		offsetInput := fmt.Sprintf("%s-%d", original_url, i)
		short_code := encodeUrl(offsetInput)

		// Checking for collision
		existingURL, err := s.DB.GetOriginalURL(short_code)
		if err != nil {
			fmt.Printf("Error when checking for collision: %s", err)
			return "", err
		}

		// No collision found; save to db and cache
		if existingURL == "" {
			err := s.DB.SaveShortCode(short_code, original_url)
			if err != nil {
				fmt.Printf("Error when no collision was found: %s", err)
				return "", err
			}
			_ = s.Cache.Set(short_code, original_url)
			return fmt.Sprintf("%s/%s", s.Domain, short_code), nil
		}

		// Found collision but checking if it maps to same original URL
		if existingURL == original_url {
			return fmt.Sprintf("%s/%s", s.Domain, short_code), nil
		}

		fmt.Printf("Collision detected for shortCode %s. Retrying... (attempt %d)\n", short_code, i+1)
	}
	return "", fmt.Errorf("collision resoltuion failed. Please try with a different URL")
}

func (s *Service) GetOriginalURL(short_url string) (string, error) {
	url, err := s.Cache.Get(short_url)
	if err == nil {
		return url, nil
	}
	return s.DB.GetOriginalURL(short_url)
}
