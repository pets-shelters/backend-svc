package authorization

import (
	"crypto/rand"
	"encoding/base64"
	"github.com/pets-shelters/backend-svc/internal/structs"
	"github.com/pkg/errors"
)

func (uc *UseCase) Login() (*structs.LoginResult, error) {
	state := generateState()
	cookieSession := generateCookieSession()
	err := uc.cache.Set(cookieSession, state, uc.oauth.GetStateLifetime()).Err()
	if err != nil {
		return nil, errors.Wrap(err, "failed to set session to cache")
	}

	authUrl := uc.oauth.AuthCodeURL(state)

	return &structs.LoginResult{
		CookieSession: cookieSession,
		Url:           authUrl,
	}, nil
}

func generateState() string {
	b := make([]byte, 64)
	rand.Read(b)
	state := base64.URLEncoding.EncodeToString(b)

	return state
}

func generateCookieSession() string {
	b := make([]byte, 32)
	rand.Read(b)
	session := base64.URLEncoding.EncodeToString(b)

	return session
}
