package animals

import (
	"github.com/gin-gonic/gin"
	"github.com/pets-shelters/backend-svc/internal/controller/helpers"
	"github.com/pets-shelters/backend-svc/internal/exceptions"
	"github.com/pkg/errors"
	"net/http"
	"strconv"
)

func (r *routes) delete(ctx *gin.Context) {
	animalIdString := ctx.Param("id")
	animalId, err := strconv.Atoi(animalIdString)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helpers.FormBadRequestError(err.Error()))
		return
	}

	userId, ok := ctx.Get(helpers.JwtIdCtx)
	if !ok {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, helpers.FormCustomError(helpers.Unauthorized, ""))
		return
	}

	err = r.useCase.Delete(
		ctx.Request.Context(),
		userId.(int64),
		int64(animalId),
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
		r.log.Error(err.Error(), "failed to process usecase - delete animal")
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, helpers.FormInternalError())
		return
	}

	ctx.Status(http.StatusOK)
}
