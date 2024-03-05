package authorization

import (
	"github.com/gin-gonic/gin"
	"github.com/pets-shelters/backend-svc/internal/usecase"
	"github.com/pets-shelters/backend-svc/pkg/logger"
)

type routes struct {
	useCase usecase.Authorization
	log     logger.Interface
}

func NewRoutes(handler *gin.RouterGroup, useCase usecase.Authorization, log logger.Interface) {
	r := &routes{useCase, log}

	handler.POST("/sign-up", r.registration)
}
