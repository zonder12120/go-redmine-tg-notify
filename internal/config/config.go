package config

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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
	err := loadEnv()
	if err != nil {
		return nil, err
	}

	cfg := &Config{
		RedmineBaseURL: os.Getenv("REDMINE_BASE_URL"),
		RedmineAPIKey:  os.Getenv("REDMINE_API_KEY"),
		TelegramToken:  os.Getenv("TELEGRAM_TOKEN"),
		ChatID:         os.Getenv("CHAT_ID"),
	}

	return cfg, nil
}

func (c *Config) CheckAfterInit() error {
	if c.RedmineBaseURL == "" || c.RedmineAPIKey == "" || c.TelegramToken == "" || c.ChatID == "" {
		return fmt.Errorf(".env don't have requried value, check .env file")
	}

	return nil
}

func loadEnv() error {
	file, err := os.Open(".env")
	if os.IsNotExist(err) {
		return fmt.Errorf("no .env file found, using system environment variables: %s", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "#") || len(strings.TrimSpace(line)) == 0 {
			continue
		}

		keyValue := strings.SplitN(line, "=", 2)
		if len(keyValue) != 2 {
			return fmt.Errorf("invalid line in .env file: %s", line)
		}
		key := strings.TrimSpace(keyValue[0])
		value := strings.TrimSpace(keyValue[1])

		os.Setenv(key, value)
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error reading .env file: %s", err)
	}

	return nil
}
