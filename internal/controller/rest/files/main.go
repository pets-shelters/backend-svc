package files

import (
	"github.com/gin-gonic/gin"
	"github.com/pets-shelters/backend-svc/configs"
	"github.com/pets-shelters/backend-svc/internal/controller/middlewares"
	"github.com/pets-shelters/backend-svc/internal/usecase"
	"github.com/pets-shelters/backend-svc/pkg/logger"
)

type routes struct {
	filesUseCase      usecase.IFiles
	log               logger.Interface
	temporaryFilesCfg configs.TemporaryFiles
}

func NewRoutes(handler *gin.RouterGroup, filesUseCase usecase.IFiles, jwtUseCase usecase.IJwt,
	log logger.Interface, tempFilesCfg configs.TemporaryFiles) {
	r := &routes{
		filesUseCase,
		log,
		tempFilesCfg,
	}

	handler.POST("/upload", middlewares.ValidateAccessJwt(jwtUseCase), r.upload)
}
