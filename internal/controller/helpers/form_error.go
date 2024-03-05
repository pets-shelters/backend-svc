package helpers

type JsonError struct {
	Code   string `json:"code"`
	Detail string `json:"detail"`
}

const (
	UserAlreadyExists = "user_already_exists"
)

func FormCustomError(code string, detail string) JsonError {
	return JsonError{
		Code:   code,
		Detail: detail,
	}
}

func FormBadRequestError(detail string) JsonError {
	return JsonError{
		Code:   "bad_request",
		Detail: detail,
	}
}

func FormInternalError(detail string) JsonError {
	return JsonError{
		Code:   "internal",
		Detail: detail,
	}
}
