package repo

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/pets-shelters/backend-svc/internal/exceptions"
	"github.com/pets-shelters/backend-svc/internal/usecase"
	"github.com/pets-shelters/backend-svc/internal/usecase/repo/entity"
	"github.com/pets-shelters/backend-svc/pkg/postgres"
	"github.com/pkg/errors"
	"log"
)

const (
	adoptersTableName           = "adopters"
	phoneNumberUniqueConstraint = "unique_adopter_phone"
)

type AdoptersRepo struct {
	*postgres.Postgres
}

func NewAdoptersRepo(pg *postgres.Postgres) *AdoptersRepo {
	return &AdoptersRepo{pg}
}

func (r *AdoptersRepo) CreateWithConn(ctx context.Context, conn usecase.IConnection, adopter entity.Adopter) (int64, error) {
	sql, args, err := r.Builder.
		Insert(adoptersTableName).
		Columns("name", "phone_number").
		Values(adopter.Name, adopter.PhoneNumber).
		Suffix("returning id").
		ToSql()
	if err != nil {
		return 0, errors.Wrap(err, "failed to build adopter insert query")
	}

	var id int64
	err = conn.QueryRow(ctx, sql, args...).Scan(&id)
	log.Printf("repo %+v, %+v", id, err)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.ConstraintName == phoneNumberUniqueConstraint {
			return 0, exceptions.NewAdopterExistsException()
		}

		return 0, errors.Wrap(err, "failed to QueryRow adopter insert query")
	}

	return id, nil
}

func (r *AdoptersRepo) Create(ctx context.Context, adopter entity.Adopter) (int64, error) {
	return r.CreateWithConn(ctx, r.Pool, adopter)
}

func (r *AdoptersRepo) Get(ctx context.Context, id int64) (*entity.Adopter, error) {
	sql, args, err := r.Builder.
		Select("*").
		From(adoptersTableName).
		Where(squirrel.Eq{"id": id}).
		ToSql()
	if err != nil {
		return nil, errors.Wrap(err, "failed to build get adopter query")
	}

	var adopter entity.Adopter
	err = r.Pool.QueryRow(ctx, sql, args...).Scan(&adopter.ID, &adopter.Name, &adopter.PhoneNumber)
	if err != nil {
		if errors.As(err, &pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, errors.Wrap(err, "failed to Query get adopter query")
	}

	return &adopter, nil
}

func (r *AdoptersRepo) Select(ctx context.Context, filterPhoneNumber string) ([]entity.Adopter, error) {
	sql, args, err := r.Builder.
		Select("*").
		From(adoptersTableName).
		Where(squirrel.Like{"phone_number": "%" + filterPhoneNumber + "%"}).
		ToSql()
	if err != nil {
		return nil, errors.Wrap(err, "failed to build select adopters query")
	}

	rows, err := r.Pool.Query(ctx, sql, args...)
	if err != nil {
		return nil, errors.Wrap(err, "failed to Query select adopters query")
	}

	adopters := make([]entity.Adopter, 0)
	defer rows.Close()
	for rows.Next() {
		var adopter entity.Adopter
		err = rows.Scan(&adopter.ID, &adopter.Name, &adopter.PhoneNumber)
		if err != nil {
			return nil, errors.Wrap(err, "failed to scan user entity")
		}

		adopters = append(adopters, adopter)
	}

	return adopters, nil
}
