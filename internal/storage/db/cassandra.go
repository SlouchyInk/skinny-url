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
