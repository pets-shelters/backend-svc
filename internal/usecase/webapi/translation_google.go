package webapi

import (
	translator "github.com/Conight/go-googletrans"
)

// TranslationWebAPI -.
type TranslationWebAPI struct {
	conf translator.Config
}

// New -.
func New() *TranslationWebAPI {
	conf := translator.Config{
		UserAgent:   []string{"Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:15.0) Gecko/20100101 Firefox/15.0.1"},
		ServiceUrls: []string{"translate.google.com"},
	}

	return &TranslationWebAPI{
		conf: conf,
	}
}
