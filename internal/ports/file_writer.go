package ports

import (
	"context"

	"github.com/AndresRamirez9912/pattern-of-the-day/internal/domain"
)

// FileWriter defines the interface for writing generated challenge content
// to durable files (e.g. markdown) that a user can read outside the CLI.
type FileWriter interface {
	// WriteChallenge writes the challenge's title, description and details
	// to outDir.
	WriteChallenge(ctx context.Context, outDir string, challenge *domain.Challenge) error
	// WriteClues writes every clue generated so far for the challenge to
	// outDir, replacing whatever was written there before.
	WriteClues(ctx context.Context, outDir string, challenge *domain.Challenge) error
	// WriteFeedback writes the feedback for one attempt to outDir. Each
	// attempt gets its own file (keyed by the attempt's SequenceOrder)
	WriteFeedback(ctx context.Context, outDir string, attempt *domain.Attempt, challenge *domain.Challenge, feedback *domain.Feedback) error
}
