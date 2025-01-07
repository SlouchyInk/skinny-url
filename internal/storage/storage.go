package storage

type DBStorage interface {
	SaveShortCode(short_code string, original_url string) error
	GetOriginalURL(short_code string) (string, error)
}

type CacheStorage interface {
	Set(key, value string) error
	Get(key string) (string, error)
}
