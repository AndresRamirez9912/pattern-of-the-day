package filesystem

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/AndresRamirez9912/pattern-of-the-day/internal/domain"
	"github.com/AndresRamirez9912/pattern-of-the-day/internal/ports"
)

// Type assertion MarkdownWriter must implement the FileWriter port
var _ ports.FileWriter = &MarkdownWriter{}

// MarkdownWriter is a FileWriter implementation that writes generated
// challenge content as markdown files on the local filesystem.
type MarkdownWriter struct{}

// NewMarkdownWriter creates a new instance of MarkdownWriter.
func NewMarkdownWriter() *MarkdownWriter {
	return &MarkdownWriter{}
}

// WriteChallenge writes challenge.md. Target is deliberately omitted so it
// stays a surprise for whoever is solving the challenge.
func (w *MarkdownWriter) WriteChallenge(ctx context.Context, outDir string, challenge *domain.Challenge) error {
	var sb strings.Builder

	fmt.Fprintf(&sb, "# %s\n\n", challenge.Name)
	fmt.Fprintf(&sb, "- **ID:** %d\n", challenge.Id)
	fmt.Fprintf(&sb, "- **Dificultad:** %s\n", challenge.Difficulty)
	fmt.Fprintf(&sb, "- **Tipo:** %s\n\n", challenge.Type)
	fmt.Fprintf(&sb, "## Descripción\n\n%s\n", challenge.Description)

	return writeFile(outDir, "challenge.md", sb.String())
}

// WriteClues writes clues.md with every clue generated so far for the
// challenge, replacing whatever was there before — the file always
// reflects the full, current set of clues, not just the newest one.
func (w *MarkdownWriter) WriteClues(ctx context.Context, outDir string, challenge *domain.Challenge) error {
	var sb strings.Builder

	fmt.Fprintf(&sb, "# Pistas — %s\n", challenge.Name)

	for _, clue := range challenge.Clues {
		fmt.Fprintf(&sb, "\n## Pista %d\n\n%s\n", clue.SequenceOrder, clue.Description)
	}

	return writeFile(outDir, "clues.md", sb.String())
}

// WriteFeedback writes attempt-<N>-feedback.md, where N is the attempt's
// SequenceOrder, so feedback from an earlier attempt on the same challenge
// is never overwritten by a later one.
func (w *MarkdownWriter) WriteFeedback(ctx context.Context, outDir string, attempt *domain.Attempt, challenge *domain.Challenge, feedback *domain.Feedback) error {
	var sb strings.Builder

	fmt.Fprintf(&sb, "# Feedback — Intento %d\n\n", attempt.SequenceOrder)
	fmt.Fprintf(&sb, "- **Challenge:** %s\n", challenge.Name)
	fmt.Fprintf(&sb, "- **Target:** %s\n", challenge.Target)
	fmt.Fprintf(&sb, "- **Score:** %d/100\n\n", feedback.Score)
	fmt.Fprintf(&sb, "## Resumen\n\n%s\n", feedback.Summary)

	if len(feedback.Suggestions) > 0 {
		sb.WriteString("\n## Sugerencias\n\n")
		for _, suggestion := range feedback.Suggestions {
			fmt.Fprintf(&sb, "- %s\n", suggestion)
		}
	}

	filename := fmt.Sprintf("attempt-%d-feedback.md", attempt.SequenceOrder)
	return writeFile(outDir, filename, sb.String())
}

// writeFile creates outDir if needed and writes content to name inside it.
func writeFile(outDir, name, content string) error {
	if err := os.MkdirAll(outDir, 0o755); err != nil {
		return fmt.Errorf("filesystem: creating output directory %q: %w", outDir, err)
	}

	path := filepath.Join(outDir, name)
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		return fmt.Errorf("filesystem: writing %q: %w", path, err)
	}

	return nil
}
