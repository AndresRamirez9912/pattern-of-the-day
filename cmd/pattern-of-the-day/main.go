package main

import "github.com/AndresRamirez9912/pattern-of-the-day/cmd/pattern-of-the-day/cmd"

func main() {
	// Initialize root command
	rootCmd := cmd.NewRootCmd()

	// Execute root command
	err := rootCmd.Execute()
	if err != nil {
		panic(err)
	}
}
