package db

import (
	"time"

	"github.com/gocql/gocql"
)

type CassandraDB struct {
	session *gocql.Session
}

// Initializes the cluster configuration and creates a sesison.
func NewCassandraDB(hosts []string, keyspace string) (*CassandraDB, error) {
	// Initialize the Cassandra cluster
	cluster := gocql.NewCluster(hosts...)
	cluster.Keyspace = keyspace

	// Create a session
	session, err := cluster.CreateSession()
	if err != nil {
		return nil, err
	}
	return &CassandraDB{session: session}, nil
}

func (db *CassandraDB) SaveShortCode(short_code, original_url, user_id string) error {
	expiration_date := time.Now().AddDate(1, 0, 0)
	if user_id == "" {
		err := db.session.Query(
			"INSERT INTO url_mapping (short_code, long_url, expiration_date, creation_date, user_id) VALUES (?, ?, ?, ?, null)",
			short_code, original_url, expiration_date, time.Now()).Exec()
		return err
	}
	err := db.session.Query(
		"INSERT INTO url_mapping (short_code, long_url, expiration_date, creation_date, user_id) VALUES (?, ?, ?, ?, ?)",
		short_code, original_url, expiration_date, time.Now(), user_id).Exec()
	return err
}

func (db *CassandraDB) GetOriginalURL(short_code string) (string, error) {
	var url string
	err := db.session.Query(
		"SELECT long_url FROM url_mapping WHERE short_code = ?",
		short_code).Scan(&url)
	if err != nil {
		if err == gocql.ErrNotFound {
			return "", nil // No collision
		}
	}
	return url, err
}

func (db *CassandraDB) SaveClickCount(key string, count int) error {
	err := db.session.Query(
		"INSERT INTO click_counter (short_code, click_count) VALUES (?, ?)",
		key, count).Exec()
	return err
}

func (db *CassandraDB) GetUser(short_code string) (string, error) {
	var user string
	err := db.session.Query(
		"SELECT user_id FROM url_mapping where short_code = ?",
		short_code).Scan(&user)
	return user, err

}
