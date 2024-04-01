package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/pets-shelters/backend-svc/internal/controller/helpers"
	"github.com/pets-shelters/backend-svc/internal/controller/rest/authorization"
	"github.com/pets-shelters/backend-svc/internal/controller/rest/files"
	"github.com/pets-shelters/backend-svc/internal/controller/rest/shelters"
	"github.com/pets-shelters/backend-svc/internal/usecase"
	"github.com/pets-shelters/backend-svc/pkg/logger"
)

func NewRouter(handler *gin.Engine, log logger.Interface, useCases usecase.UseCases, routerConfigs helpers.RouterConfigs) {
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())

	handler.StaticFile("/dist.json", "./docs/dist.json")

	authorizationGroup := handler.Group("/authorization")
	{
		authorization.NewRoutes(authorizationGroup, useCases.Authorization, useCases.Jwt, log, routerConfigs)
	}
	sheltersGroup := handler.Group("/shelters")
	{
		shelters.NewRoutes(sheltersGroup, useCases.Shelters, useCases.Jwt, log, routerConfigs)
	}
	filesGroup := handler.Group("/files")
	{
		files.NewRoutes(filesGroup, useCases.Files, useCases.Jwt, log, routerConfigs.TemporaryFilesCfg)
	}
}
