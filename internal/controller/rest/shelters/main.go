package shelters

import (
	"github.com/gin-gonic/gin"
	"github.com/pets-shelters/backend-svc/internal/controller/helpers"
	"github.com/pets-shelters/backend-svc/internal/controller/middlewares"
	"github.com/pets-shelters/backend-svc/internal/usecase"
	"github.com/pets-shelters/backend-svc/pkg/logger"
)

type routes struct {
	useCase usecase.IShelters
	log     logger.Interface
}

func NewRoutes(handler *gin.RouterGroup, sheltersUseCase usecase.IShelters,
	jwtUseCase usecase.IJwt, log logger.Interface, routerConfigs helpers.RouterConfigs) {
	r := &routes{
		sheltersUseCase,
		log,
	}

	handler.POST("/", middlewares.ValidateAccessJwt(jwtUseCase), r.create)
}
