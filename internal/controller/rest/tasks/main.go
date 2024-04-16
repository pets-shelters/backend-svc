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
	handler.DELETE("/:id", middlewares.ValidateAccessJwt(jwtUseCase), r.delete)
	handler.PUT("/:id/done", middlewares.ValidateAccessJwt(jwtUseCase), r.setTaskDone)
	handler.GET("/", middlewares.ValidateAccessJwt(jwtUseCase), r.getListWithExecutions)
}

func NewRoutesWithAnimalId(handler *gin.RouterGroup, tasksUseCase usecase.ITasks,
	jwtUseCase usecase.IJwt, log logger.Interface) {
	r := &routes{
		tasksUseCase,
		log,
	}

	handler.GET("/", middlewares.ValidateAccessJwt(jwtUseCase), r.getListForAnimal)
}
