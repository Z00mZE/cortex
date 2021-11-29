package app

import (
	"fmt"
	"github.com/Z00mZE/cortex/auth/config"
	v1 "github.com/Z00mZE/cortex/auth/internal/controller/http/v1"
	"github.com/Z00mZE/cortex/auth/pkg/httpserver"
	"github.com/Z00mZE/cortex/auth/pkg/postgres"
	"github.com/labstack/echo/v4"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func Run(cfg *config.AuthConfig) {
	// Repository
	pg, err := postgres.New(cfg.Database.Host)
	if err != nil {
		log.Fatal(fmt.Errorf("app - Run - postgres.New: %w", err))
	}
	defer pg.Close()

	handler := echo.New()
	httpServer := httpserver.New(handler, httpserver.Port(cfg.Http.Port))
	fmt.Println("Ok")

	v1.NewRouter(handler, true)
	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		log.Println("app - Run - signal: " + s.String())
	case err = <-httpServer.Notify():
		log.Fatal(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
	}

	// Shutdown
	if err = httpServer.Shutdown(); err != nil {
		log.Fatal(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	}
}
