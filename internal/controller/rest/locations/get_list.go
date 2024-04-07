package locations

import (
	"github.com/gin-gonic/gin"
	"github.com/pets-shelters/backend-svc/internal/controller/helpers"
	"github.com/pets-shelters/backend-svc/internal/structs/responses"
	"net/http"
	"strconv"
)

func (r *routes) getList(ctx *gin.Context) {
	shelterIdString := ctx.Param("id")
	shelterId, err := strconv.Atoi(shelterIdString)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helpers.FormBadRequestError(err.Error()))
		return
	}

	locations, err := r.useCase.GetList(
		ctx.Request.Context(),
		int64(shelterId),
	)
	if err != nil {
		r.log.Error(err.Error(), "failed to process usecase - get locations list")
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, helpers.FormInternalError())
		return
	}

	ctx.JSON(http.StatusOK, helpers.JsonData[[]responses.Location]{
		Data: locations,
	})
}
