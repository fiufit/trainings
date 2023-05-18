package repositories

import (
	"context"
	"io/ioutil"
	"net/http"

	"github.com/fiufit/trainings/contracts"
	"github.com/fiufit/trainings/models"
	"github.com/fiufit/trainings/utils"
	"go.uber.org/zap"
)

//go:generate mockery --name Users
type Users interface {
	GetUserByID(ctx context.Context, userID string) (models.User, error)
}

type UserRepository struct {
	url     string
	logger  *zap.Logger
	version string
}

func NewUserRepository(url string, logger *zap.Logger, version string) UserRepository {
	return UserRepository{url: url, logger: logger, version: version}
}

func (repo UserRepository) GetUserByID(ctx context.Context, userID string) (models.User, error) {
	url := repo.url + "/" + repo.version + "/users/" + userID
	res, err := utils.MakeRequest(http.MethodGet, url, nil)
	if err != nil {
		return models.User{}, err
	}
	defer res.Body.Close()
	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return models.User{}, err
	}
	statusCode := res.StatusCode

	if statusCode >= 400 {
		err := contracts.UnwrapError(resBody)
		return models.User{}, err
	}

	var userResponse models.User
	user, err := contracts.UnwrapOkResponse(resBody, &userResponse)
	if err != nil {
		return models.User{}, err
	}
	userResponse = *user.(*models.User)
	return userResponse, nil
}
