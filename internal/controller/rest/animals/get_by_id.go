package animals

import (
	"github.com/gin-gonic/gin"
	"github.com/pets-shelters/backend-svc/internal/controller/helpers"
	"github.com/pets-shelters/backend-svc/internal/exceptions"
	"github.com/pets-shelters/backend-svc/internal/structs/responses"
	"github.com/pkg/errors"
	"net/http"
	"strconv"
)

func (r *routes) getById(ctx *gin.Context) {
	animalIdString := ctx.Param("id")
	animalId, err := strconv.Atoi(animalIdString)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helpers.FormBadRequestError(err.Error()))
		return
	}

	var userId *int64
	accessToken, _ := ctx.Cookie(helpers.AccessTokenCookieName)
	if accessToken != "" {
		userIdString, err := r.jwtUseCase.VerifyAccessToken(accessToken)
		if err != nil && !errors.As(err, &exceptions.InvalidJwtException{}) {
			r.log.Error(err.Error(), "failed to verify access token - get animal")
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, helpers.FormInternalError())
			return
		}
		userIdInt, _ := strconv.Atoi(userIdString)
		userIdInt64 := int64(userIdInt)
		userId = &userIdInt64
	}

	animal, err := r.useCase.GetById(
		ctx.Request.Context(),
		int64(animalId),
		userId,
	)
	if err != nil {
		if errors.As(err, &exceptions.NotFoundException{}) {
			ctx.AbortWithStatusJSON(http.StatusNotFound, helpers.FormCustomError(helpers.EntityNotFound, ""))
			return
		}
		r.log.Error(err.Error(), "failed to process usecase - get animal")
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, helpers.FormInternalError())
		return
	}

	ctx.JSON(http.StatusOK, helpers.JsonData[responses.Animal]{
		Data: *animal,
	})
}
