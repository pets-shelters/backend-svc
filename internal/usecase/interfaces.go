package usecase

import (
	"context"
	"github.com/jackc/pgx/v4"
	"github.com/pets-shelters/backend-svc/internal/usecase/repo"

	"github.com/pets-shelters/backend-svc/internal/entity"
)

type (
	Authorization interface {
		Registration(context.Context, entity.Shelter, string) error
	}

	DBRepo interface {
		GetSheltersRepo() *repo.SheltersRepo
		GetUsersRepo() *repo.UsersRepo
		Transaction(ctx context.Context, f func(pgx.Tx) error) error
	}
)
