package helpers

import (
	"github.com/Masterminds/squirrel"
	"github.com/pets-shelters/backend-svc/internal/usecase/repo/entity"
)

func ApplyPagination(builder squirrel.SelectBuilder, orderByQuery string, pagination entity.Pagination) squirrel.SelectBuilder {
	offset := (pagination.Page - 1) * pagination.Limit
	return builder.
		OrderBy(orderByQuery).
		Offset(offset).
		Limit(pagination.Limit)
}
