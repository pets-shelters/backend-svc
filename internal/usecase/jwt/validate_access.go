package jwt

import (
	"github.com/golang-jwt/jwt"
	"github.com/pets-shelters/backend-svc/internal/exceptions"
	"github.com/pkg/errors"
	"time"
)

func (uc *UseCase) VerifyAccessToken(accessTokenString string) (string, error) {
	accessToken, err := jwt.ParseWithClaims(accessTokenString, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(uc.cfg.AccessSecret), nil
	})
	if err != nil {
		return "", exceptions.NewInvalidJwtException()
	}

	claims, ok := accessToken.Claims.(*jwt.StandardClaims)
	if !ok {
		return "", exceptions.NewInvalidJwtException()
	}
	if claims.ExpiresAt < time.Now().Unix() {
		return "", exceptions.NewInvalidJwtException()
	}

	return claims.Id, nil
}
