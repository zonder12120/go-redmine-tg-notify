package notify

import (
	"fmt"
	"strconv"
	"strings"

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

func AddPriorityTxt(oldPriorityId, newPriorityId int) (string, error) {
	str, err := utils.ConcatStrings(
		"\\\n\\-изменился приоритет c ",
		"*", oldPriorString(oldPriorityId), "*",
		" на ",
		"*", newPriorString(newPriorityId), "*",
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

func CreateMsg(issueId int, priorityId int, trackerId int, title string, text string, assignToName string) (string, error) {

	title = markDownFilter(title)

	assignStr, err := utils.ConcatStrings("\\\nИсполнитель *", assignToName, "*")
	if err != nil {
		return "", fmt.Errorf("error concat strings to create assigned str in msg-builder: %s", err)
	}

	if assignToName == "" {
		assignStr = ""
	}

	str, err := utils.ConcatStrings(
		markTracker(trackerId),
		markPriority(priorityId),
		"В задаче [", strconv.Itoa(issueId), "]",
		"(", cfg.RedmineBaseURL, "/issues/", strconv.Itoa(issueId), ")",
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
func markDownFilter(str string) string {
	str = strings.ReplaceAll(str, "*", "\\*")
	str = strings.ReplaceAll(str, "_", "\\_")
	str = strings.ReplaceAll(str, "[", "\\[")
	str = strings.ReplaceAll(str, "]", "\\]")
	str = strings.ReplaceAll(str, "(", "\\(")
	str = strings.ReplaceAll(str, ")", "\\)")
	str = strings.ReplaceAll(str, "~", "\\~")
	str = strings.ReplaceAll(str, ">", "\\>")
	str = strings.ReplaceAll(str, "<", "\\<")
	str = strings.ReplaceAll(str, "#", "\\#")
	str = strings.ReplaceAll(str, "+", "\\+")
	str = strings.ReplaceAll(str, "-", "\\-")
	str = strings.ReplaceAll(str, "=", "\\=")
	str = strings.ReplaceAll(str, "|", "\\|")
	str = strings.ReplaceAll(str, ".", "\\.")
	str = strings.ReplaceAll(str, "!", "\\!")

	return str
}
