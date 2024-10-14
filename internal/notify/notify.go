package notify

import (
	"strconv"

	"github.com/zonder12120/go-redmine-tg-notify/internal/config"
	"github.com/zonder12120/go-redmine-tg-notify/internal/telegram"
	"github.com/zonder12120/go-redmine-tg-notify/pkg/utils"
)

var cfg, _ = config.LoadConfig()

var tgClient = telegram.NewClient(cfg.TelegramToken, cfg.ChatID)

func NotifyNewTask(issueId int, priorityId int, title string, assignToName string) error {
	msg, err := utils.ConcatStrings(
		markPriority(priorityId),
		" Добавлена новая задача ",
		"(", cfg.RedmineBaseURL, "/issues/", strconv.Itoa(issueId), ")",
		" \\- ", title,
		" для ",
		"*",
		assignToName,
		"*",
	)
	if err != nil {
		return utils.HadleError("Error concat strings for notify new task: ", err)
	}
	return Notify(msg)

}

func Notify(msg string) error {
	return tgClient.SendMsg(msg)
}
