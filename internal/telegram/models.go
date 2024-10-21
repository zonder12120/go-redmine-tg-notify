package telegram

type message struct {
	ChatID     string `json:"chat_id"`
	Text       string `json:"text"`
	Parse_mode string `json:"parse_mode"`
}

func newMessage(chatID, txt string) *message {
	return &message{
		ChatID:     chatID,
		Text:       txt,
		Parse_mode: "MarkdownV2",
	}
}

type response struct {
	Ok          bool   `json:"ok"`
	ErrorCode   int    `json:"error_code,omitempty"`
	Description string `json:"description,omitempty"`
}
