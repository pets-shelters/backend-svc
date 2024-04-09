package requests

type Pagination struct {
	Page  *int64 `form:"page" validate:"required_with=Limit"`
	Limit *int64 `form:"limit" validate:"required_with=Page"`
}
