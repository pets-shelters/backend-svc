package entity

import (
	"database/sql"
)

type UserRole string

const (
	EmployeeUserRole UserRole = "employee"
	ManagerUserRole  UserRole = "manager"
)

type User struct {
	ID        int64         `db:"id" structs:"-" json:"id"`
	Email     string        `db:"email" structs:"email" json:"email"`
	ShelterID sql.NullInt64 `db:"shelter_id" structs:"shelter_id" json:"shelter_id"`
	Role      UserRole      `db:"role" structs:"role" json:"role"`
}
