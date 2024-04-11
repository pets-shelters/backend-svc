package repo

import (
	"context"
	"github.com/jackc/pgx/v4"
	"github.com/pets-shelters/backend-svc/internal/usecase"
	"github.com/pets-shelters/backend-svc/pkg/postgres"
	"github.com/pkg/errors"
)

type DBRepo struct {
	db *postgres.Postgres
}

func NewDBRepo(pg *postgres.Postgres) *DBRepo {
	return &DBRepo{pg}
}

func (r *DBRepo) GetSheltersRepo() usecase.ISheltersRepo {
	return NewSheltersRepo(r.db)
}

func (r *DBRepo) GetUsersRepo() usecase.IUsersRepo {
	return NewUsersRepo(r.db)
}

func (r *DBRepo) GetFilesRepo() usecase.IFilesRepo {
	return NewFilesRepo(r.db)
}

func (r *DBRepo) GetTemporaryFilesRepo() usecase.ITemporaryFilesRepo {
	return NewTemporaryFilesRepo(r.db)
}

func (r *DBRepo) GetLocationsRepo() usecase.ILocationsRepo {
	return NewLocationsRepo(r.db)
}

func (r *DBRepo) GetAnimalsRepo() usecase.IAnimalsRepo {
	return NewAnimalsRepo(r.db)
}

func (r *DBRepo) GetAnimalTypesEnumRepo() usecase.IAnimalTypesEnumRepo {
	return NewAnimalTypesEnumRepo(r.db)
}

func (r *DBRepo) GetAdoptersRepo() usecase.IAdoptersRepo {
	return NewAdoptersRepo(r.db)
}

func (r *DBRepo) Transaction(ctx context.Context, f func(pgx.Tx) error) error {
	err := r.db.Pool.BeginTxFunc(ctx, pgx.TxOptions{}, f)
	if err != nil {
		return errors.Wrap(err, "failed to BeginTxFunc")
	}
	return nil
}
