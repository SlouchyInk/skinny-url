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

func (s *Service) ShortenURL(original_url string, user_id string) (string, error) {
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
			if user_id == "" {
				err := s.DB.SaveShortCode(short_code, original_url, "")
				if err != nil {
					fmt.Printf("Error when no collision was found: %s", err)
					return "", err
				}
			}
			err := s.DB.SaveShortCode(short_code, original_url, user_id)
			if err != nil {
				fmt.Printf("Error when no collision was found: %s", err)
				return "", err
			}
			_ = s.Cache.Set(short_code, original_url, user_id)
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

func (s *Service) IncrementClickCount(short_code string) error {
	return s.Cache.IncrementClickCount(short_code)
}

func (s *Service) GetOriginalURL(short_code string) (string, error) {
	cacheCh := make(chan string)
	dbCh := make(chan string)
	errCh := make(chan error)

	go func() {
		url, err := s.Cache.Get(short_code)
		if err != nil {
			errCh <- err
			return
		}
		cacheCh <- url
	}()

	go func() {
		url, err := s.DB.GetOriginalURL(short_code)
		if err != nil {
			errCh <- err
			return
		}
		dbCh <- url
	}()

	select {
	case url := <-cacheCh:
		return url, nil
	case url := <-dbCh:
		user_id, err := s.DB.GetUser(short_code)
		if err != nil {
			fmt.Println("Error getting user from short_code...")
		}
		_ = s.Cache.Set(short_code, url, user_id)
		return url, nil
	case err := <-errCh:
		fmt.Println("Error found when getting original url: ", err)
		return "", err
	}

}
