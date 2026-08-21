package app

import (
	"context"
	"database/sql"

	"github.com/AndresRamirez9912/pattern-of-the-day/internal/app/config"
)

// App represents the application and its dependencies
type App struct {
	cfg      *config.AppConfig
	ctx      context.Context
	db       *sql.DB
	services *Services
}

// NewApp creates a new instance of the application with the provided configuration and context
func NewApp(cfg *config.AppConfig, ctx context.Context) (*App, error) {
	// Create a new instance of the application
	app := &App{
		cfg: cfg,
		ctx: ctx,
	}

	// Initialize the connection with the DB
	err := app.bootstrapDatabase()
	if err != nil {
		return nil, err
	}

	// Initialize services with the provided configuration
	app.services = NewServices(cfg)

	return app, nil
}
