package ollama

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

// skippedSolutionDirs are directories never worth reading as part of a
// submitted solution (dependency trees and VCS metadata, not the user's code).
var skippedSolutionDirs = map[string]bool{
	"vendor": true,
	".git":   true,
}

// readSolutionSource reads the code submitted for an attempt. path may point
// either to a single source file or to a directory containing a whole Go
// project (several .go files across packages) — the exercise is meant to be
// solved as an actual project, not necessarily a single file.
func readSolutionSource(path string) (string, error) {
	info, err := os.Stat(path)
	if err != nil {
		return "", fmt.Errorf("ollama: reading solution path: %w", err)
	}

	if !info.IsDir() {
		content, err := os.ReadFile(path)
		if err != nil {
			return "", fmt.Errorf("ollama: reading solution file: %w", err)
		}
		return string(content), nil
	}

	return readSolutionDir(path)
}

// readSolutionDir walks a directory and concatenates every .go file it finds,
// each preceded by a comment naming its path relative to the given root so
// the model can tell which file each snippet came from.
func readSolutionDir(root string) (string, error) {
	var sb strings.Builder

	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() {
			if skippedSolutionDirs[d.Name()] {
				return filepath.SkipDir
			}
			return nil
		}

		content, err := os.ReadFile(path)
		if err != nil {
			return fmt.Errorf("ollama: reading solution file %q: %w", path, err)
		}

		relPath, err := filepath.Rel(root, path)
		if err != nil {
			relPath = path
		}

		fmt.Fprintf(&sb, "// file: %s\n%s\n\n", relPath, content)
		return nil
	})
	if err != nil {
		return "", fmt.Errorf("ollama: walking solution directory: %w", err)
	}

	if sb.Len() == 0 {
		return "", fmt.Errorf("ollama: no .go files found in solution directory %q", root)
	}

	return sb.String(), nil
}
