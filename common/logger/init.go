package logger

import (
	"fmt"
	"io"
	"log/slog"
	"os"
	"path/filepath"
	"time"
)

var (
	Logger  *slog.Logger
	Handler slog.Handler
)

func InitLogger(rootDir string) {
	logFile := fmt.Sprintf("%s.log", time.Now().Format("20060102"))
	logPath := filepath.Join(rootDir, "log", logFile)

	file, err := os.OpenFile(logPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}

	mw := io.MultiWriter(os.Stdout, file)

	Handler = slog.NewJSONHandler(mw, nil)
	Logger = slog.New(Handler)

	slog.SetDefault(Logger)
}
