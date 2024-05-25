package entity

import (
	"database/sql"
	"github.com/pets-shelters/backend-svc/internal/structs"
)

type User struct {
	ID        int64            `db:"id" structs:"-"`
	Email     string           `db:"email" structs:"email"`
	ShelterID sql.NullInt64    `db:"shelter_id" structs:"shelter_id"`
	Role      structs.UserRole `db:"role" structs:"role"`
}

type UserWithShelterName struct {
	User
	ShelterName sql.NullString `db:"name" structs:"name"`
}

type UsersFilters struct {
	Email     *string
	ShelterID *int64
}
