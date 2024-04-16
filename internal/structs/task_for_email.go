package structs

type TaskForEmail struct {
	Description string `json:"description"`
	AnimalType  string `json:"animal_type"`
	AnimalName  string `json:"animal_name"`
	Time        string `json:"time"`
}
