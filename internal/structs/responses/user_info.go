package responses

import "encoding/json"

type UserInfo struct {
	ID          int64  `json:"id"`
	Registered  bool   `json:"registered"`
	ShelterID   int64  `json:"shelter_id"`
	ShelterName string `json:"shelter_name"`
}

func (ui UserInfo) MarshalBinary() ([]byte, error) {
	return json.Marshal(ui)
}
