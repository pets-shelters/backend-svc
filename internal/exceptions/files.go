package exceptions

import "github.com/pkg/errors"

type FilesOverloadException struct {
	err error
}

func NewFilesOverloadException() FilesOverloadException {
	return FilesOverloadException{
		err: errors.New("files_overload_error"),
	}
}
func (g FilesOverloadException) Error() string {
	return g.err.Error()
}

type InvalidFileTypeException struct {
	err error
}

func NewInvalidFileTypeException() InvalidFileTypeException {
	return InvalidFileTypeException{
		err: errors.New("invalid_filetype_error"),
	}
}
func (g InvalidFileTypeException) Error() string {
	return g.err.Error()
}

type FileNotFoundException struct {
	err error
}

func NewFileNotFoundException() FileNotFoundException {
	return FileNotFoundException{
		err: errors.New("file_not_found_error"),
	}
}
func (g FileNotFoundException) Error() string {
	return g.err.Error()
}
