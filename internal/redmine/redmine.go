package redmine

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"strings"

	"github.com/zonder12120/go-redmine-tg-notify/internal/http"
	"github.com/zonder12120/go-redmine-tg-notify/pkg/config"
)

var builder strings.Builder

func GetIssuesList(cfg config.Config) IssuesList {
	url := makeURL(cfg)

	resp, err := http.Client.Get(url)
	if err != nil {
		log.Fatalf("Error sending GET request %s,\nerr: %s", url, err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error read response body from GET req: %s,\nerr: %s", url, err)
	}

	var issuesList IssuesList

	err = json.Unmarshal(body, &issuesList)
	if err != nil {
		log.Fatalf("Error encoding body from GET req: %s,\nerr: %s", url, err)
	}

	return issuesList
}

func getProjectsFilter() string {
	pId := config.ProjectsId

	for _, id := range pId {
		builder.WriteString(fmt.Sprintf("&project_id=%v", id))
	}
	return builder.String()
}

func connectStrings(s []string) string {
	builder.Reset()
	for _, string := range s {
		builder.WriteString(string)
	}
	return builder.String()
}

func makeURL(cfg config.Config) string {
	projectsFilter := getProjectsFilter()

	return connectStrings(
		[]string{
			cfg.RedmineBaseURL,
			"/issues.json?key=",
			cfg.RedmineAPIKey,
			projectsFilter,
		})
}
