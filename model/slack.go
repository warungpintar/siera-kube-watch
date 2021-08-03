package model

import (
	"encoding/json"
	"fmt"
)

type SlackModel struct {
	Text     string `json:"text"`
	Channel  string `json:"channel"`
	Username string `json:"username"`
}

func (model *SlackModel) New(text string, channel string, username string) {
	model.Text = text
	model.Channel = channel
	model.Username = username
}

func (model *SlackModel) Send(url string) (err error) {
	buffer, err := json.Marshal(model)
	fmt.Println(string(buffer))
	if err != nil {
		return
	}

	err = postRequest(buffer, url)
	return
}
