package main

import (
	"context"
	"log"
	"os"

	"github.com/GoogleCloudPlatform/functions-framework-go/funcframework"
	index "teyvat.dev/scraper-go"
)

func main() {
	ctx := context.Background()

	// Character Scrape
	if err := funcframework.RegisterHTTPFunctionContext(ctx, "/WikiScrapeCharacters", index.WikiScrapeCharacters); err != nil {
		log.Fatalf("funcframework.RegisterHTTPFunctionContext: %v\n", err)
	}

	// Use PORT environment variable, or default to 8080.
	port := "8080"
	if envPort := os.Getenv("PORT"); envPort != "" {
		port = envPort
	}
	if err := funcframework.Start(port); err != nil {
		log.Fatalf("funcframework.Start: %v\n", err)
	}
}
