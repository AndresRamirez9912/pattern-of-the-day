package usecases

import (
	"github.com/spf13/cobra"
)

// NewUsersUseCaseCmd creates the command for executing user-related use cases.
func NewUsersUseCaseCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "users",
		Short: "Execute user-related use cases",
		Long:  "This command allows you to execute user-related use cases within the application.",
		Run: func(cmd *cobra.Command, args []string) {
			err := cmd.Help()
			if err != nil {
				panic(err)
			}
		},
	}

	// Add subcommands for specific user-related use cases here
	cmd.AddCommand(CreateUserUseCaseCmd())

	return cmd
}

// CreateUserUseCaseCmd creates the command for executing the create user use case.
// This commands fetches the
func CreateUserUseCaseCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create",
		Short: "Execute the create user use case",
		Long:  "This command allows you to execute the create user use case within the application.",
		Run: func(cmd *cobra.Command, args []string) {
			// Initialize the application
			app := InitApp()
			defer app.GracefulShutdown()

			// Read username and email from command flags
			username, _ := cmd.Flags().GetString("username")
			email, _ := cmd.Flags().GetString("email")

			// Implement the use case for creating a user
			_, err := app.Services.User.CreateUser.Execute(app.Ctx, username, email)
			if err != nil {
				panic(err)
			}

			// Print a success message or perform any other necessary actions
			println("User created successfully")
		},
	}
	// Add flags for the create user use case
	cmd.Flags().String("username", "", "Username of the new user")
	cmd.Flags().String("email", "", "Email of the new user")

	return cmd
}
