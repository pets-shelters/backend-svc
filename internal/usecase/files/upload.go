package files

import (
	"bytes"
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"github.com/pets-shelters/backend-svc/internal/exceptions"
	"github.com/pets-shelters/backend-svc/internal/usecase/repo/entity"
	"github.com/pets-shelters/backend-svc/internal/usecase/s3"
	"github.com/pkg/errors"
	"time"
)

func (uc *UseCase) Upload(ctx context.Context, userId int64, fileContent []byte, fileType string) (int64, error) {
	tempFilesNumber, err := uc.repo.GetTemporaryFilesRepo().CountForUserId(ctx, userId)
	if err != nil {
		return 0, errors.Wrap(err, "failed to get temporary files number for user")
	}
	if tempFilesNumber >= 10 {
		return 0, exceptions.NewFilesOverloadException()
	}

	fileExtension := s3.ContentTypes[fileType]
	if fileExtension == "" {
		return 0, exceptions.NewInvalidFileTypeException()
	}

	var tempFileId int64

	filePath := fmt.Sprintf("/%d/%d.%s", userId, time.Now().Unix(), fileExtension)
	err = uc.repo.Transaction(ctx, func(tx pgx.Tx) error {
		fileId, err := uc.repo.GetFilesRepo().CreateWithConn(ctx, tx, entity.File{
			Bucket: uc.publicBucketName,
			Path:   filePath,
		})
		if err != nil {
			return errors.Wrap(err, "failed to create file entity")
		}

		tempFileId, err = uc.repo.GetTemporaryFilesRepo().CreateWithConn(ctx, tx, entity.TemporaryFile{
			FileID:    fileId,
			UserID:    userId,
			CreatedAt: time.Now().UTC(),
		})
		if err != nil {
			return errors.Wrap(err, "failed to create temporary_file entity")
		}

		fileBytesReader := bytes.NewReader(fileContent)
		err = uc.s3Provider.UploadFile(ctx, fileBytesReader, uc.publicBucketName, filePath, fileType)
		if err != nil {
			return errors.Wrap(err, "failed to upload file to s3")
		}

		return nil
	})
	if err != nil {
		return 0, errors.Wrap(err, "failed to process transaction")
	}

	return tempFileId, nil
}
