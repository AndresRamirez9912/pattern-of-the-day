package ollama

import (
	"bufio"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

// maxSolutionChars caps how much submitted solution source gets sent to the
// model in one request. Without a cap, pointing solutionPath at a large
// project could produce a prompt far beyond what the model's context
// window can hold, causing Ollama to return an empty or garbled response.
const maxSolutionChars = 100_000

// solutionFileExtensions are the source file extensions read from a
// solution directory — covers the technologies challenges currently target
// (Go code, Terraform HCL).
var solutionFileExtensions = map[string]bool{
	".go": true,
	".tf": true,
}

// skippedSolutionDirs are directories never worth reading as part of a
// submitted solution: dependency trees, tool caches, not the user's own
// code. Hidden directories (starting with ".", e.g. .git, .terraform) are
// always skipped too.
var skippedSolutionDirs = map[string]bool{
	"vendor":       true,
	"node_modules": true,
}

// readSolutionSource reads the code submitted for an attempt. path may
// point either to a single source file or to a directory containing a
// whole project (several files across packages/modules) — the exercise is
// meant to be solved as an actual project, not necessarily a single file.
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

// readSolutionDir walks a directory and concatenates every recognized
// source file it finds, each preceded by a comment naming its path
// relative to the given root so the model can tell which file each snippet
// came from. Generated code is skipped, since it isn't the user's own
// work, and the result is capped at maxSolutionChars.
func readSolutionDir(root string) (string, error) {
	var sb strings.Builder

	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() {
			if path != root && (strings.HasPrefix(d.Name(), ".") || skippedSolutionDirs[d.Name()]) {
				return filepath.SkipDir
			}
			return nil
		}

		if sb.Len() >= maxSolutionChars {
			return nil
		}

		if !solutionFileExtensions[filepath.Ext(path)] || isGeneratedFile(path) {
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
		return "", fmt.Errorf("ollama: no source files found in solution directory %q", root)
	}

	source := sb.String()
	if len(source) > maxSolutionChars {
		source = source[:maxSolutionChars] + "\n\n// [... contenido truncado: el proyecto es más grande de lo que se pudo incluir ...]"
	}

	return source, nil
}

// isGeneratedFile reports whether a file starts with Go's standard
// "generated code" marker (https://go.dev/s/generatedcode), so tool output
// (e.g. sqlc, protobuf) isn't mistaken for the user's own solution.
func isGeneratedFile(path string) bool {
	file, err := os.Open(path)
	if err != nil {
		return false
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for i := 0; i < 5 && scanner.Scan(); i++ {
		line := scanner.Text()
		if strings.HasPrefix(line, "// Code generated ") && strings.HasSuffix(line, "DO NOT EDIT.") {
			return true
		}
	}

	if scanner.Err() != nil {
		return false
	}

	return false
}
