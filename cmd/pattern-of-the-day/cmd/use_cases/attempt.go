package usecases

import (
	"fmt"

	"github.com/spf13/cobra"
)

// NewAttemptUseCaseCmd creates the command for executing attempt-related use cases.
func NewAttemptUseCaseCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "attempt",
		Short: "Execute attempt-related use cases",
		Long:  "This command allows you to execute attempt-related use cases within the application.",
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
	}

	cmd.AddCommand(CreateAttemptUseCaseCmd())

	return cmd
}

// CreateAttemptUseCaseCmd creates the command for executing the create attempt use case.
func CreateAttemptUseCaseCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create <challenge-id>",
		Short: "Start a new attempt for a challenge",
		Long: `Crea un nuevo intento (attempt) para un reto existente. Solo se puede crear
un intento nuevo si todos los intentos previos de ese reto ya están cerrados
(completed o failed) — un reto no puede tener dos intentos abiertos a la vez.

Nota: crear un challenge con "challenge create" ya crea automáticamente su
primer intento. Usa este comando solo para reintentar un reto después de
cerrar el intento anterior (generando su feedback).

Uso:
  patternd use-cases attempt create <challenge-id>

Ejemplo:
  patternd use-cases attempt create 4`,
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			challengeId, err := ParseStringId(args[0])
			if err != nil {
				return fmt.Errorf("challenge-id inválido %q: %w", args[0], err)
			}

			app := InitApp()
			defer app.GracefulShutdown()

			attempt, err := app.Services.Attempts.CreateAttempt.Execute(app.Ctx, challengeId)
			if err != nil {
				return err
			}

			fmt.Printf("Intento creado (id=%d, challenge_id=%d, status=%s)\n", attempt.Id, challengeId, attempt.Status)

			return nil
		},
	}

	return cmd
}
