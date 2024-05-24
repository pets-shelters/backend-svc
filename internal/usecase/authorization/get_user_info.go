package authorization

import (
	"context"
	"github.com/pets-shelters/backend-svc/internal/structs/responses"
	"github.com/pkg/errors"
)

func (uc *UseCase) GetUserInfo(ctx context.Context, userId int64) (*responses.UserInfo, error) {
	user, err := uc.repo.GetUsersRepo().Get(ctx, userId)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get user entity")
	}

	return &responses.UserInfo{
		ID:         userId,
		Registered: user.ShelterID.Valid,
		ShelterID:  user.ShelterID.Int64,
	}, nil
}
