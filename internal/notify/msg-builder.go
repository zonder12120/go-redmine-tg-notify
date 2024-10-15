package notify

import (
	"fmt"
	"strconv"

	"github.com/zonder12120/go-redmine-tg-notify/pkg/utils"
)

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

func CreateMsg(issueID int, priorityID int, trackerID int, title string, text string, assignToName string) (string, error) {

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
		"(", cfg.RedmineBaseURL, "/issues/", strconv.Itoa(issueID), ")",
		" \\- ", title,
		text,
		assignStr,
	)
	if err != nil {
		return "", fmt.Errorf("error create message in msg-builder: %s", err)
	}

	return str, nil
}

// Добавляем экранирование для спец символов MarkdownV2, чтобы telegram смог распарсить текст
