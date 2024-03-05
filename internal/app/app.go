package app

import (
	"github.com/pets-shelters/backend-svc/configs"
	"github.com/pets-shelters/backend-svc/internal/controller/rest"
	"github.com/pets-shelters/backend-svc/internal/usecase"
	"github.com/pets-shelters/backend-svc/internal/usecase/authorization"
	"github.com/pets-shelters/backend-svc/internal/usecase/repo"
	"github.com/pets-shelters/backend-svc/pkg/httpserver"
	"github.com/pets-shelters/backend-svc/pkg/logger"
	"github.com/pets-shelters/backend-svc/pkg/postgres"
	"github.com/pkg/errors"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
)

func Run(cfg *configs.Config) {
	log := logger.New(cfg.Log.Level)

	pg, err := postgres.New(cfg.PG.URL)
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed to define postgres db"))
	}
	defer pg.Close()

	err = migrateUp(cfg.PG.URL)
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed to migrateUp"))
	}

	dbRepo := repo.NewDBRepo(pg)
	useCases := usecase.UseCases{
		Authorization: authorization.NewUseCase(dbRepo),
	}

	handler := gin.New()
	rest.NewRouter(handler, log, useCases)
	httpServer := httpserver.New(handler, cfg.HTTP.Port)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		log.Info("app - Run - signal: " + s.String())
	case err = <-httpServer.Notify():
		log.Error(errors.Wrap(err, "app - Run - httpServer.Notify"))
	}

	err = httpServer.Shutdown()
	if err != nil {
		log.Error(errors.Wrap(err, "failed to shutdown httpServer"))
	}
}
