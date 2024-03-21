package postgres

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/pets-shelters/backend-svc/internal/usecase"
	"github.com/pets-shelters/backend-svc/internal/usecase/postgres/entity"
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

func (r *TemporaryFilesRepo) CreateWithConn(ctx context.Context, conn usecase.IConnection, file entity.File) (int64, error) {
	sql, args, err := r.Builder.
		Insert(temporaryFilesTableName).
		Columns("bucket", "path").
		Values(file.Bucket, file.Path).
		Suffix("returning id").
		ToSql()
	if err != nil {
		return 0, errors.Wrap(err, "failed to build file insert query")
	}

	var id int64
	err = conn.QueryRow(ctx, sql, args...).Scan(&id)
	if err != nil {
		return 0, errors.Wrap(err, "failed to QueryRow file insert query")
	}

	return id, nil
}

func (r *FilesRepo) Create(ctx context.Context, file entity.File) (int64, error) {
	return r.CreateWithConn(ctx, r.Pool, file)
}

func (r *FilesRepo) GetWithConn(ctx context.Context, conn usecase.IConnection, id int64) (*entity.File, error) {
	sql, args, err := r.Builder.
		Select("*").
		From(filesTableName).
		Where(squirrel.Eq{"id": id}).
		ToSql()
	if err != nil {
		return nil, errors.Wrap(err, "failed to build get file query")
	}

	var file entity.File
	err = conn.QueryRow(ctx, sql, args...).Scan(&file)
	if err != nil {
		return nil, errors.Wrap(err, "failed to QueryRow get file query")
	}

	return &file, nil
}

func (r *FilesRepo) Get(ctx context.Context, id int64) (*entity.File, error) {
	return r.GetWithConn(ctx, r.Pool, id)
}

func (r *FilesRepo) DeleteWithConn(ctx context.Context, conn usecase.IConnection, id int64) (int64, error) {
	sql, args, err := r.Builder.
		Delete(filesTableName).
		Where(squirrel.Eq{"id": id}).
		ToSql()
	if err != nil {
		return 0, errors.Wrap(err, "failed to build delete file query")
	}

	commandTag, err := conn.Exec(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "failed to Exec get file query")
	}

	return commandTag.RowsAffected(), nil
}

func (r *FilesRepo) Delete(ctx context.Context, id int64) (int64, error) {
	return r.DeleteWithConn(ctx, r.Pool, id)
}
