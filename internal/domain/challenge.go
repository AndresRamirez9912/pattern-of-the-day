package domain

import (
	"fmt"
	"strings"
)

// maxClues is the maximum number of clues a challenge can have.
const maxClues = 3

// Challenge defines the properties of a challenge
type Challenge struct {
	Id          int64
	Name        string
	Description string
	Difficulty  Difficulty
	// Type indicates the subject area the challenge covers
	Type ChallengeType
	// Target is the specific subject the challenge evaluates
	Target string
	Clues  []Clue
	UserId int64
}

// NewChallenge creates a new challenge with the given parameters. UserId is
// left unset — a freshly generated challenge has no owner yet, that's
// assigned separately once the requesting user has been resolved.
func NewChallenge(id int64, name, description string, difficulty Difficulty, challengeType ChallengeType, target string) *Challenge {
	return &Challenge{
		Id:          id,
		Name:        name,
		Description: description,
		Difficulty:  difficulty,
		Type:        challengeType,
		Target:      target,
		Clues:       []Clue{},
	}
}

// Difficulty is a custom type to define the supported challenge difficulty
type Difficulty string

// Supported challenge difficulty
const (
	ChallengeDifficultyEasy   Difficulty = "easy"
	ChallengeDifficultyMedium Difficulty = "medium"
	ChallengeDifficultyHard   Difficulty = "hard"
)

// AllDifficulties returns every supported challenge difficulty.
func AllDifficulties() []Difficulty {
	return []Difficulty{ChallengeDifficultyEasy, ChallengeDifficultyMedium, ChallengeDifficultyHard}
}

// IsValidDifficulty reports whether d is one of the supported difficulties.
func IsValidDifficulty(d Difficulty) bool {
	for _, valid := range AllDifficulties() {
		if d == valid {
			return true
		}
	}
	return false
}

// ParseDifficulty validates and normalizes a raw string into a Difficulty.
func ParseDifficulty(s string) (Difficulty, error) {
	d := Difficulty(strings.ToLower(strings.TrimSpace(s)))
	if !IsValidDifficulty(d) {
		return "", fmt.Errorf("invalid difficulty %q: must be one of %v", s, AllDifficulties())
	}
	return d, nil
}

// ChallengeType is a custom type to define which subject area a challenge covers.
type ChallengeType string

// Supported challenge types
const (
	TerraformChallengeType      ChallengeType = "terraform"
	DesignPatternsChallengeType ChallengeType = "design-patterns"
	DataAnalyticsChallengeType  ChallengeType = "data-analytics"
)

// AllChallengeTypes returns every supported challenge type.
func AllChallengeTypes() []ChallengeType {
	return []ChallengeType{TerraformChallengeType, DesignPatternsChallengeType, DataAnalyticsChallengeType}
}

// IsValidChallengeType reports whether t is one of the supported challenge types.
func IsValidChallengeType(t ChallengeType) bool {
	for _, valid := range AllChallengeTypes() {
		if t == valid {
			return true
		}
	}
	return false
}

// ParseChallengeType validates and normalizes a raw string into a ChallengeType.
func ParseChallengeType(s string) (ChallengeType, error) {
	t := ChallengeType(strings.ToLower(strings.TrimSpace(s)))
	if !IsValidChallengeType(t) {
		return "", fmt.Errorf("invalid challenge type %q: must be one of %v", s, AllChallengeTypes())
	}
	return t, nil
}

// CanAddMoreClues reports whether the challenge can still receive another clue
func (c *Challenge) CanAddMoreClues() bool {
	return len(c.Clues) < maxClues
}
