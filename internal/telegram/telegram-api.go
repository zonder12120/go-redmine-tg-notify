package telegram

import (
	"encoding/json"
	"fmt"

	httpreq "github.com/zonder12120/go-redmine-tg-notify/internal/http-req"
	"github.com/zonder12120/go-redmine-tg-notify/pkg/utils"
)

const TG_BASE_URL = "https://api.telegram.org/bot"

type Client struct {
	TelegramToken string
	ChatID        string
}

func NewClient(tkn string, id string) *Client {
	return &Client{
		TelegramToken: tkn,
		ChatID:        id,
	}
}

func (c *Client) SendMsg(txt string) error {
	fmt.Println("Отправляем сообщение: ", txt)

	jsonData, err := json.Marshal(newMessage(c.ChatID, txt))
	if err != nil {
		return fmt.Errorf("error marshalling data for send message req: %s", err)
	}

	url, err := utils.ConcatStrings(TG_BASE_URL, c.TelegramToken, "/sendMessage")
	if err != nil {
		return fmt.Errorf("error concat string for send message req: %s", err)
	}

	err = httpreq.PostReq(url, jsonData)
	if err != nil {
		return fmt.Errorf("error send message req %s", err)
	}
	return nil
}
