package locations

import (
	"github.com/gin-gonic/gin"
	"github.com/pets-shelters/backend-svc/internal/controller/middlewares"
	"github.com/pets-shelters/backend-svc/internal/usecase"
	"github.com/pets-shelters/backend-svc/pkg/logger"
)

type routes struct {
	useCase usecase.ILocations
	log     logger.Interface
}

func NewRoutes(handler *gin.RouterGroup, locationsUseCase usecase.ILocations,
	jwtUseCase usecase.IJwt, log logger.Interface) {
	r := &routes{
		locationsUseCase,
		log,
	}

	handler.POST("/", middlewares.ValidateAccessJwt(jwtUseCase), r.create)
	handler.DELETE("/:id", middlewares.ValidateAccessJwt(jwtUseCase), r.delete)
	handler.GET("/cities", r.getCities)
}

func NewRoutesWithId(handler *gin.RouterGroup, locationsUseCase usecase.ILocations,
	jwtUseCase usecase.IJwt, log logger.Interface) {
	r := &routes{
		locationsUseCase,
		log,
	}

	handler.GET("/", r.getList)
}
