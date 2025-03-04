package api

import (
	"fmt"
	"net/http"
)

// Client represents the API client
type Client struct {
	APIKey  string
	BaseURL string
}

// NewClient initializes a new API client
func NewClient(apiKey string) *Client {
	return &Client{
		APIKey:  apiKey,
		BaseURL: "https://api.yonoma.io/v1/", // Replace with actual API URL
	}
}

// Get sends a GET request to the API
func (c *Client) Get(path string) (*http.Response, error) {
	url := c.BaseURL + path
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+c.APIKey)

	client := &http.Client{}
	return client.Do(req)
}

// Post sends a POST request to the API
func (c *Client) Post(path string, body *http.Request) (*http.Response, error) {
	url := c.BaseURL + path
	req, err := http.NewRequest("POST", url, body.Body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+c.APIKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	return client.Do(req)
}

// Delete sends a DELETE request to the API
func (c *Client) Delete(path string) (*http.Response, error) {
	url := c.BaseURL + path
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+c.APIKey)

	client := &http.Client{}
	return client.Do(req)
}

// PrintClientInfo prints client details (for debugging)
func (c *Client) PrintClientInfo() {
	fmt.Println("API Client Initialized:")
	fmt.Println("Base URL:", c.BaseURL)
	fmt.Println("API Key:", c.APIKey)
}
