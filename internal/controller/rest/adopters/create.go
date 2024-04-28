package adopters

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/pets-shelters/backend-svc/internal/controller/helpers"
	"github.com/pets-shelters/backend-svc/internal/exceptions"
	"github.com/pets-shelters/backend-svc/internal/structs/requests"
	"github.com/pets-shelters/backend-svc/internal/structs/responses"
	"github.com/pkg/errors"
	"net/http"
)

func (r *routes) create(ctx *gin.Context) {
	var request helpers.JsonData[requests.CreateAdopter]
	err := ctx.BindJSON(&request)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helpers.FormBadRequestError(err.Error()))
		return
	}
	err = validator.New().Struct(request)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helpers.FormBadRequestError(err.Error()))
		return
	}

	id, err := r.useCase.Create(
		ctx.Request.Context(),
		request.Data,
	)
	if err != nil {
		if errors.As(err, &exceptions.AdopterExistsException{}) {
			ctx.AbortWithStatusJSON(http.StatusConflict, helpers.FormCustomError(helpers.AdopterAlreadyExists, ""))
			return
		}
		r.log.Error(err.Error(), "failed to process usecase - create adopter")
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, helpers.FormInternalError())
		return
	}

	ctx.JSON(http.StatusCreated, helpers.JsonData[responses.CreatedID]{
		Data: responses.CreatedID{
			ID: id,
		},
	})
}
