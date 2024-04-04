package requests

type CreateShelter struct {
	Name        string  `json:"name" validate:"required"`
	Logo        int64   `json:"logo" validate:"required"`
	PhoneNumber string  `json:"phone_number" validate:"required,len=12"`
	Instagram   *string `json:"instagram" validate:"omitempty,url"`
	Facebook    *string `json:"facebook" validate:"omitempty,url"`
}
