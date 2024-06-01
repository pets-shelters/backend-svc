package authorization

import (
	"context"
	"github.com/pets-shelters/backend-svc/internal/structs/responses"
	"github.com/pkg/errors"
)

func (uc *UseCase) GetUserInfo(ctx context.Context, userId int64) (*responses.UserInfo, error) {
	userInfo, err := uc.cache.GetUserInfo(userId)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read user_info from cache")
	}
	if userInfo != nil {
		return userInfo, nil
	}

	user, err := uc.repo.GetUsersRepo().GetWithShelterName(ctx, userId)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get user with shelter_name entity")
	}

	userInfo = &responses.UserInfo{
		ID:          userId,
		Email:       user.Email,
		Role:        user.Role,
		Registered:  user.ShelterID.Valid,
		ShelterID:   user.ShelterID.Int64,
		ShelterName: user.ShelterName.String,
	}
	err = uc.cache.SetUserInfo(userId, *userInfo)
	if err != nil {
		return nil, errors.Wrap(err, "failed to set user_info to cache")
	}

	return userInfo, nil
}
