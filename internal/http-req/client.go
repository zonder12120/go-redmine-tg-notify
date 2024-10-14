package httpreq

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

var client = http.DefaultClient

func GetRespBody(url string) ([]byte, error) {
	resp, err := client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error sending GET request: %s", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("response by server non-OK: %v", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error read response body from GET request: %s", err)
	}

	return body, nil
}

func PostReq(url string, jsonData []byte) error {
	resp, err := client.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("error sending POST request: %s", err)
	}

	fmt.Println("Отправили вот такое тело: ", string(jsonData))

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("response by server non-OK: %v", resp.StatusCode)
	}

	return nil
}

func PostRespBody(url string, jsonData []byte) ([]byte, error) {
	resp, err := client.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("error sending POST request: %s", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("response by server non-OK: %v", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error read response body from GET request: %s", err)
	}

	return body, nil
}
