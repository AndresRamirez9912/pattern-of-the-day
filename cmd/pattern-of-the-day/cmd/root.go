package cmd

import (
	"github.com/spf13/cobra"
)

func NewRootCmd() *cobra.Command {
	// Create root command
	rootCmd := &cobra.Command{
		Use:   "patternd",
		Short: "Pattern of the Day CLI",
		Long:  `Patern of the Day CLI is a command line interface that generates challenges to practice design patterns`,
	}

	// Add sub-commands
	rootCmd.AddCommand(NewVersionInfoCmd())

	return rootCmd
}
