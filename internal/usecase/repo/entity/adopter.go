package entity

type Adopter struct {
	ID          int64  `db:"id" structs:"-"`
	Name        string `db:"name" structs:"name"`
	PhoneNumber string `db:"phone_number" structs:"phone_number"`
}

type AdoptersFilters struct {
	PhoneNumber *string
}
