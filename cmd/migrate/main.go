package main

import (
	"context"
	"log"
	"os"

	"github.com/zeiss/fiber-goth/adapters"
	"github.com/zeiss/service-lens/internal/adapters/db"
	"github.com/zeiss/service-lens/internal/configs"
	"github.com/zeiss/service-lens/internal/models"

	"github.com/katallaxie/pkg/logger"
	"github.com/spf13/cobra"
	seed "github.com/zeiss/gorm-seed"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var seeds = []seed.Seed{
	{
		Name: "create teams",
		Run: func(db *gorm.DB) error {
			return db.Create([]adapters.GothTeam{
				{
					Name:        "Super Admins",
					Slug:        "superadmins",
					Description: "Super Admins have access to all features and can manage all resources.",
				},
			}).Error
		},
	},
	{
		Name: "create profile questions",
		Run: func(db *gorm.DB) error {
			return db.Create([]models.ProfileQuestion{
				{
					Title:         "What is the current cloud adoption phase for the organization architecting or operating the workloads in this profile?",
					Description:   "The cloud adoption phase is a measure of the maturity of the organization in adopting cloud technologies.",
					MulipleChoice: false,
					Choices: []models.ProfileQuestionChoice{
						{
							Title: "Envision Phase",
						},
						{
							Title: "Align Adoption Phase",
						},
						{
							Title: "Launch Adoption Phase",
						},
						{
							Title: "Scale Adoption Phase",
						},
						{
							Title: "Post-Adoption Phase",
						},
					},
				},
				{
					Title:         "What is the business value that workloads in this profile represent for your team, organization, or company?",
					Description:   "The business value is a measure of the value that the workloads in this profile provide to the organization.",
					MulipleChoice: false,
					Choices: []models.ProfileQuestionChoice{
						{
							Title: "Business-Critical Workloads",
						},
						{
							Title: "Strategic External-facing Workloads",
						},
						{
							Title: "Strategic Internal-facing Workloads",
						},
						{
							Title: "Internal Business Workloads",
						},
						{
							Title: "General Use Workloads",
						},
						{
							Title: "Experimentation or Testing Workloads",
						},
					},
				},
			}).Error
		},
	},
}

var cfg = configs.New()

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
	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: "service_lens_",
		},
	})
	if err != nil {
		return err
	}

	db, err := db.NewDB(conn)
	if err != nil {
		return err
	}

	err = db.Migrate(ctx)
	if err != nil {
		return err
	}

	seeder := seed.NewSeeder(conn)
	err = seeder.Seed(ctx, seeds...)
	if err != nil {
		panic(err)
	}

	return nil
}
