package requests

type CreateEmployee struct {
	Email string `json:"email" validate:"required,email"`
}
