package http

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

var Client = http.DefaultClient

func GetRespBody(url string) ([]byte, error) {
	resp, err := Client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error sending GET request %s: %v", url, err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("response by server non-OK: %v", resp.StatusCode)
	}

	log.Printf("Request successfully sent: %s", url)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error read response body from GET request %s: %s", url, err)
	}

	return body, nil
}
