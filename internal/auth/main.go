package main

import (
	"github.com/Z00mZE/cortex/internal/auth/config"
	"github.com/Z00mZE/cortex/internal/auth/internal/app"
	"github.com/labstack/gommon/log"
)

func main() {
	config, configError := config.NewAuthConfig()
	if configError != nil {
		log.Fatalf("Config error: %s", configError)
	}
	app.Run(config)
}
