package redis

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/pets-shelters/backend-svc/internal/structs/responses"
	"github.com/pkg/errors"
)

func (r *Redis) SetUserInfo(userId int64, info responses.UserInfo) error {
	err := r.client.Set(formUserInfoKey(userId), info, r.userInfoLifetime).Err()
	if err != nil {
		return errors.Wrap(err, "failed to set user_info to cache")
	}

	return nil
}

func (r *Redis) DeleteUserInfo(userId int64) error {
	err := r.client.Del(formUserInfoKey(userId)).Err()
	if err != nil {
		return errors.Wrap(err, "failed to delete user_info from cache")
	}

	return nil
}

func (r *Redis) GetUserInfo(userId int64) (*responses.UserInfo, error) {
	bytes, err := r.client.Get(formUserInfoKey(userId)).Bytes()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return nil, nil
		}
		return nil, errors.Wrap(err, "failed to get from cache")
	}

	var result responses.UserInfo
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal bytes")
	}

	return &result, nil
}

func formUserInfoKey(userId int64) string {
	return fmt.Sprintf("user-info %d", userId)
}
