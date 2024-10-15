package config

import (
	"fmt"
	"os"

	"github.com/zonder12120/go-redmine-tg-notify/pkg/env"
)

type Config struct {
	RedmineBaseURL string // URL вашего Redmine
	RedmineAPIKey  string // Ваш api-key для Redmine
	TelegramToken  string // Токен для бота (BotFather)
	ChatID         string // id чата, где будет спамить бот (Get My ID)
	ProjectsID     []int  // слайс id проектов, по которым ты хочешь получать оповещения
}

func LoadConfig() (Config, error) {
	err := env.LoadEnv()
	if err != nil {
		return Config{}, err
	}

	cfg := Config{
		RedmineBaseURL: os.Getenv("REDMINE_BASE_URL"),
		RedmineAPIKey:  os.Getenv("REDMINE_API_KEY"),
		TelegramToken:  os.Getenv("TELEGRAM_TOKEN"),
		ChatID:         os.Getenv("CHAT_ID"),
		ProjectsID:     env.GetSliceIntFromEnv("PROJECTS_LIST"),
	}

	return cfg, nil
}

func (c *Config) CheckAfterInit() error {
	if c.RedmineBaseURL == "" || c.RedmineAPIKey == "" || c.TelegramToken == "" || c.ChatID == "" || c.ProjectsID == nil {
		return fmt.Errorf(".env don't have requried value, check .env file")
	}

	return nil
}
