package checktime

type Holidays struct {
	Items []Item `json:"items"`
}

type Item struct {
	Start Start `json:"start"`
}

type Start struct {
	Date string `json:"date"`
}
