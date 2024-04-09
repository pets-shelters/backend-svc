package responses

type Location struct {
	ID            int64  `json:"id"`
	City          string `json:"city"`
	Address       string `json:"address"`
	AnimalsNumber int64  `json:"animals_number"`
}
