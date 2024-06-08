package main

import (
	"context"
	"log"
	"os"

	authz "github.com/zeiss/fiber-authz"
	"github.com/zeiss/service-lens/internal/adapters"
	"github.com/zeiss/service-lens/internal/configs"
	"github.com/zeiss/service-lens/internal/utils"

	"github.com/katallaxie/pkg/logger"
	"github.com/spf13/cobra"
	seed "github.com/zeiss/gorm-seed"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var seeds = []seed.Seed{
	{
		Name: "create permissions",
		Run: func(db *gorm.DB) error {
			return db.Create([]authz.Permission{
				// {
				// 	Slug: utils.PermissionSuperAdmin.String(),
				// },
				// {
				// 	Slug: utils.PermissionAdmin.String(),
				// },
				// {
				// 	Slug: utils.PermissionCreate.String(),
				// },
				// {
				// 	Slug: utils.PermissionEdit.String(),
				// },
				// {
				// 	Slug: utils.PermissionDelete.String(),
				// },
				// {
				// 	Slug: utils.PermissionView.String(),
				// },
			}).Error
		},
	},
	{
		Name: "create roles",
		Run: func(db *gorm.DB) error {
			var permissions *[]authz.Permission
			if err := db.Where("slug IN ?", []string{utils.PermissionSuperAdmin.String(), utils.PermissionAdmin.String(), utils.PermissionCreate.String(), utils.PermissionEdit.String(), utils.PermissionDelete.String(), utils.PermissionView.String()}).Find(&permissions).Error; err != nil {
				return err
			}

			role := &authz.Role{
				Name:        "Super Admin",
				Permissions: permissions,
			}

			if err := db.Save(role).Error; err != nil {
				return err
			}

			permissions = &[]authz.Permission{}
			if err := db.Where("slug IN ?", []string{utils.PermissionAdmin.String(), utils.PermissionCreate.String(), utils.PermissionEdit.String(), utils.PermissionDelete.String(), utils.PermissionView.String()}).Find(&permissions).Error; err != nil {
				return err
			}

			role = &authz.Role{
				Name:        "Admin",
				Permissions: permissions,
			}

			if err := db.Save(role).Error; err != nil {
				return err
			}

			permissions = &[]authz.Permission{}
			if err := db.Where("slug IN ?", []string{utils.PermissionAdmin.String(), utils.PermissionCreate.String(), utils.PermissionEdit.String(), utils.PermissionDelete.String(), utils.PermissionView.String()}).Find(&permissions).Error; err != nil {
				return err
			}

			role = &authz.Role{
				Name:        "Owner",
				Permissions: permissions,
			}

			if err := db.Save(role).Error; err != nil {
				return err
			}

			permissions = &[]authz.Permission{}
			if err := db.Where("slug IN ?", []string{utils.PermissionCreate.String(), utils.PermissionEdit.String(), utils.PermissionDelete.String(), utils.PermissionView.String()}).Find(&permissions).Error; err != nil {
				return err
			}

			role = &authz.Role{
				Name:        "Editor",
				Permissions: permissions,
			}

			if err := db.Save(role).Error; err != nil {
				return err
			}

			permissions = &[]authz.Permission{}
			if err := db.Where("slug IN ?", []string{utils.PermissionView.String()}).Find(&permissions).Error; err != nil {
				return err
			}

			role = &authz.Role{
				Name:        "Viewer",
				Permissions: permissions,
			}

			if err := db.Save(role).Error; err != nil {
				return err
			}

			return nil
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
	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	db := adapters.NewDB(conn)
	err = db.RunMigration()
	if err != nil {
		return err
	}

	seeder := seed.NewSeeder(seed.WithDatabase(conn))
	err = seeder.Seed(ctx, seeds...)
	if err != nil {
		panic(err)
	}

	return nil
}
