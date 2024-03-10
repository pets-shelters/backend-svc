package requests

type CreateShelter struct {
	Name        string `json:"name" binding:"required"`
	Logo        string `json:"logo" binding:"required"`
	City        string `json:"city" binding:"required"`
	PhoneNumber string `json:"phone_number" binding:"required,len=12"`
	Instagram   string `json:"instagram" binding:"-"`
	Facebook    string `json:"facebook" binding:"-"`
}
