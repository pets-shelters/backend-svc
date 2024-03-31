package files

import (
	"github.com/pets-shelters/backend-svc/internal/usecase"
)

type UseCase struct {
	repo             usecase.IDBRepo
	s3Provider       usecase.IS3Provider
	publicBucketName string
}

func NewUseCase(repo usecase.IDBRepo, s3Provider usecase.IS3Provider, publicReadBucketName string) *UseCase {
	return &UseCase{
		repo:             repo,
		s3Provider:       s3Provider,
		publicBucketName: publicReadBucketName,
	}
}
