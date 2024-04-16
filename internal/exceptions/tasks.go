package exceptions

import "github.com/pkg/errors"

type InvalidTaskExecutionDateException struct {
	err error
}

func NewInvalidTaskExecutionDateException() InvalidTaskExecutionDateException {
	return InvalidTaskExecutionDateException{
		err: errors.New("invalid_task_execution_date_error"),
	}
}
func (g InvalidTaskExecutionDateException) Error() string {
	return g.err.Error()
}
