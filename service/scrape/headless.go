package scrape

import (
	"context"

	hl "github.com/zzzgydi/webscraper/service/headless"
)

func (s *Scrape) headlessScrape(ctx context.Context, rawUrl string) (*ScrapeResult, error) {
	title, content, err := hl.Headless(rawUrl, s.readability, s.rewiseDomain)
	if err != nil {
		return nil, err
	}

	return &ScrapeResult{
		Url:     rawUrl,
		Title:   title,
		Content: content,
	}, nil
}
