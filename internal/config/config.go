package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// Здесь прописываем проекты, которые должны отслеживаться
// 25 - СБС, 34 – ВСК, 134 - Zetta, 143 – OneFront
var ProjectsId = []int{25}

type Config struct {
	RedmineBaseURL string
	RedmineAPIKey  string
	TelegramToken  string
	ChatID         string
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("no .env file found, using system environment variables %s", err)
	}

	cfg := &Config{
		RedmineBaseURL: getEnv("REDMINE_BASE_URL", ""),
		RedmineAPIKey:  getEnv("REDMINE_API_KEY", ""),
		TelegramToken:  getEnv("TELEGRAM_TOKEN", ""),
		ChatID:         getEnv("CHAT_ID", ""),
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

func (c *Config) CheckAfterInit() error {
	if c.RedmineBaseURL == "" || c.RedmineAPIKey == "" || c.TelegramToken == "" || c.ChatID == "" {
		return fmt.Errorf(".env don't have requried value, check .env file")
	}

	return nil
}
