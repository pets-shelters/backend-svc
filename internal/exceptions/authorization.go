package exceptions

import "errors"

type UserExistsException struct {
	err error
}

func NewUserExistsException() UserExistsException {
	return UserExistsException{
		err: errors.New("user_exists_error"),
	}
}
func (g UserExistsException) Error() string {
	return g.err.Error()
}

type InvalidStateException struct {
	err error
}

func NewInvalidStateException() InvalidStateException {
	return InvalidStateException{
		err: errors.New("invalid_state_error"),
	}
}
func (g InvalidStateException) Error() string {
	return g.err.Error()
}
