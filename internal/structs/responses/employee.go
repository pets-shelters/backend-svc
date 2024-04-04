package responses

import (
	"github.com/pets-shelters/backend-svc/internal/structs"
)

type Employee struct {
	ID    int64            `json:"id"`
	Email string           `json:"email"`
	Role  structs.UserRole `json:"role"`
}
