package config

import (
	"fmt"
	"log/slog"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

var AppConf *AppConfig

func InitConfig(rootDir string, env string) {
	AppConf = &AppConfig{}

	fullConfigFilePath := filepath.Join(rootDir, "config", fmt.Sprintf("%s.yaml", env))
	slog.Info("Init config file", "file", fullConfigFilePath)

	data, err := os.ReadFile(fullConfigFilePath)
	if err != nil {
		slog.Error("Init config file error", "file", fullConfigFilePath, "error", err)
		panic(err)
	}
	configContent := []byte(os.ExpandEnv(string(data)))
	err = yaml.Unmarshal(configContent, AppConf)
	if err != nil {
		slog.Error("Parse config file error", "file", fullConfigFilePath, "error", err)
		panic(err)
	}
}

func GetEnv() string {
	devEnv := "dev"
	env := os.Getenv("CONFIG_ENV")
	if env == "" {
		slog.Warn(fmt.Sprintf("CONFIG_ENV is NOT set, default to [%s].\n", devEnv))
		env = devEnv
	}
	return env
}

func GetRootDir() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(err)
	}
	return dir
}
