package oauth

import (
	"github.com/pets-shelters/backend-svc/configs"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"time"
)

const (
	callbackEndpoint = "/authorization/callback"
	googleEmailUrl   = "https://www.googleapis.com/auth/userinfo.email"
)

type OAuth struct {
	*oauth2.Config
	stateLifetime time.Duration
}

func NewOAuth(cfg configs.OAuth, httpAddr string) *OAuth {
	oauthCfg := &oauth2.Config{
		ClientID:     cfg.ClientID,
		ClientSecret: cfg.ClientSecret,
		RedirectURL:  httpAddr + callbackEndpoint,
		Endpoint:     google.Endpoint,
		Scopes:       []string{googleEmailUrl},
	}
	return &OAuth{
		oauthCfg,
		cfg.StateLifetime,
	}
}

func (oa *OAuth) GetStateLifetime() time.Duration {
	return oa.stateLifetime
}
