package main

import (
	"github.com/zzzgydi/webscraper/common/config"
	"github.com/zzzgydi/webscraper/common/initializer"
	"github.com/zzzgydi/webscraper/common/logger"
	"github.com/zzzgydi/webscraper/router"
)

func main() {
	env := config.GetEnv()
	rootDir := config.GetRootDir()
	logger.InitLogger(rootDir)
	config.InitConfig(rootDir, env)
	initializer.InitInitializer()
	router.InitHttpServer()
}
