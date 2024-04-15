package router

import (
	"log/slog"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/zzzgydi/webscraper/common/config"
	L "github.com/zzzgydi/webscraper/common/logger"
)

func InitHttpServer() {
	r := gin.New()
	r.Use(gin.Recovery())

	// cors
	r.Use(cors.New(cors.Config{
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return true
		},
		MaxAge: 12 * time.Hour,
	}))

	// register routers
	DemoRouter(r)
	HealthRouter(r)
	InnerRouter(r)

	logger := slog.NewLogLogger(L.Handler, slog.LevelError)
	srv := &http.Server{
		Addr:     ":" + strconv.FormatInt(int64(config.AppConf.HttpPort), 10),
		Handler:  r,
		ErrorLog: logger,
	}
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		panic(err)
	}
}
