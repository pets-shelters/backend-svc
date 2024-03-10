package shelters

import (
	"github.com/pets-shelters/backend-svc/internal/controller/helpers"
	"github.com/pets-shelters/backend-svc/internal/exceptions"
	"github.com/pets-shelters/backend-svc/internal/structs/requests"
	"github.com/pkg/errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (r *routes) create(ctx *gin.Context) {
	var request helpers.JsonData[requests.CreateShelter]
	err := ctx.BindJSON(&request)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helpers.FormBadRequestError(err.Error()))
		return
	}

	userEmail, ok := ctx.Get(helpers.JwtEmailCtx)
	if !ok {
		ctx.AbortWithStatusJSON(http.StatusForbidden, helpers.FormCustomError("code", err.Error()))
		return
	}
	err = r.useCase.Create(
		ctx.Request.Context(),
		request.Data,
		userEmail.(string),
	)
	if err != nil {
		if errors.As(err, &exceptions.UserHasShelterException{}) {
			ctx.AbortWithStatusJSON(http.StatusConflict, helpers.FormCustomError(helpers.UserAlreadyHasShelter, ""))
			return
		}
		r.log.Error(err, "failed to process usecase - create shelter")
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, helpers.FormInternalError(err.Error()))
		return
	}

	ctx.Status(http.StatusCreated)
}
