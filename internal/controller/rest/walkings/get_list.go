package walkings

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/pets-shelters/backend-svc/internal/controller/helpers"
	"github.com/pets-shelters/backend-svc/internal/structs/requests"
	"github.com/pets-shelters/backend-svc/internal/structs/responses"
	"net/http"
)

type GetWalkingsListRequest struct {
	requests.WalkingsFilters
	*requests.Pagination
}

func (r *routes) getList(ctx *gin.Context) {
	var queryParams GetWalkingsListRequest
	err := ctx.BindQuery(&queryParams)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helpers.FormBadRequestError(err.Error()))
		return
	}
	err = validator.New().Struct(queryParams)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helpers.FormBadRequestError(err.Error()))
		return
	}

	userId, ok := ctx.Get(helpers.JwtIdCtx)
	if !ok {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, helpers.FormCustomError(helpers.Unauthorized, ""))
		return
	}

	walkings, paginationMetadata, err := r.useCase.GetList(
		ctx.Request.Context(),
		queryParams.WalkingsFilters,
		queryParams.Pagination,
		userId.(int64),
	)
	if err != nil {
		r.log.Error(err.Error(), "failed to process usecase - get walkings list")
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, helpers.FormInternalError())
		return
	}

	ctx.JSON(http.StatusOK, helpers.JsonData[[]responses.Walking]{
		Data:               walkings,
		PaginationMetadata: paginationMetadata,
	})
}
