package ollama

import (
	"encoding/json"
	"fmt"
	"strings"
)

// jsonResponse is implemented by every response shape we expect back from
// the model, so decodeJSON can validate right after unmarshaling.
type jsonResponse interface {
	validate() error
}

// decodeJSON unmarshals raw model content (expected to be a JSON object,
// since callers set ChatRequest.Format = "json") into T and validates it.
func decodeJSON[T jsonResponse](content string) (T, error) {
	var v T

	if strings.TrimSpace(content) == "" {
		return v, fmt.Errorf("ollama: model returned an empty response — the prompt may be too large for the model's context window, or the model failed to generate output")
	}

	if err := json.Unmarshal([]byte(content), &v); err != nil {
		return v, fmt.Errorf("ollama: decoding json response: %w (raw: %q)", err, content)
	}
	if err := v.validate(); err != nil {
		return v, err
	}
	return v, nil
}

// challengeResponse is the JSON shape expected back from the model for GenerateChallenge.
type challengeResponse struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (r challengeResponse) validate() error {
	if strings.TrimSpace(r.Name) == "" {
		return fmt.Errorf("ollama: challenge response missing \"name\"")
	}
	if strings.TrimSpace(r.Description) == "" {
		return fmt.Errorf("ollama: challenge response missing \"description\"")
	}
	return nil
}

// clueResponse is the JSON shape expected back from the model for GenerateClue.
type clueResponse struct {
	Clue string `json:"clue"`
}

func (r clueResponse) validate() error {
	if strings.TrimSpace(r.Clue) == "" {
		return fmt.Errorf("ollama: clue response missing \"clue\"")
	}
	return nil
}

// feedbackResponse is the JSON shape expected back from the model for EvaluateSolution.
type feedbackResponse struct {
	Score       int      `json:"score"`
	Summary     string   `json:"summary"`
	Suggestions []string `json:"suggestions"`
}

func (r feedbackResponse) validate() error {
	if strings.TrimSpace(r.Summary) == "" {
		return fmt.Errorf("ollama: feedback response missing \"summary\"")
	}
	if r.Score < 0 || r.Score > 100 {
		return fmt.Errorf("ollama: feedback response score out of range: %d", r.Score)
	}
	return nil
}
