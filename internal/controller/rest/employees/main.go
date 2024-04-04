package employees

import (
	"github.com/gin-gonic/gin"
	"github.com/pets-shelters/backend-svc/internal/controller/middlewares"
	"github.com/pets-shelters/backend-svc/internal/usecase"
	"github.com/pets-shelters/backend-svc/pkg/logger"
)

type routes struct {
	useCase usecase.IEmployees
	log     logger.Interface
}

func NewRoutes(handler *gin.RouterGroup, usersUseCase usecase.IEmployees,
	jwtUseCase usecase.IJwt, log logger.Interface) {
	r := &routes{
		usersUseCase,
		log,
	}

	handler.POST("/", middlewares.ValidateAccessJwt(jwtUseCase), r.create)
	handler.DELETE("/:id", middlewares.ValidateAccessJwt(jwtUseCase), r.delete)
	handler.GET("/", middlewares.ValidateAccessJwt(jwtUseCase), r.getList)
}
