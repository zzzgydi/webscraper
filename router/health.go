package router

import (
	"github.com/gin-gonic/gin"
)

func HealthRouter(r *gin.Engine) {
	r.GET("/health", func(c *gin.Context) {
		c.String(200, "ok")
	})
}
