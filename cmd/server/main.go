package main

import (
	"context"
	"log"
	"os"

	"github.com/zeiss/service-lens/internal/adapters"
	"github.com/zeiss/service-lens/internal/configs"
	"github.com/zeiss/service-lens/internal/controllers"
	"github.com/zeiss/service-lens/internal/services"

	"github.com/katallaxie/pkg/logger"
	"github.com/katallaxie/pkg/server"
	"github.com/spf13/cobra"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var cfg = configs.New()

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

var rootCmd = &cobra.Command{
	RunE: func(cmd *cobra.Command, args []string) error {
		return run(cmd.Context())
	},
}

func init() {
	rootCmd.PersistentFlags().StringVar(&cfg.Flags.Addr, "addr", ":8080", "addr")
	rootCmd.PersistentFlags().StringVar(&cfg.Flags.DB.Database, "db-database", cfg.Flags.DB.Database, "Database name")
	rootCmd.PersistentFlags().StringVar(&cfg.Flags.DB.Username, "db-username", cfg.Flags.DB.Username, "Database user")
	rootCmd.PersistentFlags().StringVar(&cfg.Flags.DB.Password, "db-password", cfg.Flags.DB.Password, "Database password")
	rootCmd.PersistentFlags().IntVar(&cfg.Flags.DB.Port, "db-port", cfg.Flags.DB.Port, "Database port")

	rootCmd.SilenceUsage = true
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}

func run(ctx context.Context) error {
	log.SetFlags(0)
	log.SetOutput(os.Stderr)

	logger.RedirectStdLog(logger.LogSink)

	dsn := "host=host.docker.internal user=example password=example dbname=example port=5432 sslmode=disable"
	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	db := adapters.NewDB(conn)
	err = db.RunMigration()
	if err != nil {
		return err
	}

	pc := controllers.NewProfilesController(db)
	lc := controllers.NewLensesController(db)

	srv, _ := server.WithContext(ctx)
	webSrv := services.New(cfg, pc, lc, db)

	srv.Listen(webSrv, true)
	if err := srv.Wait(); err != nil {
		return err
	}

	return nil
}
