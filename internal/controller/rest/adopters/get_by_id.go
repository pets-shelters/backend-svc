package adopters

import (
	"github.com/gin-gonic/gin"
	"github.com/pets-shelters/backend-svc/internal/controller/helpers"
	"github.com/pets-shelters/backend-svc/internal/exceptions"
	"github.com/pets-shelters/backend-svc/internal/structs/responses"
	"github.com/pkg/errors"
	"net/http"
	"strconv"
)

func (r *routes) getById(ctx *gin.Context) {
	adopterIdString := ctx.Param("id")
	adopterId, err := strconv.Atoi(adopterIdString)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helpers.FormBadRequestError(err.Error()))
		return
	}

	userId, ok := ctx.Get(helpers.JwtIdCtx)
	if !ok {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, helpers.FormCustomError(helpers.Unauthorized, ""))
		return
	}

	adopter, err := r.useCase.GetById(
		ctx.Request.Context(),
		userId.(int64),
		int64(adopterId),
	)
	if err != nil {
		if errors.As(err, &exceptions.NotFoundException{}) {
			ctx.AbortWithStatusJSON(http.StatusNotFound, helpers.FormCustomError(helpers.EntityNotFound, ""))
			return
		}
		if errors.As(err, &exceptions.PermissionDeniedException{}) {
			ctx.AbortWithStatusJSON(http.StatusForbidden, helpers.FormCustomError(helpers.PermissionDenied, ""))
			return
		}
		r.log.Error(err.Error(), "failed to process usecase - get adopter")
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, helpers.FormInternalError())
		return
	}

	ctx.JSON(http.StatusOK, helpers.JsonData[responses.Adopter]{
		Data: *adopter,
	})
}
