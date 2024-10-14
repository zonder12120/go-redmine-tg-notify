package redmine

type IssuesList struct {
	Issues []Issue `json:"issues"`
}

type IssueInfo struct {
	Issue Issue `json:"issue"`
}

// Поля структуры упорядочены согласно принципам выравнивания
type Issue struct {
	Journals    []Journal `json:"journals,omitempty"`
	Title       string    `json:"subject"`
	CreateTime  string    `json:"created_on"`
	UpdateTime  string    `json:"updated_on"`
	ClosedTime  string    `json:"closed_on"`
	Description string    `json:"description"`
	AssignedTo  IdField   `json:"assigned_to"`
	Project     IdField   `json:"project"`
	Tracker     IdField   `json:"tracker"`
	Status      IdField   `json:"status"`
	Priority    IdField   `json:"priority"`
	Id          int       `json:"id"`
}

type IdField struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Journal struct {
	Id    int    `json:"id"`
	Notes string `json:"notes"`
}

type ProjectsList struct {
	Projects []IdField `json:"projects"`
}
