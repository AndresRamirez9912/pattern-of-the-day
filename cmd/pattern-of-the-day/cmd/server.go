package cmd

import (
	"context"

	"github.com/AndresRamirez9912/pattern-of-the-day/internal/adapters/inbound/rest"
	"github.com/AndresRamirez9912/pattern-of-the-day/internal/app"
	"github.com/AndresRamirez9912/pattern-of-the-day/internal/app/config"
	"github.com/spf13/cobra"
)

// StartServerCmd creates a new Cobra command to start the server.
func StartServerCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "start-server",
		Short: "Starts the server",
		Long:  "Starts the server and listens for incoming requests",
		Run: func(cmd *cobra.Command, args []string) {
			// Load the configuration and initialize the application
			cfg, err := config.LoadConfig(".")
			if err != nil {
				panic(err)
			}

			// Initialize the application
			app, err := app.NewApp(cfg, context.Background())
			if err != nil {
				panic(err)
			}

			// Create the server instance with app information
			server := rest.NewServer(
				*cfg,
				app.Logger.WithSection("section", "rest-server"),
				app.Services,
			)

			// Start the server
			err = server.Start()
			if err != nil {
				panic(err)
			}

			// Gracefully handle server shutdown
			defer func() {
				err := server.Shutdown(app.Ctx)
				if err != nil {
					panic(err)
				}
			}()

		},
	}
}
