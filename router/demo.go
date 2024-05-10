package router

import (
	"github.com/gin-gonic/gin"
	"github.com/zzzgydi/webscraper/router/controller"
	"github.com/zzzgydi/webscraper/router/middleware"
)

func init() {
	RegisterRoute(func(r *gin.Engine) {
		demoRouter(r)
	})
}

func demoRouter(r *gin.Engine) {
	r.LoadHTMLGlob("static/*")
	r.GET("/", middleware.LoggerMiddleware, controller.GetScrapeHTML)
}
