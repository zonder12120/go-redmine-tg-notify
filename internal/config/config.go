package config

import (
	"fmt"
	"os"

	"github.com/zonder12120/go-redmine-tg-notify/pkg/env"
)

type Config struct {
	ProjectsID      []int  // Слайс ID проектов, по которым ты хочешь получать оповещения
	RedmineBaseURL  string // URL вашего Redmine
	RedmineAPIKey   string // Ваш api-key для Redmine
	GoogleDevApiKey string // Ваш api-key из Google Developer Account
	TelegramToken   string // Токен для бота (BotFather)
	ChatID          string // ID чата, где будет спамить бот (Get My ID)
	TimeZone        string // Часовой пояс, чтобы не испытывать проблем с определением времени на сервере
}

func LoadConfig() (Config, error) {
	err := env.LoadEnv()
	if err != nil {
		return Config{}, err
	}

	cfg := Config{
		ProjectsID:      env.GetSliceIntFromEnv("PROJECTS_LIST"),
		RedmineBaseURL:  os.Getenv("REDMINE_BASE_URL"),
		RedmineAPIKey:   os.Getenv("REDMINE_API_KEY"),
		GoogleDevApiKey: os.Getenv("GOOGLE_DEV_API_KEY"),
		TelegramToken:   os.Getenv("TELEGRAM_TOKEN"),
		ChatID:          os.Getenv("CHAT_ID"),
		TimeZone:        os.Getenv("TIME_ZONE"),
	}

	return cfg, nil
}

func (c *Config) CheckAfterInit() error {
	if len(c.ProjectsID) == 0 || c.RedmineBaseURL == "" ||
		c.RedmineAPIKey == "" || c.TelegramToken == "" ||
		c.ChatID == "" || c.GoogleDevApiKey == "" ||
		c.TimeZone == "" {

		return fmt.Errorf(".env don't have requried value, check .env file")
	}

	return nil
}
