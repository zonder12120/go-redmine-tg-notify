package redmine

import (
	"encoding/json"
	"fmt"

	"github.com/zonder12120/go-redmine-tg-notify/internal/config"
	httpreq "github.com/zonder12120/go-redmine-tg-notify/internal/http-req"
	"github.com/zonder12120/go-redmine-tg-notify/pkg/utils"
)

type Client struct {
	RedmineBaseURL string
	RedmineAPIKey  string
}

func NewClient(url string, key string) *Client {
	return &Client{
		RedmineBaseURL: url,
		RedmineAPIKey:  key,
	}
}

func (c *Client) GetIssuesList() (IssuesList, error) {
	var issuesList IssuesList

	url, err := utils.ConcatStrings(c.RedmineBaseURL, "/issues.json?key=", c.RedmineAPIKey, getProjectsFilter(), "&limit=100")
	if err != nil {
		return issuesList, fmt.Errorf("error concat strings for get issues request %s", err)
	}

	body, err := httpreq.GetRespBody(url)
	if err != nil {
		return issuesList, fmt.Errorf("error get issues req %s", err)
	}

	err = json.Unmarshal(body, &issuesList)
	if err != nil {
		return issuesList, fmt.Errorf("error encoding body from get issues req %s", err)
	}
	return issuesList, nil
}

func (c *Client) GetIssueInfo(issueId int) (IssueInfo, error) {
	var issueInfo IssueInfo

	url, err := utils.ConcatStrings(c.RedmineBaseURL, "/issues/", fmt.Sprintf("%v.json", issueId), "?include=journals&key=", c.RedmineAPIKey, "&limit=100")
	if err != nil {
		return issueInfo, err
	}

	body, err := httpreq.GetRespBody(url)
	if err != nil {
		return issueInfo, fmt.Errorf("error get issue info req %s", err)
	}

	err = json.Unmarshal(body, &issueInfo)
	if err != nil {
		return issueInfo, fmt.Errorf("error encoding body from get issue info req %s", err)
	}

	return issueInfo, nil
}

func (c *Client) GetProjectsList() (ProjectsList, error) {
	var projectsList ProjectsList

	url, err := utils.ConcatStrings(c.RedmineBaseURL, "/projects.json?key=", c.RedmineAPIKey)
	if err != nil {
		return projectsList, utils.HadleError("Error get project list req", err)
	}

	body, err := httpreq.GetRespBody(url)
	if err != nil {
		return projectsList, err
	}

	err = json.Unmarshal(body, &projectsList)
	if err != nil {
		return projectsList, utils.HadleError("Error encoding body from get project list req", err)
	}

	return projectsList, nil
}

func getProjectsFilter() string {
	for _, id := range config.ProjectsId {
		builder.WriteString(fmt.Sprintf("&project_id=%v", id))
	}

	filter := builder.String()

	builder.Reset()

	return filter
}
