package trainings

import (
	"context"

	"github.com/fiufit/trainings/contracts/trainings"
	"github.com/fiufit/trainings/models"
	"github.com/fiufit/trainings/repositories"
	"go.uber.org/zap"
)

type TrainingGetter interface {
	GetTrainingPlans(ctx context.Context, req trainings.GetTrainingsRequest) (trainings.GetTrainingsResponse, error)
	GetTrainingByID(ctx context.Context, trainingID uint) (models.TrainingPlan, error)
}

type TrainingGetterImpl struct {
	trainings repositories.TrainingPlans
	firebase  repositories.Firebase
	logger    *zap.Logger
}

func NewTrainingGetterImpl(trainings repositories.TrainingPlans, firebase repositories.Firebase, logger *zap.Logger) TrainingGetterImpl {
	return TrainingGetterImpl{trainings: trainings, firebase: firebase, logger: logger}
}

func (uc *TrainingGetterImpl) GetTrainingPlans(ctx context.Context, req trainings.GetTrainingsRequest) (trainings.GetTrainingsResponse, error) {
	res, err := uc.trainings.GetTrainingPlans(ctx, req)
	if err != nil {
		return res, err
	}
	for i := range res.TrainingPlans {
		uc.fillTrainingPicture(ctx, &res.TrainingPlans[i])
		uc.calculateMeanScore(ctx, &res.TrainingPlans[i])
	}
	return res, nil
}

func (uc *TrainingGetterImpl) GetTrainingByID(ctx context.Context, trainingID uint) (models.TrainingPlan, error) {
	training, err := uc.trainings.GetTrainingByID(ctx, trainingID)
	if err != nil {
		return training, err
	}
	uc.fillTrainingPicture(ctx, &training)
	uc.calculateMeanScore(ctx, &training)
	return training, nil
}

func (uc *TrainingGetterImpl) fillTrainingPicture(ctx context.Context, training *models.TrainingPlan) {
	trainingPictureUrl := uc.firebase.GetTrainingPictureUrl(ctx, training.ID, training.TrainerID)
	(*training).PictureUrl = trainingPictureUrl
}

func (uc *TrainingGetterImpl) calculateMeanScore(ctx context.Context, training *models.TrainingPlan) {
	var sum float32
	for i := range training.Reviews {
		sum += float32(training.Reviews[i].Score)
	}
	reviewCount := float32(len(training.Reviews))
	if reviewCount != 0 {
		(*training).MeanScore = sum / reviewCount
	} else {
		(*training).MeanScore = 0
	}

}
