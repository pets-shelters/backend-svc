package responses

type Location struct {
	ID            int64  `json:"id"`
	City          string `json:"city"`
	Address       string `json:"address"`
	ShelterID     int64  `json:"shelter_id"`
	AnimalsNumber int64  `json:"animals_number"`
}
