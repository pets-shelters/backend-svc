package jwt

import (
	"github.com/golang-jwt/jwt"
	"github.com/pets-shelters/backend-svc/internal/structs"
	"github.com/pkg/errors"
	"time"
)

func (uc *UseCase) CreateTokensPair(userEmail string) (*structs.TokensPair, error) {
	accessToken, err := uc.createAccessToken(userEmail)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create access token")
	}

	refreshToken, err := uc.createRefreshToken(userEmail)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create refresh token")
	}

	return &structs.TokensPair{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (uc *UseCase) createAccessToken(userEmail string) (string, error) {
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.StandardClaims{
			Id:        userEmail,
			ExpiresAt: time.Now().Add(uc.cfg.AccessLifetime).Unix(),
		})
	accessTokenString, err := accessToken.SignedString([]byte(uc.cfg.AccessSecret))
	if err != nil {
		return "", errors.Wrap(err, "failed to sign token with string")
	}

	return accessTokenString, nil
}

func (uc *UseCase) createRefreshToken(userEmail string) (string, error) {
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.StandardClaims{
			Id:        userEmail,
			ExpiresAt: time.Now().Add(uc.cfg.RefreshLifetime).Unix(),
		})
	refreshTokenString, err := refreshToken.SignedString([]byte(uc.cfg.RefreshSecret))
	if err != nil {
		return "", errors.Wrap(err, "failed to sign token with string")
	}

	return refreshTokenString, nil
}
