package shelters

import (
	"github.com/gin-gonic/gin"
	"github.com/pets-shelters/backend-svc/internal/controller/helpers"
	"github.com/pets-shelters/backend-svc/internal/structs/responses"
	"net/http"
)

func (r *routes) getNames(ctx *gin.Context) {
	filterName := ctx.Query("filter[name]")

	sheltersNames, err := r.useCase.GetNames(
		ctx.Request.Context(),
		filterName,
	)
	if err != nil {
		r.log.Error(err.Error(), "failed to process usecase - get shelters' names")
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, helpers.FormInternalError())
		return
	}

	ctx.JSON(http.StatusOK, helpers.JsonData[[]responses.ShelterName]{
		Data: sheltersNames,
	})
}
