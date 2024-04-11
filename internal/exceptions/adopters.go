package exceptions

import "github.com/pkg/errors"

type AdopterExistsException struct {
	err error
}

func NewAdopterExistsException() AdopterExistsException {
	return AdopterExistsException{
		err: errors.New("adopter_exists_error"),
	}
}
func (g AdopterExistsException) Error() string {
	return g.err.Error()
}
