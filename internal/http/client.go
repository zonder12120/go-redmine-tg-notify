package http

import (
	"io"
	"log"
	"net/http"
)

var Client = http.DefaultClient

func GetRespBody(url string) []byte {
	resp, err := Client.Get(url)
	if err != nil {
		log.Fatalf("Error sending GET request %s,\nerr: %s\n", url, err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Response by server non-OK: %v\n", resp.StatusCode)
	}

	log.Printf("Request successfully sent: %s", url)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error read response body from GET req: %s,\nerr: %s\n", url, err)
	}

	return body
}
