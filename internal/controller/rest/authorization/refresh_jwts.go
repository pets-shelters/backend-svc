package authorization

import (
	"github.com/gin-gonic/gin"
	"github.com/pets-shelters/backend-svc/internal/controller/helpers"
	"github.com/pets-shelters/backend-svc/internal/exceptions"
	"github.com/pkg/errors"
	"net/http"
)

func (r *routes) refreshJwts(ctx *gin.Context) {
	refreshToken, err := ctx.Cookie(helpers.RefreshTokenCookieName)
	if err != nil {
		ctx.AbortWithStatus(http.StatusForbidden)
		return
	}

	userEmail, err := r.jwtUseCase.VerifyRefreshToken(refreshToken)
	if err != nil {
		if errors.As(err, &exceptions.InvalidJwtException{}) {
			ctx.AbortWithStatusJSON(http.StatusForbidden, helpers.FormCustomError("code", err.Error()))
			return
		}
		r.log.Error(err, "failed to process usecase - verify refresh token")
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, helpers.FormInternalError(err.Error()))
		return
	}

	tokensPair, err := r.jwtUseCase.CreateTokensPair(userEmail)
	if err != nil {
		r.log.Error(err, "failed to process usecase - create tokens pair")
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, helpers.FormInternalError(err.Error()))
		return
	}

	ctx.SetCookie(helpers.AccessTokenCookieName, tokensPair.AccessToken, r.accessTokenLifetime, "/", r.domain, false, true)
	ctx.SetCookie(helpers.RefreshTokenCookieName, tokensPair.RefreshToken, r.refreshTokenLifetime, "/", r.domain, false, true)
}
