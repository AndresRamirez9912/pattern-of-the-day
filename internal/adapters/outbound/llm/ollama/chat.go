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
	Options  *ChatOptions  `json:"options,omitempty"`
}

// ChatOptions carries Ollama model runtime options for a single request.
type ChatOptions struct {
	// NumCtx sets the context window size, in tokens. Ollama defaults this
	// to a small value (historically 2048) regardless of what the model
	// itself supports, so large prompts get silently truncated unless this
	// is set explicitly.
	NumCtx int `json:"num_ctx,omitempty"`
	// NumPredict caps how many tokens the model may generate in the
	// response. Ollama's own default for this is small in many versions
	// (e.g. 128), which silently cuts the response off mid-sentence before
	// the JSON object is closed — that looks like a parsing bug but is
	// actually generation being stopped early, so it needs to be set
	// explicitly for anything longer than a short reply.
	NumPredict int `json:"num_predict,omitempty"`
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

// Context window sizes considered when sizing a request for its prompt.
// minNumCtx matches (and makes explicit) Ollama's own historical default,
// so small prompts don't pay for a larger window than they need.
const (
	minNumCtx = 2048
	maxNumCtx = 32768
	// numPredict is a generous, fixed cap on response length. Every prompt
	// in this package asks for a single short-to-medium JSON object (a
	// challenge, a clue, or feedback), so one fixed budget covers all of
	// them without needing per-call tuning.
	numPredict = 4096
)

// contextOptionsFor sizes the context window (num_ctx) for a request based
// on the combined length of its prompts, so large inputs — e.g. a whole
// project's source code submitted for evaluation — aren't silently
// truncated by Ollama's small default context window. It also sets
// num_predict explicitly, since Ollama's own default for it is too small
// for a multi-paragraph response and would truncate it well before num_ctx
// becomes the limiting factor.
func contextOptionsFor(prompts ...string) *ChatOptions {
	chars := 0
	for _, p := range prompts {
		chars += len(p)
	}

	// Rough heuristic: ~4 characters per token, plus headroom for the
	// model's own output (bounded by numPredict) and a safety margin.
	estimatedTokens := chars/4 + numPredict
	estimatedTokens += estimatedTokens / 5

	numCtx := minNumCtx
	for numCtx < estimatedTokens && numCtx < maxNumCtx {
		numCtx *= 2
	}

	return &ChatOptions{NumCtx: numCtx, NumPredict: numPredict}
}
