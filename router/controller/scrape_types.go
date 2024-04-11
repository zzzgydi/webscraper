package controller

type PostScrapeRequest struct {
	UrlList []string `json:"url_list"`
}

type ScrapeResult struct {
	Url     string `json:"url"`
	Content string `json:"content"`
	Error   string `json:"error,omitempty"`
}
