package exceptions

import "github.com/pkg/errors"

type AnimalUnavailableForWalkingException struct {
	err error
}

func NewAnimalUnavailableForWalkingException() AnimalUnavailableForWalkingException {
	return AnimalUnavailableForWalkingException{
		err: errors.New("animal_unavailable_for_walking_error"),
	}
}
func (g AnimalUnavailableForWalkingException) Error() string {
	return g.err.Error()
}
