package cmd

import (
	"context"
	"os/signal"
	"syscall"
	"time"

	"github.com/AndresRamirez9912/pattern-of-the-day/internal/adapters/inbound/rest"
	"github.com/AndresRamirez9912/pattern-of-the-day/internal/app"
	"github.com/AndresRamirez9912/pattern-of-the-day/internal/app/config"
	"github.com/spf13/cobra"
)

// defaultShutdownTimeout is used when the config doesn't set app.shutdown_timeout.
const defaultShutdownTimeout = 10 * time.Second

// StartServerCmd creates a new Cobra command to start the server.
func StartServerCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "start-server",
		Short: "Starts the server",
		Long:  "Starts the server and listens for incoming requests until an interrupt or termination signal is received.",
		RunE: func(cmd *cobra.Command, args []string) error {
			// ctx is canceled the moment the process receives SIGINT/SIGTERM,
			// which is what triggers the shutdown sequence below.
			ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
			defer stop()

			// Load the configuration and initialize the application
			cfg, err := config.LoadConfig(".")
			if err != nil {
				return err
			}

			application, err := app.NewApp(cfg, ctx)
			if err != nil {
				return err
			}

			// Create the server instance with app information
			server := rest.NewServer(
				*cfg,
				application.Logger.WithSection("section", "rest-server"),
				application.Services,
			)

			// Start the server in the background so this goroutine is free to
			// wait for either a shutdown signal or the server failing to start.
			serverErr := make(chan error, 1)
			go func() {
				serverErr <- server.Start()
			}()

			select {
			case <-ctx.Done():
				// Signal received: fall through to the shutdown sequence below.
			case err := <-serverErr:
				if err != nil {
					return err
				}
			}

			// Stop accepting new connections and let in-flight requests
			// finish, then close the DB — in that order, since closing the
			// DB first could break requests still being served.
			shutdownTimeout := cfg.App.ShutdownTimeout
			if shutdownTimeout <= 0 {
				shutdownTimeout = defaultShutdownTimeout
			}

			shutdownCtx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
			defer cancel()

			if err := server.Shutdown(shutdownCtx); err != nil {
				return err
			}

			return application.GracefulShutdown()
		},
	}
}
