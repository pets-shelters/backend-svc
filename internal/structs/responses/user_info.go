package responses

import (
	"encoding/json"
	"github.com/pets-shelters/backend-svc/internal/structs"
)

type UserInfo struct {
	ID          int64            `json:"id"`
	Email       string           `json:"email"`
	Role        structs.UserRole `json:"role"`
	Registered  bool             `json:"registered"`
	ShelterID   int64            `json:"shelter_id"`
	ShelterName string           `json:"shelter_name"`
}

func (ui UserInfo) MarshalBinary() ([]byte, error) {
	return json.Marshal(ui)
}
