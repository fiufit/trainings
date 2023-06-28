package repositories

import (
	"context"
	"errors"
	"testing"

	tcontracts "github.com/fiufit/trainings/contracts/training_sessions"
	"github.com/fiufit/trainings/models"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap/zaptest"
)

func TestTrainingSessionRepository_Create_OK(t *testing.T) {
	defer testSuite.TruncateModels()
	ctx := context.Background()
	db := testSuite.DB
	repository := NewTrainingSessionsRepository(db, zaptest.NewLogger(t))

	testTrainingPlan := models.TrainingPlan{
		Name: "test",
	}
	_ = db.Create(&testTrainingPlan)
	testTrainingSession := models.TrainingSession{
		TrainingPlanID:      1,
		TrainingPlanVersion: 1,
		TrainingPlan:        testTrainingPlan,
	}

	createdTrainingSession, err := repository.Create(ctx, testTrainingSession)
	assert.NoError(t, err)
	assert.Equal(t, createdTrainingSession.ID, uint(1))

}

func TestTrainingSessionRepository_Create_DBError(t *testing.T) {
	defer testSuite.TruncateModels()
	ctx := context.Background()
	db := testSuite.DB
	repository := NewTrainingSessionsRepository(db, zaptest.NewLogger(t))

	testTrainingPlan := models.TrainingPlan{
		Name: "test",
	}
	_ = db.Create(&testTrainingPlan)
	testTrainingSession := models.TrainingSession{
		TrainingPlanID:      1,
		TrainingPlanVersion: 1,
		TrainingPlan:        testTrainingPlan,
	}

	_ = repository.db.AddError(errors.New("test error"))
	_, err := repository.Create(ctx, testTrainingSession)
	assert.Error(t, err)
	db.Error = nil //overwrite the db Error so that TruncateModels() doesn't panic
}

func TestTrainingSessionRepository_CreateWithTrainingIdNotFound_DBError(t *testing.T) {
	defer testSuite.TruncateModels()
	ctx := context.Background()
	db := testSuite.DB
	repository := NewTrainingSessionsRepository(db, zaptest.NewLogger(t))

	testTrainingSession := models.TrainingSession{
		TrainingPlanID:      1,
		TrainingPlanVersion: 1,
	}

	_, err := repository.Create(ctx, testTrainingSession)
	assert.Error(t, err)
	db.Error = nil //overwrite the db Error so that TruncateModels() doesn't panic
}

func TestTrainingSessionRepository_GetById_OK(t *testing.T) {
	defer testSuite.TruncateModels()
	ctx := context.Background()
	db := testSuite.DB
	repository := NewTrainingSessionsRepository(db, zaptest.NewLogger(t))

	testTrainingPlan := models.TrainingPlan{
		Name: "test",
	}
	err := db.Create(&testTrainingPlan).Error
	assert.NoError(t, err)
	testTrainingSession := models.TrainingSession{
		TrainingPlanID:      testTrainingPlan.ID,
		TrainingPlanVersion: testTrainingPlan.Version,
		TrainingPlan:        testTrainingPlan,
	}

	err = db.Create(&testTrainingSession).Error
	assert.NoError(t, err)

	trainingSession, err := repository.GetByID(ctx, testTrainingSession.ID)
	assert.NoError(t, err)
	assert.Equal(t, trainingSession.ID, testTrainingSession.ID)
	assert.Equal(t, trainingSession.TrainingPlanID, testTrainingPlan.ID)
	assert.Equal(t, trainingSession.TrainingPlanVersion, testTrainingPlan.Version)
	assert.Equal(t, trainingSession.TrainingPlan.Name, "test")
}

func TestTrainingSessionRepository_GetById_DBError(t *testing.T) {
	defer testSuite.TruncateModels()
	ctx := context.Background()
	db := testSuite.DB

	repository := NewTrainingSessionsRepository(db, zaptest.NewLogger(t))

	_ = repository.db.AddError(errors.New("test error"))
	_, err := repository.GetByID(ctx, uint(1))
	assert.Error(t, err)
	db.Error = nil //overwrite the db Error so that TruncateModels() doesn't panic
}

func TestTrainingSessionRepository_Get_OK(t *testing.T) {
	defer testSuite.TruncateModels()
	ctx := context.Background()
	db := testSuite.DB
	repository := NewTrainingSessionsRepository(db, zaptest.NewLogger(t))
	user_id := "user_id"

	req := tcontracts.GetTrainingSessionsRequest{
		UserID: user_id,
	}

	testTrainingPlan := models.TrainingPlan{
		Name: "test",
	}
	err := db.Create(&testTrainingPlan).Error
	assert.NoError(t, err)
	testTrainingSession := models.TrainingSession{
		TrainingPlanID:      testTrainingPlan.ID,
		TrainingPlanVersion: testTrainingPlan.Version,
		TrainingPlan:        testTrainingPlan,
		UserID:              user_id,
	}

	err = db.Create(&testTrainingSession).Error
	assert.NoError(t, err)

	trainingSession, err := repository.Get(ctx, req)
	assert.NoError(t, err)
	assert.Equal(t, trainingSession[0].ID, testTrainingSession.ID)
	assert.Equal(t, trainingSession[0].TrainingPlanID, testTrainingPlan.ID)
	assert.Equal(t, trainingSession[0].TrainingPlanVersion, testTrainingPlan.Version)
	assert.Equal(t, trainingSession[0].TrainingPlan.Name, "test")
}

func TestTrainingSessionRepository_Get_DBError(t *testing.T) {
	defer testSuite.TruncateModels()
	ctx := context.Background()
	db := testSuite.DB

	repository := NewTrainingSessionsRepository(db, zaptest.NewLogger(t))

	_ = repository.db.AddError(errors.New("test error"))
	_, err := repository.Get(ctx, tcontracts.GetTrainingSessionsRequest{})
	assert.Error(t, err)
	db.Error = nil //overwrite the db Error so that TruncateModels() doesn't panics
}

func TestTrainingSessionRepository_Update_OK(t *testing.T) {
	defer testSuite.TruncateModels()
	ctx := context.Background()
	db := testSuite.DB
	repository := NewTrainingSessionsRepository(db, zaptest.NewLogger(t))

	testTrainingPlan := models.TrainingPlan{
		Name: "test",
	}
	err := db.Create(&testTrainingPlan).Error
	assert.NoError(t, err)
	testTrainingSession := models.TrainingSession{
		TrainingPlanID:      testTrainingPlan.ID,
		TrainingPlanVersion: testTrainingPlan.Version,
		TrainingPlan:        testTrainingPlan,
	}

	err = db.Create(&testTrainingSession).Error
	assert.NoError(t, err)

	testTrainingSession.Done = true
	updatedTrainingSession, err := repository.Update(ctx, testTrainingSession)
	assert.NoError(t, err)
	assert.Equal(t, updatedTrainingSession.TrainingPlan.ID, testTrainingPlan.ID)
	assert.Equal(t, updatedTrainingSession.TrainingPlan.Version, testTrainingPlan.Version)
	assert.Equal(t, updatedTrainingSession.Done, true)
}

func TestTrainingSessionRepository_Update_DBError(t *testing.T) {
	defer testSuite.TruncateModels()
	ctx := context.Background()
	db := testSuite.DB

	repository := NewTrainingSessionsRepository(db, zaptest.NewLogger(t))

	_ = repository.db.AddError(errors.New("test error"))
	_, err := repository.Update(ctx, models.TrainingSession{})
	assert.Error(t, err)
	db.Error = nil //overwrite the db Error so that TruncateModels() doesn't panic
}
