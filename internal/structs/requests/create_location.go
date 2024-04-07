package requests

type CreateLocation struct {
	City    string `json:"city" validate:"required"`
	Address string `json:"address" validate:"required"`
}
