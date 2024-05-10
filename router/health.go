package router

import (
	"github.com/gin-gonic/gin"
)

func init() {
	RegisterRoute(func(r *gin.Engine) {
		healthRouter(r)
	})
}

func healthRouter(r *gin.Engine) {
	r.GET("/health", func(c *gin.Context) {
		c.String(200, "ok")
	})
}
