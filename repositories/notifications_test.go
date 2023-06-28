package repositories

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"github.com/fiufit/trainings/models"
	"github.com/fiufit/trainings/utils"
	"github.com/stretchr/testify/assert"
	"github.com/undefinedlabs/go-mpatch"
	"go.uber.org/zap/zaptest"
)

var mockNotificationsRespose map[string]string = map[string]string{
	"data": "ok",
}

func MakeNotificationRequestMock200(method string, url string, body io.Reader) (*http.Response, error) {
	userJSON, err := json.Marshal(mockNotificationsRespose)
	if err != nil {
		return nil, err
	}
	return &http.Response{
		Body:       io.NopCloser(bytes.NewBuffer(userJSON)),
		StatusCode: 200,
	}, nil
}

func MakeNotificationRequestMock400(method string, url string, body io.Reader) (*http.Response, error) {
	return &http.Response{
		Body:       io.NopCloser(bytes.NewBuffer([]byte("error"))),
		StatusCode: 400,
	}, nil
}

func MakeNotificationRequestErr(method string, url string, body io.Reader) (*http.Response, error) {
	return nil, assert.AnError
}

func TestNewNotificationRepository(t *testing.T) {
	url := "http://localhost:8080"
	logger := zaptest.NewLogger(t)
	version := "v1"
	notificationRepository := NewNotificationRepository(url, logger, version)
	assert.NotNil(t, notificationRepository)
	assert.Equal(t, url, notificationRepository.url)
	assert.Equal(t, logger, notificationRepository.logger)
	assert.Equal(t, version, notificationRepository.version)
}

func TestSendGoalNotificationOk(t *testing.T) {
	ctx := context.Background()
	patch, err := mpatch.PatchMethod(utils.MakeRequest, MakeNotificationRequestMock200)
	if err != nil {
		t.Fatal(err)
	}
	url := "http://localhost:8080"
	logger := zaptest.NewLogger(t)
	version := "v1"
	notificationRepository := NewNotificationRepository(url, logger, version)
	goal := models.Goal{
		ID:    1,
		Title: "test",
	}

	err = notificationRepository.SendGoalNotification(ctx, goal)
	assert.NoError(t, err)
	err = patch.Unpatch()
	if err != nil {
		t.Fatal(err)
	}
}

func TestSendGoalNotificationErr(t *testing.T) {
	ctx := context.Background()
	patch, err := mpatch.PatchMethod(utils.MakeRequest, MakeNotificationRequestErr)
	if err != nil {
		t.Fatal(err)
	}
	url := "http://localhost:8080"
	logger := zaptest.NewLogger(t)
	version := "v1"
	notificationRepository := NewNotificationRepository(url, logger, version)
	goal := models.Goal{
		ID:    1,
		Title: "test",
	}

	err = notificationRepository.SendGoalNotification(ctx, goal)
	assert.Error(t, err)
	err = patch.Unpatch()
	if err != nil {
		t.Fatal(err)
	}
}

func TestSendGoalNotificationBadRequest(t *testing.T) {
	ctx := context.Background()
	patch, err := mpatch.PatchMethod(utils.MakeRequest, MakeNotificationRequestMock400)
	if err != nil {
		t.Fatal(err)
	}
	url := "http://localhost:8080"
	logger := zaptest.NewLogger(t)
	version := "v1"
	notificationRepository := NewNotificationRepository(url, logger, version)
	goal := models.Goal{
		ID:    1,
		Title: "test",
	}

	err = notificationRepository.SendGoalNotification(ctx, goal)
	assert.Error(t, err)
	err = patch.Unpatch()
	if err != nil {
		t.Fatal(err)
	}
}
