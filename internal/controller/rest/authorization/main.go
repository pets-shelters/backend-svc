package authorization

import (
	"github.com/gin-gonic/gin"
	"github.com/pets-shelters/backend-svc/internal/controller/helpers"
	"github.com/pets-shelters/backend-svc/internal/controller/middlewares"
	"github.com/pets-shelters/backend-svc/internal/usecase"
	"github.com/pets-shelters/backend-svc/pkg/logger"
)

type routes struct {
	authUseCase          usecase.IAuthorization
	jwtUseCase           usecase.IJwt
	log                  logger.Interface
	domain               string
	loginCookieLifetime  int
	accessTokenLifetime  int
	refreshTokenLifetime int
	webClientUrl         string
	oauthWebRedirect     string
}

func NewRoutes(handler *gin.RouterGroup, authUseCase usecase.IAuthorization,
	jwtUseCase usecase.IJwt, log logger.Interface, routerConfigs helpers.RouterConfigs) {
	r := &routes{
		authUseCase,
		jwtUseCase,
		log,
		routerConfigs.Domain,
		routerConfigs.LoginCookieLifetime,
		routerConfigs.AccessTokenLifetime,
		routerConfigs.RefreshTokenLifetime,
		routerConfigs.WebClientUrl,
		routerConfigs.OAuthWebRedirect,
	}

	handler.GET("/refresh", r.refreshJwts)
	handler.GET("/login", r.login)
	handler.GET("/callback", r.callback)
	handler.GET("/user-info", middlewares.ValidateAccessJwt(jwtUseCase), r.getUserInfo)
}
