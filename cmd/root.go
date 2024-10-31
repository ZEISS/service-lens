package cmd

import (
	"context"
	"fmt"
	"log"

	"github.com/kelseyhightower/envconfig"
	"github.com/zeiss/fiber-goth/adapters"
	"github.com/zeiss/fiber-goth/providers"
	"github.com/zeiss/fiber-goth/providers/entraid"
	"github.com/zeiss/fiber-goth/providers/github"
	"github.com/zeiss/pkg/dbx"
	"github.com/zeiss/service-lens/internal/adapters/db"
	"github.com/zeiss/service-lens/internal/adapters/handlers"
	"github.com/zeiss/service-lens/internal/cfg"
	"github.com/zeiss/service-lens/internal/models"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	logger "github.com/gofiber/fiber/v2/middleware/logger"
	requestid "github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/katallaxie/pkg/server"
	"github.com/spf13/cobra"
	goth "github.com/zeiss/fiber-goth"
	adapter "github.com/zeiss/fiber-goth/adapters/gorm"
	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/toasts"
	reload "github.com/zeiss/fiber-reload"
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
	Root.PersistentFlags().StringVar(&config.Flags.Environment, "environment", config.Flags.Environment, "environment")
	Root.PersistentFlags().StringVar(&config.Flags.DatabaseURI, "db-rul", config.Flags.DatabaseURI, "Database URI")
	Root.PersistentFlags().StringVar(&config.Flags.DatabaseTablePrefix, "db-table-prefix", config.Flags.DatabaseTablePrefix, "Database table prefix")
	Root.PersistentFlags().BoolVar(&config.Flags.GitHubEnabled, "github-enabled", config.Flags.GitHubEnabled, "GitHub enabled")
	Root.PersistentFlags().BoolVar(&config.Flags.EntraIDEnabled, "entraid-enabled", config.Flags.EntraIDEnabled, "EntraID enabled")

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
		if s.cfg.Flags.GitHubEnabled {
			providers.RegisterProvider(github.New(s.cfg.Flags.GitHubClientID, s.cfg.Flags.GitHubClientSecret, s.cfg.Flags.GitHubCallbackURL))
		}

		if s.cfg.Flags.EntraIDEnabled {
			providers.RegisterProvider(entraid.New(s.cfg.Flags.EntraIDClientID, s.cfg.Flags.EntraIDClientSecret, s.cfg.Flags.EntraIDCallbackURL, entraid.TenantType(s.cfg.Flags.EntraIDTenantID)))
		}

		conn, err := gorm.Open(postgres.Open(config.Flags.DatabaseURI), &gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				TablePrefix: s.cfg.Flags.DatabaseTablePrefix,
			},
		})
		if err != nil {
			return err
		}

		store, err := dbx.NewDatabase(conn, db.NewReadTx(), db.NewWriteTx())
		if err != nil {
			return err
		}

		err = store.Migrate(ctx,
			&adapters.GothUser{},
			&adapters.GothAccount{},
			&adapters.GothSession{},
			&adapters.GothVerificationToken{},
			&models.Template{},
			&models.Ownable{},
			&models.Workflow{},
			&models.WorkflowState{},
			&models.WorkflowTransition{},
			&models.Workable{},
			&models.Reaction{},
			&models.ProfileQuestion{},
			&models.ProfileQuestionChoice{},
			&models.ProfileQuestionAnswer{},
			&models.Design{},
			&models.DesignRevision{},
			&models.DesignComment{},
			&models.DesignCommentRevision{},
			&models.Environment{},
			&models.Profile{},
			&models.Tag{},
			&models.Lens{},
			&models.Pillar{},
			&models.Question{},
			&models.Resource{},
			&models.Choice{},
			&models.Risk{},
			&models.Workload{},
			&models.WorkloadLensQuestionAnswer{},
			&models.Setting{},
		)
		if err != nil {
			return err
		}

		gorm := adapter.New(conn)

		gothConfig := goth.Config{
			Adapter:        gorm,
			Secret:         goth.GenerateKey(),
			CookieHTTPOnly: true,
		}

		userHandler := handlers.NewUserHandler()
		settingsHandlers := handlers.NewSettingsHandler()
		previewHandlers := handlers.NewPreviewHandler()

		handlers := handlers.New(store)

		app := fiber.New(
			fiber.Config{
				ErrorHandler: toasts.DefaultErrorHandler,
			},
		)
		app.Use(requestid.New())
		app.Use(helmet.New())
		app.Use(logger.New())
		app.Use(reload.Environment(s.cfg.Flags.Environment))

		if s.cfg.Flags.Environment == reload.Development {
			reload.WithHotReload(app)
		}

		app.Use(goth.NewProtectMiddleware(gothConfig))

		compFuncConfig := htmx.Config{
			ErrorHandler: toasts.DefaultErrorHandler,
		}

		app.Get("/", handlers.Dashboard())
		app.Get("/login", htmx.NewCompFuncHandler(userHandler.Login(), compFuncConfig))
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
		designs.Get("/search/workflows", handlers.SearchWorkflows())
		designs.Get("/search/templates", handlers.SearchTemplates())
		designs.Get("/:id", handlers.ShowDesign())
		designs.Put("/:id", handlers.UpdateDesign())
		designs.Delete("/:id", handlers.DeleteDesign())
		designs.Post("/:id/tags", handlers.AddTagDesign())
		designs.Delete("/:id/tags/:tag_id", handlers.RemoveTagDesign())
		designs.Post("/:id/comments", handlers.CreateDesignComment())
		designs.Delete("/:id/comments/:comment_id", handlers.DeleteDesignComment())
		designs.Get("/:id/revisions", handlers.ListDesignRevisions())
		designs.Get("/:id/body/edit", handlers.EditBodyDesign())
		designs.Put("/:id/body/edit", handlers.UpdateBodyDesign())
		designs.Get("/:id/title/edit", handlers.EditTitleDesign())
		designs.Put("/:id/title/edit", handlers.UpdateTitleDesign())
		designs.Post("/:id/reactions", handlers.DesignReactions())
		designs.Post("/:id/tasks", handlers.Task())
		designs.Delete("/:id/reactions/:reaction_id", handlers.DesignReactions())
		designs.Post("/:id/comments/:comment_id/reactions", handlers.CreateDesignCommentReaction())
		designs.Delete("/:id/comments/:comment_id/reactions/:reaction_id", handlers.DeleteDesignCommentReaction())

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
		lenses.Post("/:id/publish", handlers.PublishLens())
		lenses.Delete("/:id/publish", handlers.UnpublishLens())

		// Workloads ...
		workloads := app.Group("/workloads")
		workloads.Get("/", handlers.ListWorkloads())
		workloads.Get("/new", handlers.NewWorkload())
		workloads.Post("/new", handlers.CreateWorkload())
		workloads.Get("/search/lenses", handlers.SearchLenses())
		workloads.Get("/search/environments", handlers.SearchEnvironments())
		workloads.Get("/search/profiles", handlers.SearchProfiles())
		workloads.Get("/:id", handlers.ShowWorkload())
		workloads.Get("/:id/edit", handlers.EditWorkload())
		workloads.Post("/:id/edit", handlers.EditWorkload())
		workloads.Post("/:id/tags", handlers.AddTagWorkload())
		workloads.Delete("/:id/tags/:tag_id", handlers.RemoveTagWorkload())
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
		workflows.Post("/new", handlers.CreateWorkflow())
		workflows.Get("/:id", handlers.ShowWorkflow())
		workflows.Post("/:id/steps", handlers.CreateWorkflowStep())
		workflows.Delete("/:id/steps/:step_id", handlers.DeleteWorkflowStep())
		workflows.Put("/:id/steps", handlers.UpdateWorkflowSteps())
		workflows.Delete("/:id", handlers.DeleteWorkflow())

		// Templates ...
		templates := app.Group("/templates")
		templates.Get("/", handlers.ListTemplates())
		templates.Get("/new", handlers.NewTemplate())
		templates.Get("/:id", handlers.ShowTemplate())
		templates.Delete("/:id", handlers.DeleteTemplate())
		templates.Get("/:id/edit/body", handlers.EditTemplateBody())
		templates.Put("/:id/edit/body", handlers.EditTemplateBody())
		templates.Get("/:id/edit/title", handlers.EditTemplateTitle())
		templates.Put("/:id/edit/title", handlers.EditTemplateTitle())
		templates.Post("/new", handlers.CreateTemplate())

		// Me ...
		app.Get("/me", handlers.Me())

		// Settings ...
		app.Get("/settings", htmx.NewCompFuncHandler(settingsHandlers.ListSettings, compFuncConfig))

		// Preview ...
		app.Post("/preview", htmx.NewCompFuncHandler(previewHandlers.Preview, compFuncConfig))

		err = app.Listen(s.cfg.Flags.Addr)
		if err != nil {
			return err
		}

		return nil
	}
}
