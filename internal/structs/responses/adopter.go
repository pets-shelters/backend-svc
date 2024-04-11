package responses

type Adopter struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
}
