package tasks

import (
	"github.com/gin-gonic/gin"
	"github.com/pets-shelters/backend-svc/internal/controller/helpers"
	"github.com/pets-shelters/backend-svc/internal/structs/requests"
	"github.com/pets-shelters/backend-svc/internal/structs/responses"
	"github.com/pets-shelters/backend-svc/pkg/bind"
	"net/http"
)

func (r *routes) getListWithExecutions(ctx *gin.Context) {
	var filters requests.TasksWithExecutionsFilters
	err := bind.BindQueryWithSlices[requests.TasksWithExecutionsFilters](ctx, &filters)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helpers.FormBadRequestError(err.Error()))
		return
	}

	userId, ok := ctx.Get(helpers.JwtIdCtx)
	if !ok {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, helpers.FormCustomError(helpers.Unauthorized, ""))
		return
	}

	tasksWithExecutions, err := r.useCase.GetListWithExecutions(
		ctx.Request.Context(),
		userId.(int64),
		filters,
	)
	if err != nil {
		r.log.Error(err.Error(), "failed to process usecase - get tasks with executions")
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, helpers.FormInternalError())
		return
	}

	ctx.JSON(http.StatusOK, helpers.JsonData[[]responses.TaskWithExecutions]{
		Data: tasksWithExecutions,
	})
}
