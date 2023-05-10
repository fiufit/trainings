package repositories

import (
	"context"
	"errors"
	"fmt"
	"time"

	"cloud.google.com/go/storage"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"go.uber.org/zap"
	"google.golang.org/api/option"
)

type Firebase interface {
	GetTrainingPictureUrl(ctx context.Context, trainingID string, trainerID string) string
}

type FirebaseRepository struct {
	logger            *zap.Logger
	app               *firebase.App
	auth              *auth.Client
	storageBucketName string
	storageBucket     *storage.BucketHandle
}

func NewFirebaseRepository(logger *zap.Logger, sdkJson []byte, storageBucketName string) (FirebaseRepository, error) {
	opt := option.WithCredentialsJSON(sdkJson)
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return FirebaseRepository{}, err
	}

	auth, err := app.Auth(context.Background())
	if err != nil {
		return FirebaseRepository{}, err
	}

	storageClient, err := app.Storage(context.Background())
	if err != nil {
		return FirebaseRepository{}, err
	}

	storageBucket, err := storageClient.Bucket(storageBucketName)
	if err != nil {
		return FirebaseRepository{}, err
	}

	repo := FirebaseRepository{
		logger:            logger,
		app:               app,
		auth:              auth,
		storageBucketName: storageBucketName,
		storageBucket:     storageBucket,
	}

	return repo, nil
}

func (repo FirebaseRepository) GetTrainingPictureUrl(ctx context.Context, trainingID string, trainerID string) string {
	defaultPictureUrl := fmt.Sprintf("https://storage.cloud.google.com/%v/training_pictures/default.png", repo.storageBucketName)
	trainingPicturePath := "training_pictures/" + trainerID + "/" + trainingID + "/training.png"
	pictureHandle := repo.storageBucket.Object(trainingPicturePath)
	_, err := pictureHandle.Attrs(ctx)
	if err != nil {
		if !errors.Is(err, storage.ErrObjectNotExist) {
			repo.logger.Error("Unable to retrieve training picture from firebase storage", zap.String("trainingID", trainingID))
		}
		trainingPicturePath = defaultPictureUrl
	}

	opts := storage.SignedURLOptions{
		Method:  "GET",
		Expires: time.Now().Add(time.Hour * 24),
	}
	pictureUrl, err := repo.storageBucket.SignedURL(trainingPicturePath, &opts)
	if err != nil {
		pictureUrl = ""
		repo.logger.Error("Unable to sign training picture from firebase storage", zap.String("trainingID", trainingID))
	}
	return pictureUrl
}
