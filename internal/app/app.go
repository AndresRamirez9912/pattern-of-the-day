package app

import (
	"context"
	"database/sql"

	"github.com/AndresRamirez9912/pattern-of-the-day/internal/app/config"
)

// App represents the application and its dependencies
type App struct {
	// cfg contains the application configuration
	cfg *config.AppConfig
	// Ctx is the context for managing application lifecycle
	Ctx context.Context
	// db is the database connection
	db *sql.DB
	// Services contains the services layer of the application
	Services *Services
	// Logger is the application logger
	Logger *Logger
}

// NewApp creates a new instance of the application with the provided configuration and context
func NewApp(cfg *config.AppConfig, ctx context.Context) (*App, error) {
	// Create a new instance of the application
	app := &App{
		cfg: cfg,
		Ctx: ctx,
	}

	// Create the app logger
	app.Logger = NewLogger("app", INFO, false, false)

	// Initialize the connection with the DB
	err := app.bootstrapDatabase()
	if err != nil {
		return nil, err
	}

	// Initialize services with the provided configuration
	app.Services = app.newServices(cfg, app.Logger)

	return app, nil
}

// GracefulShutdown performs a graceful shutdown of the application, closing the database connection and performing any necessary cleanup tasks
func (a *App) GracefulShutdown() error {
	// Close the database connection if it exists
	if a.db != nil {
		err := a.db.Close()
		if err != nil {
			a.Logger.Error("error closing DB connection", "error", err.Error())
			return err
		}
	}

	// Perform any other cleanup tasks here if necessary

	return nil
}
