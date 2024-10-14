package telegram

type Message struct {
	ChatId     string `json:"chat_id"`
	Text       string `json:"text"`
	Parse_mode string `json:"parse_mode"`
}
