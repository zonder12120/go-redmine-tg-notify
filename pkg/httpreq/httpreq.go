// Пакет включает в себя дефолтный http клиент с установленным таймаутом на 10 секунд
// и функции, отправляющие запросы GET и POST

package httpreq

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"time"
)

var client = &http.Client{
	Timeout: 10 * time.Second,
}

// GET запрос, возвращающий тело
func GetReqBody(url string) ([]byte, error) {
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

// POST запрос без обработки тела, возвращает только ошибку
func PostReq(url string, jsonData []byte) error {
	resp, err := client.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("error sending POST request: %s", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("response by server non-OK: %v", resp.StatusCode)
	}

	return nil
}

// POST запрос, возвращающий тело
func PostReqBody(url string, jsonData []byte) ([]byte, error) {
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
