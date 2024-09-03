package cfg

import (
	"os"
)

// Flags contains the command line flags.
type Flags struct {
	Environment             string `envconfig:"SERVICE_LENS_ENV" default:"production"`
	Addr                    string `envconfig:"SERVICE_LENS_ADDR" default:":8084"`
	DatabaseURI             string `envconfig:"SERVICE_LENS_DATABASE_URI" default:""`
	DatabaseTablePrefix     string `envconfig:"SERVICE_LENS_DATABASE_TABLE_PREFIX" default:"service_lens_"`
	FGAApiUrl               string `envconfig:"SERVICE_LENS_FGA_API_URL" default:""`
	FGAStoreID              string `envconfig:"SERVICE_LENS_FGA_STORE_ID" default:""`
	FGAAuthorizationModelID string `envconfig:"SERVICE_LENS_FGA_AUTHORIZATION_MODEL_ID" default:""`
	OIDCIssuer              string `envconfig:"SERVICE_LENS_OIDC_ISSUER" default:""`
	OIDCAudience            string `envconfig:"SERVICE_LENS_OIDC_AUDIENCE" default:""`
	GitHubCallbackURL       string `envconfig:"SERVICE_LENS_GITHUB_CALLBACK_URL" default:""`
	GitHubClientID          string `envconfig:"SERVICE_LENS_GITHUB_CLIENT_ID" default:""`
	GitHubClientSecret      string `envconfig:"SERVICE_LENS_GITHUB_CLIENT_SECRET" default:""`
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
