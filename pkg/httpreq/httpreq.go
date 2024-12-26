// Пакет включает в себя дефолтный http клиент с установленным таймаутом на 10 секунд
// и функции, отправляющие запросы GET и POST

package httpreq

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"time"
)

var client = &http.Client{
	Timeout: 10 * time.Second,
}

func doRequestBody(req *http.Request) ([]byte, error) {
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %s", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("server returned non-OK status: %v", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %s", err)
	}

	return body, nil
}

// GET запрос, возвращающий тело
func GetReqBody(url string) ([]byte, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating GET request: %s", err)
	}

	return doRequestBody(req)
}

// POST запрос, возвращающий тело
func PostReqBody(url string, jsonData []byte) ([]byte, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("error creating POST request: %s", err)
	}
	req.Header.Set("Content-Type", "application/json")

	return doRequestBody(req)
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
