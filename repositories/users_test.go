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

var mockUser models.User = models.User{
	ID:       "1",
	Nickname: "test",
}

var mockUserResponse map[string]models.User = map[string]models.User{
	"data": mockUser,
}

func MakeRequestMockStatus200(method string, url string, body io.Reader) (*http.Response, error) {
	userJSON, err := json.Marshal(mockUserResponse)
	if err != nil {
		return nil, err
	}
	return &http.Response{
		Body:       io.NopCloser(bytes.NewBuffer(userJSON)),
		StatusCode: 200,
	}, nil
}

func MakeRequestMockStatus400(method string, url string, body io.Reader) (*http.Response, error) {
	return &http.Response{
		Body:       io.NopCloser(bytes.NewBuffer([]byte("error"))),
		StatusCode: 400,
	}, nil
}

func MakeRequestMockErr(method string, url string, body io.Reader) (*http.Response, error) {
	return nil, assert.AnError
}

func TestNewGetUserRepository(t *testing.T) {
	url := "http://localhost:8080"
	logger := zaptest.NewLogger(t)
	version := "v1"
	repo := NewUserRepository(url, logger, version)
	assert.Equal(t, url, repo.url)
	assert.Equal(t, logger, repo.logger)
	assert.Equal(t, version, repo.version)
}

func TestGetUserByIDOk(t *testing.T) {

	ctx := context.Background()
	patch, err := mpatch.PatchMethod(utils.MakeRequest, MakeRequestMockStatus200)
	if err != nil {
		t.Fatal(err)
	}

	url := "http://localhost:8080"
	logger := zaptest.NewLogger(t)
	version := "v1"
	repo := NewUserRepository(url, logger, version)
	user, err := repo.GetUserByID(ctx, "1")
	assert.NoError(t, err)
	assert.Equal(t, "1", user.ID)
	assert.Equal(t, "test", user.Nickname)

	err = patch.Unpatch()
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetUserByID400Error(t *testing.T) {

	ctx := context.Background()
	patch, err := mpatch.PatchMethod(utils.MakeRequest, MakeRequestMockStatus400)
	if err != nil {
		t.Fatal(err)
	}

	url := "http://localhost:8080"
	logger := zaptest.NewLogger(t)
	version := "v1"
	repo := NewUserRepository(url, logger, version)
	_, err = repo.GetUserByID(ctx, "1")
	assert.Error(t, err)

	err = patch.Unpatch()
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetUserByIDError(t *testing.T) {

	ctx := context.Background()
	patch, err := mpatch.PatchMethod(utils.MakeRequest, MakeRequestMockErr)
	if err != nil {
		t.Fatal(err)
	}

	url := "http://localhost:8080"
	logger := zaptest.NewLogger(t)
	version := "v1"
	repo := NewUserRepository(url, logger, version)
	_, err = repo.GetUserByID(ctx, "1")
	assert.Error(t, err)

	err = patch.Unpatch()
	if err != nil {
		t.Fatal(err)
	}
}
