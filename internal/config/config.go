package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Здесь прописываем проекты, которые должны отслеживаться
// 25 - СБС, 34 – ВСК, 134 - Zetta, 143 – OneFront
var ProjectsId = []int{25}

type Config struct {
	RedmineBaseURL  string
	RedmineAPIKey   string
	TelegramBaseURL string
	TelegramToken   string
	ChatID          string
}

func LoadConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, logError("No .env file found, using system environment variables")
	}

	cfg := &Config{
		RedmineBaseURL:  getEnv("REDMINE_BASE_URL", ""),
		RedmineAPIKey:   getEnv("REDMINE_API_KEY", ""),
		TelegramBaseURL: getEnv("TELEGRAM_BASE_URL", ""),
		TelegramToken:   getEnv("TELEGRAM_TOKEN", ""),
		ChatID:          getEnv("CHAT_ID", ""),
	}

	return cfg, nil
}

func getEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if exists {
		return value
	}
	return defaultValue
}

func logError(message string) error {
	err := fmt.Errorf("%s", message)
	log.Println(err)
	return err
}
