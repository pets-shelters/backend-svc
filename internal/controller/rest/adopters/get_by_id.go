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

	adopter, err := r.useCase.GetById(
		ctx.Request.Context(),
		int64(adopterId),
	)
	if err != nil {
		if errors.As(err, &exceptions.NotFoundException{}) {
			ctx.AbortWithStatusJSON(http.StatusNotFound, helpers.FormCustomError(helpers.EntityNotFound, ""))
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
