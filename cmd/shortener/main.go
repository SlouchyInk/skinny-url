package main

import (
	"fmt"
	"log"
	net "net/http"
	"os"
	"strings"

	"github.com/hdurham99/skinny-url/internal/http"
	"github.com/hdurham99/skinny-url/internal/shortener"
	"github.com/hdurham99/skinny-url/internal/storage/cache"
	"github.com/hdurham99/skinny-url/internal/storage/db"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env
	godotenv.Load("../../.env")

	// Parse the env varaibles
	redisHost := os.Getenv("CACHE")
	cassHosts := strings.Split(os.Getenv("DB"), ",")
	keyspace := os.Getenv("KEYSPACE")
	domain := os.Getenv("DOMAIN")

	// Initialize storage
	cassDB, err := db.NewCassandraDB(cassHosts, keyspace)
	if err != nil {
		fmt.Println("Error creating Cassandra session:", err)
		return
	}
	redisCache := cache.NewRedisCache(redisHost)

	// Create shortener service
	shortService := shortener.NewShortenerService(cassDB, redisCache, domain)

	// Create new handler for shortener service
	handler := http.NewHandler(shortService)

	// Create HTTP router
	router := http.NewRouter(handler)

	// Start the server
	log.Fatal(net.ListenAndServe(":8080", router))
}
