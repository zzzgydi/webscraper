package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/zzzgydi/webscraper/service/scrape"
)

func PostScrape(c *gin.Context) {
	req := &PostScrapeRequest{}
	if err := c.ShouldBindJSON(req); err != nil {
		ReturnBadRequest(c, err)
		return
	}
	if len(req.UrlList) == 0 {
		ReturnBadRequest(c, fmt.Errorf("url list is empty"))
		return
	}

	res := make([]ScrapeResult, len(req.UrlList))
	s := scrape.NewScrape(true)

	for idx, url := range req.UrlList {
		res[idx] = ScrapeResult{Url: url}
		content, err := s.Run(c.Request.Context(), url)
		if err != nil {
			res[idx].Error = err.Error()
		} else {
			res[idx].Content = content
		}
	}

	ReturnSuccess(c, res)
}
