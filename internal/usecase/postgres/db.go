package postgres

import (
	"context"
	"github.com/jackc/pgx/v4"
	"github.com/pets-shelters/backend-svc/internal/usecase"
	"github.com/pets-shelters/backend-svc/pkg/postgres"
	"github.com/pkg/errors"
)

type DBRepo struct {
	*postgres.Postgres
}

func NewDBRepo(pg *postgres.Postgres) *DBRepo {
	return &DBRepo{pg}
}

func (r *DBRepo) GetSheltersRepo() usecase.ISheltersRepo {
	return NewSheltersRepo(r.Postgres)
}

func (r *DBRepo) GetUsersRepo() usecase.IUsersRepo {
	return NewUsersRepo(r.Postgres)
}

func (r *DBRepo) Transaction(ctx context.Context, f func(pgx.Tx) error) error {
	err := r.Pool.BeginTxFunc(ctx, pgx.TxOptions{}, f)
	if err != nil {
		return errors.Wrap(err, "failed to BeginTxFunc")
	}
	return nil
}
