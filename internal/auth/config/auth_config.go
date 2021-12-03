package config

import (
	"github.com/kelseyhightower/envconfig"
)

type AuthConfig struct {
	Http struct {
		Port string `envconfig:"HTTP_PORT" required:"true"`
	}
	Database struct {
		Host string `envconfig:"DB_DSN" required:"true"`
	}
}

func NewAuthConfig() (*AuthConfig, error) {
	cfg := new(AuthConfig)
	if readError := envconfig.Process("", cfg); readError != nil {
		return cfg, readError
	}
	return cfg, nil
}
