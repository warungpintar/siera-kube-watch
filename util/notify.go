package util

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/warungpintar/siera-kube-watch/config"
	"log"
	"net/http"
)

type SieraRequest struct {
	Text string `json:"text"`
}

func postEvent(message string) {
	log.Println(message)
	var urls []string

	if config.GlobalConfig.Webhook.Enabled {
		urls = append(urls, config.GlobalConfig.Webhook.Url)
	}

	if config.GlobalConfig.Slack.Enabled {
		urls = append(urls, config.GlobalConfig.Slack.Url)
	}

	for _, url := range urls {
		err := postSieraRequest(message, url)
		if err != nil {
			log.Println(err)
		}
	}
}

func postSieraRequest(text string, url string) (err error) {
	requestModel := &SieraRequest{
		Text: text,
	}

	reqBuffer, err := json.Marshal(requestModel)
	if err != nil {
		return
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(reqBuffer))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return errors.New(fmt.Sprintf("Failed to post event with status code %d", resp.StatusCode))
	}

	return nil
}
