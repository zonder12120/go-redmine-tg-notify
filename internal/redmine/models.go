package redmine

type IssuesList struct {
	Issues []Issue `json:"issues"`
}

// Поля структуры упорядочены согласно принципам выравнивания
type Issue struct {
	Title       string  `json:"subject"`
	CreateTime  string  `json:"created_on"`
	UpdateTime  string  `json:"updated_on"`
	ClosedTime  string  `json:"closed_on"`
	Description string  `json:"description"`
	Id          int     `json:"id"`
	AssignedTo  IdField `json:"assigned_to"`
	Project     IdField `json:"project"`
	Tracker     IdField `json:"tracker"`
	Status      IdField `json:"status"`
	Priority    IdField `json:"priority"`
}

type IdField struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type ProjectsList struct {
	Projects []Project `json:"projects"`
}

type Project struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
