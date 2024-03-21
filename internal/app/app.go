package app

import (
	"github.com/pets-shelters/backend-svc/configs"
	"github.com/pets-shelters/backend-svc/internal/controller/helpers"
	"github.com/pets-shelters/backend-svc/internal/controller/rest"
	"github.com/pets-shelters/backend-svc/internal/usecase"
	"github.com/pets-shelters/backend-svc/internal/usecase/authorization"
	"github.com/pets-shelters/backend-svc/internal/usecase/jwt"
	"github.com/pets-shelters/backend-svc/internal/usecase/oauth"
	"github.com/pets-shelters/backend-svc/internal/usecase/postgres"
	"github.com/pets-shelters/backend-svc/internal/usecase/redis"
	"github.com/pets-shelters/backend-svc/internal/usecase/shelters"
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

	stateLifetimeSecs := cfg.OAuth.StateLifetime.Seconds()
	dbRepo := postgres.NewDBRepo(pg)
	oauth := oauth.NewOAuth(cfg.OAuth, cfg.Infrastructure.ServiceUrl)
	cache := redis.NewRedis(cfg.Redis)
	jwt := jwt.NewUseCase(cfg.Jwt)
	useCases := usecase.UseCases{
		Authorization: authorization.NewUseCase(dbRepo, *oauth, *cache, jwt),
		Jwt:           jwt,
		Shelters:      shelters.NewUseCase(dbRepo),
	}

	handler := gin.New()
	routerConfigs := helpers.RouterConfigs{
		LoginCookieLifetime:  int(stateLifetimeSecs),
		AccessTokenLifetime:  int(cfg.Jwt.AccessLifetime.Seconds()),
		RefreshTokenLifetime: int(cfg.Jwt.RefreshLifetime.Seconds()),
		Domain:               cfg.Infrastructure.Domain,
	}
	rest.NewRouter(handler, log, useCases, routerConfigs)
	httpServer := httpserver.New(handler, cfg.HTTP.Addr)

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
