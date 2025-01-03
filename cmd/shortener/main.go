package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/hdurham99/skinny-url/internal/storage/db"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env
	err := godotenv.Load("../../.env")

	// Parse the comma-separated HOSTS variable
	cassHosts := strings.Split(os.Getenv("HOSTS"), ",")
	keyspace := os.Getenv("KEYSPACE")

	// Initialize Cassandra DB session
	cassDB, err := db.NewCassandraDB(cassHosts, keyspace)
	if err != nil {
		fmt.Println("Error creating Cassandra session:", err)
		return
	}
	fmt.Println("Cassandra DB session initialized successfully:", cassDB)
}
