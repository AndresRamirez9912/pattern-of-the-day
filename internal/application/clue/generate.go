package clue

import (
	"context"

	"github.com/AndresRamirez9912/pattern-of-the-day/internal/app"
	"github.com/AndresRamirez9912/pattern-of-the-day/internal/domain"
	"github.com/AndresRamirez9912/pattern-of-the-day/internal/ports"
)

// GenerateClueUseCase is responsible for generating a new clue using the
// LLM provider and saving it to the ClueRepository
type GenerateClueUseCase struct {
	Logger         app.Logger
	LLMProvider    ports.LLMProvider
	ClueRepository ports.ClueRepository
}

// NewGenerateClueUseCase creates a new instance of GenerateClueUseCase with the provided LLMProvider and ClueRepository
func NewGenerateClueUseCase(logger app.Logger, llmProvider ports.LLMProvider, clueRepository ports.ClueRepository) *GenerateClueUseCase {
	return &GenerateClueUseCase{
		Logger:         logger,
		LLMProvider:    llmProvider,
		ClueRepository: clueRepository,
	}
}

// Execute generates a new clue using the LLM provider and saves it to the ClueRepository
func (g *GenerateClueUseCase) Execute(ctx context.Context, challenge *domain.Challenge) (*domain.Clue, error) {
	// Generate a new clue using the LLM provider
	clue, err := g.LLMProvider.GenerateClue(ctx, challenge)
	if err != nil {
		g.Logger.Error("error generating clue through the LLM provider", "error", err.Error())

		return nil, err
	}

	// Save the generated clue to the ClueRepository
	err = g.ClueRepository.SaveClue(ctx, challenge.Id, clue)
	if err != nil {
		g.Logger.Error("error saving clue to the repository", "error", err.Error())
		return nil, err
	}

	return clue, nil
}
