package headless

import (
	"context"
	"fmt"
	"net/url"
	"strings"
	"time"

	md "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/cdproto/runtime"
	"github.com/chromedp/chromedp"
)

func Headless(rawUrl string, rewiseDomain bool) (string, string, error) {
	ctx, cancel := chromedp.NewContext(allocCtx)
	defer cancel()

	// init
	_ = chromedp.Run(ctx)

	timeoutCtx, cancel := context.WithTimeout(ctx, 10*time.Second) // 10s
	defer cancel()

	var result map[string]any

	err := chromedp.Run(timeoutCtx,
		chromedp.Navigate(rawUrl),
		chromedp.Evaluate(
			readabilityJS+`;new Readability(document.cloneNode(true)).parse()`,
			&result,
			func(p *runtime.EvaluateParams) *runtime.EvaluateParams {
				return p.WithAwaitPromise(true)
			},
		),
	)
	if err != nil {
		return "", "", err
	}

	domain := ""
	u, err := url.Parse(rawUrl)
	if err == nil {
		domain = u.Scheme + "://" + u.Host
	}

	options := &md.Options{
		GetAbsoluteURL: func(selec *goquery.Selection, rawURL, _ string) string {
			if !rewiseDomain {
				return rawURL
			}
			if strings.HasPrefix(rawURL, "/") {
				return domain + rawURL
			}
			return rawURL
		},
	}

	converter := md.NewConverter("", true, options)
	content, ok := result["content"].(string)
	if !ok {
		return "", "", fmt.Errorf("content not found")
	}

	content, err = converter.ConvertString(content)
	if err != nil {
		return "", "", err
	}

	title, ok := result["title"].(string)
	if !ok {
		title = "No Title"
	}

	return title, content, nil
}
