package locations

import "github.com/pets-shelters/backend-svc/internal/usecase"

type UseCase struct {
	repo usecase.IDBRepo
}

func NewUseCase(repo usecase.IDBRepo) *UseCase {
	return &UseCase{
		repo: repo,
	}
}
