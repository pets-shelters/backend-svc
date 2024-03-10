package exceptions

import "github.com/pkg/errors"

type InvalidJwtException struct {
	err error
}

func NewInvalidJwtException() InvalidJwtException {
	return InvalidJwtException{
		err: errors.New("invalid_jwt_error"),
	}
}
func (g InvalidJwtException) Error() string {
	return g.err.Error()
}
