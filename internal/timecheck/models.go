package timecheck

type holidays struct {
	Items []item `json:"items"`
}

type item struct {
	Start start `json:"start"`
}

type start struct {
	Date string `json:"date"`
}
