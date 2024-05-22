package animals

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
	var request helpers.JsonData[requests.CreateAnimal]
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

	userId, ok := ctx.Get(helpers.JwtIdCtx)
	if !ok {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, helpers.FormCustomError(helpers.Unauthorized, ""))
		return
	}

	id, err := r.useCase.Create(
		ctx.Request.Context(),
		request.Data,
		userId.(int64),
	)
	if err != nil {
		if errors.As(err, &exceptions.PermissionDeniedException{}) {
			ctx.AbortWithStatusJSON(http.StatusForbidden, helpers.FormCustomError(helpers.PermissionDenied, ""))
			return
		}
		if errors.As(err, &exceptions.LocationNotFoundException{}) {
			ctx.AbortWithStatusJSON(http.StatusNotFound, helpers.FormCustomError(helpers.LocationNotFound, ""))
			return
		}
		if errors.As(err, &exceptions.FileNotFoundException{}) {
			ctx.AbortWithStatusJSON(http.StatusNotFound, helpers.FormCustomError(helpers.FileNotFound, ""))
			return
		}
		r.log.Error(err.Error(), "failed to process usecase - create animal")
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, helpers.FormInternalError())
		return
	}

	ctx.JSON(http.StatusCreated, helpers.JsonData[responses.CreatedID]{
		Data: responses.CreatedID{
			ID: id,
		},
	})
}
