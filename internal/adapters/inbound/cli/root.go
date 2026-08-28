package cli

import (
	"strconv"

	"github.com/spf13/cobra"
)

// NewUseCasesRootCmd creates the root command for executing cases.
func NewUseCasesRootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "use-cases",
		Short: "Execute the specific use case",
		Long:  "This command allows you to execute a specific use case within the application.",
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
	}

	// Add subcommands for specific use cases here
	rootCmd.AddCommand(NewUsersUseCaseCmd())
	rootCmd.AddCommand(NewChallengeUseCaseCmd())
	rootCmd.AddCommand(NewClueUseCaseCmd())
	rootCmd.AddCommand(NewAttemptUseCaseCmd())
	rootCmd.AddCommand(NewFeedbackUseCaseCmd())

	return rootCmd
}

// ParseStringId parses a string ID into an int64. It returns an error if the conversion fails.
func ParseStringId(id string) (int64, error) {
	return strconv.ParseInt(id, 10, 64)
}
