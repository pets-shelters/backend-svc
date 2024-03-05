package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/pets-shelters/backend-svc/internal/controller/rest/authorization"

	"github.com/pets-shelters/backend-svc/internal/usecase"
	"github.com/pets-shelters/backend-svc/pkg/logger"
)

func NewRouter(handler *gin.Engine, log logger.Interface, useCases usecase.UseCases) {
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())

	handlerGroup := handler.Group("/authorization")
	{
		authorization.NewRoutes(handlerGroup, useCases.Authorization, log)
	}
}
