package notify

import (
	"strconv"
	"strings"

	"github.com/zonder12120/go-redmine-tg-notify/pkg/utils"
)

func AddStatusTxt(oldStatusName string, newStatusName string) (string, error) {
	str, err := utils.ConcatStrings(
		"\\\nизменился статус c ",
		"*", oldStatusName, "*",
		" на ",
		"*", newStatusName, "*",
	)
	if err != nil {
		return "", err
	}

	return str, nil
}

func AddPriorityTxt(oldPriorityId int, newPriorityId int) (string, error) {
	str, err := utils.ConcatStrings(
		"\\\nизменился приоритет c ",
		"*", oldPriorString(oldPriorityId), "*",
		"на ",
		"*", newPriorString(newPriorityId), "*",
	)
	if err != nil {
		return "", err
	}

	return str, nil
}

func AddAssignedTxt(oldAssigned string, newAssigned string) (string, error) {
	str, err := utils.ConcatStrings(
		"\\\nизменился исполнитель c ",
		"*", oldAssigned, "*",
		"на ",
		"*", newAssigned, "*",
	)
	if err != nil {
		return "", err
	}

	return str, nil
}

func AddNewCommentTxt(str string) (string, error) {
	str, err := utils.ConcatStrings(
		"\\\nбыл добавлен комментарий: ",
		"*", str, "*",
	)
	if err != nil {
		return "", err
	}

	return str, nil
}

func CreateMsg(issueId int, priorityId int, trackerId int, title string, text string, assignToName string) (string, error) {

	title = fixMarkDown(title)

	str, err := utils.ConcatStrings(
		markTracker(trackerId),
		markPriority(priorityId),
		"В задаче [", strconv.Itoa(issueId), "]",
		"(", cfg.RedmineBaseURL, "/issues/", strconv.Itoa(issueId), ")",
		" \\- ", title,
		text,
		"\\\nИсполнитель *", assignToName, "*",
	)
	if err != nil {
		return "", err
	}

	return str, nil
}

func fixMarkDown(str string) string {
	str = strings.ReplaceAll(str, "*", "\\*")
	str = strings.ReplaceAll(str, "_", "\\_")
	str = strings.ReplaceAll(str, "[", "\\[")
	str = strings.ReplaceAll(str, "]", "\\]")
	str = strings.ReplaceAll(str, "(", "\\(")
	str = strings.ReplaceAll(str, ")", "\\)")
	str = strings.ReplaceAll(str, "~", "\\~")
	str = strings.ReplaceAll(str, ">", "\\>")
	str = strings.ReplaceAll(str, "#", "\\#")
	str = strings.ReplaceAll(str, "+", "\\+")
	str = strings.ReplaceAll(str, "-", "\\-")
	str = strings.ReplaceAll(str, "=", "\\=")
	str = strings.ReplaceAll(str, "|", "\\|")
	str = strings.ReplaceAll(str, ".", "\\.")
	str = strings.ReplaceAll(str, "!", "\\!")

	return str
}
