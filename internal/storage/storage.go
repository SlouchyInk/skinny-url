package storage

type DBStorage interface {
	SaveShortURL(short_url string, original_url string) error
	GetOriginalURL(short_url string) (string, error)
}
