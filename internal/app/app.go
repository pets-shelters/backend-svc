package app

import (
	"github.com/pets-shelters/backend-svc/configs"
	"github.com/pets-shelters/backend-svc/internal/controller/helpers"
	"github.com/pets-shelters/backend-svc/internal/controller/rest"
	"github.com/pets-shelters/backend-svc/internal/usecase"
	"github.com/pets-shelters/backend-svc/internal/usecase/animals"
	"github.com/pets-shelters/backend-svc/internal/usecase/authorization"
	"github.com/pets-shelters/backend-svc/internal/usecase/employees"
	"github.com/pets-shelters/backend-svc/internal/usecase/files"
	"github.com/pets-shelters/backend-svc/internal/usecase/jwt"
	"github.com/pets-shelters/backend-svc/internal/usecase/locations"
	"github.com/pets-shelters/backend-svc/internal/usecase/mailjet"
	"github.com/pets-shelters/backend-svc/internal/usecase/oauth"
	"github.com/pets-shelters/backend-svc/internal/usecase/redis"
	"github.com/pets-shelters/backend-svc/internal/usecase/repo"
	"github.com/pets-shelters/backend-svc/internal/usecase/s3"
	"github.com/pets-shelters/backend-svc/internal/usecase/schedulers"
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
		log.Fatal(errors.Wrap(err, "failed to define repo db").Error())
	}
	defer pg.Close()

	err = migrateUp(cfg.PG.URL)
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed to migrateUp").Error())
	}

	stateLifetimeSecs := cfg.OAuth.StateLifetime.Seconds()
	dbRepo := repo.NewDBRepo(pg)
	oauth := oauth.NewOAuth(cfg.OAuth, cfg.Infrastructure.ServiceUrl)
	cache := redis.NewRedis(cfg.Redis)
	jwt := jwt.NewUseCase(cfg.Jwt)
	s3Provider, err := s3.NewProvider(cfg.S3)
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed to init s3 provider").Error())
	}
	emailsProvider := mailjet.NewMailjet(cfg.Mailjet)
	useCases := usecase.UseCases{
		Authorization: authorization.NewUseCase(dbRepo, *oauth, *cache, jwt),
		Jwt:           jwt,
		Shelters:      shelters.NewUseCase(dbRepo, cfg.S3.Endpoint),
		Files:         files.NewUseCase(dbRepo, s3Provider, cfg.S3.PublicReadBucket),
		Employees:     employees.NewUseCase(dbRepo, emailsProvider),
		Locations:     locations.NewUseCase(dbRepo),
		Animals:       animals.NewUseCase(dbRepo, cfg.S3.Endpoint),
	}

	jobsScheduler, err := schedulers.NewJobsScheduler(log, dbRepo)
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed to init jobs scheduler").Error())
	}
	err = jobsScheduler.WithCleanTemporaryFilesJob(s3Provider, cfg.TemporaryFiles)
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed to add clean_temporary_files job").Error())
	}
	jobsScheduler.Start()
	defer jobsScheduler.Shutdown()

	handler := gin.New()
	routerConfigs := helpers.RouterConfigs{
		LoginCookieLifetime:  int(stateLifetimeSecs),
		AccessTokenLifetime:  int(cfg.Jwt.AccessLifetime.Seconds()),
		RefreshTokenLifetime: int(cfg.Jwt.RefreshLifetime.Seconds()),
		Domain:               cfg.Infrastructure.Domain,
		TemporaryFilesCfg:    cfg.TemporaryFiles,
	}
	rest.NewRouter(handler, log, useCases, routerConfigs)
	httpServer := httpserver.New(handler, cfg.HTTP.Addr)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		log.Info("app - Run - signal: " + s.String())
	case err = <-httpServer.Notify():
		log.Error(errors.Wrap(err, "app - Run - httpServer.Notify").Error())
	}

	err = httpServer.Shutdown()
	if err != nil {
		log.Error(errors.Wrap(err, "failed to shutdown httpServer").Error())
	}
}
