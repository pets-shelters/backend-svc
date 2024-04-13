package tasks

import (
	"github.com/gin-gonic/gin"
	"github.com/pets-shelters/backend-svc/internal/controller/middlewares"
	"github.com/pets-shelters/backend-svc/internal/usecase"
	"github.com/pets-shelters/backend-svc/pkg/logger"
)

type routes struct {
	useCase usecase.ITasks
	log     logger.Interface
}

func NewRoutes(handler *gin.RouterGroup, tasksUseCase usecase.ITasks,
	jwtUseCase usecase.IJwt, log logger.Interface) {
	r := &routes{
		tasksUseCase,
		log,
	}

	handler.POST("/", middlewares.ValidateAccessJwt(jwtUseCase), r.create)
}
