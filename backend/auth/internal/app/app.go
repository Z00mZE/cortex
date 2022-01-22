package app

import (
	"fmt"
	"github.com/Z00mZE/cortex/backend/auth/config"
	"github.com/Z00mZE/cortex/backend/auth/internal/adapter/postgre"
	routesV1 "github.com/Z00mZE/cortex/backend/auth/internal/controller/http/v1"
	"github.com/Z00mZE/cortex/backend/auth/internal/cqrs/user"
	"github.com/Z00mZE/cortex/backend/auth/internal/cqrs/user/command"
	"github.com/Z00mZE/cortex/backend/auth/internal/cqrs/user/query"
	"github.com/Z00mZE/cortex/backend/auth/pkg/httpserver"
	"github.com/Z00mZE/cortex/backend/auth/pkg/postgres"
	"github.com/Z00mZE/cortex/backend/auth/usecase"
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

	app := echo.New()
	httpServer := httpserver.NewHttpServer(app, httpserver.Port(cfg.Http.Port))
	fmt.Println("Ok")

	//	иницализация роутинга
	go initRoutes(app, pg)

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		log.Println("app - Run - signal: " + s.String())
		if shutdownError := httpServer.Shutdown(); shutdownError != nil {
			log.Panicln(shutdownError)
		}
	case err = <-httpServer.Notify():
		log.Fatal(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
	}

	// Shutdown
	if err = httpServer.Shutdown(); err != nil {
		log.Fatal(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	}
}

func initRoutes(app *echo.Echo, pg *postgres.Postgres) {
	userRepository := postgre.NewUserRepository(pg)
	routesV1.NewRouter(
		app,
		usecase.NewAuthUseCase(
			user.User{
				Query: user.Query{
					IsEmailRegistered:             query.NewIsEmailRegistered(userRepository),
					FindByEmailAndPasswordHandler: query.NewFindByEmailAndPasswordHandler(userRepository),
				},
				Command: user.Command{
					CreateUserHandler: command.NewCreateUser(userRepository),
				},
			},
		),
	)
}
