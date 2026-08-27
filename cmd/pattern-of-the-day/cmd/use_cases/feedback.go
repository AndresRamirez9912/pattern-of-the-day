package usecases

import (
	"fmt"

	"github.com/spf13/cobra"
)

// NewFeedbackUseCaseCmd creates the command for executing feedback-related use cases.
func NewFeedbackUseCaseCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "feedback",
		Short: "Execute feedback-related use cases",
		Long:  "This command allows you to execute feedback-related use cases within the application.",
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
	}

	cmd.AddCommand(CreateFeedbackUseCaseCmd())

	return cmd
}

// CreateFeedbackUseCaseCmd creates the command for executing the create feedback use case.
func CreateFeedbackUseCaseCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create <attempt-id> <solution-path>",
		Short: "Evaluate a submitted solution and close out an attempt",
		Long: `Evalúa la solución entregada para un intento (attempt) usando el modelo de
lenguaje configurado, guarda el feedback generado, y marca el intento como
completado.

Argumentos (obligatorios, en este orden):
  attempt-id      ID numérico del intento que se está evaluando
  solution-path   Ruta a la solución: un archivo .go o una carpeta con un
                  proyecto Go completo

Uso:
  patternd use-cases feedback create <attempt-id> <solution-path>

Ejemplos:
  patternd use-cases feedback create 7 ./mi-solucion
  patternd use-cases feedback create 7 ./mi-solucion/main.go`,
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			attemptId, err := ParseStringId(args[0])
			if err != nil {
				return fmt.Errorf("attempt-id inválido %q: %w", args[0], err)
			}
			solutionPath := args[1]

			app := InitApp()
			defer app.GracefulShutdown()

			attempt, err := app.Services.Attempts.GetAttempt.Execute(app.Ctx, attemptId)
			if err != nil {
				return err
			}
			if attempt.ChallengeId == nil {
				return fmt.Errorf("el intento %d no tiene un challenge asociado", attemptId)
			}

			challenge, err := app.Services.Challenge.GetChallenge.Execute(app.Ctx, *attempt.ChallengeId)
			if err != nil {
				return err
			}

			fmt.Printf("Evaluando solución en %q para el intento %d...\n", solutionPath, attemptId)

			generatedFeedback, err := app.Services.Feedback.CreateFeedback.Execute(app.Ctx, attempt, challenge, solutionPath)
			if err != nil {
				return err
			}

			fmt.Printf("Feedback creado (id=%d, score=%d): %s\n", generatedFeedback.Id, generatedFeedback.Score, generatedFeedback.Summary)
			if len(generatedFeedback.Suggestions) > 0 {
				fmt.Println("Sugerencias:")
				for _, suggestion := range generatedFeedback.Suggestions {
					fmt.Printf("  - %s\n", suggestion)
				}
			}
			fmt.Printf("Intento %d marcado como %s\n", attempt.Id, attempt.Status)

			return nil
		},
	}

	return cmd
}
