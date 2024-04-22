package controller

type PostScrapeRequest struct {
	UrlList     []string `json:"url_list"`
	Headless    *bool    `json:"headless,omitempty"`
	Readability *bool    `json:"readability,omitempty"`
}
