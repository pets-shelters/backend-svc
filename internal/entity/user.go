package entity

type UserRole string

const (
	EmployeeUserRole UserRole = "employee"
	ManagerUserRole  UserRole = "manager"
)

type User struct {
	ID        int64    `db:"id" json:"id"`
	Email     string   `db:"email" json:"email"`
	ShelterID int64    `db:"shelter_id" json:"shelter_id"`
	Role      UserRole `db:"role" json:"role"`
}
