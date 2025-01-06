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
		short_url := fmt.Sprintf("%s/%s", s.Domain, encodeUrl(offsetInput))

		// Checking for collision
		existingURL, err := s.DB.GetOriginalURL(short_url)
		if err != nil {
			return "", err
		}

		// No collision found; save to db and cache
		if existingURL == "" {
			err := s.DB.SaveShortURL(short_url, original_url)
			if err != nil {
				return "", err
			}
			_ = s.Cache.Set(short_url, original_url)
			return short_url, nil
		}

		// Found collision but checking if it maps to same original URL
		if existingURL == original_url {
			return short_url, nil
		}
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
