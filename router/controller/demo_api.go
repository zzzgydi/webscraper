package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zzzgydi/webscraper/router/utils"
	"github.com/zzzgydi/webscraper/service/scrape"
)

func GetScrapeHTML(c *gin.Context) {
	urlParam := c.Query("u")
	noHeadless := c.Query("headless") == "false"
	noReadability := c.Query("readability") == "false"

	if urlParam == "" {
		c.Redirect(http.StatusFound, "https://github.com/zzzgydi/webscraper")
		return
	}

	trace := utils.GetTraceLogger(c)
	if trace != nil {
		trace.Trace("url", urlParam)
	}

	s := scrape.NewScrape(!noHeadless, !noReadability, true)
	res, err := s.Run(c.Request.Context(), urlParam)
	if err != nil {
		c.HTML(http.StatusOK, "error.html", gin.H{
			"error": err.Error(),
		})
		return
	}

	c.HTML(http.StatusOK, "index.html", gin.H{
		"title":   res.Title,
		"content": res.Content,
	})
}
