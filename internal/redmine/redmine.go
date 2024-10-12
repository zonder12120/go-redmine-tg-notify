package redmine

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/zonder12120/go-redmine-tg-notify/internal/config"
	"github.com/zonder12120/go-redmine-tg-notify/internal/http"
	"github.com/zonder12120/go-redmine-tg-notify/pkg/utils"
)

var builder strings.Builder

func GetIssuesList(cfg config.Config) IssuesList {
	url := utils.ConcatStrings([]string{cfg.RedmineBaseURL, "/issues.json?key=", cfg.RedmineAPIKey, getProjectsFilter(), "&limit=100"})

	body := http.GetRespBody(url)

	var issuesList IssuesList

	err := json.Unmarshal(body, &issuesList)
	if err != nil {
		log.Fatalf("Error encoding body from GET req: %s,\nerr: %s\n", url, err)
	}

	return issuesList
}

func GetIssueInfo(cfg config.Config, issueId int) Issue {
	url := utils.ConcatStrings([]string{cfg.RedmineBaseURL, "/issues/", fmt.Sprintf("%v.json", issueId), "?include=journals?key=", cfg.RedmineAPIKey, getProjectsFilter()})
	body := http.GetRespBody(url)

	var issue Issue

	err := json.Unmarshal(body, &issue)
	if err != nil {
		log.Fatalf("Error encoding body from GET req: %s,\nerr: %s\n", url, err)
	}

	return issue
}

func getProjectsFilter() string {
	for _, id := range config.ProjectsId {
		builder.WriteString(fmt.Sprintf("&project_id=%v", id))
	}

	filter := builder.String()

	builder.Reset()

	return filter
}
