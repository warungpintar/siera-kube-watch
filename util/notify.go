package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/warungpintar/siera-kube-watch/config"
	"log"
	"net/http"
)

type SieraRequest struct {
	Text string `json:"text"`
}

func postEvent(text string) {
	if config.GlobalConfig.Webhook.Enabled {
		url := config.GlobalConfig.Webhook.Url

		requestModel := &SieraRequest{
			Text: text,
		}

		reqBuffer, err := json.Marshal(requestModel)
		if err != nil {
			fmt.Println(err)
			return
		}
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(reqBuffer))
		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			log.Println(err)
			return
		}
		defer resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			log.Println(err)
		}
	}

	return
}
