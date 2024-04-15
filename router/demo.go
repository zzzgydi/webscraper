package router

import (
	"github.com/gin-gonic/gin"
	"github.com/zzzgydi/webscraper/router/controller"
)

func DemoRouter(r *gin.Engine) {
	r.LoadHTMLGlob("static/*")
	r.GET("/", controller.GetScrapeHTML)
}
