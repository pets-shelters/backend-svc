package authorization

import (
	"github.com/pets-shelters/backend-svc/internal/usecase"
	"github.com/pets-shelters/backend-svc/internal/usecase/oauth"
	"github.com/pets-shelters/backend-svc/internal/usecase/redis"
)

type UseCase struct {
	repo  usecase.IDBRepo
	oauth oauth.OAuth
	cache redis.Redis
	jwt   usecase.IJwt
}

func NewUseCase(repo usecase.IDBRepo, oauth oauth.OAuth, cache redis.Redis, jwt usecase.IJwt) *UseCase {
	return &UseCase{
		repo:  repo,
		oauth: oauth,
		cache: cache,
		jwt:   jwt,
	}
}
