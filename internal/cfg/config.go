package cfg

import (
	"os"
)

// Flags contains the command line flags.
type Flags struct {
	Environment         string `envconfig:"SERVICE_LENS_ENV" default:"production"`
	Addr                string `envconfig:"SERVICE_LENS_ADDR" default:":8084"`
	DatabaseURI         string `envconfig:"POSTGRES_URL" default:""`
	DatabaseTablePrefix string `envconfig:"SERVICE_LENS_DATABASE_TABLE_PREFIX" default:"service_lens_"`
	GitHubEnabled       bool   `envconfig:"SERVICE_LENS_GITHUB_ENABLED" default:"false"`
	GitHubCallbackURL   string `envconfig:"SERVICE_LENS_GITHUB_CALLBACK_URL" default:""`
	GitHubClientID      string `envconfig:"SERVICE_LENS_GITHUB_CLIENT_ID" default:""`
	GitHubClientSecret  string `envconfig:"SERVICE_LENS_GITHUB_CLIENT_SECRET" default:""`
	EntraIDEnabled      bool   `envconfig:"SERVICE_LENS_ENTRAID_ENABLED" default:"true"`
	EntraIDClientID     string `envconfig:"SERVICE_LENS_ENTRAID_CLIENT_ID" default:""`
	EntraIDClientSecret string `envconfig:"SERVICE_LENS_ENTRAID_CLIENT_SECRET" default:""`
	EntraIDCallbackURL  string `envconfig:"SERVICE_LENS_ENTRAID_CALLBACK_URL" default:""`
	EntraIDTenantID     string `envconfig:"SERVICE_LENS_ENTRAID_TENANT_ID" default:""`
}

// NewFlags returns a new instance of Flags.
func NewFlags() *Flags {
	return &Flags{}
}

// New returns a new instance of Config.
func New() *Config {
	return &Config{
		Flags: NewFlags(),
	}
}

// Config contains the configuration.
type Config struct {
	Flags *Flags
}

// Cwd returns the current working directory.
func (c *Config) Cwd() (string, error) {
	return os.Getwd()
}
