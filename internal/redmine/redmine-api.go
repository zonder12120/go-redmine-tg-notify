package redmine

import (
	"encoding/json"
	"fmt"
	"strings"

	httpreq "github.com/zonder12120/go-redmine-tg-notify/internal/httpreq"
	"github.com/zonder12120/go-redmine-tg-notify/pkg/utils"
)

type Client struct {
	RedmineBaseURL string
	RedmineAPIKey  string
	ProjectsID     []int
}

func NewClient(url string, key string, projectsID []int) *Client {
	return &Client{
		RedmineBaseURL: url,
		RedmineAPIKey:  key,
		ProjectsID:     projectsID,
	}
}

func (c *Client) GetIssuesList() (IssuesList, error) {
	var issuesList IssuesList

	url, err := utils.ConcatStrings(c.RedmineBaseURL, "/issues.json?key=", c.RedmineAPIKey, getProjectsFilter(c.ProjectsID), "&limit=100")
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

func (c *Client) GetIssueInfo(issueID int) (IssueInfo, error) {
	var issueInfo IssueInfo

	url, err := utils.ConcatStrings(c.RedmineBaseURL, "/issues/", fmt.Sprintf("%v.json", issueID), "?include=journals&key=", c.RedmineAPIKey, "&limit=100")
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

func (c *Client) GetProjectsList() error {
	var projectsList ProjectsList

	url, err := utils.ConcatStrings(c.RedmineBaseURL, "/projects.json?key=", c.RedmineAPIKey)
	if err != nil {
		return fmt.Errorf("еrror get project list req: %s", err)
	}

	body, err := httpreq.GetRespBody(url)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &projectsList)
	if err != nil {
		return fmt.Errorf("error encoding body from get project list req %s", err)
	}

	outputProjectList(projectsList)

	return nil
}

func outputProjectList(pl ProjectsList) {
	fmt.Println("Projects List:")
	for index, p := range pl.Projects {
		fmt.Printf("id: %d, name: %s\n", p.ID, p.Name)

		if index == len(pl.Projects)-1 {
			fmt.Println("")
		}
	}
}

func getProjectsFilter(projectsID []int) string {
	var builder strings.Builder

	for _, id := range projectsID {
		builder.WriteString(fmt.Sprintf("&project_id=%v", id))
	}

	filter := builder.String()

	builder.Reset()

	return filter
}
