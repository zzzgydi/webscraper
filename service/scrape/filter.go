package scrape

import (
	"net/url"
	"strings"

	"github.com/zzzgydi/webscraper/common/config"
)

type Filter struct {
	hostMap     map[string]bool
	blockString []string
}

func NewFilter(conf config.FilterRule) *Filter {
	hostMap := make(map[string]bool)
	for _, h := range conf.Host {
		hostMap[h] = true
	}

	return &Filter{
		hostMap:     hostMap,
		blockString: conf.BlockString,
	}
}

// Pass checks if the url can be scraped
func (f *Filter) Pass(urlStr string) bool {
	u, err := url.Parse(urlStr)
	if err != nil {
		return false
	}

	if _, ok := f.hostMap[u.Host]; ok {
		return false
	}

	return true
}

func (f *Filter) ContentFilter(content string) string {
	// remove block string
	for _, b := range f.blockString {
		if b != "" {
			content = strings.Replace(content, b, "", -1)
		}
	}

	return content
}
