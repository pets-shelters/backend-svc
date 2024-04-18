package requests

type UpdateShelter struct {
	Name        *string `json:"name"`
	Logo        *int64  `json:"logo"`
	PhoneNumber *string `json:"phone_number" validate:"omitempty,len=13"`
	Instagram   *string `json:"instagram" validate:"omitempty,url"`
	Facebook    *string `json:"facebook" validate:"omitempty,url"`
}
