package grafana

import (
	"fmt"
	"net/http"
)

type Client struct {
	token   string
	BaseURL string
	client  *http.Client
}

func New(baseURL string, token string) (*Client, error) {
	if len(baseURL) == 0 {
		return nil, fmt.Errorf("base url cannot be null")
	}

	if len(token) == 0 {
		return nil, fmt.Errorf("token is required")
	}

	return &Client{
		BaseURL: baseURL,

		token:  token,
		client: &http.Client{},
	}, nil
}

func (c *Client) Do(req *http.Request) (*http.Response, error) {
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.token))
	req.Header.Add("Content-Type", "application/json")
	return c.client.Do(req)
}
