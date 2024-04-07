package entity

type Location struct {
	ID        int64  `db:"id" structs:"-"`
	City      string `db:"city" structs:"city"`
	Address   string `db:"address" structs:"address"`
	ShelterID int64  `db:"shelter_id" structs:"shelter_id"`
}

type LocationsAnimals struct {
	ID            int64
	City          string
	Address       string
	ShelterID     int64
	AnimalsNumber int64
}
