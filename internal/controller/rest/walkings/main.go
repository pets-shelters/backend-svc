package walkings

import (
	"github.com/gin-gonic/gin"
	"github.com/pets-shelters/backend-svc/internal/controller/middlewares"
	"github.com/pets-shelters/backend-svc/internal/usecase"
	"github.com/pets-shelters/backend-svc/pkg/logger"
)

type routes struct {
	useCase usecase.IWalkings
	log     logger.Interface
}

func NewRoutes(handler *gin.RouterGroup, walkingsUseCase usecase.IWalkings,
	jwtUseCase usecase.IJwt, log logger.Interface) {
	r := &routes{
		walkingsUseCase,
		log,
	}

	handler.PUT("/:id", middlewares.ValidateAccessJwt(jwtUseCase), r.approve)
	handler.DELETE("/:id", middlewares.ValidateAccessJwt(jwtUseCase), r.delete)
	handler.GET("/", middlewares.ValidateAccessJwt(jwtUseCase), r.getList)
}

func NewRoutesWithAnimalId(handler *gin.RouterGroup, walkingsUseCase usecase.IWalkings,
	jwtUseCase usecase.IJwt, log logger.Interface) {
	r := &routes{
		walkingsUseCase,
		log,
	}

	handler.POST("/request", r.createPending)
	handler.POST("/", middlewares.ValidateAccessJwt(jwtUseCase), r.createApproved)
}
