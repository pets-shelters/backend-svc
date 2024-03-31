package repo

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
	"github.com/pets-shelters/backend-svc/internal/usecase"
	"github.com/pets-shelters/backend-svc/internal/usecase/repo/entity"
	"github.com/pets-shelters/backend-svc/pkg/postgres"
	"github.com/pkg/errors"
)

const temporaryFilesTableName = "temporary_files"

type TemporaryFilesRepo struct {
	*postgres.Postgres
}

func NewTemporaryFilesRepo(pg *postgres.Postgres) *TemporaryFilesRepo {
	return &TemporaryFilesRepo{pg}
}

func (r *TemporaryFilesRepo) CreateWithConn(ctx context.Context, conn usecase.IConnection, temporaryFile entity.TemporaryFile) (int64, error) {
	sql, args, err := r.Builder.
		Insert(temporaryFilesTableName).
		Columns("file_id", "user_id", "created_at").
		Values(temporaryFile.FileID, temporaryFile.UserID, temporaryFile.CreatedAt).
		Suffix("returning id").
		ToSql()
	if err != nil {
		return 0, errors.Wrap(err, "failed to build temporary_file insert query")
	}

	var id int64
	err = conn.QueryRow(ctx, sql, args...).Scan(&id)
	if err != nil {
		return 0, errors.Wrap(err, "failed to QueryRow temporary_file insert query")
	}

	return id, nil
}

func (r *TemporaryFilesRepo) Create(ctx context.Context, temporaryFile entity.TemporaryFile) (int64, error) {
	return r.CreateWithConn(ctx, r.Pool, temporaryFile)
}

func (r *TemporaryFilesRepo) GetWithConn(ctx context.Context, conn usecase.IConnection, id int64) (*entity.TemporaryFile, error) {
	sql, args, err := r.Builder.
		Select("*").
		From(temporaryFilesTableName).
		Where(squirrel.Eq{"id": id}).
		ToSql()
	if err != nil {
		return nil, errors.Wrap(err, "failed to build get temporary_file query")
	}

	var temporaryFile entity.TemporaryFile
	err = conn.QueryRow(ctx, sql, args...).Scan(&temporaryFile.ID, &temporaryFile.FileID, &temporaryFile.UserID, &temporaryFile.CreatedAt)
	if err != nil {
		if errors.As(err, &pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, errors.Wrap(err, "failed to QueryRow get temporary_file query")
	}

	return &temporaryFile, nil
}

func (r *TemporaryFilesRepo) Get(ctx context.Context, id int64) (*entity.TemporaryFile, error) {
	return r.GetWithConn(ctx, r.Pool, id)
}

func (r *TemporaryFilesRepo) DeleteWithConn(ctx context.Context, conn usecase.IConnection, id int64) (*entity.TemporaryFile, error) {
	sql, args, err := r.Builder.
		Delete(temporaryFilesTableName).
		Where(squirrel.Eq{"id": id}).
		Suffix("returning *").
		ToSql()
	if err != nil {
		return nil, errors.Wrap(err, "failed to build delete temporary_file query")
	}

	var temporaryFile entity.TemporaryFile
	err = conn.QueryRow(ctx, sql, args...).Scan(&temporaryFile.ID, &temporaryFile.FileID, &temporaryFile.UserID, &temporaryFile.CreatedAt)
	if err != nil {
		if errors.As(err, &pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, errors.Wrap(err, "failed to QueryRow delete temporary_file query")
	}

	return &temporaryFile, nil
}

func (r *TemporaryFilesRepo) CountForUserId(ctx context.Context, userId int64) (int64, error) {
	sql, args, err := r.Builder.
		Select("count(*)").
		From(temporaryFilesTableName).
		Where(squirrel.Eq{"user_id": userId}).
		ToSql()
	if err != nil {
		return 0, errors.Wrap(err, "failed to build count temporary_file query")
	}

	var count int64
	err = r.Pool.QueryRow(ctx, sql, args...).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "failed to QueryRow temporary_file count query")
	}

	return count, nil
}
