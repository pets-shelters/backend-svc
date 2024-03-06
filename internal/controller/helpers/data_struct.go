package helpers

type JsonData[T any] struct {
	Data T `json:"data" binding:"required"`
}
