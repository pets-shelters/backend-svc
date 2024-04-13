package authorization

import (
	"github.com/gin-gonic/gin"
	"github.com/pets-shelters/backend-svc/internal/controller/helpers"
	"github.com/pets-shelters/backend-svc/internal/structs/responses"
	"net/http"
)

func (r *routes) getUserInfo(ctx *gin.Context) {
	userId, ok := ctx.Get(helpers.JwtIdCtx)
	if !ok {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, helpers.FormCustomError(helpers.Unauthorized, ""))
		return
	}

	userinfo, err := r.authUseCase.GetUserInfo(
		ctx.Request.Context(),
		userId.(int64),
	)
	if err != nil {
		r.log.Error(err.Error(), "failed to process usecase - get user info")
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, helpers.FormInternalError())
		return
	}

	ctx.JSON(http.StatusOK, helpers.JsonData[responses.UserInfo]{
		Data: *userinfo,
	})
}
