package repo

import (
	"context"
	"github.com/fatih/structs"
	"github.com/lib/pq"
	"github.com/pets-shelters/backend-svc/internal/entity"
	"github.com/pets-shelters/backend-svc/internal/exceptions"
	"github.com/pets-shelters/backend-svc/pkg/postgres"
	"github.com/pkg/errors"
)

const (
	usersTableName   = "users"
	uniqueConstraint = "unique_user_email"
)

type UsersRepo struct {
	*postgres.Postgres
}

func NewUsersRepo(pg *postgres.Postgres) *UsersRepo {
	return &UsersRepo{pg}
}

func (r *UsersRepo) CreateWithConn(ctx context.Context, conn Connection, user entity.User) (int64, error) {
	sql, args, err := r.Builder.
		Insert(usersTableName).
		SetMap(structs.Map(user)).
		Suffix("returning id").
		ToSql()
	if err != nil {
		return 0, errors.Wrap(err, "failed to build user insert query")
	}

	var id int64
	err = conn.QueryRow(ctx, sql, args...).Scan(&id)
	if err != nil {
		return 0, errors.Wrap(err, "failed to QueryRow user insert query")
	}

	if err, ok := err.(*pq.Error); ok && err.Constraint == uniqueConstraint {
		return 0, exceptions.NewUserExistsException("")
	}

	return id, nil
}

func (r *UsersRepo) Create(ctx context.Context, user entity.User) (int64, error) {
	return r.CreateWithConn(ctx, r.Pool, user)
}

func (r *UsersRepo) SelectUsersWithConn(ctx context.Context, conn Connection) ([]entity.User, error) {
	sql, _, err := r.Builder.
		Select("*").
		From(usersTableName).
		ToSql()
	if err != nil {
		return nil, errors.Wrap(err, "failed to build select users query")
	}

	var shelters []entity.User
	rows, err := conn.Query(ctx, sql)
	if err != nil {
		return nil, errors.Wrap(err, "failed to Query select users query")
	}

	err = rows.Scan(&shelters)
	if err != nil {
		return nil, errors.Wrap(err, "failed to scan users")
	}

	return shelters, nil
}

func (r *UsersRepo) SelectUsers(ctx context.Context) ([]entity.User, error) {
	return r.SelectUsersWithConn(ctx, r.Pool)
}
