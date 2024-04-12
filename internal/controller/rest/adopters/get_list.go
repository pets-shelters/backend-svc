package adopters

import (
	"github.com/gin-gonic/gin"
	"github.com/pets-shelters/backend-svc/internal/controller/helpers"
	"github.com/pets-shelters/backend-svc/internal/structs/responses"
	"net/http"
)

func (r *routes) getList(ctx *gin.Context) {
	filterPhoneNumber := ctx.Query("filter[phone_number]")

	adopters, err := r.useCase.GetList(
		ctx.Request.Context(),
		filterPhoneNumber,
	)
	if err != nil {
		r.log.Error(err.Error(), "failed to process usecase - get adopters list")
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, helpers.FormInternalError())
		return
	}

	ctx.JSON(http.StatusOK, helpers.JsonData[[]responses.Adopter]{
		Data: adopters,
	})
}
