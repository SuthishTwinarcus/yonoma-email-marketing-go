package contacts

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Contact represents a contact in the system
type ContactsYonomaClient struct {
	apiKey  string
	baseURL string
	client  *http.Client
}

func NewContactsYonomaClient(apiKey string) *ContactsYonomaClient {
	return &ContactsYonomaClient{
		apiKey:  apiKey,
		baseURL: "http://localhost:8080/v1/",
		client:  &http.Client{},
	}
}

func (yc *ContactsYonomaClient) Request(method, endpoint string, data interface{}) (map[string]interface{}, error) {
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

type Contacts struct {
	client *ContactsYonomaClient
}

func NewContacts(client *ContactsYonomaClient) *Contacts {
	return &Contacts{client: client}
}

func (c *Contacts) Create(groupId string, contactData map[string]interface{}) (map[string]interface{}, error) {
	endpoint := fmt.Sprintf("contacts/%s/create", groupId)
	return c.client.Request("POST", endpoint, contactData)
}

func (c *Contacts) Update(groupId, contactId string, contactData map[string]interface{}) (map[string]interface{}, error) {
	endpoint := fmt.Sprintf("contacts/%s/status/%s", groupId, contactId)
	return c.client.Request("POST", endpoint, contactData)
}

func (c *Contacts) AddTag(contactId string, contactData map[string]interface{}) (map[string]interface{}, error) {
	endpoint := fmt.Sprintf("contacts/tags/%s/add", contactId)
	return c.client.Request("POST", endpoint, contactData)
}

func (c *Contacts) RemoveTag(contactId string, contactData map[string]interface{}) (map[string]interface{}, error) {
	endpoint := fmt.Sprintf("contacts/tags/%s/remove", contactId)
	return c.client.Request("POST", endpoint, contactData)
}
