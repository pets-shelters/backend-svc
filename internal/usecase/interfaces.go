package usecase

import (
	"context"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/pets-shelters/backend-svc/internal/structs"
	"github.com/pets-shelters/backend-svc/internal/structs/requests"
	"github.com/pets-shelters/backend-svc/internal/usecase/postgres/entity"
)

type (
	IConnection interface {
		Exec(ctx context.Context, sql string, arguments ...interface{}) (commandTag pgconn.CommandTag, err error)
		Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
		QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
		QueryFunc(ctx context.Context, sql string, args []interface{}, scans []interface{}, f func(pgx.QueryFuncRow) error) (pgconn.CommandTag, error)
	}

	IDBRepo interface {
		GetSheltersRepo() ISheltersRepo
		GetUsersRepo() IUsersRepo
		Transaction(ctx context.Context, f func(pgx.Tx) error) error
	}

	ISheltersRepo interface {
		CreateWithConn(ctx context.Context, conn IConnection, shelter entity.Shelter) (int64, error)
		Create(ctx context.Context, shelter entity.Shelter) (int64, error)
		SelectSheltersWithConn(ctx context.Context, conn IConnection) ([]entity.Shelter, error)
		SelectShelters(ctx context.Context) ([]entity.Shelter, error)
	}

	IUsersRepo interface {
		CreateWithConn(ctx context.Context, conn IConnection, user entity.User) (int64, error)
		Create(ctx context.Context, user entity.User) (int64, error)
		SelectUsersWithConn(ctx context.Context, conn IConnection) ([]entity.User, error)
		SelectUsers(ctx context.Context) ([]entity.User, error)
		UpdateShelterIDWithConn(ctx context.Context, conn IConnection, userEmail string, shelterId int64) (int64, error)
	}

	IJwt interface {
		CreateTokensPair(userEmail string) (*structs.TokensPair, error)
		VerifyAccessToken(accessTokenString string) (string, error)
		VerifyRefreshToken(refreshTokenString string) (string, error)
	}

	IAuthorization interface {
		Login() (*structs.LoginResult, error)
		Callback(ctx context.Context, cookie string, googleState string, googleCode string) (*structs.TokensPair, error)
	}

	IShelters interface {
		Create(ctx context.Context, req requests.CreateShelter, userEmail string) error
	}
)
