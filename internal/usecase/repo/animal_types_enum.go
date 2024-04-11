package repo

import (
	"context"
	"fmt"
	"github.com/pets-shelters/backend-svc/pkg/postgres"
	"github.com/pkg/errors"
)

const (
	animalTypesEnumName = "animal_type"
)

type AnimalTypesEnumRepo struct {
	*postgres.Postgres
}

func NewAnimalTypesEnumRepo(pg *postgres.Postgres) *AnimalTypesEnumRepo {
	return &AnimalTypesEnumRepo{pg}
}

func (r *AnimalTypesEnumRepo) Create(ctx context.Context, newValue string) error {
	sql := "SELECT add_animal_type($1)"
	args := []interface{}{newValue}

	_, err := r.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return errors.Wrap(err, "failed to Exec add animal_type value query")
	}

	return nil
}

func (r *AnimalTypesEnumRepo) Select(ctx context.Context) ([]string, error) {
	sql, args, err := r.Builder.
		Select(fmt.Sprintf("unnest(enum_range(null::%s));", animalTypesEnumName)).
		ToSql()
	if err != nil {
		return nil, errors.Wrap(err, "failed to build select animal_type query")
	}

	rows, err := r.Pool.Query(ctx, sql, args...)
	if err != nil {
		return nil, errors.Wrap(err, "failed to QueryRow select animal_type query")
	}

	var result []string
	defer rows.Close()
	for rows.Next() {
		var row string
		err = rows.Scan(&row)
		if err != nil {
			return nil, errors.Wrap(err, "failed to scan animal_type")
		}

		result = append(result, row)
	}

	return result, nil
}
