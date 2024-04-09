package animals

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/pets-shelters/backend-svc/internal/controller/helpers"
	"github.com/pets-shelters/backend-svc/internal/structs/requests"
	"github.com/pets-shelters/backend-svc/internal/structs/responses"
	"github.com/pets-shelters/backend-svc/pkg/bind"
	"log"
	"net/http"
)

type GetAnimalsListRequest struct {
	requests.AnimalsFilters
	*requests.Pagination
}

func (r *routes) getList(ctx *gin.Context) {
	var queryParams GetAnimalsListRequest
	err := bind.BindQueryWithSlices[GetAnimalsListRequest](ctx, &queryParams)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helpers.FormBadRequestError(err.Error()))
		return
	}
	log.Printf("%+v", queryParams.Pagination)
	log.Printf("%+v", validator.New().Struct(queryParams.Pagination))
	err = validator.New().Struct(queryParams)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helpers.FormBadRequestError(err.Error()))
		return
	}

	animals, paginationMetadata, err := r.useCase.GetList(
		ctx.Request.Context(),
		queryParams.AnimalsFilters,
		queryParams.Pagination,
	)
	if err != nil {
		r.log.Error(err.Error(), "failed to process usecase - get animals list")
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, helpers.FormInternalError())
		return
	}

	ctx.JSON(http.StatusOK, helpers.JsonData[[]responses.Animal]{
		Data:               animals,
		PaginationMetadata: paginationMetadata,
	})
}
