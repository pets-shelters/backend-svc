package locations

import (
	"github.com/gin-gonic/gin"
	"github.com/pets-shelters/backend-svc/internal/controller/helpers"
	"net/http"
)

func (r *routes) getCities(ctx *gin.Context) {
	cities, err := r.useCase.GetCities(
		ctx.Request.Context(),
	)
	if err != nil {
		r.log.Error(err.Error(), "failed to process usecase - get locations' cities list")
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, helpers.FormInternalError())
		return
	}

	ctx.JSON(http.StatusOK, helpers.JsonData[[]string]{
		Data: cities,
	})
}
