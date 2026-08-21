package domain

// Challenge defines the properties of a challenge
type Challenge struct {
	Id          string
	Name        string
	Description string
	Dificulty   Difficulty
	Type        ChallengeType
	Pattern     Pattern
	Status      ChallengeStatus
	Clues       []Clue
}

// NewChallenge creates a new challenge with the given parameters
func NewChallenge(id, name, description string, difficulty Difficulty, challengeType ChallengeType, pattern Pattern) *Challenge {
	return &Challenge{
		Id:          id,
		Name:        name,
		Description: description,
		Dificulty:   difficulty,
		Type:        challengeType,
		Pattern:     pattern,
		Status:      ChallengeStatusPending,
		Clues:       []Clue{},
	}
}

// Difficulty is a custom type to define the supported challenge difficulty
type Difficulty int

// Supported challenge difficulty
const (
	Easy = iota
	Medimum
	Hard
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

// Status is a custom type to define the status of a challenge
type ChallengeStatus string

// List of supported challenge status
const (
	ChallengeStatusPending   ChallengeStatus = "pending"
	ChallengeStatusCompleted ChallengeStatus = "completed"
)

// IsFinished checks if the challenge is finished
func (c Challenge) IsFinished() bool {
	return c.Status == ChallengeStatusCompleted
}

// Complete marks the challenge as completed
func (c *Challenge) Complete() {
	c.Status = ChallengeStatusCompleted
}

// AddClue adds a clue to the challenge
func (c *Challenge) AddClue(clue Clue) error {
	// Only 3 clues can be generated for a challenge
	if len(c.Clues) >= 3 {
		return ErrMaxCluesReached
	}

	c.Clues = append(c.Clues, clue)
	return nil
}
