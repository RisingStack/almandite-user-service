package config

import (
	"log"
	// ...
	_ "github.com/joho/godotenv/autoload"
	"github.com/kelseyhightower/envconfig"
)

// Configuration ..
type Configuration struct {
	PostgresURL   string `envconfig:"pg_url" required:"true"`
	DebugSQL      bool   `envconfig:"debug_sql" default:"false"`
	JwtSigningKey string `envonfig:"jwt_signing_key" default:"jwtsigningkey"`
}

var configuration Configuration

// GetConfiguration returns the configuration settings
func GetConfiguration() *Configuration {
	return &configuration
}

func init() {
	err := envconfig.Process("", &configuration)
	if err != nil {
		log.Fatal("Configuration can not be processsed: ", err)
	}
}
