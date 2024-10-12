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
	url, err := utils.ConcatStrings([]string{cfg.RedmineBaseURL, "/issues.json?key=", cfg.RedmineAPIKey, getProjectsFilter(), "&limit=100"})
	utils.FatalOnError(err)

	body, err := http.GetRespBody(url)
	utils.FatalOnError(err)

	var issuesList IssuesList

	err = json.Unmarshal(body, &issuesList)
	if err != nil {
		log.Fatalf("Error encoding body from GET req: %s, err: %s", url, err)
	}

	return issuesList
}

func GetIssueInfo(cfg config.Config, issueId int) Issue {
	url, err := utils.ConcatStrings([]string{cfg.RedmineBaseURL, "/issues/", fmt.Sprintf("%v.json", issueId), "?include=journals?key=", cfg.RedmineAPIKey, getProjectsFilter()})
	utils.FatalOnError(err)

	body, err := http.GetRespBody(url)
	utils.FatalOnError(err)

	var issue Issue

	err = json.Unmarshal(body, &issue)
	if err != nil {
		log.Fatalf("Error encoding body from GET req: %s,err: %s", url, err)
	}

	return issue
}

func GetProjectsList(cfg config.Config) ProjectsList {
	url, err := utils.ConcatStrings([]string{cfg.RedmineBaseURL, "/projects.json?key=", cfg.RedmineAPIKey})
	utils.FatalOnError(err)

	body, err := http.GetRespBody(url)
	utils.FatalOnError(err)

	var projectsList ProjectsList

	err = json.Unmarshal(body, &projectsList)
	if err != nil {
		log.Fatalf("Error encoding body from GET req: %s,err: %s", url, err)
	}

	return projectsList
}

func getProjectsFilter() string {
	for _, id := range config.ProjectsId {
		builder.WriteString(fmt.Sprintf("&project_id=%v", id))
	}

	filter := builder.String()

	builder.Reset()

	return filter
}
