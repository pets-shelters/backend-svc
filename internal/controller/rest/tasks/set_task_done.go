package tasks

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/pets-shelters/backend-svc/internal/controller/helpers"
	"github.com/pets-shelters/backend-svc/internal/exceptions"
	"github.com/pets-shelters/backend-svc/internal/structs/requests"
	"github.com/pkg/errors"
	"net/http"
	"strconv"
)

func (r *routes) setTaskDone(ctx *gin.Context) {
	var request helpers.JsonData[requests.SetTaskDone]
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

	taskIdString := ctx.Param("id")
	taskId, err := strconv.Atoi(taskIdString)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helpers.FormBadRequestError(err.Error()))
		return
	}

	err = r.useCase.SetExecution(
		ctx.Request.Context(),
		request.Data,
		int64(taskId),
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
		if errors.As(err, &exceptions.InvalidTaskExecutionDateException{}) {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, helpers.FormCustomError(helpers.InvalidTaskExecutionDate, ""))
			return
		}
		r.log.Error(err.Error(), "failed to process usecase - set task done")
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, helpers.FormInternalError())
		return
	}

	ctx.Status(http.StatusCreated)
}
