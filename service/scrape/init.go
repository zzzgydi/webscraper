package scrape

import (
	"net/http"
	"net/url"
	"time"

	"github.com/zzzgydi/webscraper/common/config"
	"github.com/zzzgydi/webscraper/common/initializer"
)

var (
	scrapeClient *http.Client
	filter       *Filter
)

func initScrape() error {
	scrapeClient = &http.Client{
		Timeout: 10 * time.Second,
		Transport: &http.Transport{
			MaxIdleConns:          30,
			IdleConnTimeout:       3600 * time.Second,
			ResponseHeaderTimeout: 360 * time.Second,
			ExpectContinueTimeout: 360 * time.Second,
			DisableCompression:    false,
		},
	}

	// set http proxy
	if config.AppConf.HttpProxy != "" {
		proxyUrl, err := url.Parse(config.AppConf.HttpProxy)
		if err != nil {
			return err
		}
		scrapeClient.Transport = &http.Transport{
			Proxy: http.ProxyURL(proxyUrl),
		}
	}

	filter = NewFilter(config.AppConf.Filter)

	return nil
}

func init() {
	initializer.Register("scrape", initScrape)
}
