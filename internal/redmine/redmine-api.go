package redmine

import (
	"encoding/json"
	"fmt"

	httpreq "github.com/zonder12120/go-redmine-tg-notify/internal/http-req"
	"github.com/zonder12120/go-redmine-tg-notify/pkg/utils"
)

type Client struct {
	RedmineBaseURL string
	RedmineAPIKey  string
	ProjectsId     []int
}

func NewClient(url string, key string, projectsId []int) *Client {
	return &Client{
		RedmineBaseURL: url,
		RedmineAPIKey:  key,
		ProjectsId:     projectsId,
	}
}

func (c *Client) GetIssuesList() (IssuesList, error) {
	var issuesList IssuesList

	url, err := utils.ConcatStrings(c.RedmineBaseURL, "/issues.json?key=", c.RedmineAPIKey, getProjectsFilter(c.ProjectsId), "&limit=100")
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

func (c *Client) GetProjectsList() error {
	var projectsList ProjectsList

	url, err := utils.ConcatStrings(c.RedmineBaseURL, "/projects.json?key=", c.RedmineAPIKey)
	if err != nil {
		return utils.HadleError("Error get project list req", err)
	}

	body, err := httpreq.GetRespBody(url)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &projectsList)
	if err != nil {
		return utils.HadleError("Error encoding body from get project list req", err)
	}

	outputProjectList(projectsList)

	return nil
}

func outputProjectList(pl ProjectsList) {
	fmt.Println("Projects List:")
	for index, p := range pl.Projects {
		fmt.Printf("id: %d, name: %s\n", p.Id, p.Name)

		if index == len(pl.Projects)-1 {
			fmt.Println("")
		}
	}
}

func getProjectsFilter(projectsId []int) string {
	for _, id := range projectsId {
		builder.WriteString(fmt.Sprintf("&project_id=%v", id))
	}

	filter := builder.String()

	builder.Reset()

	return filter
}
