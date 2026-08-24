package app

import (
	"database/sql"

	_ "github.com/glebarez/go-sqlite"
)

// bootstrapDatabase initializes the database connection with SQLite
func (a *App) bootstrapDatabase() error {
	// Open the database connection
	// The file "app.db" is automatically created if it does not exist.
	db, err := sql.Open("sqlite", "./app.db")
	if err != nil {
		a.logger.Error("error openning DB connection", "error", err.Error())

		return err
	}

	// Configure DB connection from settings
	db.SetMaxOpenConns(a.cfg.Database.MaxOpenConns)
	db.SetMaxIdleConns(a.cfg.Database.MinIdleConns)

	// Close connection when the application is shutting down
	defer db.Close()

	// Validate the connection to the database by pinging it
	err = db.Ping()
	if err != nil {
		a.logger.Error("error pinging DB connection", "error", err.Error())

		return err
	}

	// Assign the database connection to the app
	a.db = db

	// If we reach here, the database connection is successfully established
	return nil
}
