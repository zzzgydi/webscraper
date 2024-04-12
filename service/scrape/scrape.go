package scrape

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"sync"

	md "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/PuerkitoBio/goquery"
	"github.com/zzzgydi/webscraper/common/utils"
	"golang.org/x/net/html/charset"
)

type Scrape struct {
	rewiseDomain bool
	pipeline     []func(string) string
}

type ScrapeResult struct {
	Url     string `json:"url"`
	Content string `json:"content"`
	Error   string `json:"error,omitempty"`
}

func NewScrape(rewiseDomain bool) *Scrape {
	return &Scrape{
		rewiseDomain: rewiseDomain,
	}
}

func (s *Scrape) AddPipeline(fn func(string) string) {
	s.pipeline = append(s.pipeline, fn)
}

func (s *Scrape) request(ctx context.Context, rawUrl string) (string, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", rawUrl, nil)
	if err != nil {
		return "", err
	}

	headers := map[string]string{
		"Accept": "text/html;q=0.9, application/xhtml+xml;q=0.8",
		// "Accept-Encoding":           "gzip, deflate",
		"Cache-Control": "no-cache",
		"Connection":    "keep-alive",
		"Pragma":        "no-cache",
		"User-Agent":    utils.RandomUserAgent(),
	}
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	resp, err := scrapeClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("request status: %d", resp.StatusCode)
	}

	contentType := resp.Header.Get("Content-Type")
	if !strings.Contains(contentType, "text/html") {
		return "", fmt.Errorf("invalid Content-Type: %s", contentType)
	}

	reader, err := charset.NewReader(resp.Body, contentType)
	if err != nil {
		return "", err
	}

	domain := ""
	u, err := url.Parse(rawUrl)
	if err == nil {
		domain = u.Scheme + "://" + u.Host
	}

	options := &md.Options{
		GetAbsoluteURL: func(selec *goquery.Selection, rawURL, _ string) string {
			if !s.rewiseDomain {
				return rawURL
			}
			// 如果是相对路径，拼接成绝对路径
			if strings.HasPrefix(rawURL, "/") {
				return domain + rawURL
			}
			return rawURL
		},
	}

	converter := md.NewConverter("", true, options)
	markdown, err := converter.ConvertReader(reader)
	if err != nil {
		return "", err
	}

	return markdown.String(), nil
}

func (s *Scrape) Run(ctx context.Context, rawUrl string) (string, error) {
	if !filter.Pass(rawUrl) {
		return "", fmt.Errorf("url not allowed: %s", rawUrl)
	}

	ret, err := s.request(ctx, rawUrl)
	if err != nil {
		return "", err
	}

	for _, fn := range s.pipeline {
		ret = fn(ret)
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

			content, err := s.Run(ctx, url)
			if err != nil {
				res[idx].Error = err.Error()
			} else {
				res[idx].Content = content
			}
		}(idx, url)
	}

	wg.Wait()

	return res, nil
}
