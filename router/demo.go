package router

import (
	"github.com/gin-gonic/gin"
	"github.com/zzzgydi/webscraper/router/controller"
	"github.com/zzzgydi/webscraper/router/middleware"
)

func DemoRouter(r *gin.Engine) {
	r.LoadHTMLGlob("static/*")
	r.GET("/", middleware.LoggerMiddleware, controller.GetScrapeHTML)
}
