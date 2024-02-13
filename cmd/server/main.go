package main

import (
	"context"
	"log"
	"os"

	"github.com/katallaxie/pkg/logger"
	"github.com/spf13/cobra"
	"github.com/zeiss/service-lens/internal/adapters"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB ...
type DB struct {
	Username string
	Password string
	Port     int
	Database string
	Host     string
}

// Flags contains the command line flags.
type Flags struct {
	Addr string
	DB   *DB
}

// New ...
func New() *Config {
	return &Config{
		Flags: &Flags{
			DB: &DB{
				Username: "example",
				Password: "example",
				Port:     5432,
				Database: "example",
				Host:     "host.docker.internal",
			},
		},
	}
}

// Config ...
type Config struct {
	Flags *Flags
}

// Cwd returns the current working directory.
func (c *Config) Cwd() (string, error) {
	return os.Getwd()
}

var cfg = New()

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

	return nil
}
