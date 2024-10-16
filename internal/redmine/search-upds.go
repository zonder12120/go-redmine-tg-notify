package redmine

import (
	"fmt"
	"log"
	"strings"

	"github.com/zonder12120/go-redmine-tg-notify/internal/notify"
	"github.com/zonder12120/go-redmine-tg-notify/internal/timecheck"
	"github.com/zonder12120/go-redmine-tg-notify/pkg/utils"
)

func InitIgnoredIssuesMap(issuesArr []int) map[int]struct{} {
	ignoredIssuesMap := make(map[int]struct{}, len(issuesArr))

	for _, issueId := range issuesArr {
		ignoredIssuesMap[issueId] = struct{}{}
	}

	return ignoredIssuesMap
}

func MakeMapIssuesList(i *IssuesList) map[int]Issue {
	m := make(map[int]Issue, len(i.Issues))
	for _, issue := range i.Issues {
		m[issue.ID] = issue
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

func (c *Client) NotifyUpdates(oldIssueMap, newIssueMap map[int]Issue, ignoredIssuesMap map[int]struct{}) {

	// Сравниваем старые задачи с новыми
	for newIssueID, newIssue := range newIssueMap {

		_, exists := ignoredIssuesMap[newIssueID]

		if exists {
			continue
		}

		oldIssue, exists := oldIssueMap[newIssueID]

		// Если есть новая задача, сразу создаём оповещение
		if !exists && timecheck.IsWorkTime(c.GoogleDevApiKey) {
			err := notify.NotifyNewTask(newIssueID, newIssue.Priority.ID, newIssue.Title, newIssue.AssignedTo.Name)
			utils.LogErr(fmt.Sprintf("Error notify new task number %v", newIssueID), err)

			continue
		}

		msg, err := createDiffMessage(oldIssue, newIssue)
		utils.LogErr(fmt.Sprintf("Error create msg for task number %v", newIssueID), err)

		if msg != "" && timecheck.IsWorkTime(c.GoogleDevApiKey) {
			err := notify.Notify(msg)
			if err != nil {
				log.Println("Error send message to chat: ", err)
			}
		}
	}
}

func createDiffMessage(oldIssue, newIssue Issue) (string, error) {
	var builder strings.Builder

	if oldIssue.Status.ID != newIssue.Status.ID {
		str, err := notify.AddStatusTxt(oldIssue.Status.Name, newIssue.Status.Name)
		utils.LogErr("Error add issueses status text", err)

		builder.WriteString(str)
	}

	if oldIssue.Priority.ID != newIssue.Priority.ID {
		str, err := notify.AddPriorityTxt(oldIssue.Priority.ID, newIssue.Priority.ID)
		utils.LogErr("Error add issueses priority text", err)

		builder.WriteString(str)
	}

	newComment := compareIssuesJournals(oldIssue, newIssue)

	if strings.TrimSpace(newComment) != "" {
		str, err := notify.AddNewCommentTxt(newComment)
		utils.LogErr("Error concat strings on compare issueses comments", err)

		builder.WriteString(str)
	}

	assignedToName := newIssue.AssignedTo.Name

	if oldIssue.AssignedTo.ID != newIssue.AssignedTo.ID {
		str, err := notify.AddAssignedTxt(oldIssue.AssignedTo.Name, newIssue.AssignedTo.Name)
		utils.LogErr("Error concat strings on compare issueses priority", err)

		assignedToName = ""

		builder.WriteString(str)
	}

	text := builder.String()

	if text == "" {
		return "", nil
	}

	msg, err := notify.CreateMsg(newIssue.ID, newIssue.Priority.ID, newIssue.Tracker.ID, newIssue.Title, text, assignedToName)
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
