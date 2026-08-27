package ollama

import (
	"context"
	"io"
	"net/http"
	"strings"
)

// Supported paths for the Ollama API endpoints
const (
	ChatPath = "/api/chat"
)

// Client represents the Ollama client for interacting with the Ollama HTTP API.
type Client struct {
	BaseUrl string
	ApiKey  *string
	Client  *http.Client
}

// NewClient creates a new instance of the Ollama client with the provided base URL and API key.
func NewClient(baseUrl string, apiKey *string) *Client {
	if strings.TrimSpace(baseUrl) == "" {
		panic("Ollama base URL must be provided")
	}

	return &Client{
		BaseUrl: baseUrl,
		ApiKey:  apiKey,
		Client:  &http.Client{},
	}
}

// SendRequest sends an HTTP request to the Ollama API with the specified method, path, and body
func (c *Client) SendRequest(ctx context.Context, method string, path string, body []byte) ([]byte, error) {
	// Create the full URL by combining the base URL and the path
	url := c.BaseUrl + path

	// Create a new HTTP request with the provided method, URL, and body
	req, err := http.NewRequestWithContext(ctx, method, url, strings.NewReader(string(body)))
	if err != nil {
		return nil, err
	}

	// Set the necessary headers for the request
	req.Header.Set("Content-Type", "application/json")
	if c.ApiKey != nil {
		req.Header.Set("Authorization", "Bearer "+*c.ApiKey)
	}

	// Send the HTTP request using the client
	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Read the response body
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return responseBody, nil
}
