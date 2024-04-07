package repo

import (
	"context"
	"github.com/pets-shelters/backend-svc/internal/usecase"
	"github.com/pets-shelters/backend-svc/internal/usecase/repo/entity"
	"github.com/pets-shelters/backend-svc/pkg/postgres"
	"github.com/pkg/errors"
)

const (
	animalsTableName = "animals"
)

type AnimalsRepo struct {
	*postgres.Postgres
}

func NewAnimalsRepo(pg *postgres.Postgres) *AnimalsRepo {
	return &AnimalsRepo{pg}
}

func (r *AnimalsRepo) CreateWithConn(ctx context.Context, conn usecase.IConnection, animal entity.Animal) (int64, error) {
	sql, args, err := r.Builder.
		Insert(animalsTableName).
		Columns("location_id", "photo", "name", "birth_date", "type", "gender", "sterilized",
			"private_description", "public_description").
		Values(animal.LocationID, animal.Photo, animal.Name, animal.BirthDate, animal.Type, animal.Gender, animal.Sterilized,
			animal.PrivateDescription, animal.PublicDescription).
		Suffix("returning id").
		ToSql()
	if err != nil {
		return 0, errors.Wrap(err, "failed to build animal insert query")
	}

	var id int64
	err = conn.QueryRow(ctx, sql, args...).Scan(&id)
	if err != nil {
		return 0, errors.Wrap(err, "failed to QueryRow animal insert query")
	}

	return id, nil
}

func (r *AnimalsRepo) Create(ctx context.Context, animal entity.Animal) (int64, error) {
	return r.CreateWithConn(ctx, r.Pool, animal)
}
