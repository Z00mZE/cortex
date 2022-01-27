package app

import (
	"fmt"
	"github.com/Z00mZE/cortex/backend/auth/config"
	"github.com/Z00mZE/cortex/backend/auth/ent"
	routesV1 "github.com/Z00mZE/cortex/backend/auth/internal/controller/http/v1"
	"github.com/Z00mZE/cortex/backend/auth/pkg/hash"
	"github.com/Z00mZE/cortex/backend/auth/pkg/mailer"
	"github.com/Z00mZE/cortex/backend/auth/usecase"
	"golang.org/x/crypto/bcrypt"
	"time"

	//"github.com/Z00mZE/cortex/backend/auth/internal/cqrs/user"
	//"github.com/Z00mZE/cortex/backend/auth/internal/cqrs/user/command"
	//"github.com/Z00mZE/cortex/backend/auth/internal/cqrs/user/query"

	//"github.com/Z00mZE/cortex/backend/auth/internal/cqrs/user"
	//"github.com/Z00mZE/cortex/backend/auth/internal/cqrs/user/command"
	//"github.com/Z00mZE/cortex/backend/auth/internal/cqrs/user/query"

	//"github.com/Z00mZE/cortex/backend/auth/internal/cqrs/user/command"
	//"github.com/Z00mZE/cortex/backend/auth/internal/cqrs/user/query"
	"github.com/Z00mZE/cortex/backend/auth/pkg/httpserver"
	"github.com/Z00mZE/cortex/backend/auth/pkg/postgres"
	"github.com/labstack/echo/v4"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func Run(cfg *config.AuthConfig) {
	// Repository
	orm, err := postgres.NewPostgresORM(
		cfg.Database.Host,
		postgres.MaxPoolSize(10),
		postgres.SetConnMaxLifetime(time.Minute*15),
		postgres.SetConnMaxIdleTime(time.Minute*5),
	)
	if err != nil {
		log.Fatal(fmt.Errorf("app - Run - postgres.NewPostgresORM: %w", err))
	}

	defer func() {
		if err := orm.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	app := echo.New()
	httpServer := httpserver.NewHttpServer(app, httpserver.Port(cfg.Http.Port))
	fmt.Println("Ok")

	//	иницализация роутинга
	go initRoutes(app, orm)

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

func initRoutes(app *echo.Echo, orm *ent.Client) {
	authCases, authCasesError := usecase.NewAuthUseCase(
		hash.NewBCryptHash(bcrypt.MaxCost),
		mailer.NewDummyMailSender(),
		orm.User,
		orm.Session,
	)

	if authCasesError != nil {
		fmt.Println("init AuthCases service has error:", authCasesError)
	} else {
		routesV1.NewRouter(
			app,
			authCases,
		)
	}
}
