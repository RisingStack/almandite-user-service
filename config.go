package main

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

// Configuration ..
type Configuration struct {
	PostgresURL string `envconfig:"pg_url" required:"true"`
	DebugSQL    bool   `envconfig:"debug_sql" default:"false"`
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
