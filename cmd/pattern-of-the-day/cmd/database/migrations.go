package database

import (
	"database/sql"

	"github.com/pressly/goose/v3"
	"github.com/spf13/cobra"

	_ "github.com/glebarez/go-sqlite"
)

// migrationsDir is the directory containing the goose migration files.
const migrationsDir = "migrations"

// openDB opens the SQLite database at dbPath and configures goose to use the
// sqlite3 dialect for tracking applied migrations.
func openDB() (*sql.DB, error) {
	// Open the SQLite database
	db, err := sql.Open("sqlite", dbPath)
	if err != nil {
		return nil, err
	}

	// Set the goose dialect to sqlite3 for tracking applied migrations
	err = goose.SetDialect("sqlite3")
	if err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}

// NewUpCmd applies all pending migrations.
func NewUpCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "up",
		Short: "Apply all pending migrations",
		RunE: func(cmd *cobra.Command, args []string) error {
			db, err := openDB()
			if err != nil {
				return err
			}
			defer db.Close()

			err = goose.Up(db, migrationsDir)
			if err != nil {
				return err
			}

			cmd.Println("migrations applied")
			return nil
		},
	}
}

// NewDownCmd rolls back the last applied migration.
func NewDownCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "down",
		Short: "Roll back the last migration",
		RunE: func(cmd *cobra.Command, args []string) error {
			db, err := openDB()
			if err != nil {
				return err
			}
			defer db.Close()

			err = goose.Down(db, migrationsDir)
			if err != nil {
				return err
			}

			cmd.Println("last migration rolled back")
			return nil
		},
	}
}

// NewStatusCmd prints the migration status.
func NewStatusCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "status",
		Short: "Show migration status",
		RunE: func(cmd *cobra.Command, args []string) error {
			db, err := openDB()
			if err != nil {
				return err
			}
			defer db.Close()

			return goose.Status(db, migrationsDir)
		},
	}
}

// NewResetCmd rolls back all applied migrations.
func NewResetCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "reset",
		Short: "Roll back all migrations",
		RunE: func(cmd *cobra.Command, args []string) error {
			db, err := openDB()
			if err != nil {
				return err
			}
			defer db.Close()

			err = goose.Reset(db, migrationsDir)
			if err != nil {
				return err
			}

			cmd.Println("all migrations rolled back")
			return nil
		},
	}
}
