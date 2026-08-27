package main

import (
	"os"

	"github.com/AndresRamirez9912/pattern-of-the-day/cmd/pattern-of-the-day/cmd"
)

func main() {
	// Initialize root command
	rootCmd := cmd.NewRootCmd()

	// Execute root command. Cobra already prints the error so on failure
	// we just need to exit with a non-zero status instead of panicking
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
