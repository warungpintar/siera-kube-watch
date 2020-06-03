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

	if config.GlobalConfig.Telegram.Enabled {
		url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage?chat_id=%s", config.GlobalConfig.Telegram.Token, config.GlobalConfig.Telegram.ChatId)
		urls = append(urls, url)
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

	buffer, err := json.Marshal(requestModel)
	if err != nil {
		return
	}

	err = postRequest(buffer, url)
	return
}

func postRequest(buffer []byte, url string) (err error) {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(buffer))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		log.Println(url)
		return errors.New(fmt.Sprintf("Failed to post event to %s with status code %d", url, resp.StatusCode))
	}

	return nil
}
