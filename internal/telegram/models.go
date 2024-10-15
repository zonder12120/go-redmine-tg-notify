package telegram

type Message struct {
	ChatID     string `json:"chat_id"`
	Text       string `json:"text"`
	Parse_mode string `json:"parse_mode"`
}

func newMessage(chatID, txt string) *Message {
	return &Message{
		ChatID:     chatID,
		Text:       txt,
		Parse_mode: "MarkdownV2",
	}
}
