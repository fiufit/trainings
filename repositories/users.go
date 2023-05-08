package repositories

import (
	"context"
	"io/ioutil"
	"net/http"

	"github.com/fiufit/trainings/contracts"
	"github.com/fiufit/trainings/contracts/users"
	"github.com/fiufit/trainings/utils"
	"go.uber.org/zap"
)

//go:generate mockery --name Users
type Users interface {
	GetUserByID(ctx context.Context, userID string) (users.GetUserResponse, error)
}

type UserRepository struct {
	url     string
	logger  *zap.Logger
	version string
}

func NewUserRepository(url string, logger *zap.Logger, version string) UserRepository {
	return UserRepository{url: url, logger: logger, version: version}
}

func (repo UserRepository) GetUserByID(ctx context.Context, userID string) (users.GetUserResponse, error) {
	url := repo.url + "/" + repo.version + "/users/" + userID
	res, err := utils.MakeRequest(http.MethodGet, url, nil)
	if err != nil {
		return users.GetUserResponse{}, err
	}
	defer res.Body.Close()
	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return users.GetUserResponse{}, err
	}
	statusCode := res.StatusCode

	if statusCode >= 400 {
		err := contracts.UnwrapError(resBody)
		return users.GetUserResponse{}, err
	}

	var userResponse users.GetUserResponse
	user, err := contracts.UnwrapOkResponse(resBody, &userResponse)
	if err != nil {
		return users.GetUserResponse{}, err
	}
	userResponse = *user.(*users.GetUserResponse)
	return userResponse, nil
}
