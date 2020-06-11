package model

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"net/http"
)

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
