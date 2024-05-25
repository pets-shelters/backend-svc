package authorization

import (
	"context"
	"encoding/json"
	"github.com/bradfitz/gomemcache/memcache"
	"github.com/pets-shelters/backend-svc/internal/exceptions"
	"github.com/pets-shelters/backend-svc/internal/structs"
	"github.com/pets-shelters/backend-svc/internal/usecase/repo/entity"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"strconv"
)

const googleUserinfoUrl = "https://www.googleapis.com/oauth2/v2/userinfo"

type callbackInfo struct {
	Email string `json:"email"`
}

func (uc *UseCase) Callback(ctx context.Context, cookie string, googleState string, googleCode string) (*structs.TokensPair, error) {
	err := uc.validateUsersGoogleState(cookie, googleState)
	if err != nil {
		return nil, errors.Wrap(err, "failed to validate user's state")
	}

	userinfo, err := uc.getGoogleUserInfo(googleCode)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get google userinfo")
	}

	id, err := uc.repo.GetUsersRepo().Create(ctx, entity.User{
		Email: userinfo.Email,
		Role:  structs.ManagerUserRole,
	})
	if err != nil {
		if !errors.As(err, &exceptions.UserExistsException{}) {
			return nil, errors.Wrap(err, "failed to create user entity")
		}
	}

	if id == 0 {
		users, err := uc.repo.GetUsersRepo().Select(ctx, entity.UsersFilters{Email: &userinfo.Email})
		if err != nil {
			return nil, errors.Wrap(err, "failed to select user")
		}
		id = users[0].ID
	}
	tokensPair, err := uc.jwt.CreateTokensPair(strconv.Itoa(int(id)))
	if err != nil {
		return nil, errors.Wrap(err, "failed to create tokens pair")
	}

	return tokensPair, nil
}

func (uc *UseCase) validateUsersGoogleState(cookie string, googleState string) error {
	state, err := uc.cache.GetGoogleState(cookie)
	if err != nil {
		if errors.Is(err, memcache.ErrCacheMiss) {
			return exceptions.NewInvalidStateException()
		}
		return errors.Wrap(err, "failed to get googleState from cache")
	}
	if state != googleState {
		return exceptions.NewInvalidStateException()
	}

	return nil
}

func (uc *UseCase) getGoogleUserInfo(googleCode string) (*callbackInfo, error) {
	token, err := uc.oauth.Exchange(context.Background(), googleCode)
	if err != nil {
		return nil, errors.Wrap(err, "failed to exchange googleCode - token")
	}

	resp, err := http.Get(googleUserinfoUrl + "?access_token=" + token.AccessToken)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get google userinfo")
	}
	userinfoBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read google userinfo body")
	}

	var userinfo callbackInfo
	err = json.Unmarshal(userinfoBytes, &userinfo)
	if err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal google userinfo response body")
	}

	return &userinfo, nil
}
