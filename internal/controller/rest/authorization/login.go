package authorization

import (
	"github.com/gin-gonic/gin"
	"github.com/pets-shelters/backend-svc/internal/controller/helpers"
	"net/http"
)

func (r *routes) login(ctx *gin.Context) {
	loginResult, err := r.authUseCase.Login()
	if err != nil {
		r.log.Error(err, "failed to process usecase - login")
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, helpers.FormInternalError(err.Error()))
		return
	}

	ctx.SetCookie(helpers.LoginCookieName, loginResult.CookieSession, r.loginCookieLifetime, "/", r.domain, false, true)
	ctx.Redirect(http.StatusTemporaryRedirect, loginResult.Url)
}
