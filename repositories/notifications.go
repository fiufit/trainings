package repositories

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/fiufit/trainings/contracts"
	"github.com/fiufit/trainings/models"
	"github.com/fiufit/trainings/utils"
	"go.uber.org/zap"
)

//go:generate mockery --name Notifications
type Notifications interface {
	SendGoalNotification(ctx context.Context, goal models.Goal) error
}

type NotificationRepository struct {
	url     string
	logger  *zap.Logger
	version string
}

func NewNotificationRepository(url string, logger *zap.Logger, version string) NotificationRepository {
	return NotificationRepository{url: url, logger: logger, version: version}
}

func (repo NotificationRepository) SendGoalNotification(ctx context.Context, goal models.Goal) error {
	url := repo.url + "/api/" + repo.version + "/notifications/push"
	body := notificationBody{
		ToUserID: []string{goal.UserID},
		Title:    "FiuFit",
		Body:     fmt.Sprintf("Goal %s completed", goal.Title),
		Subtitle: "Congrats, you have completed your goal!",
		Sound:    "default",
		Data: map[string]interface{}{
			"redirectTo": "Profile",
			"params": map[string]interface{}{
				"forceRefresh": true,
			},
		},
	}
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return err
	}
	bodyReader := bytes.NewBuffer(jsonBody)
	res, err := utils.MakeRequest(http.MethodPost, url, bodyReader)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	statusCode := res.StatusCode

	if statusCode >= 400 {
		err := contracts.UnwrapError(resBody)
		return err
	}
	return nil
}

type notificationBody struct {
	ToUserID []string               `json:"to_user_id"`
	Title    string                 `json:"title"`
	Subtitle string                 `json:"subtitle"`
	Body     string                 `json:"body"`
	Sound    string                 `json:"sound"`
	Data     map[string]interface{} `json:"data"`
}
