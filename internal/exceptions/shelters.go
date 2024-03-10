package exceptions

import "github.com/pkg/errors"

type UserHasShelterException struct {
	err error
}

func NewUserHasShelterException() UserHasShelterException {
	return UserHasShelterException{
		err: errors.New("user_has_shelter_error"),
	}
}
func (g UserHasShelterException) Error() string {
	return g.err.Error()
}
