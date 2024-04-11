package requests

type UpdateAnimal struct {
	LocationID         *int64  `json:"location_id"`
	Photo              *int64  `json:"photo"`
	Sterilized         *bool   `json:"sterilized"`
	AdopterID          *int64  `json:"adopter_id"`
	PublicDescription  *string `json:"public_description"`
	PrivateDescription *string `json:"private_description"`
}
