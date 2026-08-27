package domain

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

// NewChallenge creates a new challenge with the given parameters
func NewChallenge(id int64, name, description string, difficulty Difficulty, challengeType ChallengeType, target string, userId int64) *Challenge {
	return &Challenge{
		Id:          id,
		Name:        name,
		Description: description,
		Difficulty:  difficulty,
		Type:        challengeType,
		Target:      target,
		Clues:       []Clue{},
		UserId:      userId,
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

// CanAddMoreClues reports whether the challenge can still receive another clue
func (c *Challenge) CanAddMoreClues() bool {
	return len(c.Clues) < maxClues
}
