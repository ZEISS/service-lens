package cmd

import (
	"context"
	"log"
	"os"

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
		providers.RegisterProvider(github.New(os.Getenv("GITHUB_KEY"), os.Getenv("GITHUB_SECRET"), "http://localhost:3000/auth/github/callback"))

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

		// Site ...
		site := app.Group("/site")
		site.Get("/teams", handlers.ListTeams())
		site.Get("/teams/new", handlers.NewTeam())
		site.Post("/teams/new", handlers.CreateTeam())
		site.Get("/teams/:id", handlers.ShowTeam())

		// Team ...
		team := app.Group("/teams/:t_slug")

		team.Get("/profiles", handlers.ListProfiles())
		team.Get("/profiles/new", handlers.NewProfile())
		team.Post("/profiles/new", handlers.CreateProfile())
		team.Get("/profiles/:id", handlers.ShowProfile())
		team.Put("/profiles/:id", handlers.EditProfile())
		team.Delete("/profiles/:id", handlers.DeleteProfile())

		// Environments ...
		team.Get("/environments", handlers.ListEnvironments())
		team.Get("/environments/new", handlers.NewEnvironment())
		team.Post("/environments/new", handlers.CreateEnvironment())
		team.Get("/environments/:id", handlers.ShowEnvironment())
		team.Get("/environments/:id/edit", handlers.EditEnvironment())
		team.Put("/environments/:id", handlers.UpdateEnvironment())
		team.Delete("/environments/:id", handlers.DeleteEnvironment())

		// Lenses ...
		team.Get("/lenses", handlers.ListLenses())
		team.Get("/lenses/new", handlers.NewLens())
		team.Post("/lenses/new", handlers.CreateLens())
		team.Get("/lenses/:id", handlers.ShowLens())
		team.Get("/lenses/:id/edit", handlers.EditLens())
		team.Put("/lenses/:id", handlers.UpdateLens())
		team.Delete("/lenses/:id", handlers.DeleteLens())

		// Workloads ...
		team.Get("/workloads", handlers.ListWorkloads())
		team.Get("/workloads/new", handlers.NewWorkload())
		team.Post("/workloads/new", handlers.CreateWorkload())
		team.Get("/workloads/:id", handlers.ShowWorkload())
		team.Get("/workloads/:id/edit", handlers.EditWorkload())
		// app.Put("/workloads/:id", handlers.UpdateWorkload())
		team.Delete("/workloads/:id", handlers.DeleteWorkload())
		team.Get("/workloads/partials/environments", handlers.ListEnvironmentsPartial())
		team.Get("/workloads/partials/profiles", handlers.ListProfilesPartial())
		team.Get("/workloads/:id/lenses/:lens", handlers.ShowWorkloadLens())
		team.Get("/workloads/:id/lenses/:lens/edit", handlers.EditWorkloadLens())
		team.Get("/workloads/:workload/lenses/:lens/question/:question", handlers.ShowLensQuestion())
		team.Put("/workloads/:workload/lenses/:lens/question/:question", handlers.UpdateWorkloadAnswer())

		// Me ...
		app.Get("/me", handlers.Me())

		// Settings ...
		app.Get("/settings", handlers.ShowSettings())

		err = app.Listen(s.cfg.Flags.Addr)
		if err != nil {
			return err
		}

		return nil
	}
}
