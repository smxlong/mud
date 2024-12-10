package mud

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kelseyhightower/envconfig"
	"github.com/smxlong/kit/logger"
	"github.com/smxlong/kit/webserver"
	"github.com/smxlong/mud/ent"
	"github.com/smxlong/mud/ent/player"
	"github.com/smxlong/mud/password"
	"github.com/spf13/pflag"
)

// Service is the mud's main service
type Service struct {
	DBDriver      string `envconfig:"DB_DRIVER" default:"sqlite3"`
	DBDataSource  string `envconfig:"DB_DATASOURCE" default:"file:ent?mode=memory&cache=shared&_fk=1"`
	ListenAddress string `envconfig:"LISTEN_ADDRESS" default:":3000"`
	JWTAudience   string `envconfig:"JWT_AUDIENCE" default:"mud"`
	JWTIssuer     string `envconfig:"JWT_ISSUER" default:"mud"`
	JWTSigningKey string `envconfig:"JWT_SIGNING_KEY" default:"CHANGEME"`
	g             *gin.Engine
	l             logger.Logger
	entcli        *ent.Client
}

// New creates a new mud service
func New() *Service {
	s := &Service{
		ListenAddress: ":3000",
		g:             gin.New(),
	}
	s.setupMiddleware()
	s.setupRoutes()
	return s
}

// Run the mud service
func (s *Service) Run(ctx context.Context, l logger.Logger) error {
	s.l = l
	entcli, err := ent.Open(s.DBDriver, s.DBDataSource)
	if err != nil {
		return err
	}
	if err = initSchema(ctx, entcli); err != nil {
		return err
	}
	s.entcli = entcli
	return webserver.ListenAndServe(ctx, &http.Server{
		Addr:    s.ListenAddress,
		Handler: s.g,
	})
}

// BindFlags binds the command line flags to the service
func (s *Service) BindFlags(flags *pflag.FlagSet) {
	flags.StringVar(&s.ListenAddress, "listen-address", s.ListenAddress, "listen address")
}

// BindEnvironment binds the environment variables to the service
func (s *Service) BindEnvironment() error {
	return envconfig.Process("", s)
}

// initSchema initializes the schema for the service
func initSchema(ctx context.Context, entcli *ent.Client) error {
	err := entcli.Schema.Create(ctx)
	if err != nil {
		return err
	}
	// find a user with email admin@example.com - create if not present
	_, err = entcli.Player.Query().Where(player.EmailEQ("admin@example.com")).Only(ctx)
	if err != nil {
		_, err = entcli.Player.Create().
			SetName("admin").
			SetEmail("admin@example.com").
			SetPassword(password.Hash("password")).
			Save(ctx)
		if err != nil {
			return err
		}
	}
	return nil
}
