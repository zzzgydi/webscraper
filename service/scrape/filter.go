package scrape

import "net/url"

type Filter struct {
	hostMap map[string]bool
}

func NewFilter(host []string) *Filter {
	hostMap := make(map[string]bool)
	for _, h := range host {
		hostMap[h] = true
	}

	return &Filter{
		hostMap: hostMap,
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
