package cmd

import (
	"context"
	"os"

	"github.com/kelseyhightower/envconfig"
	"github.com/zeiss/fiber-goth/providers"
	"github.com/zeiss/fiber-goth/providers/github"
	"github.com/zeiss/service-lens/internal/adapters/db"
	"github.com/zeiss/service-lens/internal/adapters/handlers"

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

func init() {
	Root.PersistentFlags().StringVar(&cfg.Flags.Addr, "addr", ":3000", "addr")
	Root.PersistentFlags().StringVar(&cfg.Flags.DB.Database, "db-database", cfg.Flags.DB.Database, "Database name")
	Root.PersistentFlags().StringVar(&cfg.Flags.DB.Username, "db-username", cfg.Flags.DB.Username, "Database user")
	Root.PersistentFlags().StringVar(&cfg.Flags.DB.Password, "db-password", cfg.Flags.DB.Password, "Database password")
	Root.PersistentFlags().IntVar(&cfg.Flags.DB.Port, "db-port", cfg.Flags.DB.Port, "Database port")
	Root.PersistentFlags().StringVar(&cfg.Flags.DB.Addr, "db-host", cfg.Flags.DB.Addr, "Database host")

	Root.SilenceUsage = true
}

var Root = &cobra.Command{
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		err := envconfig.Process("", cfg.Flags)
		if err != nil {
			return err
		}

		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		srv := NewWebSrv(cfg)

		s, _ := server.WithContext(cmd.Context())
		s.Listen(srv, false)

		return s.Wait()
	},
}

var _ server.Listener = (*WebSrv)(nil)

// WebSrv is the server that implements the Noop interface.
type WebSrv struct {
	cfg *Config
}

// NewWebSrv returns a new instance of NoopSrv.
func NewWebSrv(cfg *Config) *WebSrv {
	return &WebSrv{cfg}
}

// Start starts the server.
func (s *WebSrv) Start(ctx context.Context, ready server.ReadyFunc, run server.RunFunc) func() error {
	return func() error {
		providers.RegisterProvider(github.New(os.Getenv("GITHUB_KEY"), os.Getenv("GITHUB_SECRET"), "http://localhost:3000/auth/github/callback"))

		conn, err := gorm.Open(postgres.Open(cfg.DSN()), &gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				TablePrefix: "service_lens_",
			},
		})
		if err != nil {
			return err
		}

		store, err := db.NewDB(conn)
		if err != nil {
			return err
		}

		err = store.Migrate(ctx)
		if err != nil {
			return err
		}

		gorm, err := adapter.New(conn)
		if err != nil {
			return err
		}

		gothConfig := goth.Config{
			Adapter:        gorm,
			Secret:         goth.GenerateKey(),
			CookieHTTPOnly: true,
			ResponseFilter: func(c *fiber.Ctx) error {
				return c.Redirect("/")
			},
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
		site.Get("/teams/:id", handlers.ShowTeam())

		// Me ...
		app.Get("/me", handlers.Me())

		// Settings ...
		app.Get("/settings", handlers.ShowSettings())

		// Profiles ...
		app.Get("/profiles", handlers.ListProfiles())
		app.Get("/profiles/new", handlers.NewProfile())
		app.Post("/profiles/new", handlers.CreateProfile())
		app.Get("/profiles/:id", handlers.ShowProfile())
		app.Put("/profiles/:id", handlers.EditProfile())
		app.Delete("/profiles/:id", handlers.DeleteProfile())

		// Environments ...
		app.Get("/environments", handlers.ListEnvironments())
		app.Get("/environments/new", handlers.NewEnvironment())
		app.Post("/environments/new", handlers.CreateEnvironment())
		app.Get("/environments/:id", handlers.ShowEnvironment())
		app.Get("/environments/:id/edit", handlers.EditEnvironment())
		app.Put("/environments/:id", handlers.UpdateEnvironment())
		app.Delete("/environments/:id", handlers.DeleteEnvironment())

		// Lenses ...
		app.Get("/lenses", handlers.ListLenses())
		app.Get("/lenses/new", handlers.NewLens())
		app.Post("/lenses/new", handlers.CreateLens())
		app.Get("/lenses/:id", handlers.ShowLens())
		app.Get("/lenses/:id/edit", handlers.EditLens())
		app.Put("/lenses/:id", handlers.UpdateLens())
		app.Delete("/lenses/:id", handlers.DeleteLens())

		// Workloads ...
		app.Get("/workloads", handlers.ListWorkloads())
		app.Get("/workloads/new", handlers.NewWorkload())
		app.Post("/workloads/new", handlers.CreateWorkload())
		app.Get("/workloads/:id", handlers.ShowWorkload())
		app.Get("/workloads/:id/edit", handlers.EditWorkload())
		// app.Put("/workloads/:id", handlers.UpdateWorkload())
		app.Delete("/workloads/:id", handlers.DeleteWorkload())
		app.Get("/workloads/partials/environments", handlers.ListEnvironmentsPartial())
		app.Get("/workloads/partials/profiles", handlers.ListProfilesPartial())
		app.Get("/workloads/:id/lenses/:lens", handlers.ShowWorkloadLens())
		app.Get("/workloads/:id/lenses/:lens/edit", handlers.EditWorkloadLens())
		app.Get("/workloads/:workload/lenses/:lens/question/:question", handlers.ShowLensQuestion())
		app.Put("/workloads/:workload/lenses/:lens/question/:question", handlers.UpdateWorkloadAnswer())

		err = app.Listen(s.cfg.Flags.Addr)
		if err != nil {
			return err
		}

		return nil
	}
}
