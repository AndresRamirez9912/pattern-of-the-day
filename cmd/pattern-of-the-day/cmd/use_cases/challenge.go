package usecases

import (
	"fmt"

	"github.com/AndresRamirez9912/pattern-of-the-day/internal/domain"
	"github.com/AndresRamirez9912/pattern-of-the-day/internal/ports"
	"github.com/spf13/cobra"
)

// NewChallengeUseCaseCmd creates the command for executing challenge-related use cases.
func NewChallengeUseCaseCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "challenge",
		Short: "Execute challenge-related use cases",
		Long:  "This command allows you to execute challenge-related use cases within the application.",
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
	}

	// Add subcommands for specific challenge-related use cases here
	cmd.AddCommand(CreateChallengeUseCaseCmd())

	return cmd
}

// CreateChallengeUseCaseCmd creates the command for executing the create challenge use case.
func CreateChallengeUseCaseCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create <username> <difficulty>",
		Short: "Generate a new challenge and its initial attempt",
		Long: `Genera un nuevo reto usando el modelo de lenguaje configurado, y crea
automáticamente el primer intento (attempt) pendiente para ese reto.
Escribe los detalles del reto en challenge.md dentro de --out.

Argumentos (obligatorios, en este orden):
  username    Nombre de usuario dueño del reto
  difficulty  Dificultad del reto: easy, medium o hard
  type        Tipo de reto: terraform, design-patterns o data-analytics

Todo lo demás es opcional vía flags. Si no usas --topic o --target,
el caso de uso elige un valor aleatorio para cada uno que omitas, así puedes
generar retos variados sin tener que pensarlos. --target se elige acorde al
--type resultante (por ejemplo, un patrón de diseño si el tipo es design-patterns).

Ejemplos:
  patternd use-cases challenge create andres medium
  patternd use-cases challenge create andres hard --type design-patterns --topic "sistema de pagos" --out ./mis-retos`,
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			username := args[0]

			// Extract difficulty from args
			difficulty, err := domain.ParseDifficulty(args[1])
			if err != nil {
				return err
			}

			// Extract challenge type from args
			challengeType := domain.ChallengeType(args[2])
			if !domain.IsValidChallengeType(challengeType) {
				return fmt.Errorf("invalid challenge type %q", challengeType)
			}

			// Extract output directory from flags
			outDir := mustGetString(cmd, "out")

			// Initialize the application context and services
			app := InitApp()
			defer app.GracefulShutdown()

			// Execute the create challenge use case
			createdChallenge, createdAttempt, err := app.Services.Challenge.CreateChallenge.Execute(
				app.Ctx,
				username,
				ports.ChallengeGenerationRequest{
					Topic:      mustGetString(cmd, "topic"),
					Difficulty: difficulty,
					Target:     mustGetString(cmd, "target"),
					Type:       domain.ChallengeType(mustGetString(cmd, "type")),
				},
				outDir,
			)
			if err != nil {
				return err
			}

			// Log the created challenge and attempt
			app.Logger.Info("challenge creado", "id", createdChallenge.Id, "name", createdChallenge.Name)
			app.Logger.Info("intento inicial creado", "id", createdAttempt.Id, "status", createdAttempt.Status)

			return nil
		},
	}

	cmd.Flags().String("topic", "", "Topic of the challenge (random if omitted)")
	cmd.Flags().String("target", "", "Specific subject to evaluate within the type, e.g. facade (random if omitted)")
	cmd.Flags().String("out", ".", "Output directory for challenge.md")

	return cmd
}

// mustGetString reads a string flag, ignoring the (never-populated) error
// cobra returns for a flag that was registered on the same command.
func mustGetString(cmd *cobra.Command, name string) string {
	value, _ := cmd.Flags().GetString(name)
	return value
}
