package db

import (
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

func (db *CassandraDB) SaveShortURL(short_url string, original_url string) error {
	return db.session.Query(
		"INSERT INTO url_mapping (short_url, long_url) VALUES (?, ?)",
		short_url, original_url).Exec()
}

func (db *CassandraDB) GetOriginalURL(short_url string) (string, error) {
	var url string
	err := db.session.Query("SELECT long_url FROM short_url WHERE short_url = ?", short_url).Scan(&url)
	return url, err
}
