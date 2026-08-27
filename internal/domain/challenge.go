package domain

// Challenge defines the properties of a challenge
type Challenge struct {
	Id          int64
	Name        string
	Description string
	Dificulty   Difficulty
	Type        ChallengeType
	Pattern     Pattern
	Clues       []Clue
	UserId      int64
}

// NewChallenge creates a new challenge with the given parameters
func NewChallenge(id int64, name, description string, difficulty Difficulty, challengeType ChallengeType, pattern Pattern, userId int64) *Challenge {
	return &Challenge{
		Id:          id,
		Name:        name,
		Description: description,
		Dificulty:   difficulty,
		Type:        challengeType,
		Pattern:     pattern,
		Clues:       []Clue{},
		UserId:      userId,
	}
}

// Difficulty is a custom type to define the supported challenge difficulty
type Difficulty string

// Supported challenge difficulty
const (
	ChallengeDifficultyEasy   = "easy"
	ChallengeDifficultyMedium = "medium"
	ChallengeDifficultyHard   = "hard"
)

// ChallengeType is a custom type to define which challenge types are supported
type ChallengeType string

// Supported challenge types
const (
	CreationalChallengeType  ChallengeType = "creational"
	StructuralChallengeType  ChallengeType = "structural"
	BehaviouralChallengeType ChallengeType = "behavioural"
)

// Pattern is a custom type to define which patterns are supported
type Pattern string

// List of supported patterns
const (
	FactoryMethodPattern         Pattern = "factory-method"
	AbstractFactoryPattern       Pattern = "abstract-factory"
	BuilderPattern               Pattern = "builder"
	PrototypePattern             Pattern = "prototype"
	SingletonPattern             Pattern = "singleton"
	AdapterPattern               Pattern = "adapter"
	BridgePattern                Pattern = "bridge"
	CompositePattern             Pattern = "composite"
	DecoratorPattern             Pattern = "decorator"
	FacadePattern                Pattern = "facade"
	FlyweightPattern             Pattern = "flyweight"
	ProxyPattern                 Pattern = "proxy"
	ChainOfResponsibilityPattern Pattern = "chain-of-responsibility"
	CommandPattern               Pattern = "command"
	IteratorPattern              Pattern = "iterator"
	MediatorPattern              Pattern = "mediator"
	MementoPattern               Pattern = "memento"
	ObserverPattern              Pattern = "observer"
	StatePattern                 Pattern = "state"
	StrategyPattern              Pattern = "strategy"
	TemplateMethodPattern        Pattern = "template-method"
	VisitorPattern               Pattern = "visitor"
)

// AddClue adds a clue to the challenge
func (c *Challenge) AddClue(clue Clue) error {
	// Only 3 clues can be generated for a challenge
	if len(c.Clues) >= 3 {
		return ErrMaxCluesReached
	}

	c.Clues = append(c.Clues, clue)
	return nil
}
