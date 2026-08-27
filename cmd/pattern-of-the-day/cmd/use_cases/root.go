package usecases

import (
	"context"
	"strconv"

	"github.com/AndresRamirez9912/pattern-of-the-day/internal/app"
	"github.com/AndresRamirez9912/pattern-of-the-day/internal/app/config"
	"github.com/spf13/cobra"
)

// NewUseCasesRootCmd creates the root command for executing cases.
func NewUseCasesRootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "use-cases",
		Short: "Execute the specific use case",
		Long:  "This command allows you to execute a specific use case within the application.",
		Run: func(cmd *cobra.Command, args []string) {
			err := cmd.Help()
			if err != nil {
				panic(err)
			}
		},
	}

	// Add subcommands for specific use cases here
	rootCmd.AddCommand(NewUsersUseCaseCmd())
	rootCmd.AddCommand(NewChallengeUseCaseCmd())

	return rootCmd
}

// InitApp initializes the application with the configuration and context.
func InitApp() *app.App {
	cfg, err := config.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	app, err := app.NewApp(cfg, context.Background())
	if err != nil {
		panic(err)
	}

	return app
}

// ParseStringId parses a string ID into an int64. It returns an error if the conversion fails.
func ParseStringId(id string) (int64, error) {
	return strconv.ParseInt(id, 10, 64)
}
