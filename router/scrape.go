package router

import (
	"github.com/gin-gonic/gin"
	ctl "github.com/zzzgydi/webscraper/router/controller"
	"github.com/zzzgydi/webscraper/router/middleware"
)

func InnerRouter(r *gin.Engine) {
	v1 := r.Group("/v1")
	v1.Use(middleware.LoggerMiddleware)

	v1.POST("/scrape", ctl.PostScrape)
}
