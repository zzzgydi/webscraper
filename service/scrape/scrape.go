package scrape

import (
	"context"
	"fmt"
	"sync"
)

type Scrape struct {
	rewiseDomain bool
	headless     bool
	readability  bool
	pipeline     []func(string) string
}

func NewScrape(headless, readability, rewiseDomain bool) *Scrape {
	return &Scrape{
		headless:     headless,
		readability:  readability,
		rewiseDomain: rewiseDomain,
	}
}

func (s *Scrape) AddPipeline(fn func(string) string) {
	s.pipeline = append(s.pipeline, fn)
}

func (s *Scrape) Run(ctx context.Context, rawUrl string) (*ScrapeResult, error) {
	if !filter.Pass(rawUrl) {
		return nil, fmt.Errorf("url not allowed: %s", rawUrl)
	}

	var ret *ScrapeResult
	var err error

	if s.headless {
		ret, err = s.headlessScrape(ctx, rawUrl)
	} else {
		ret, err = s.directScrape(ctx, rawUrl)
	}
	if err != nil {
		return nil, err
	}

	ret.Content = filter.ContentFilter(ret.Content)

	for _, fn := range s.pipeline {
		ret.Content = fn(ret.Content)
	}

	return ret, nil
}

func (s *Scrape) BatchRun(ctx context.Context, urlList []string) ([]ScrapeResult, error) {
	if len(urlList) == 0 {
		return nil, fmt.Errorf("url list is empty")
	}

	res := make([]ScrapeResult, len(urlList))

	var wg sync.WaitGroup

	for idx, url := range urlList {
		wg.Add(1)
		go func(idx int, url string) {
			defer wg.Done()
			res[idx] = ScrapeResult{Url: url}

			// add recover
			defer func() {
				if r := recover(); r != nil {
					res[idx].Error = fmt.Sprintf("panic: %v", r)
				}
			}()

			ret, err := s.Run(ctx, url)
			if err != nil {
				res[idx].Error = err.Error()
			} else {
				res[idx].Title = ret.Title
				res[idx].Content = ret.Content
			}
		}(idx, url)
	}

	wg.Wait()

	return res, nil
}
