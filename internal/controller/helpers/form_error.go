package helpers

type JsonError struct {
	Code   string `json:"code"`
	Detail string `json:"detail"`
}

const (
	UserAlreadyExists           = "user_already_exists"
	UserAlreadyHasShelter       = "user_already_has_shelter"
	FilesOverload               = "files_overload"
	InvalidFileType             = "invalid_filetype"
	FileNotFound                = "file_not_found"
	LocationNotFound            = "location_not_found"
	Unauthorized                = "unauthorized"
	PermissionDenied            = "permission_denied"
	EntityNotFound              = "entity_not_found"
	LocationHasAnimals          = "location_has_animals"
	AdopterAlreadyExists        = "adopter_already_exists"
	InvalidTaskExecutionDate    = "invalid_task_execution_date"
	AnimalUnavailableForWalking = "animal_unavailable_for_walking"
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

func FormInternalError() JsonData[JsonError] {
	return FormCustomError("internal", "")
}
