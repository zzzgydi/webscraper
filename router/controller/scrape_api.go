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

	s := scrape.NewScrape(true)
	res, err := s.BatchRun(c.Request.Context(), req.UrlList)
	if err != nil {
		ReturnServerError(c, err)
		return
	}

	ReturnSuccess(c, res)
}
