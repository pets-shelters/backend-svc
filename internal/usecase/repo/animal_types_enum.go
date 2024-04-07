package repo

import (
	"context"
	"github.com/pets-shelters/backend-svc/internal/usecase"
	"github.com/pets-shelters/backend-svc/pkg/postgres"
	"github.com/pkg/errors"
	"log"
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

func (r *AnimalTypesEnumRepo) CreateWithConn(ctx context.Context, conn usecase.IConnection, newValue string) error {
	sql := "SELECT add_animal_type($1)"
	args := []interface{}{newValue}

	log.Printf("%+v", sql)
	log.Printf("%+v", args)
	_, err := conn.Exec(ctx, sql, args...)
	if err != nil {
		return errors.Wrap(err, "failed to Exec add animal_type value query")
	}

	return nil
}
