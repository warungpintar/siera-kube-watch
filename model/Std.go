package model

import "encoding/json"

type StdModel struct {
	Text string `json:"text"`
}

func (model *StdModel) New(text string) {
	model.Text = text
}

func (model *StdModel) Send(url string) (err error) {
	buffer, err := json.Marshal(model)
	if err != nil {
		return
	}

	err = postRequest(buffer, url)
	return
}
