package scraper

import "context"

type Scraper interface {
	Scrape(ctx context.Context, url string) (result string, err error)
}