package helpers

import "github.com/pets-shelters/backend-svc/configs"

// RouterConfigs TODO consists of small routes structs
type RouterConfigs struct {
	LoginCookieLifetime  int
	AccessTokenLifetime  int
	RefreshTokenLifetime int
	Domain               string
	TemporaryFilesCfg    configs.TemporaryFiles
}
