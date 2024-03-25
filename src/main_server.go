package main

import (
	"context"
	"html/template"
	"log"
	"net/http"

	"golang.org/x/text/message"

	"github.com/go-redis/redis/v8"
)

// PageContent struct for rendering.
type PageContent struct {
	Counter string
}

// Redis Connection
var (
	client        *redis.Client
	page_template *template.Template
)

func clientHandler(w http.ResponseWriter, r *http.Request) {

	// Get visitor count.
	counter, err := getAndincrementVisitorCount()
	if err != nil {
		log.Println("[x] Error fetching visitor count: ", err)
	}

	// Format the countter value (add comma separator).
	formatted_counter := message.NewPrinter(message.MatchLanguage("en")).Sprintf("%d", counter)

	page := PageContent{Counter: formatted_counter}

	if err := page_template.Execute(w, page); err != nil {
		log.Fatal(err)
	}
}

func getAndincrementVisitorCount() (int64, error) {
	// Increment total visitors while called.
	ctx := context.Background()
	count, err := client.Incr(ctx, "total_visitors").Result()
	if err != nil {
		return -1, err
	}
	return count, nil
}

func main() {

	client = redis.NewClient(&redis.Options{
		Addr:     "redis:6379", // Assuming Redis is running locally
		Password: "",           // No password
		DB:       0,            // Default DB
	})

	// Test the Redis connection.
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		log.Fatal("[x] Error connecting to Redis:", err)
	}

	// Load template.
	page_template = template.Must(template.ParseFiles("/static/html/page_template.html"))
	if page_template == nil {
		log.Fatal("[x] Error loading template.")
	}

	http.HandleFunc("/", clientHandler)
	log.Println("[+] Collecting God Boris Power...")
	log.Println("[.] Server starting up and listening on port 8080.")
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil))
	log.Println("[.] Server Shutting Down.")
}
