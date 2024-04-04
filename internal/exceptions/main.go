package exceptions

import "github.com/pkg/errors"

type NotFoundException struct {
	err error
}

func NewNotFoundException() NotFoundException {
	return NotFoundException{
		err: errors.New("not_found_error"),
	}
}
func (g NotFoundException) Error() string {
	return g.err.Error()
}
