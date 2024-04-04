package responses

import "time"

type Shelter struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	Logo        string    `json:"logo"`
	PhoneNumber string    `json:"phone_number"`
	Instagram   *string   `json:"instagram,omitempty"`
	Facebook    *string   `json:"facebook,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
}
