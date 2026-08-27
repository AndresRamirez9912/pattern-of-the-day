package usecases

import (
	"fmt"

	"github.com/spf13/cobra"
)

// NewUsersUseCaseCmd creates the command for executing user-related use cases.
func NewUsersUseCaseCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "users",
		Short: "Execute user-related use cases",
		Long:  "This command allows you to execute user-related use cases within the application.",
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
	}

	// Add subcommands for specific user-related use cases here
	cmd.AddCommand(CreateUserUseCaseCmd())

	return cmd
}

// CreateUserUseCaseCmd creates the command for executing the create user use case.
func CreateUserUseCaseCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create <username> <email>",
		Short: "Create a new user",
		Long: `Crea un nuevo usuario. El ID que devuelve es el que después vas a necesitar
para crear challenges a nombre de ese usuario.

Argumentos (obligatorios, en este orden):
  username  Nombre de usuario
  email     Correo del usuario

Uso:
  patternd use-cases users create <username> <email>

Ejemplo:
  patternd use-cases users create andres andres@example.com`,
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			username := args[0]
			email := args[1]

			// Initialize the application
			app := InitApp()
			defer app.GracefulShutdown()

			// Execute the use case for creating a user
			createdUser, err := app.Services.User.CreateUser.Execute(app.Ctx, username, email)
			if err != nil {
				return err
			}

			fmt.Printf("Usuario creado (id=%d): %s <%s>\n", createdUser.Id, createdUser.UserName, createdUser.Email)

			return nil
		},
	}

	return cmd
}
