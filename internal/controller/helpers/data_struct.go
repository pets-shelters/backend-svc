package helpers

import "github.com/pets-shelters/backend-svc/internal/structs/responses"

type JsonData[T any] struct {
	Data               T                             `json:"data" binding:"required"`
	PaginationMetadata *responses.PaginationMetadata `json:"pagination_metadata,omitempty" binding:"-"`
}
