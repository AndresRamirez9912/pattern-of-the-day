package handlers

// CreateUserRequest represents the payload for creating a new user.
type CreateUserRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

// CreateUserResponse represents the response payload after creating a new user.
type CreateUserResponse struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

// CreateChallengeRequest represents the payload for creating a new challenge.
type CreateChallengeRequest struct {
	UserName   string `json:"username"`
	Difficulty string `json:"difficulty"`
	Type       string `json:"type"`
	Target     string `json:"target"`
	Topic      string `json:"topic"`
}

// CreateChallengeResponse represents the response payload after creating a new challenge.
type CreateChallengeResponse struct {
	Challenge Challenge `json:"challenge"`
	Attempt   Attempt   `json:"attempt"`
}

// CreateClueRequest represents the payload for creating a new clue.
type CreateClueRequest struct {
	ChallengeId int64 `json:"challenge_id"`
}

// CreateClueResponse represents the response payload after creating a new clue.
type CreateClueResponse struct {
	Clue Clue `json:"clue"`
}

// Challenge represents a coding challenge.
type Challenge struct {
	ChallengeId int64  `json:"challenge_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Difficulty  string `json:"difficulty"`
	Type        string `json:"type"`
	Target      string `json:"target"`
}

// Attempt represents an attempt made by a user on a challenge.
type Attempt struct {
	AttemptId int64  `json:"attempt_id"`
	Status    string `json:"status"`
	Sequence  int    `json:"sequence"`
}

// Clue represents a clue related to a challenge.
type Clue struct {
	ClueId      int64  `json:"clue_id"`
	ChallengeId int64  `json:"challenge_id"`
	Description string `json:"description"`
}
