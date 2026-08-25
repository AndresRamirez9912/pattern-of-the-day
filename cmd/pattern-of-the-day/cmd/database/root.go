package database

import "github.com/spf13/cobra"

// dbPath is the path to the SQLite database file, shared by all database subcommands.
var dbPath string

// NewDatabaseRootCmd creates a new root command for the database subcommands.
func NewDatabaseRootCmd() *cobra.Command {
	// Create a new root command for the database subcommands
	cmd := &cobra.Command{
		Use:   "database",
		Short: "database migration root command",
		Long:  "Database migration root commands for managing schema changes",
		Run: func(cmd *cobra.Command, args []string) {
			// This is the root command so execute the help command
			err := cmd.Help()
			if err != nil {
				panic(err)
			}
		},
	}

	cmd.PersistentFlags().StringVar(&dbPath, "db", "./app.db", "path to the SQLite database file")

	// Add subcommands for database operations
	cmd.AddCommand(NewUpCmd())
	cmd.AddCommand(NewDownCmd())
	cmd.AddCommand(NewStatusCmd())
	cmd.AddCommand(NewResetCmd())

	return cmd
}
