package storage

import (
	"github.com/hdurham99/skinny-url/internal/storage/db"
)

type DBStorage interface {
	SaveShortCode(short_code, original_url, user_id string) error
	GetOriginalURL(short_code string) (string, error)
	SaveClickCount(key string, count int) error
	GetUser(short_code string) (string, error)
}

type CacheStorage interface {
	Set(short_code, long_url, user_id string) error
	Get(key string) (string, error)
	FlushClickCountsToDB(c *db.CassandraDB)
	IncrementClickCount(short_code string) error
}
