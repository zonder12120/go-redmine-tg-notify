package telegram

import (
	"encoding/json"
	"log"

	httpreq "github.com/zonder12120/go-redmine-tg-notify/pkg/httpreq"
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
	jsonDataReq, err := json.Marshal(newMessage(c.ChatID, txt))
	if err != nil {
		return err
	}

	url, err := utils.ConcatStrings(TG_BASE_URL, c.TelegramToken, "/sendMessage")
	if err != nil {
		return err
	}

	body, err := httpreq.PostReqBody(url, jsonDataReq)
	if err != nil {
		log.Println(err)
	}

	var jsonDataResp response

	if len(body) != 0 {
		err = json.Unmarshal(body, &jsonDataResp)
		if err != nil {
			return err
		}

		if !jsonDataResp.Ok {
			log.Println("Error response description: ", jsonDataResp.Description)
		}
	} else {
		log.Println("Empty body by sendMessage")
	}

	return nil
}
