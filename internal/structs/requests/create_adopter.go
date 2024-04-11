package requests

type CreateAdopter struct {
	Name        string `json:"name" validate:"required"`
	PhoneNumber string `json:"phone_number" validate:"required,len=12"`
}
