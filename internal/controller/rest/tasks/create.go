package tasks

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/pets-shelters/backend-svc/internal/controller/helpers"
	"github.com/pets-shelters/backend-svc/internal/exceptions"
	"github.com/pets-shelters/backend-svc/internal/structs/requests"
	"github.com/pkg/errors"
	"net/http"
	"time"
)

func (r *routes) create(ctx *gin.Context) {
	var request helpers.JsonData[requests.CreateTask]
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
	if time.Time(request.Data.StartDate).After(time.Time(request.Data.EndDate)) {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helpers.FormBadRequestError("end_date must be greater or equal start_date"))
		return
	}

	userId, ok := ctx.Get(helpers.JwtIdCtx)
	if !ok {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, helpers.FormCustomError(helpers.Unauthorized, ""))
		return
	}

	err = r.useCase.Create(
		ctx.Request.Context(),
		request.Data,
		userId.(int64),
	)
	if err != nil {
		if errors.As(err, &exceptions.PermissionDeniedException{}) {
			ctx.AbortWithStatusJSON(http.StatusForbidden, helpers.FormCustomError(helpers.PermissionDenied, ""))
			return
		}
		if errors.As(err, &exceptions.NotFoundException{}) {
			ctx.AbortWithStatusJSON(http.StatusNotFound, helpers.FormCustomError(helpers.EntityNotFound, ""))
			return
		}
		r.log.Error(err.Error(), "failed to process usecase - create task")
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, helpers.FormInternalError())
		return
	}

	ctx.Status(http.StatusCreated)
}
