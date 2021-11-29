package main

import (
	"github.com/Z00mZE/cortex/auth/config"
	"github.com/Z00mZE/cortex/auth/internal/app"
	"github.com/kelseyhightower/envconfig"
	"log"
)

func main() {
	var cfg config.AuthConfig
	if err := envconfig.Process("", &cfg); err != nil {
		log.Fatalf("Config error: %s", err)
	}
	app.Run(&cfg)
}
