package main

import (
	"github.com/hdurham99/skinny-url/internal/http"
	"github.com/hdurham99/skinny-url/internal/storage"
	"github.com/hdurham99/skinny-url/internal/shortener"
)

func main() {
	// Connect to the redis server
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "",
		DB: 0,
		Protocol: 2,
	})
	currentTime := time.Now()
	// Take in a url from the user
	var long_url string
	fmt.Scanln(&url)
	short_url := encodedUrl(url)
	fmt.Println(short_url)
	// Store and retrive the url and encoded url
	hashFields := []string{
		"long_url", long_url,
		"short_url", short_url,
		"creation_date", currentTime,
		"expiration_date", currentTime.Add(time.Day * 30)
	}
	ctx := context.Background()
	err := client.Set(ctx, )
}
