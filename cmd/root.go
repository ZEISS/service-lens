package cmd

import (
	"context"
	"fmt"
	"log"

	"github.com/kelseyhightower/envconfig"
	"github.com/zeiss/fiber-goth/providers"
	"github.com/zeiss/fiber-goth/providers/github"
	seed "github.com/zeiss/gorm-seed"
	"github.com/zeiss/service-lens/internal/adapters/db"
	"github.com/zeiss/service-lens/internal/adapters/handlers"
	"github.com/zeiss/service-lens/internal/cfg"

	"github.com/gofiber/fiber/v2"
	logger "github.com/gofiber/fiber/v2/middleware/logger"
	requestid "github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/katallaxie/pkg/server"
	"github.com/spf13/cobra"
	goth "github.com/zeiss/fiber-goth"
	adapter "github.com/zeiss/fiber-goth/adapters/gorm"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
	os      = "unknown"
	arch    = "unknown"
)

var versionFmt = fmt.Sprintf("%s-%s (%s) %s/%s", version, commit, date, os, arch)

var config *cfg.Config

func init() {
	config = cfg.New()

	err := envconfig.Process("", config.Flags)
	if err != nil {
		log.Fatal(err)
	}

	Root.AddCommand(Migrate)

	Root.PersistentFlags().StringVar(&config.Flags.Addr, "addr", config.Flags.Addr, "addr")
	Root.PersistentFlags().StringVar(&config.Flags.DatabaseURI, "db-uri", config.Flags.DatabaseURI, "Database URI")
	Root.PersistentFlags().StringVar(&config.Flags.DatabaseTablePrefix, "db-table-prefix", config.Flags.DatabaseTablePrefix, "Database table prefix")
	Root.PersistentFlags().StringVar(&config.Flags.FGAApiUrl, "fga-api-url", config.Flags.FGAApiUrl, "FGA API URL")
	Root.PersistentFlags().StringVar(&config.Flags.FGAStoreID, "fga-store-id", config.Flags.FGAStoreID, "FGA Store ID")
	Root.PersistentFlags().StringVar(&config.Flags.FGAAuthorizationModelID, "fga-authorization-model-id", config.Flags.FGAAuthorizationModelID, "FGA Authorization Model ID")
	Root.PersistentFlags().StringVar(&config.Flags.OIDCIssuer, "oidc-issuer", config.Flags.OIDCIssuer, "OIDC Issuer")
	Root.PersistentFlags().StringVar(&config.Flags.OIDCAudience, "oidc-audience", config.Flags.OIDCAudience, "OIDC Audience")

	Root.SilenceUsage = true
}

var Root = &cobra.Command{
	Version: versionFmt,
	RunE: func(cmd *cobra.Command, args []string) error {
		srv := NewWebSrv(config)

		s, _ := server.WithContext(cmd.Context())
		s.Listen(srv, false)

		return s.Wait()
	},
}

var _ server.Listener = (*WebSrv)(nil)

// WebSrv is the server that implements the Noop interface.
type WebSrv struct {
	cfg *cfg.Config
}

// NewWebSrv returns a new instance of NoopSrv.
func NewWebSrv(cfg *cfg.Config) *WebSrv {
	return &WebSrv{cfg}
}

