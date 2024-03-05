package exceptions

import "errors"

type UserExistsException struct {
	err error
}

func NewUserExistsException(msg string) UserExistsException {
	return UserExistsException{
		err: errors.New(msg),
	}
}

func (g UserExistsException) Error() string {
	return g.err.Error()
}
