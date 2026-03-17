package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"ecmn/models"
)

type Client struct {
	baseURL string
	token   string
	timeout time.Duration
	http    *http.Client
}

func New(baseURL, token string, timeout time.Duration) *Client {
	return &Client{
		baseURL: baseURL,
		token:   token,
		timeout: timeout,
		http: &http.Client{
			Timeout: timeout,
		},
	}
}

func (c *Client) GetResource(uri string, result interface{}) error {
	url := c.baseURL + uri

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+c.token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.http.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var apiResp models.Response
	if err := json.NewDecoder(resp.Body).Decode(&apiResp); err != nil {
		return fmt.Errorf("failed to decode response: %w", err)
	}

	if apiResp.Code != 1 {
		return fmt.Errorf("api error: code=%d, message=%s", apiResp.Code, apiResp.Message)
	}

	data, err := json.Marshal(apiResp.Data)
	if err != nil {
		return fmt.Errorf("failed to marshal data: %w", err)
	}

	if err := json.Unmarshal(data, result); err != nil {
		return fmt.Errorf("failed to unmarshal resource: %w", err)
	}

	return nil
}
