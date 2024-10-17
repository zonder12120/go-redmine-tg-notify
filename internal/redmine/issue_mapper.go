package redmine

import (
	"fmt"
)

func InitIgnoredIssuesMap(issuesArr []int) map[int]struct{} {
	ignoredIssuesMap := make(map[int]struct{}, len(issuesArr))

	for _, issueId := range issuesArr {
		ignoredIssuesMap[issueId] = struct{}{}
	}

	return ignoredIssuesMap
}

// Мапим список задач
func MakeMapIssuesList(i *IssuesList) map[int]Issue {
	m := make(map[int]Issue, len(i.Issues))
	for _, issue := range i.Issues {
		m[issue.ID] = issue
	}

	return m
}

// Добавление комментариев в мапу
func (c *Client) AddJournalsIssuesMap(issueMap map[int]Issue) error {
	for id, issue := range issueMap {
		IssueInfo, err := c.GetIssueInfo(id)
		if err != nil {
			return fmt.Errorf("error get old issue info %s", err)
		}

		issue.Journals = IssueInfo.Issue.Journals

		issueMap[id] = issue
	}

	return nil
}
