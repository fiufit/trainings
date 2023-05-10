package trainings

import (
	"context"
	"strconv"

	"github.com/fiufit/trainings/contracts/training"
	"github.com/fiufit/trainings/models"
	"github.com/fiufit/trainings/repositories"
	"go.uber.org/zap"
)

type TrainingGetter interface {
	GetTrainingPlans(ctx context.Context, req training.GetTrainingsRequest) (training.GetTrainingsResponse, error)
	GetTrainingByID(ctx context.Context, id string) (models.TrainingPlan, error)
}

type TrainingGetterImpl struct {
	trainings repositories.TrainingPlans
	firebase  repositories.Firebase
	logger    *zap.Logger
}

func NewTrainingGetterImpl(trainings repositories.TrainingPlans, firebase repositories.Firebase, logger *zap.Logger) TrainingGetterImpl {
	return TrainingGetterImpl{trainings: trainings, firebase: firebase, logger: logger}
}

func (uc *TrainingGetterImpl) GetTrainingPlans(ctx context.Context, req training.GetTrainingsRequest) (training.GetTrainingsResponse, error) {
	res, err := uc.trainings.GetTrainingPlans(ctx, req)
	if err != nil {
		return res, err
	}
	for i := range res.TrainingPlans {
		uc.fillTrainingPicture(ctx, &res.TrainingPlans[i])
	}
	return res, nil
}

func (uc *TrainingGetterImpl) GetTrainingByID(ctx context.Context, id string) (models.TrainingPlan, error) {
	training, err := uc.trainings.GetTrainingByID(ctx, id)
	if err != nil {
		return training, err
	}
	uc.fillTrainingPicture(ctx, &training)
	return training, nil
}

func (uc *TrainingGetterImpl) fillTrainingPicture(ctx context.Context, training *models.TrainingPlan) {
	trainingID := strconv.FormatUint(uint64(training.ID), 10)
	trainingPictureUrl := uc.firebase.GetTrainingPictureUrl(ctx, trainingID, training.TrainerID)
	(*training).PictureUrl = trainingPictureUrl
}
