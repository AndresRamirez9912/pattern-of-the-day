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
completado. Escribe attempt-<N>-feedback.md dentro de --out, donde N es el
número de intento de este challenge (no se sobreescribe el feedback de
intentos anteriores).

Argumentos (obligatorios, en este orden):
  attempt-id      ID numérico del intento que se está evaluando
  solution-path   Ruta a la solución: un archivo o una carpeta con un
                  proyecto completo

Uso:
  patternd use-cases feedback create <attempt-id> <solution-path>

Ejemplos:
  patternd use-cases feedback create 7 ./mi-solucion
  patternd use-cases feedback create 7 ./mi-solucion/main.go --out ./mis-retos`,
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			attemptId, err := ParseStringId(args[0])
			if err != nil {
				return fmt.Errorf("attempt-id inválido %q: %w", args[0], err)
			}
			solutionPath := args[1]
			outDir := mustGetString(cmd, "out")

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

			app.Logger.Info("evaluando solución", "attempt_id", attemptId, "solution_path", solutionPath)

			generatedFeedback, err := app.Services.Feedback.CreateFeedback.Execute(app.Ctx, attempt, challenge, solutionPath, outDir)
			if err != nil {
				return err
			}

			app.Logger.Info("feedback creado", "id", generatedFeedback.Id, "score", generatedFeedback.Score, "summary", generatedFeedback.Summary)

			for _, suggestion := range generatedFeedback.Suggestions {
				app.Logger.Info("sugerencia", "text", suggestion)
			}

			app.Logger.Info("intento actualizado", "id", attempt.Id, "status", attempt.Status)

			return nil
		},
	}

	cmd.Flags().String("out", ".", "Output directory for attempt-<N>-feedback.md")

	return cmd
}
