package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/hdurham99/skinny-url/internal/storage/db"
	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load()
	rawHosts := os.Getenv("HOSTS")
	cassHosts := strings.Split(rawHosts, ",")
	keyspace := os.Getenv("KEYSPACE")

	// CassDB is a Cassandra DB session
	cassDB, err := db.NewCassandraDB(cassHosts, keyspace)
	fmt.Println(cassDB, err)

}
