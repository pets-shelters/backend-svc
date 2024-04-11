package animals

import (
	"github.com/gin-gonic/gin"
	"github.com/pets-shelters/backend-svc/internal/controller/helpers"
	"github.com/pets-shelters/backend-svc/internal/structs/responses"
	"net/http"
)

func (r *routes) getTypes(ctx *gin.Context) {
	animalTypes, err := r.useCase.GetTypes(
		ctx.Request.Context(),
	)
	if err != nil {
		r.log.Error(err.Error(), "failed to process usecase - get animals' types list")
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, helpers.FormInternalError())
		return
	}

	ctx.JSON(http.StatusOK, helpers.JsonData[responses.AnimalTypes]{
		Data: *animalTypes,
	})
}
