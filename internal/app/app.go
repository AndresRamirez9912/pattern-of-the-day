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
	// ctx is the context for managing application lifecycle
	ctx context.Context
	// db is the database connection
	db *sql.DB
	// services contains the services layer of the application
	services *Services
	// logger is the application logger
	logger *Logger
}

// NewApp creates a new instance of the application with the provided configuration and context
func NewApp(cfg *config.AppConfig, ctx context.Context) (*App, error) {
	// Create a new instance of the application
	app := &App{
		cfg: cfg,
		ctx: ctx,
	}

	// Create the app logger
	app.logger = NewLogger("app", INFO, false, false)

	// Initialize the connection with the DB
	err := app.bootstrapDatabase()
	if err != nil {
		return nil, err
	}

	// Initialize services with the provided configuration
	app.services = NewServices(cfg, app.logger)

	return app, nil
}
