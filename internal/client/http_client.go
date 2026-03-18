package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"route-nn/internal/config"
	"time"

	"github.com/go-resty/resty/v2"
)

type Client struct {
	HTTPClient *resty.Client
}

func NewClient(cfg *config.Config) *Client {
	return &Client{
		HTTPClient: resty.New().
			SetTimeout(cfg.Timeout*time.Second).
			SetBasicAuth(cfg.Username, cfg.Password),
	}
}

// Универсальный метод для получения данных
func (c *Client) GetJSON(url string, result interface{}) error {
	resp, err := c.HTTPClient.R().Get(url)
	if err != nil {
		return fmt.Errorf("network error: %w", err)
	}

	// Чистим BOM (1С style)
	body := bytes.TrimPrefix(resp.Body(), []byte("\xef\xbb\xbf"))

	if resp.IsError() {
		return fmt.Errorf("server error %d: %s", resp.StatusCode(), string(body))
	}

	// Декодируем очищенные данные
	if err := json.Unmarshal(body, result); err != nil {
		return fmt.Errorf("json decode error: %w", err)
	}

	return nil
}
