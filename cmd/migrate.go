package cmd

import (
	"errors"
	"fmt"
	"log/slog"

	"github.com/ecojuntak/laklak/internal/config"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/spf13/cobra"
)

// migrateCmd represents the migrate command
var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "To run database migration",
	RunE: func(cmd *cobra.Command, args []string) error {
		db := config.DatabaseConfig()
		dbConnection := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", db.Username, db.Password, db.Host, db.Port, db.Name)
		m, err := migrate.New("file://internal/migrations", dbConnection)
		if err != nil {
			slog.Error("error creating database migrator", "error", err)
			return err
		}

		err = m.Up()
		if errors.Is(err, migrate.ErrNoChange) {
			slog.Info("no database migration changes")
			return nil
		}

		return err
	},
}

func init() {
	rootCmd.AddCommand(migrateCmd)
}
