package app

import (
	"fmt"
	"github.com/Z00mZE/cortex/internal/auth/config"
	"github.com/Z00mZE/cortex/internal/auth/internal/adapter/postgre"
	routesV1 "github.com/Z00mZE/cortex/internal/auth/internal/controller/http/v1"
	"github.com/Z00mZE/cortex/internal/auth/internal/cqrs/user"
	"github.com/Z00mZE/cortex/internal/auth/internal/cqrs/user/query"
	"github.com/Z00mZE/cortex/internal/auth/pkg/httpserver"
	"github.com/Z00mZE/cortex/internal/auth/pkg/postgres"
	"github.com/Z00mZE/cortex/internal/auth/usecase"
	"github.com/labstack/echo/v4"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func Run(cfg *config.AuthConfig) {
	// Repository
	pg, err := postgres.NewPostgres(cfg.Database.Host)
	if err != nil {
		log.Fatal(fmt.Errorf("app - Run - postgres.NewPostgres: %w", err))
	}
	defer pg.Close()

	handler := echo.New()
	httpServer := httpserver.NewHttpServer(handler, httpserver.Port(cfg.Http.Port))
	fmt.Println("Ok")

	//	иницализация роутинга
	go initRoutes(handler, pg)

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

func initRoutes(handler *echo.Echo, pg *postgres.Postgres) {
	userRepository := postgre.NewUserRepository(pg)
	routesV1.NewRouter(
		handler,
		usecase.NewAuthUseCase(
			user.User{
				Query: user.Query{
					FindByEmailAndPasswordHandler: query.NewFindByEmailAndPasswordHandler(userRepository),
				},
				Command: user.Command{},
			},
		),
	)
}
