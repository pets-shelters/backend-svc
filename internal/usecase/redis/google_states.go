package redis

import "github.com/pkg/errors"

func (r *Redis) SetGoogleState(cookieSession string, state string) error {
	err := r.client.Set(cookieSession, state, r.googleStateLifetime).Err()
	if err != nil {
		return errors.Wrap(err, "failed to set google_state to cache")
	}

	return nil
}

func (r *Redis) GetGoogleState(cookieSession string) (string, error) {
	bytes, err := r.client.Get(cookieSession).Bytes()
	if err != nil {
		return "", errors.Wrap(err, "failed to get google_state from cache")
	}

	return string(bytes), nil
}
