package model

import (
	"encoding/json"
)

type WorkplaceModel struct {
	Recipient *RecipientModel `json:"recipient"`
	Message   *MessageModel   `json:"message"`
}

type RecipientModel struct {
	ThreadKey string `json:"thread_key"`
}

type MessageModel struct {
	Text string `json:"text"`
}

func (model *WorkplaceModel) New(text string, threadKey string) {
	model.Recipient = &RecipientModel{
		ThreadKey: threadKey,
	}

	model.Message = &MessageModel{
		Text: text,
	}
}

func (model *WorkplaceModel) Send(url string) (err error) {
	buffer, err := json.Marshal(model)
	if err != nil {
		return
	}

	err = postRequest(buffer, url)
	return
}
