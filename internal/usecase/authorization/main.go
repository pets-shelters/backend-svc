package authorization

import (
	"github.com/pets-shelters/backend-svc/internal/usecase"
)

type UseCase struct {
	repo usecase.DBRepo
}

func NewUseCase(repo usecase.DBRepo) *UseCase {
	return &UseCase{
		repo: repo,
	}
}
