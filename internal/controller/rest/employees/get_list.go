package employees

import (
	"github.com/gin-gonic/gin"
	"github.com/pets-shelters/backend-svc/internal/controller/helpers"
	"github.com/pets-shelters/backend-svc/internal/structs/responses"
	"net/http"
)

func (r *routes) getList(ctx *gin.Context) {
	userId, ok := ctx.Get(helpers.JwtIdCtx)
	if !ok {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, helpers.FormCustomError(helpers.Unauthorized, ""))
		return
	}

	shelterEmployees, err := r.useCase.GetList(
		ctx.Request.Context(),
		userId.(int64),
	)
	if err != nil {
		r.log.Error(err.Error(), "failed to process usecase - get employees list")
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, helpers.FormInternalError())
		return
	}

	ctx.JSON(http.StatusOK, helpers.JsonData[[]responses.Employee]{
		Data: shelterEmployees,
	})
}
