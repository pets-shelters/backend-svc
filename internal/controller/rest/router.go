package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/pets-shelters/backend-svc/internal/controller/helpers"
	"github.com/pets-shelters/backend-svc/internal/controller/rest/adopters"
	"github.com/pets-shelters/backend-svc/internal/controller/rest/animals"
	"github.com/pets-shelters/backend-svc/internal/controller/rest/authorization"
	"github.com/pets-shelters/backend-svc/internal/controller/rest/employees"
	"github.com/pets-shelters/backend-svc/internal/controller/rest/files"
	"github.com/pets-shelters/backend-svc/internal/controller/rest/locations"
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
		sheltersSubgroup := sheltersGroup.Group("")
		{
			shelters.NewRoutes(sheltersSubgroup, useCases.Shelters, useCases.Jwt, log, routerConfigs)
		}
		employeesSubgroup := sheltersGroup.Group("/employees")
		{
			employees.NewRoutes(employeesSubgroup, useCases.Employees, useCases.Jwt, log)
		}
		locationsSubgroup := sheltersGroup.Group("/locations")
		{
			locations.NewRoutes(locationsSubgroup, useCases.Locations, useCases.Jwt, log)
		}
		locationsWithIdSubgroup := sheltersGroup.Group("/:id/locations")
		{
			locations.NewRoutesWithId(locationsWithIdSubgroup, useCases.Locations, useCases.Jwt, log)
		}
		animalsSubgroup := sheltersGroup.Group("/animals")
		{
			animals.NewRoutes(animalsSubgroup, useCases.Animals, useCases.Jwt, log)
		}
		adoptersSubgroup := sheltersGroup.Group("/adopters")
		{
			adopters.NewRoutes(adoptersSubgroup, useCases.Adopters, useCases.Jwt, log)
		}
	}
	filesGroup := handler.Group("/files")
	{
		files.NewRoutes(filesGroup, useCases.Files, useCases.Jwt, log, routerConfigs.TemporaryFilesCfg)
	}
}
