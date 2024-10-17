// В этой части пакета находятся функции для сборки сообщения для оповещений

package createmsg

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/zonder12120/go-redmine-tg-notify/pkg/utils"
)

func OffHoursChanges(redmineBaseURL string, issuesIDSlc []int) (string, error) {
	var sb strings.Builder

	for index, issueID := range issuesIDSlc {
		markdownString, err := utils.ConcatStrings(
			"[", strconv.Itoa(issueID), "]",
			"(", redmineBaseURL, "/issues/", strconv.Itoa(issueID), ")",
		)
		if err != nil {
			return "", fmt.Errorf("error concat strings for message off hours changes, task number %d: %s", issueID, err)
		}

		sb.WriteString(markdownString)

		if index != len(issuesIDSlc) {
			sb.WriteString(", ")
		}
	}

	issuesStr := sb.String()

	msg, err := utils.ConcatStrings("Вне рабочего времени произошли изменения в задачах: ", issuesStr)
	if err != nil {
		return "", fmt.Errorf("error concat strings for create message off hours changes: %s", err)
	}
	return msg, nil
}

func NewTask(redmineBaseURL string, issueID int, priorityID int, title string, assignToName string) (string, error) {
	msg, err := utils.ConcatStrings(
		markPriority(priorityID),
		" Добавлена новая задача ",
		"[", strconv.Itoa(issueID), "]",
		"(", redmineBaseURL, "/issues/", strconv.Itoa(issueID), ")",
		" \\- ", title,
		" для ",
		"*",
		assignToName,
		"*",
	)
	if err != nil {
		return "", fmt.Errorf("error concat strings for create message new task: %s", err)
	}
	return msg, nil

}

func AddStatusTxt(oldStatusName, newStatusName string) (string, error) {
	str, err := utils.ConcatStrings(
		"\\\n\\-изменился статус c ",
		"*", oldStatusName, "*",
		" на ",
		"*", newStatusName, "*",
	)
	if err != nil {
		return "", err
	}

	return str, nil
}

func AddPriorityTxt(oldPriorityID, newPriorityID int) (string, error) {
	str, err := utils.ConcatStrings(
		"\\\n\\-изменился приоритет c ",
		"*", oldPriorString(oldPriorityID), "*",
		" на ",
		"*", newPriorString(newPriorityID), "*",
	)
	if err != nil {
		return "", err
	}

	return str, nil
}

func AddAssignedTxt(oldAssignedToName, newAssignedToName string) (string, error) {
	str, err := utils.ConcatStrings(
		"\\\n\\-изменился исполнитель c ",
		"*", oldAssignedToName, "*",
		" на ",
		"*", newAssignedToName, "*",
	)
	if err != nil {
		return "", err
	}

	return str, nil
}

func AddNewCommentTxt(str string) (string, error) {
	str, err := utils.ConcatStrings(
		"\\\n\\-был добавлен комментарий: ",
		"*\\\"", str, "\\\"*",
	)
	if err != nil {
		return "", err
	}

	return str, nil
}

func NewMsg(redmineBaseURL string, issueID int, priorityID int, trackerID int, title string, text string, assignToName string) (string, error) {

	title = utils.MarkDownFilter(title)

	assignStr, err := utils.ConcatStrings("\\\nИсполнитель *", assignToName, "*")
	if err != nil {
		return "", fmt.Errorf("error concat strings to create assigned str in msg-builder: %s", err)
	}

	if assignToName == "" {
		assignStr = ""
	}

	str, err := utils.ConcatStrings(
		markTracker(trackerID),
		markPriority(priorityID),
		"В задаче [", strconv.Itoa(issueID), "]",
		"(", redmineBaseURL, "/issues/", strconv.Itoa(issueID), ")",
		" \\- ", title,
		text,
		assignStr,
	)
	if err != nil {
		return "", fmt.Errorf("error create message in msg-builder: %s", err)
	}

	return str, nil
}
