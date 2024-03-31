package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/pets-shelters/backend-svc/internal/controller/helpers"
	"github.com/pets-shelters/backend-svc/internal/exceptions"
	"github.com/pets-shelters/backend-svc/internal/usecase"
	"github.com/pkg/errors"
	"net/http"
	"strconv"
)

func ValidateAccessJwt(jwt usecase.IJwt) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		accessToken, err := ctx.Cookie(helpers.AccessTokenCookieName)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusForbidden, helpers.FormCustomError("forb", err.Error()))
			return
		}

		userId, err := jwt.VerifyAccessToken(accessToken)
		if err != nil {
			if errors.As(err, &exceptions.InvalidJwtException{}) {
				ctx.AbortWithStatusJSON(http.StatusForbidden, helpers.FormCustomError("forb", "invalid jwt"))
				return
			}
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, helpers.FormInternalError(err.Error()))
			return
		}

		userIdInt, err := strconv.Atoi(userId)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusForbidden, helpers.FormCustomError("forb", err.Error()))
			return
		}
		ctx.Set(helpers.JwtIdCtx, int64(userIdInt))
		ctx.Next()
	}
}
