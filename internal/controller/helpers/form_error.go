package helpers

type JsonError struct {
	Code   string `json:"code"`
	Detail string `json:"detail"`
}

const (
	UserAlreadyExists = "user_already_exists"
)

func FormCustomError(code string, detail string) JsonData[JsonError] {
	return JsonData[JsonError]{
		Data: JsonError{
			Code:   code,
			Detail: detail,
		},
	}
}

func FormBadRequestError(detail string) JsonData[JsonError] {
	return FormCustomError("bad_request", detail)
}

func FormInternalError(detail string) JsonData[JsonError] {
	return FormCustomError("internal", detail)
}
