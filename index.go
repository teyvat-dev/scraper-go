package index

import (
	"fmt"
	"net/http"

	wikicharacterscraper "teyvat.dev/scraper-go/scrapers/wiki"
)

// WikiScrapeCharacters scrape characters table, Under development
func WikiScrapeCharacters(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, string(wikicharacterscraper.Scrape(w, r)))
}
