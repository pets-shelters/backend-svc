package requests

type CreateShelter struct {
	Name        string `json:"name" binding:"required"`
	Logo        int64  `json:"logo" binding:"required"`
	City        string `json:"city" binding:"required"`
	PhoneNumber string `json:"phone_number" binding:"required,len=12"`
	Instagram   string `json:"instagram,omitempty" binding:"-" validate:"omitempty,url"`
	Facebook    string `json:"facebook,omitempty" binding:"-" validate:"omitempty,url"`
}
