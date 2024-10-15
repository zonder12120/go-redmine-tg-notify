package telegram

type Message struct {
	ChatId     string `json:"chat_id"`
	Text       string `json:"text"`
	Parse_mode string `json:"parse_mode"`
}

func newMessage(chatId, txt string) *Message {
	return &Message{
		ChatId:     chatId,
		Text:       txt,
		Parse_mode: "MarkdownV2",
	}
}
