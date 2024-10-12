package redmine

type IssuesList struct {
	Issues []Issue `json:"issues"`
}

// Поля структуры упорядочены согласно принципам выравнивания
type Issue struct {
	Title      string  `json:"subject"`
	CreateTime string  `json:"created_on"`
	UpdateTime string  `json:"updated_on"`
	ClosedTime string  `json:"closed_on"`
	Id         int     `json:"id"`
	AssignedTo IdField `json:"assigned_to"`
	Project    IdField `json:"project_id"`
	Tracker    IdField `json:"tracker_id"`
	Status     IdField `json:"status_id"`
	Priority   IdField `json:"priority_id"`
}

type IdField struct {
	Id int `json:"id"`
}
