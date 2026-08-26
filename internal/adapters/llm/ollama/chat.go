package ollama

import (
	"encoding/json"
	"net/http"

	"context"
)

// ChatMessage represents a single message in an Ollama chat conversation.
type ChatMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// ChatRequest represents the request body for the Ollama /api/chat endpoint.
type ChatRequest struct {
	Model    string        `json:"model"`
	Messages []ChatMessage `json:"messages"`
	Stream   bool          `json:"stream"`
	Format   string        `json:"format,omitempty"`
}

// ChatResponse represents the response body for the Ollama /api/chat endpoint
// when Stream is false, the whole reply comes back as a single JSON object.
type ChatResponse struct {
	Model      string      `json:"model"`
	CreatedAt  string      `json:"created_at"`
	Message    ChatMessage `json:"message"`
	Done       bool        `json:"done"`
	DoneReason string      `json:"done_reason"`
}

// Chat sends a chat completion request to the Ollama /api/chat endpoint and
// returns the parsed response. Streaming is always disabled so the response
// body is a single JSON object, matching what SendRequest expects.
func (c *Client) Chat(ctx context.Context, req ChatRequest) (*ChatResponse, error) {
	req.Stream = false

	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	// Send the request to the Ollama API using the SendRequest method
	respBody, err := c.SendRequest(ctx, http.MethodPost, ChatPath, body)
	if err != nil {
		return nil, err
	}

	// Parse the response body into a ChatResponse struct
	var chatResp ChatResponse
	if err := json.Unmarshal(respBody, &chatResp); err != nil {
		return nil, err
	}

	return &chatResp, nil
}
