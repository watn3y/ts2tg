package teamspeak

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Client struct {
	http    *http.Client
	baseURL string
	apiKey  string
}

func NewClient(baseURL, apiKey string) *Client {
	return &Client{
		http: &http.Client{
			Timeout: 10 * time.Second,
		},
		baseURL: baseURL,
		apiKey:  apiKey,
	}
}

type apiResponse[T any] struct {
	Body   T `json:"body"`
	Status struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"status"`
}

func decodeResponse[T any](resp *http.Response) (T, error) {
	var zero T
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return zero, fmt.Errorf("HTTP error: %s", resp.Status)
	}

	var result apiResponse[T]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return zero, err
	}

	if result.Status.Code != 0 {
		return zero, fmt.Errorf("TeamSpeak API error %d: %s", result.Status.Code, result.Status.Message)
	}

	return result.Body, nil
}

func (c *Client) sendRequest(endpoint string) (*http.Response, error) {
	url := fmt.Sprintf("%s%s", c.baseURL, endpoint)
	req, err := http.NewRequest(http.MethodPost, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("x-api-key", c.apiKey)
	return c.http.Do(req)
}
