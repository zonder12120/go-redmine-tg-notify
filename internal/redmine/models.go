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
	AssignedTo  IDField   `json:"assigned_to"`
	Project     IDField   `json:"project"`
	Tracker     IDField   `json:"tracker"`
	Status      IDField   `json:"status"`
	Priority    IDField   `json:"priority"`
	ID          int       `json:"id"`
}

type IDField struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Journal struct {
	ID    int    `json:"id"`
	Notes string `json:"notes"`
}

type ProjectsList struct {
	Projects []IDField `json:"projects"`
}

type Holidays struct {
	Items []Item `json:"items"`
}

type Item struct {
	Start Start `json:"start"`
}

type Start struct {
	Date string `json:"date"`
}
