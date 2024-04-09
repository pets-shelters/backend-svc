package exceptions

import "github.com/pkg/errors"

type LocationHasAnimalsException struct {
	err error
}

func NewLocationHaveAnimalsException() LocationHasAnimalsException {
	return LocationHasAnimalsException{
		err: errors.New("location_have_animals_error"),
	}
}
func (g LocationHasAnimalsException) Error() string {
	return g.err.Error()
}

type LocationNotFoundException struct {
	err error
}

func NewLocationNotFoundException() LocationNotFoundException {
	return LocationNotFoundException{
		err: errors.New("location_not_found_error"),
	}
}
func (g LocationNotFoundException) Error() string {
	return g.err.Error()
}
