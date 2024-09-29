package webpse

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var (
	BaseURL = "https://webppt.webp.se"
)

type Client struct {
	httpClient *http.Client
	apiKey     string
}

func New(apiKey string) *Client {
	return &Client{
		httpClient: &http.Client{},
		apiKey:     apiKey,
	}
}

func get[T any](c Client, endpoint string) (*Response[T], error) {
	response := new(Response[T])

	req, err := http.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("api-key", c.apiKey)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if err = json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) GetUserInfo() (*Response[UserInfo], error) {
	endpoint := fmt.Sprintf("%s/v1/user/info", BaseURL)
	return get[UserInfo](*c, endpoint)
}

func (c *Client) GetUserStats() (*Response[UserStats], error) {
	endpoint := fmt.Sprintf("%s/v1/user/stats", BaseURL)
	return get[UserStats](*c, endpoint)
}

func (c *Client) GetProxiesStats() (*Response[ProxyList], error) {
	endpoint := fmt.Sprintf("%s/v1/proxy", BaseURL)
	return get[ProxyList](*c, endpoint)

}
