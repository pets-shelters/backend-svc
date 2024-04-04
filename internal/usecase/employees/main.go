package employees

import "github.com/pets-shelters/backend-svc/internal/usecase"

type UseCase struct {
	repo           usecase.IDBRepo
	emailsProvider usecase.IEmailsProvider
}

func NewUseCase(repo usecase.IDBRepo, emailsProvider usecase.IEmailsProvider) *UseCase {
	return &UseCase{
		repo:           repo,
		emailsProvider: emailsProvider,
	}
}
