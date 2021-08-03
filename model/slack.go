package model

import (
	"encoding/json"
	"fmt"

	"github.com/warungpintar/siera-kube-watch/config"
)

type SlackModel struct {
	Text     string `json:"text"`
	Channel  string `json:"channel"`
	Username string `json:"username"`
}

func (model *SlackModel) New(text string) {
	model.Text = text
	model.Channel = config.GlobalConfig.Slack.Channel
	model.Username = config.GlobalConfig.Slack.Username
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
