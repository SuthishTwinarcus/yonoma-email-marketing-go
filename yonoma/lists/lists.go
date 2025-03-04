package lists

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

type GroupsYonomaClient struct {
	apiKey  string
	baseURL string
	client  *http.Client
}

func NewGroupYonomaClient(apiKey string) *GroupsYonomaClient {
	return &GroupsYonomaClient{
		apiKey:  apiKey,
		baseURL: "http://localhost:8080/v1/",
		client:  &http.Client{},
	}
}

func (yc *GroupsYonomaClient) Request(method, endpoint string, data interface{}) (map[string]interface{}, error) {
	url := yc.baseURL + endpoint
	var requestBody []byte
	var err error

	if data != nil {
		requestBody, err = json.Marshal(data)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+yc.apiKey)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	resp, err := yc.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, errors.New(string(body))
	}

	var responseData map[string]interface{}
	err = json.Unmarshal(body, &responseData)
	if err != nil {
		return nil, err
	}

	return responseData, nil
}

type Lists struct {
	client *GroupsYonomaClient
}

// NewLists creates a new Lists instance.
func NewLists(client *GroupsYonomaClient) *Lists {
	return &Lists{client: client}
}

// Create a new list
func (g *Lists) Create(listData map[string]interface{}) (map[string]interface{}, error) {
	endpoint := "lists/create"
	return g.client.Request("POST", endpoint, listData)
}

// List all Lists
func (g *Lists) List() (map[string]interface{}, error) {
	endpoint := "lists/list"
	return g.client.Request("GET", endpoint, nil)
}

// Retrieve a specific list by ID
func (g *Lists) Retrieve(listID string) (map[string]interface{}, error) {
	endpoint := "lists/" + listID
	return g.client.Request("GET", endpoint, nil)
}

// Update a list
func (g *Lists) Update(listID string, listData map[string]interface{}) (map[string]interface{}, error) {
	endpoint := "lists/" + listID + "/update"
	return g.client.Request("POST", endpoint, listData)
}

// Delete a list
func (g *Lists) Delete(listID string) (map[string]interface{}, error) {
	endpoint := "lists/" + listID + "/delete"
	return g.client.Request("POST", endpoint, nil)
}
