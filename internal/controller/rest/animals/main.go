package animals

import (
	"github.com/gin-gonic/gin"
	"github.com/pets-shelters/backend-svc/internal/controller/middlewares"
	"github.com/pets-shelters/backend-svc/internal/usecase"
	"github.com/pets-shelters/backend-svc/pkg/logger"
)

type routes struct {
	useCase usecase.IAnimals
	log     logger.Interface
}

func NewRoutes(handler *gin.RouterGroup, animalsUseCase usecase.IAnimals,
	jwtUseCase usecase.IJwt, log logger.Interface) {
	r := &routes{
		animalsUseCase,
		log,
	}

	handler.POST("/", middlewares.ValidateAccessJwt(jwtUseCase), r.create)
	handler.GET("/", r.getList)
}
