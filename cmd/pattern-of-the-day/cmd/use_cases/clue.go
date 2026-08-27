package usecases

import (
	"fmt"

	"github.com/spf13/cobra"
)

// NewClueUseCaseCmd creates the command for executing clue-related use cases.
func NewClueUseCaseCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "clue",
		Short: "Execute clue-related use cases",
		Long:  "This command allows you to execute clue-related use cases within the application.",
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
	}

	cmd.AddCommand(CreateClueUseCaseCmd())

	return cmd
}

// CreateClueUseCaseCmd creates the command for executing the create clue use case.
func CreateClueUseCaseCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create <challenge-id>",
		Short: "Generate the next clue for a challenge",
		Long: `Genera una nueva pista (clue) para un reto existente, usando el modelo de
lenguaje configurado. Un reto admite un máximo de 3 pistas; cada pista nueva
es más específica que las anteriores.

Uso:
  patternd use-cases clue create <challenge-id>

Ejemplo:
  patternd use-cases clue create 4`,
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			challengeId, err := ParseStringId(args[0])
			if err != nil {
				return fmt.Errorf("challenge-id inválido %q: %w", args[0], err)
			}

			app := InitApp()
			defer app.GracefulShutdown()

			challenge, err := app.Services.Challenge.GetChallenge.Execute(app.Ctx, challengeId)
			if err != nil {
				return err
			}

			fmt.Printf("Generando pista #%d para el challenge %d...\n", len(challenge.Clues)+1, challengeId)

			clue, err := app.Services.Clue.CreateClue.Execute(app.Ctx, challenge)
			if err != nil {
				return err
			}

			fmt.Printf("Pista creada (id=%d, #%d): %s\n", clue.Id, clue.SequenceOrder, clue.Description)

			return nil
		},
	}

	return cmd
}