// Start starts the server.
func (s *WebSrv) Start(ctx context.Context, ready server.ReadyFunc, run server.RunFunc) func() error {
	return func() error {
		providers.RegisterProvider(github.New(s.cfg.Flags.GitHubClientID, s.cfg.Flags.GitHubClientSecret, s.cfg.Flags.GitHubCallbackURL))

		conn, err := gorm.Open(postgres.Open(s.cfg.Flags.DatabaseURI), &gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				TablePrefix: s.cfg.Flags.DatabaseTablePrefix,
			},
		})
		if err != nil {
			return err
		}

		store, err := seed.NewDatabase(conn, db.NewReadTx(), db.NewWriteTx())
		if err != nil {
			return err
		}

		err = store.Migrate(ctx)
		if err != nil {
			return err
		}

		gorm := adapter.New(conn)

		gothConfig := goth.Config{
			Adapter:        gorm,
			Secret:         goth.GenerateKey(),
			CookieHTTPOnly: true,
		}

		handlers := handlers.New(store)

		app := fiber.New()
		app.Use(requestid.New())
		app.Use(logger.New())

		app.Use(goth.NewProtectMiddleware(gothConfig))

		app.Get("/", handlers.Dashboard())
		app.Get("/login", handlers.Login())
		app.Get("/login/:provider", goth.NewBeginAuthHandler(gothConfig))
		app.Get("/auth/:provider/callback", goth.NewCompleteAuthHandler(gothConfig))
		app.Get("/logout", goth.NewLogoutHandler(gothConfig))

		// Stats ...
		stats := app.Group("/stats")
		stats.Get("/profiles", handlers.StatsTotalProfiles())
		stats.Get("/designs", handlers.StatsTotalDesigns())
		stats.Get("/workloads", handlers.StatsTotalWorkloads())

		// Designs ...
		designs := app.Group("/designs")
		designs.Get("/", handlers.ListDesigns())
		designs.Get("/new", handlers.NewDesign())
		designs.Post("/new", handlers.CreateDesign())
		designs.Get("/:id", handlers.ShowDesign())
		designs.Put("/:id", handlers.UpdateDesign())
		designs.Delete("/:id", handlers.DeleteDesign())
		designs.Post("/:id/comments", handlers.CreateDesignComment())
		designs.Get("/:id/body/edit", handlers.EditBodyDesign())
		designs.Put("/:id/body/edit", handlers.UpdateBodyDesign())
		designs.Get("/:id/title/edit", handlers.EditTitleDesign())
		designs.Put("/:id/title/edit", handlers.UpdateTitleDesign())

		// Profiles
		profiles := app.Group("/profiles")
		profiles.Get("/", handlers.ListProfiles())
		profiles.Get("/new", handlers.NewProfile())
		profiles.Post("/new", handlers.CreateProfile())
		profiles.Get("/:id", handlers.ShowProfile())
		profiles.Put("/:id", handlers.EditProfile())
		profiles.Delete("/:id", handlers.DeleteProfile())

		// Environments ...
		environments := app.Group("/environments")
		environments.Get("/", handlers.ListEnvironments())
		environments.Get("/new", handlers.NewEnvironment())
		environments.Post("/new", handlers.CreateEnvironment())
		environments.Get("/:id", handlers.ShowEnvironment())
		environments.Get("/:id/edit", handlers.EditEnvironment())
		environments.Put("/:id", handlers.UpdateEnvironment())
		environments.Delete("/:id", handlers.DeleteEnvironment())

		// Lenses ...
		lenses := app.Group("/lenses")
		lenses.Get("/", handlers.ListLenses())
		lenses.Post("/", handlers.NewLens())
		lenses.Get("/:id", handlers.ShowLens())
		lenses.Get("/:id/edit", handlers.EditLens())
		lenses.Put("/:id", handlers.UpdateLens())
		lenses.Delete("/:id", handlers.DeleteLens())

		// Workloads ...
		workloads := app.Group("/workloads")
		workloads.Get("/", handlers.ListWorkloads())
		workloads.Get("/new", handlers.NewWorkload())
		workloads.Post("/new", handlers.CreateWorkload())
		workloads.Post("/search/lenses", handlers.SearchLenses())
		workloads.Post("/search/environments", handlers.SearchEnvironments())
		workloads.Post("/search/profiles", handlers.SearchProfiles())
		workloads.Get("/:id", handlers.ShowWorkload())
		workloads.Get("/:id/edit", handlers.EditWorkload())
		// app.Put("/workloads/:id", handlers.UpdateWorkload())
		workloads.Delete("/:id", handlers.DeleteWorkload())
		workloads.Get("/partials/environments", handlers.ListEnvironmentsPartial())
		workloads.Get("/partials/profiles", handlers.ListProfilesPartial())
		workloads.Get("/:id/lenses/:lens", handlers.ShowWorkloadLens())
		workloads.Get("/:id/lenses/:lens/edit", handlers.EditWorkloadLens())
		workloads.Get("/:workload/lenses/:lens/question/:question", handlers.ShowLensQuestion())
		workloads.Put("/:workload/lenses/:lens/question/:question", handlers.UpdateWorkloadAnswer())

		// Tags ...
		tags := app.Group("/tags")
		tags.Get("/", handlers.ListTags())
		tags.Post("/new", handlers.CreateTag())
		tags.Delete("/:id", handlers.DeleteTag())

		// Workflows ...
		workflows := app.Group("/workflows")
		workflows.Get("/", handlers.ListWorkflows())

		// Templates ...
		templates := app.Group("/templates")
		templates.Get("/", handlers.ListTemplates())
		templates.Get("/new", handlers.NewTemplate())
		templates.Get("/:id", handlers.ShowTemplate())
		templates.Post("/new", handlers.CreateTemplate())

		// Me ...
		app.Get("/me", handlers.Me())

		// Settings ...
		app.Get("/settings", handlers.ShowSettings())

		// Preview ...
		app.Post("/preview", handlers.Preview())

		err = app.Listen(s.cfg.Flags.Addr)
		if err != nil {
			return err
		}

		return nil
	}
}
