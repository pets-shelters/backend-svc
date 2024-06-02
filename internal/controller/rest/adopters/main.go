package adopters

import (
	"github.com/gin-gonic/gin"
	"github.com/pets-shelters/backend-svc/internal/controller/middlewares"
	"github.com/pets-shelters/backend-svc/internal/usecase"
	"github.com/pets-shelters/backend-svc/pkg/logger"
)

type routes struct {
	useCase usecase.IAdopters
	log     logger.Interface
}

func NewRoutes(handler *gin.RouterGroup, adoptersUseCase usecase.IAdopters,
	jwtUseCase usecase.IJwt, log logger.Interface) {
	r := &routes{
		adoptersUseCase,
		log,
	}

	handler.GET("/:id", middlewares.ValidateAccessJwt(jwtUseCase), r.getById)
	handler.GET("", middlewares.ValidateAccessJwt(jwtUseCase), r.getList)
}

func NewRoutesWithAnimalId(handler *gin.RouterGroup, adoptersUseCase usecase.IAdopters,
	jwtUseCase usecase.IJwt, log logger.Interface) {
	r := &routes{
		adoptersUseCase,
		log,
	}

	handler.POST("", middlewares.ValidateAccessJwt(jwtUseCase), r.create)
}
