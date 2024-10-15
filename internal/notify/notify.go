package notify

import (
	"fmt"
	"log"
	"strconv"

	"github.com/zonder12120/go-redmine-tg-notify/internal/config"
	"github.com/zonder12120/go-redmine-tg-notify/internal/telegram"
	"github.com/zonder12120/go-redmine-tg-notify/pkg/utils"
)

var cfg, _ = config.LoadConfig()

var tgClient = telegram.NewClient(cfg.TelegramToken, cfg.ChatID)

func NotifyNewTask(issueID int, priorityID int, title string, assignToName string) error {
	msg, err := utils.ConcatStrings(
		markPriority(priorityID),
		" Добавлена новая задача ",
		"(", cfg.RedmineBaseURL, "/issues/", strconv.Itoa(issueID), ")",
		" \\- ", title,
		" для ",
		"*",
		assignToName,
		"*",
	)
	if err != nil {
		return fmt.Errorf("error concat strings for notify new task: %s", err)
	}
	return Notify(msg)

}

func Notify(msg string) error {
	log.Println("Sending message (MarkdownV2): ", msg)
	return tgClient.SendMsg(msg)
}
