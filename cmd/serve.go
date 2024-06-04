package cmd

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/ecojuntak/laklak/internal/app"
	"github.com/ecojuntak/laklak/internal/config"
	"github.com/spf13/cobra"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "To run app",
	Run: func(cmd *cobra.Command, args []string) {
		config := config.Config()
		dbConnection := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", config.Database.Username, config.Database.Password, config.Database.Host, config.Database.Port, config.Database.Name)
		db, err := gorm.Open(postgres.Open(dbConnection), &gorm.Config{})
		if err != nil {
			slog.Error("cannot connect to database: %s", "error", err.Error())
			panic(err.Error())
		}

		slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, nil)))

		app := app.New(db)
		go func() {
			app.StartGrpcServer(config.Server.GrpcPort)
		}()

		app.StartHTTPServer(config.Server.GrpcPort, config.Server.HttpPort)
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
