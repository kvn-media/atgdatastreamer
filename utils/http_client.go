// http_client_utils.go

package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// HTTPClient is a simple HTTP client for sending requests
type HTTPClient struct {
	BaseURL string
}

// NewHTTPClient creates a new instance of HTTPClient
func NewHTTPClient(baseURL string) *HTTPClient {
	return &HTTPClient{
		BaseURL: baseURL,
	}
}

// SendRequest sends an HTTP request with JSON data
func (c *HTTPClient) SendRequest(endpoint string, method string, requestData interface{}) ([]byte, error) {
	url := fmt.Sprintf("%s/%s", c.BaseURL, endpoint)
	jsonData, err := json.Marshal(requestData)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return responseBody, nil
}
