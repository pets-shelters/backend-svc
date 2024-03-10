package authorization

import (
	"github.com/gin-gonic/gin"
	"github.com/pets-shelters/backend-svc/internal/controller/helpers"
	"net/http"
)

func (r *routes) callback(ctx *gin.Context) {
	cookie, err := ctx.Cookie(helpers.LoginCookieName)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helpers.FormBadRequestError(err.Error()))
		return
	}

	tokensPair, err := r.authUseCase.Callback(ctx, cookie, ctx.Query("state"), ctx.Query("code"))
	if err != nil {
		r.log.Error(err, "failed to process usecase - callback")
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, helpers.FormInternalError(err.Error()))
		return
	}

	ctx.SetCookie(helpers.AccessTokenCookieName, tokensPair.AccessToken, r.accessTokenLifetime, "/", ctx.Request.Host, false, true)
	ctx.SetCookie(helpers.RefreshTokenCookieName, tokensPair.RefreshToken, r.refreshTokenLifetime, "/", ctx.Request.Host, false, true)
	//TODO add redirect to front
}
