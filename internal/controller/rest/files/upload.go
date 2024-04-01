package files

import (
	"github.com/gin-gonic/gin"
	"github.com/pets-shelters/backend-svc/internal/controller/helpers"
	"github.com/pets-shelters/backend-svc/internal/exceptions"
	"github.com/pets-shelters/backend-svc/internal/structs/responses"
	"github.com/pkg/errors"
	"net/http"
)

func (r *routes) upload(ctx *gin.Context) {
	var fileContent []byte
	err := ctx.Bind(&fileContent)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helpers.FormBadRequestError(err.Error()))
		return
	}

	userId, ok := ctx.Get(helpers.JwtIdCtx)
	if !ok {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, helpers.FormCustomError(helpers.Unauthorized, ""))
		return
	}

	fileType := ctx.ContentType()
	tempFileId, err := r.filesUseCase.Upload(
		ctx.Request.Context(),
		userId.(int64),
		fileContent,
		fileType,
	)
	if err != nil {
		if errors.As(err, &exceptions.FilesOverloadException{}) {
			ctx.AbortWithStatusJSON(http.StatusMethodNotAllowed, helpers.FormCustomError(helpers.FilesOverload, ""))
			return
		}
		if errors.As(err, &exceptions.InvalidFileTypeException{}) {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, helpers.FormBadRequestError("invalid file_type"))
			return
		}
		r.log.Error(err.Error(), "failed to process usecase - upload file")
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, helpers.FormInternalError())
		return
	}

	ctx.JSON(http.StatusCreated, helpers.JsonData[responses.UploadFile]{
		Data: responses.UploadFile{
			TemporaryFileID: tempFileId,
		},
	})
}
