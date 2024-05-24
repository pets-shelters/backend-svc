package responses

type UserInfo struct {
	ID         int64 `json:"id"`
	Registered bool  `json:"registered"`
	ShelterID  int64 `json:"shelter_id"`
}
