package notify

import (
	"fmt"
	"log"
	"strings"

	"github.com/zonder12120/go-redmine-tg-notify/internal/config"
	"github.com/zonder12120/go-redmine-tg-notify/internal/createmsg"
	"github.com/zonder12120/go-redmine-tg-notify/internal/redmine"
	"github.com/zonder12120/go-redmine-tg-notify/internal/telegram"
	"github.com/zonder12120/go-redmine-tg-notify/internal/timecheck"
	"github.com/zonder12120/go-redmine-tg-notify/pkg/utils"
)

var cfg, _ = config.LoadConfig()

var tgClient = telegram.NewClient(cfg.TelegramToken, cfg.ChatID)

var offHoursIssues = make(map[int]struct{})

func SendMessage(msg string) error {
	log.Println("Sending message (MarkdownV2): ", msg)
	return tgClient.SendMsg(msg)
}

func Updates(oldIssueMap, newIssueMap map[int]redmine.Issue, ignoredIssuesMap map[int]struct{}) {

	// Отправляем номера задач, по которым были изменения во внерабочее время
	if len(offHoursIssues) != 0 && timecheck.IsWorkTime(cfg.GoogleDevApiKey) {
		msg, err := createmsg.OffHoursChanges(cfg.RedmineBaseURL, offHoursIssues)
		utils.LogErr("Error create message off hours changes: ", err)

		err = SendMessage(msg)
		utils.LogErr("Error send message off hours changes: ", err)

		offHoursIssues = nil
	}

	// Сравниваем старые задачи с новыми
	for newIssueID, newIssue := range newIssueMap {

		_, exists := ignoredIssuesMap[newIssueID]

		if exists {
			continue
		}

		oldIssue, exists := oldIssueMap[newIssueID]

		// Если есть новая задача, сразу создаём оповещение (в рабочее время)
		if !exists {
			if !timecheck.IsWorkTime(cfg.GoogleDevApiKey) {
				offHoursIssues[newIssueID] = struct{}{}
			}

			msg, err := createmsg.NewTask(cfg.RedmineBaseURL, newIssueID, newIssue.Priority.ID, newIssue.Title, newIssue.AssignedTo.Name)
			utils.LogErr(fmt.Sprintf("Error create message new task, number %v", newIssueID), err)

			err = SendMessage(msg)
			utils.LogErr(fmt.Sprintf("Error notify new task, number %v", newIssueID), err)

			continue
		}

		if exists && oldIssueMap[newIssueID].UpdateTime != newIssueMap[newIssueID].UpdateTime {

			if !timecheck.IsWorkTime(cfg.GoogleDevApiKey) {
				offHoursIssues[newIssueID] = struct{}{}

				continue
			}

			msg, err := createDiffMessage(oldIssue, newIssue)
			utils.LogErr(fmt.Sprintf("Error create msg for task number %v", newIssueID), err)

			if msg != "" {
				err := SendMessage(msg)
				if err != nil {
					log.Println("Error send message to chat: ", err)
				}
			}
		}
	}
}

func createDiffMessage(oldIssue, newIssue redmine.Issue) (string, error) {
	var builder strings.Builder

	if oldIssue.Status.ID != newIssue.Status.ID {
		str, err := createmsg.AddStatusTxt(oldIssue.Status.Name, newIssue.Status.Name)
		utils.LogErr("Error add issueses status text", err)

		builder.WriteString(str)
	}

	if oldIssue.Priority.ID != newIssue.Priority.ID {
		str, err := createmsg.AddPriorityTxt(oldIssue.Priority.ID, newIssue.Priority.ID)
		utils.LogErr("Error add issueses priority text", err)

		builder.WriteString(str)
	}

	newComment := compareIssuesJournals(oldIssue, newIssue)

	if strings.TrimSpace(newComment) != "" {
		str, err := createmsg.AddNewCommentTxt(newComment)
		utils.LogErr("Error concat strings on compare issueses comments", err)

		builder.WriteString(str)
	}

	// Нужно, чтобы если исполнитель менялся, в итоговом сообщении не подписывался исполнитель
	assignedToName := newIssue.AssignedTo.Name

	if oldIssue.AssignedTo.ID != newIssue.AssignedTo.ID {
		str, err := createmsg.AddAssignedTxt(oldIssue.AssignedTo.Name, newIssue.AssignedTo.Name)
		utils.LogErr("Error concat strings on compare issueses priority", err)

		assignedToName = ""

		builder.WriteString(str)
	}

	text := builder.String()

	if text == "" {
		return "", nil
	}

	msg, err := createmsg.NewMsg(cfg.RedmineBaseURL, newIssue.ID, newIssue.Priority.ID, newIssue.Tracker.ID, newIssue.Title, text, assignedToName)
	if err != nil {
		return "", err
	}

	return msg, nil
}

// Выводит только последний добавленный коммент, но это оптимальнее, чем мапить комменты каждый интервал времени
func compareIssuesJournals(oldIssue, newIssue redmine.Issue) string {
	if len(newIssue.Journals) > len(oldIssue.Journals) {
		return newIssue.Journals[len(newIssue.Journals)-1].Notes
	}

	return ""
}
