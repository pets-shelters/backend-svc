package shelters

import (
	"github.com/pets-shelters/backend-svc/internal/usecase"
)

type UseCase struct {
	repo       usecase.IDBRepo
	s3Endpoint string
}

func NewUseCase(repo usecase.IDBRepo, s3Endpoint string) *UseCase {
	return &UseCase{
		repo:       repo,
		s3Endpoint: s3Endpoint,
	}
}
