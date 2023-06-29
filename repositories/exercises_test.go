package repositories

import (
	"context"
	"errors"
	"testing"

	"github.com/fiufit/trainings/models"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap/zaptest"
)

func TestExerciseRepository_Create_OK(t *testing.T) {
	defer testSuite.TruncateModels()
	ctx := context.Background()
	db := testSuite.DB
	repository := NewExerciseRepository(db, zaptest.NewLogger(t))

	testTrainingPlan := models.TrainingPlan{
		Name: "test",
	}

	_ = db.Create(&testTrainingPlan)

	testExercise := models.Exercise{
		TrainingPlanID:      testTrainingPlan.ID,
		TrainingPlanVersion: testTrainingPlan.Version,
		Title:               "title_test",
		Description:         "description_test",
	}

	createdExercise, err := repository.CreateExercise(ctx, testExercise)
	assert.NoError(t, err)
	assert.Equal(t, createdExercise.ID, uint(1))

}

func TestExerciseRepository_Create_DBError(t *testing.T) {
	defer testSuite.TruncateModels()
	ctx := context.Background()
	db := testSuite.DB
	repository := NewExerciseRepository(db, zaptest.NewLogger(t))

	testTrainingPlan := models.TrainingPlan{
		Name: "test",
	}

	_ = db.Create(&testTrainingPlan)

	testExercise := models.Exercise{
		TrainingPlanID:      testTrainingPlan.ID,
		TrainingPlanVersion: testTrainingPlan.Version,
		Title:               "title_test",
		Description:         "description_test",
	}

	_ = repository.db.AddError(errors.New("test error"))
	_, err := repository.CreateExercise(ctx, testExercise)
	assert.Error(t, err)
	db.Error = nil //overwrite the db Error so that TruncateModels() doesn't panic
}

func TestExerciseRepository_Delete_OK(t *testing.T) {
	defer testSuite.TruncateModels()
	ctx := context.Background()
	db := testSuite.DB
	repository := NewExerciseRepository(db, zaptest.NewLogger(t))

	testTrainingPlan := models.TrainingPlan{
		Name: "test",
	}

	_ = db.Create(&testTrainingPlan)

	testExercise := models.Exercise{
		TrainingPlanID:      testTrainingPlan.ID,
		TrainingPlanVersion: testTrainingPlan.Version,
		Title:               "title_test",
		Description:         "description_test",
	}

	_ = db.Create(&testExercise)

	err := repository.DeleteExercise(ctx, testExercise.ID)
	assert.NoError(t, err)

}

func TestExerciseRepository_Delete_DBError(t *testing.T) {
	defer testSuite.TruncateModels()
	ctx := context.Background()
	db := testSuite.DB
	repository := NewExerciseRepository(db, zaptest.NewLogger(t))

	testTrainingPlan := models.TrainingPlan{
		Name: "test",
	}

	_ = db.Create(&testTrainingPlan)

	testExercise := models.Exercise{
		TrainingPlanID:      testTrainingPlan.ID,
		TrainingPlanVersion: testTrainingPlan.Version,
		Title:               "title_test",
		Description:         "description_test",
	}

	_ = db.Create(&testExercise)

	_ = repository.db.AddError(errors.New("test error"))
	err := repository.DeleteExercise(ctx, testExercise.ID)
	assert.Error(t, err)
	db.Error = nil //overwrite the db Error so that TruncateModels() doesn't panic
}

func TestExerciseRepository_Delete_NotFound(t *testing.T) {
	defer testSuite.TruncateModels()
	ctx := context.Background()
	db := testSuite.DB
	repository := NewExerciseRepository(db, zaptest.NewLogger(t))

	err := repository.DeleteExercise(ctx, 1)
	assert.Error(t, err)

}

func TestExerciseRepository_GetByID_OK(t *testing.T) {
	defer testSuite.TruncateModels()
	ctx := context.Background()
	db := testSuite.DB
	repository := NewExerciseRepository(db, zaptest.NewLogger(t))

	testTrainingPlan := models.TrainingPlan{
		Name: "test",
	}

	_ = db.Create(&testTrainingPlan)

	testExercise := models.Exercise{
		TrainingPlanID:      testTrainingPlan.ID,
		TrainingPlanVersion: testTrainingPlan.Version,
		Title:               "title_test",
		Description:         "description_test",
	}

	_ = db.Create(&testExercise)

	exercise, err := repository.GetExerciseByID(ctx, testExercise.ID)
	assert.NoError(t, err)
	assert.Equal(t, exercise.ID, testExercise.ID)

}

func TestExerciseRepository_GetByID_DBError(t *testing.T) {
	defer testSuite.TruncateModels()
	ctx := context.Background()
	db := testSuite.DB
	repository := NewExerciseRepository(db, zaptest.NewLogger(t))

	_ = repository.db.AddError(errors.New("test error"))
	_, err := repository.GetExerciseByID(ctx, 1)
	assert.Error(t, err)
	db.Error = nil //overwrite the db Error so that TruncateModels() doesn't panic
}

func TestExerciseRepository_GetByID_NotFound(t *testing.T) {
	defer testSuite.TruncateModels()
	ctx := context.Background()
	db := testSuite.DB
	repository := NewExerciseRepository(db, zaptest.NewLogger(t))

	_, err := repository.GetExerciseByID(ctx, 1)
	assert.Error(t, err)

}

func TestExerciseRepository_Update_OK(t *testing.T) {
	defer testSuite.TruncateModels()
	ctx := context.Background()
	db := testSuite.DB
	repository := NewExerciseRepository(db, zaptest.NewLogger(t))

	testTrainingPlan := models.TrainingPlan{
		Name: "test",
	}

	_ = db.Create(&testTrainingPlan)

	testExercise := models.Exercise{
		TrainingPlanID:      testTrainingPlan.ID,
		TrainingPlanVersion: testTrainingPlan.Version,
		Title:               "title_test",
		Description:         "description_test",
	}

	_ = db.Create(&testExercise)

	testExercise.Title = "new_title"
	testExercise.Description = "new_description"

	updatedExercise, err := repository.UpdateExercise(ctx, testExercise)
	assert.NoError(t, err)
	assert.Equal(t, updatedExercise.Title, "new_title")
	assert.Equal(t, updatedExercise.Description, "new_description")

}

func TestExerciseRepository_Update_DBError(t *testing.T) {
	defer testSuite.TruncateModels()
	ctx := context.Background()
	db := testSuite.DB
	repository := NewExerciseRepository(db, zaptest.NewLogger(t))

	testTrainingPlan := models.TrainingPlan{
		Name: "test",
	}

	_ = db.Create(&testTrainingPlan)

	testExercise := models.Exercise{
		TrainingPlanID:      testTrainingPlan.ID,
		TrainingPlanVersion: testTrainingPlan.Version,
		Title:               "title_test",
		Description:         "description_test",
	}

	_ = db.Create(&testExercise)

	testExercise.Title = "new_title"
	testExercise.Description = "new_description"

	_ = repository.db.AddError(errors.New("test error"))
	_, err := repository.UpdateExercise(ctx, testExercise)
	assert.Error(t, err)
	db.Error = nil //overwrite the db Error so that TruncateModels() doesn't panic
}

func TestExerciseRepository_Update_NotFound(t *testing.T) {
	defer testSuite.TruncateModels()
	ctx := context.Background()
	db := testSuite.DB
	repository := NewExerciseRepository(db, zaptest.NewLogger(t))

	testExercise := models.Exercise{
		Title:       "title_test",
		Description: "description_test",
	}

	_, err := repository.UpdateExercise(ctx, testExercise)
	assert.Error(t, err)

}
