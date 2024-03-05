package helpers

import "go/types"

type JsonData[T types.Struct] struct {
	Data T `json:"data"`
}
