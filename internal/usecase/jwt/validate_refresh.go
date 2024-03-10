package jwt

import (
	"github.com/golang-jwt/jwt"
	"github.com/pets-shelters/backend-svc/internal/exceptions"
	"github.com/pkg/errors"
	"time"
)

func (uc *UseCase) VerifyRefreshToken(refreshTokenString string) (string, error) {
	refreshToken, err := jwt.ParseWithClaims(refreshTokenString, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(uc.cfg.RefreshSecret), nil
	})
	if err != nil {
		return "", exceptions.NewInvalidJwtException()
	}

	claims, ok := refreshToken.Claims.(*jwt.StandardClaims)
	if !ok {
		return "", exceptions.NewInvalidJwtException()
	}
	if claims.ExpiresAt < time.Now().Unix() {
		return "", exceptions.NewInvalidJwtException()
	}

	return claims.Id, nil
}
