package usecases

import (
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
		Run: func(cmd *cobra.Command, args []string) {
			err := cmd.Help()
			if err != nil {
				panic(err)
			}
		},
	}

	// Add subcommands for specific challenge-related use cases here
	cmd.AddCommand(CreateChallengeUseCaseCmd())

	return cmd
}

// CreateChallengeUseCaseCmd creates the command for executing the create challenge use case.
func CreateChallengeUseCaseCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create",
		Short: "Execute the create challenge use case",
		// Create a long description adding the information requested based on the position
		Long: `This command allows you to create a new challenge by specifying the topic, 
difficulty, pattern, type, and user ID as command-line arguments.

Order:
pattern-of-the-day challenge create [topic] [difficulty] [pattern] [type] [user ID]

Where:
[topic]      The topic of the challenge
[difficulty] The difficulty level of the challenge (e.g., easy, medium, hard)
[pattern]    The pattern associated with the challenge (e,g Builder, Singleton...)
[type]       The type of the challenge (e.g., creational, structural, behavioral)
[user ID]    The ID of the user creating the challenge

Example:
pattern-of-the-day challenge create "Design Patterns" 1 "Singleton" "creational" "user123"

Note:
Ensure that all required arguments are provided in the correct order.
		`,
		Run: func(cmd *cobra.Command, args []string) {
			// Initialize the application
			app := InitApp()
			defer app.GracefulShutdown()

			// Read topic, difficulty, pattern, type, and user ID from command positions
			topic := args[0]
			difficulty := args[1]
			pattern := args[2]
			challengeType := args[3]
			userId := args[4]

			// Parse the string userId to int64
			parsedUserId, err := ParseStringId(userId)
			if err != nil {
				panic(err)
			}

			// Implement the use case for creating a user
			_, err = app.Services.Challenge.CreateChallenge.Execute(app.Ctx, ports.ChallengeGenerationRequest{
				Topic:      topic,
				Difficulty: domain.Difficulty(difficulty),
				Pattern:    domain.Pattern(pattern),
				Type:       domain.ChallengeType(challengeType),
				UserId:     parsedUserId,
			})
			if err != nil {
				panic(err)
			}

			// Print a success message or perform any other necessary actions
			println("Challenge created successfully")
		},
	}

	return cmd
}
