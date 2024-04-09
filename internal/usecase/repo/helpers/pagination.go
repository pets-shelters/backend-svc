package helpers

import (
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/pets-shelters/backend-svc/internal/usecase/repo/entity"
)

func ApplyPagination(builder squirrel.SelectBuilder, tableName string, pagination entity.Pagination) squirrel.SelectBuilder {
	offset := (pagination.Page - 1) * pagination.Limit
	return builder.
		OrderBy(fmt.Sprintf("%s.id", tableName)).
		Offset(offset).
		Limit(pagination.Limit)
}
