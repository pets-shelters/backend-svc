package repo

import (
	"context"
	"github.com/fatih/structs"
	"github.com/pets-shelters/backend-svc/internal/entity"
	"github.com/pets-shelters/backend-svc/pkg/postgres"
	"github.com/pkg/errors"
)

const sheltersTableName = "shelters"

type SheltersRepo struct {
	*postgres.Postgres
}

func NewSheltersRepo(pg *postgres.Postgres) *SheltersRepo {
	return &SheltersRepo{pg}
}

func (r *SheltersRepo) CreateWithConn(ctx context.Context, conn Connection, shelter entity.Shelter) (int64, error) {
	sql, args, err := r.Builder.
		Insert(sheltersTableName).
		SetMap(structs.Map(shelter)).
		Suffix("returning id").
		ToSql()
	if err != nil {
		return 0, errors.Wrap(err, "failed to build shelter insert query")
	}

	var id int64
	err = conn.QueryRow(ctx, sql, args...).Scan(&id)
	if err != nil {
		return 0, errors.Wrap(err, "failed to QueryRow shelter insert query")
	}

	return id, nil
}

func (r *SheltersRepo) Create(ctx context.Context, shelter entity.Shelter) (int64, error) {
	return r.CreateWithConn(ctx, r.Pool, shelter)
}

func (r *SheltersRepo) SelectSheltersWithConn(ctx context.Context, conn Connection) ([]entity.Shelter, error) {
	sql, _, err := r.Builder.
		Select("*").
		From(sheltersTableName).
		ToSql()
	if err != nil {
		return nil, errors.Wrap(err, "failed to build select shelters query")
	}

	var shelters []entity.Shelter
	rows, err := conn.Query(ctx, sql)
	if err != nil {
		return nil, errors.Wrap(err, "failed to Query select shelters query")
	}

	err = rows.Scan(&shelters)
	if err != nil {
		return nil, errors.Wrap(err, "failed to scan shelters")
	}

	return shelters, nil
}

func (r *SheltersRepo) SelectShelters(ctx context.Context) ([]entity.Shelter, error) {
	return r.SelectSheltersWithConn(ctx, r.Pool)
}
