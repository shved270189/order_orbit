package config

import (
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/joho/godotenv"
)

type Config struct {
	Port        string
	DatabaseURL string
	AppEnv      string
	ProjectRoot string
}

func Load(postfix string) *Config {
	_, current_file_path, _, _ := runtime.Caller(0)
	projectRoot := filepath.Dir(filepath.Dir(filepath.Dir(current_file_path))) + "/"

	file_name := ".env"
	if postfix != "" {
		file_name = file_name + "." + postfix
	}

	if err := godotenv.Load(projectRoot + file_name); err != nil {
		log.Printf("No %s file found", projectRoot + file_name)
	}

	return &Config{
		Port:        getEnv("PORT", "8080"),
		DatabaseURL: getEnv("DATABASE_URL", "local.db"),
		AppEnv:      getEnv("APP_ENV", "local"),
		ProjectRoot: projectRoot,
	}
}

func getEnv(key, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}
