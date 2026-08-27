package cmd

import (
	"github.com/AndresRamirez9912/pattern-of-the-day/cmd/pattern-of-the-day/cmd/database"
	usecases "github.com/AndresRamirez9912/pattern-of-the-day/cmd/pattern-of-the-day/cmd/use_cases"
	"github.com/spf13/cobra"
)

// NewRootCmd creates a new root command for the CLI application
func NewRootCmd() *cobra.Command {
	// Create root command
	rootCmd := &cobra.Command{
		Use:   "patternd",
		Short: "Pattern of the Day CLI",
		Long:  `Pattern of the Day CLI is a command line interface that generates challenges to practice design patterns`,
		Run: func(cmd *cobra.Command, args []string) {
			// This is the root command so execute the help command
			err := cmd.Help()
			if err != nil {
				panic(err)
			}
		},
	}

	// Add sub-commands
	rootCmd.AddCommand(NewVersionInfoCmd())
	rootCmd.AddCommand(database.NewDatabaseRootCmd())
	rootCmd.AddCommand(usecases.NewUseCasesRootCmd())

	return rootCmd
}
