package redmine

import (
	"fmt"
	"log"
	"strings"

	"github.com/zonder12120/go-redmine-tg-notify/internal/notify"
	"github.com/zonder12120/go-redmine-tg-notify/pkg/utils"
)

var builder strings.Builder

func MakeMapIssuesList(i IssuesList) map[int]Issue {
	m := make(map[int]Issue, len(i.Issues))
	for _, issue := range i.Issues {
		m[issue.Id] = issue
	}

	return m
}

// Добавление комментариев в мапу
func (c *Client) AddJournalsIssuesMap(issueMap map[int]Issue) error {
	for id, issue := range issueMap {
		IssueInfo, err := c.GetIssueInfo(id)
		if err != nil {
			return fmt.Errorf("error get old issue info %s", err)
		}

		issue.Journals = IssueInfo.Issue.Journals

		issueMap[id] = issue
	}

	return nil
}

func (c *Client) NotifyUpdates(oldIssueMap, newIssueMap map[int]Issue) {

	// Сравниваем старые задачи с новыми
	for newIssueId, newIssue := range newIssueMap {

		oldIssue, exists := oldIssueMap[newIssueId]

		// Если есть новая задача, сразу создаём оповещение
		if !exists {
			err := notify.NotifyNewTask(newIssueId, newIssue.Priority.Id, newIssue.Title, newIssue.AssignedTo.Name)
			utils.LogErr(fmt.Sprintf("Error notify new task number %v", newIssueId), err)

			continue
		}

		msg, err := createDiffMessage(oldIssue, newIssue)
		utils.LogErr(fmt.Sprintf("Error create msg for task number %v", newIssueId), err)

		if msg != "" {
			err := notify.Notify(msg)
			if err != nil {
				log.Printf("Error send message to chat: %s\n", err)
			}
		}
	}
}

func createDiffMessage(oldIssue, newIssue Issue) (string, error) {
	builder.Reset()

	if oldIssue.Status.Id != newIssue.Status.Id {
		str, err := notify.AddStatusTxt(oldIssue.Status.Name, newIssue.Status.Name)
		utils.HadleError("Error add issueses status text", err)

		builder.WriteString(str)
	}

	if oldIssue.Priority.Id != newIssue.Priority.Id {
		str, err := notify.AddPriorityTxt(oldIssue.Priority.Id, newIssue.Priority.Id)
		utils.LogErr("Error add issueses priority text", err)

		builder.WriteString(str)
	}

	newComment := compareIssuesJournals(oldIssue, newIssue)

	if newComment != "" {
		str, err := notify.AddNewCommentTxt(newComment)
		utils.LogErr("Error concat strings on compare issueses comments", err)

		builder.WriteString(str)
	}

	assignedToName := newIssue.AssignedTo.Name

	if oldIssue.AssignedTo.Id != newIssue.AssignedTo.Id {
		str, err := notify.AddAssignedTxt(oldIssue.AssignedTo.Name, newIssue.AssignedTo.Name)
		utils.LogErr("Error concat strings on compare issueses priority", err)

		assignedToName = ""

		builder.WriteString(str)
	}

	text := builder.String()

	if text == "" {
		return "", nil
	}

	msg, err := notify.CreateMsg(newIssue.Id, newIssue.Priority.Id, newIssue.Tracker.Id, newIssue.Title, text, assignedToName)
	if err != nil {
		return "", err
	}

	return msg, nil
}

// Выводит только последний добавленный коммент, но это оптимальнее, чем мапить комменты каждую минуту
func compareIssuesJournals(oldIssue, newIssue Issue) string {
	if len(newIssue.Journals) > len(oldIssue.Journals) {
		return newIssue.Journals[len(newIssue.Journals)-1].Notes
	}

	return ""
}
