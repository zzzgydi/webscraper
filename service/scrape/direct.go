package scrape

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	md "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/PuerkitoBio/goquery"
	"github.com/zzzgydi/webscraper/common/utils"
	"golang.org/x/net/html/charset"
)

func (s *Scrape) directScrape(ctx context.Context, rawUrl string) (*ScrapeResult, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", rawUrl, nil)
	if err != nil {
		return nil, err
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
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("request status: %d", resp.StatusCode)
	}

	contentType := resp.Header.Get("Content-Type")
	if !strings.Contains(contentType, "text/html") {
		return nil, fmt.Errorf("invalid Content-Type: %s", contentType)
	}

	reader, err := charset.NewReader(resp.Body, contentType)
	if err != nil {
		return nil, err
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

	ret := &ScrapeResult{
		Url: rawUrl,
	}

	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		return nil, err
	}

	ret.Title = doc.Find("title").First().Text()

	converter := md.NewConverter("", true, options)
	ret.Content = converter.Convert(doc.Selection)

	return ret, nil
}
