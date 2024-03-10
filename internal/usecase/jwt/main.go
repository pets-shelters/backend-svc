package jwt

import (
	"github.com/pets-shelters/backend-svc/configs"
)

type UseCase struct {
	cfg configs.Jwt
}

func NewUseCase(cfg configs.Jwt) *UseCase {
	return &UseCase{cfg}
}
