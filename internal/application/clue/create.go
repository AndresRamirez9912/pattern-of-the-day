package clue

import (
	"context"
	"fmt"

	"github.com/AndresRamirez9912/pattern-of-the-day/internal/domain"
	"github.com/AndresRamirez9912/pattern-of-the-day/internal/ports"
)

// CreateClueUseCase is responsible for generating a new clue using the
// LLM provider and saving it to the ClueRepository
type CreateClueUseCase struct {
	Logger         ports.Logger
	LLMProvider    ports.LLMProvider
	ClueRepository ports.ClueRepository
	FileWriter     ports.FileWriter
}

// NewCreateClueUseCase creates a new instance of CreateClueUseCase with the provided dependencies
func NewCreateClueUseCase(logger ports.Logger, llmProvider ports.LLMProvider, clueRepository ports.ClueRepository, fileWriter ports.FileWriter) *CreateClueUseCase {
	return &CreateClueUseCase{
		Logger:         logger,
		LLMProvider:    llmProvider,
		ClueRepository: clueRepository,
		FileWriter:     fileWriter,
	}
}

// Execute generates a new clue using the LLM provider, saves it to the ClueRepository,
// and rewrites clues.md in outDir with the challenge's full, current list of clues.
func (g *CreateClueUseCase) Execute(ctx context.Context, challenge *domain.Challenge, outDir string) (*domain.Clue, error) {
	if !challenge.CanAddMoreClues() {
		return nil, fmt.Errorf("challenge %d already has the maximum number of clues", challenge.Id)
	}

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

	// Reflect the new clue on the in-memory challenge so clues.md includes it
	challenge.Clues = append(challenge.Clues, *clue)

	// Write the updated list of clues to the clues.md file in the output directory
	err = g.FileWriter.WriteClues(ctx, outDir, challenge)
	if err != nil {
		g.Logger.Error("error writing clues file", "error", err.Error())
		return nil, err
	}

	return clue, nil
}
