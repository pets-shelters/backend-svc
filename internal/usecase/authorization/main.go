package authorization

import (
	"github.com/pets-shelters/backend-svc/internal/usecase"
	"github.com/pets-shelters/backend-svc/internal/usecase/oauth"
)

type UseCase struct {
	repo  usecase.IDBRepo
	oauth oauth.OAuth
	cache usecase.IRedis
	jwt   usecase.IJwt
}

func NewUseCase(repo usecase.IDBRepo, oauth oauth.OAuth, cache usecase.IRedis, jwt usecase.IJwt) *UseCase {
	return &UseCase{
		repo:  repo,
		oauth: oauth,
		cache: cache,
		jwt:   jwt,
	}
}
