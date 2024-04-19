package animals

import (
	"github.com/gin-gonic/gin"
	"github.com/pets-shelters/backend-svc/internal/controller/middlewares"
	"github.com/pets-shelters/backend-svc/internal/usecase"
	"github.com/pets-shelters/backend-svc/pkg/logger"
)

type routes struct {
	useCase    usecase.IAnimals
	jwtUseCase usecase.IJwt
	log        logger.Interface
}

func NewRoutes(handler *gin.RouterGroup, animalsUseCase usecase.IAnimals,
	jwtUseCase usecase.IJwt, log logger.Interface) {
	r := &routes{
		animalsUseCase,
		jwtUseCase,
		log,
	}

	idGroup := handler.Group("/:id")
	{
		idGroup.PUT("", middlewares.ValidateAccessJwt(jwtUseCase), r.update)
		idGroup.DELETE("", middlewares.ValidateAccessJwt(jwtUseCase), r.delete)
		idGroup.GET("", r.getById)
	}

	handler.POST("", middlewares.ValidateAccessJwt(jwtUseCase), r.create)
	handler.GET("", r.getList)
	handler.GET("/types", r.getTypes)
}
