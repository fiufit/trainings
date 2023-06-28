package repositories

import (
	"context"
	"errors"
	"testing"
	"time"

	gContracts "github.com/fiufit/trainings/contracts/goals"
	"github.com/fiufit/trainings/models"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap/zaptest"
)

func TestGoalsRepository_Create_OK(t *testing.T) {
	defer testSuite.TruncateModels()
	ctx := context.Background()
	db := testSuite.DB
	repository := NewGoalsRepository(db, zaptest.NewLogger(t))
	user_id := "user_id_test"

	testGoal := models.Goal{
		Title:  "title_test",
		UserID: user_id,
	}

	createdGoal, err := repository.Create(ctx, testGoal)
	assert.NoError(t, err)
	assert.Equal(t, createdGoal.ID, uint(1))

}

func TestGoalsRepository_Create_DBError(t *testing.T) {
	defer testSuite.TruncateModels()
	ctx := context.Background()
	db := testSuite.DB
	repository := NewGoalsRepository(db, zaptest.NewLogger(t))
	user_id := "user_id_test"

	testGoal := models.Goal{
		Title:  "title_test",
		UserID: user_id,
	}

	_ = repository.db.AddError(errors.New("test error"))
	_, err := repository.Create(ctx, testGoal)
	assert.Error(t, err)
	db.Error = nil //overwrite the db Error so that TruncateModels() doesn't panic
}

func TestGoalsRepository_GetByID_OK(t *testing.T) {
	defer testSuite.TruncateModels()
	ctx := context.Background()
	db := testSuite.DB
	repository := NewGoalsRepository(db, zaptest.NewLogger(t))
	user_id := "user_id_test"

	testGoal := models.Goal{
		Title:  "title_test",
		UserID: user_id,
	}

	_ = db.Create(&testGoal)

	goal, err := repository.GetByID(ctx, testGoal.ID)
	assert.NoError(t, err)
	assert.Equal(t, goal.ID, testGoal.ID)
}

func TestGoalsRepository_GetByID_DBError(t *testing.T) {
	defer testSuite.TruncateModels()
	ctx := context.Background()
	db := testSuite.DB
	repository := NewGoalsRepository(db, zaptest.NewLogger(t))
	user_id := "user_id_test"

	testGoal := models.Goal{
		Title:  "title_test",
		UserID: user_id,
	}

	_ = db.Create(&testGoal)

	_ = repository.db.AddError(errors.New("test error"))
	_, err := repository.GetByID(ctx, testGoal.ID)
	assert.Error(t, err)
	db.Error = nil //overwrite the db Error so that TruncateModels() doesn't panic
}

func TestGoalsRepository_GetByID_NotFound(t *testing.T) {
	defer testSuite.TruncateModels()
	ctx := context.Background()
	db := testSuite.DB
	repository := NewGoalsRepository(db, zaptest.NewLogger(t))

	_, err := repository.GetByID(ctx, 1)
	assert.Error(t, err)
}

func TestGoalsRepository_GetByUserId_OK(t *testing.T) {
	defer testSuite.TruncateModels()
	ctx := context.Background()
	db := testSuite.DB
	repository := NewGoalsRepository(db, zaptest.NewLogger(t))
	user_id := "user_id_test"

	testGoal := models.Goal{
		Title:  "title_test",
		UserID: user_id,
	}

	_ = db.Create(&testGoal)

	req := gContracts.GetGoalsRequest{
		UserID: user_id,
	}

	goals, err := repository.Get(ctx, req)
	assert.NoError(t, err)
	assert.Equal(t, goals[0].ID, testGoal.ID)
}

func TestGoalsRepository_GetByGoalType_OK(t *testing.T) {
	defer testSuite.TruncateModels()
	ctx := context.Background()
	db := testSuite.DB
	repository := NewGoalsRepository(db, zaptest.NewLogger(t))
	user_id := "user_id_test"

	testGoal := models.Goal{
		Title:    "title_test",
		UserID:   user_id,
		GoalType: "goal_type_test",
	}

	_ = db.Create(&testGoal)

	req := gContracts.GetGoalsRequest{
		GoalType: "goal_type_test",
	}

	goals, err := repository.Get(ctx, req)
	assert.NoError(t, err)
	assert.Equal(t, goals[0].ID, testGoal.ID)
}

func TestGoalsRepository_GetByGoalSubtype_OK(t *testing.T) {
	defer testSuite.TruncateModels()
	ctx := context.Background()
	db := testSuite.DB
	repository := NewGoalsRepository(db, zaptest.NewLogger(t))
	user_id := "user_id_test"

	testGoal := models.Goal{
		Title:       "title_test",
		UserID:      user_id,
		GoalType:    "goal_type_test",
		GoalSubtype: "goal_subtype_test",
	}

	_ = db.Create(&testGoal)

	req := gContracts.GetGoalsRequest{
		GoalSubtype: "goal_subtype_test",
	}

	goals, err := repository.Get(ctx, req)
	assert.NoError(t, err)
	assert.Equal(t, goals[0].ID, testGoal.ID)
}

func TestGoalsRepository_Get_DBError(t *testing.T) {
	defer testSuite.TruncateModels()
	ctx := context.Background()
	db := testSuite.DB
	repository := NewGoalsRepository(db, zaptest.NewLogger(t))
	user_id := "user_id_test"

	testGoal := models.Goal{
		Title:  "title_test",
		UserID: user_id,
	}

	_ = db.Create(&testGoal)

	req := gContracts.GetGoalsRequest{
		UserID: user_id,
	}

	_ = repository.db.AddError(errors.New("test error"))
	_, err := repository.Get(ctx, req)
	assert.Error(t, err)
	db.Error = nil //overwrite the db Error so that TruncateModels() doesn't panic
}

func TestGoalsRepository_GetByUserIdNotFoundReturnsEmptyArray_OK(t *testing.T) {
	defer testSuite.TruncateModels()
	ctx := context.Background()
	db := testSuite.DB
	repository := NewGoalsRepository(db, zaptest.NewLogger(t))

	req := gContracts.GetGoalsRequest{
		UserID: "user_id_test",
	}

	goals, err := repository.Get(ctx, req)
	assert.NoError(t, err)
	assert.Equal(t, len(goals), 0)
}

func TestGoalsRepository_Update_OK(t *testing.T) {
	defer testSuite.TruncateModels()
	ctx := context.Background()
	db := testSuite.DB
	repository := NewGoalsRepository(db, zaptest.NewLogger(t))

	testGoal := models.Goal{
		Title: "title_test",
	}

	_ = db.Create(&testGoal)

	testGoal.Title = "title_test_updated"

	updatedGoal, err := repository.Update(ctx, testGoal)
	assert.NoError(t, err)
	assert.Equal(t, updatedGoal.Title, "title_test_updated")
}

func TestGoalsRepository_Update_DBError(t *testing.T) {
	defer testSuite.TruncateModels()
	ctx := context.Background()
	db := testSuite.DB
	repository := NewGoalsRepository(db, zaptest.NewLogger(t))

	testGoal := models.Goal{
		Title: "title_test",
	}

	_ = db.Create(&testGoal)

	testGoal.Title = "title_test_updated"

	_ = repository.db.AddError(errors.New("test error"))
	_, err := repository.Update(ctx, testGoal)
	assert.Error(t, err)
	db.Error = nil //overwrite the db Error so that TruncateModels() doesn't panic
}

func TestGoalsRepository_UpdateWithGoalNotFoundCreatesNewGoal_OK(t *testing.T) {
	defer testSuite.TruncateModels()
	ctx := context.Background()
	db := testSuite.DB
	repository := NewGoalsRepository(db, zaptest.NewLogger(t))

	testGoal := models.Goal{
		Title:  "title_test",
		UserID: "user_id_test",
	}

	updatedGoal, err := repository.Update(ctx, testGoal)
	assert.NoError(t, err)
	assert.Equal(t, updatedGoal.Title, "title_test")
}

func TestGoalsRepository_Delete_OK(t *testing.T) {
	defer testSuite.TruncateModels()
	ctx := context.Background()
	db := testSuite.DB
	repository := NewGoalsRepository(db, zaptest.NewLogger(t))

	testGoal := models.Goal{
		Title: "title_test",
	}

	_ = db.Create(&testGoal)

	err := repository.Delete(ctx, testGoal.ID)
	assert.NoError(t, err)
}

func TestGoalsRepository_Delete_DBError(t *testing.T) {
	defer testSuite.TruncateModels()
	ctx := context.Background()
	db := testSuite.DB
	repository := NewGoalsRepository(db, zaptest.NewLogger(t))

	testGoal := models.Goal{
		Title: "title_test",
	}

	_ = db.Create(&testGoal)

	_ = repository.db.AddError(errors.New("test error"))
	err := repository.Delete(ctx, testGoal.ID)
	assert.Error(t, err)
	db.Error = nil //overwrite the db Error so that TruncateModels() doesn't panic
}

func TestGoalsRepository_DeleteWithGoalNotFoundReturnsError(t *testing.T) {
	defer testSuite.TruncateModels()
	ctx := context.Background()
	db := testSuite.DB
	repository := NewGoalsRepository(db, zaptest.NewLogger(t))

	err := repository.Delete(ctx, 1)
	assert.Error(t, err)
	db.Error = nil //overwrite the db Error so that TruncateModels() doesn't panic
}

func TestGoalsRepository_UpdateBySessionStepCountGoalType_OK(t *testing.T) {
	defer testSuite.TruncateModels()
	ctx := context.Background()
	db := testSuite.DB
	repository := NewGoalsRepository(db, zaptest.NewLogger(t))
	user_id := "user_id_test"

	trainingPlan := models.TrainingPlan{
		Name: "test",
		Tags: []models.Tag{
			{
				Name: "speed",
			},
		},
		Difficulty: "beginner",
	}

	err := db.Create(&trainingPlan).Error
	assert.NoError(t, err)

	trainingSession := models.TrainingSession{
		TrainingPlanID:      trainingPlan.ID,
		TrainingPlanVersion: trainingPlan.Version,
		TrainingPlan:        trainingPlan,
		StepCount:           100,
		SecondsCount:        120,
		UserID:              user_id,
		Done:                true,
	}

	err = db.Create(&trainingSession).Error
	assert.NoError(t, err)

	testGoal := models.Goal{
		GoalType:          "step count",
		GoalValue:         100,
		GoalValueProgress: 0,
		UserID:            user_id,
		Deadline:          time.Now().AddDate(0, 0, 1),
	}

	_ = db.Create(&testGoal)

	updatedGoal, err := repository.UpdateBySession(ctx, trainingSession)
	assert.NoError(t, err)
	assert.Equal(t, updatedGoal[0].GoalValueProgress, uint(100))
}

func TestGoalsRepository_UpdateBySessionMinutesCountGoalType_OK(t *testing.T) {
	defer testSuite.TruncateModels()
	ctx := context.Background()
	db := testSuite.DB
	repository := NewGoalsRepository(db, zaptest.NewLogger(t))
	user_id := "user_id_test"

	trainingPlan := models.TrainingPlan{
		Name: "test",
		Tags: []models.Tag{
			{
				Name: "speed",
			},
		},
		Difficulty: "beginner",
	}

	err := db.Create(&trainingPlan).Error
	assert.NoError(t, err)

	trainingSession := models.TrainingSession{
		TrainingPlanID:      trainingPlan.ID,
		TrainingPlanVersion: trainingPlan.Version,
		TrainingPlan:        trainingPlan,
		StepCount:           100,
		SecondsCount:        120,
		UserID:              user_id,
		Done:                true,
	}

	err = db.Create(&trainingSession).Error
	assert.NoError(t, err)

	testGoal := models.Goal{
		GoalType:          "minutes count",
		GoalValue:         2,
		GoalValueProgress: 0,
		UserID:            user_id,
		Deadline:          time.Now().AddDate(0, 0, 1),
	}

	_ = db.Create(&testGoal)

	updatedGoal, err := repository.UpdateBySession(ctx, trainingSession)
	assert.NoError(t, err)
	assert.Equal(t, updatedGoal[0].GoalValueProgress, uint(2))
}

func TestGoalsRepository_UpdateBySessionSessionsCountGoalTypeDifficultySubtype_OK(t *testing.T) {
	defer testSuite.TruncateModels()
	ctx := context.Background()
	db := testSuite.DB
	repository := NewGoalsRepository(db, zaptest.NewLogger(t))
	user_id := "user_id_test"

	trainingPlan := models.TrainingPlan{
		Name: "test",
		Tags: []models.Tag{
			{
				Name: "speed",
			},
		},
		Difficulty: "beginner",
	}

	err := db.Create(&trainingPlan).Error
	assert.NoError(t, err)

	trainingSession := models.TrainingSession{
		TrainingPlanID:      trainingPlan.ID,
		TrainingPlanVersion: trainingPlan.Version,
		TrainingPlan:        trainingPlan,
		StepCount:           100,
		SecondsCount:        120,
		UserID:              user_id,
		Done:                true,
	}

	err = db.Create(&trainingSession).Error
	assert.NoError(t, err)

	testGoal := models.Goal{
		GoalType:          "sessions count",
		GoalSubtype:       "beginner",
		GoalValue:         1,
		GoalValueProgress: 0,
		UserID:            user_id,
		Deadline:          time.Now().AddDate(0, 0, 1),
	}

	_ = db.Create(&testGoal)

	updatedGoal, err := repository.UpdateBySession(ctx, trainingSession)
	assert.NoError(t, err)
	assert.Equal(t, updatedGoal[0].GoalValueProgress, uint(1))
}

func TestGoalsRepository_UpdateBySessionSessionsCountGoalTypeTagSubtype_OK(t *testing.T) {
	defer testSuite.TruncateModels()
	ctx := context.Background()
	db := testSuite.DB
	repository := NewGoalsRepository(db, zaptest.NewLogger(t))
	user_id := "user_id_test"

	trainingPlan := models.TrainingPlan{
		Name: "test",
		Tags: []models.Tag{
			{
				Name: "speed",
			},
		},
		Difficulty: "beginner",
	}

	err := db.Create(&trainingPlan).Error
	assert.NoError(t, err)

	trainingSession := models.TrainingSession{
		TrainingPlanID:      trainingPlan.ID,
		TrainingPlanVersion: trainingPlan.Version,
		TrainingPlan:        trainingPlan,
		StepCount:           100,
		SecondsCount:        120,
		UserID:              user_id,
		Done:                true,
	}

	err = db.Create(&trainingSession).Error
	assert.NoError(t, err)

	testGoal := models.Goal{
		GoalType:          "sessions count",
		GoalSubtype:       "speed",
		GoalValue:         1,
		GoalValueProgress: 0,
		UserID:            user_id,
		Deadline:          time.Now().AddDate(0, 0, 1),
	}

	_ = db.Create(&testGoal)

	updatedGoal, err := repository.UpdateBySession(ctx, trainingSession)
	assert.NoError(t, err)
	assert.Equal(t, updatedGoal[0].GoalValueProgress, uint(1))
}

func TestGoalsRepository_UpdateBySessionExpiredDeadlineDoesNotUpdateGoal_OK(t *testing.T) {
	defer testSuite.TruncateModels()
	ctx := context.Background()
	db := testSuite.DB
	repository := NewGoalsRepository(db, zaptest.NewLogger(t))
	user_id := "user_id_test"

	trainingPlan := models.TrainingPlan{
		Name: "test",
		Tags: []models.Tag{
			{
				Name: "speed",
			},
		},
		Difficulty: "beginner",
	}

	err := db.Create(&trainingPlan).Error
	assert.NoError(t, err)

	trainingSession := models.TrainingSession{
		TrainingPlanID:      trainingPlan.ID,
		TrainingPlanVersion: trainingPlan.Version,
		TrainingPlan:        trainingPlan,
		StepCount:           100,
		SecondsCount:        120,
		UserID:              user_id,
		Done:                true,
	}

	err = db.Create(&trainingSession).Error
	assert.NoError(t, err)

	testGoal := models.Goal{
		GoalType:          "sessions count",
		GoalSubtype:       "speed",
		GoalValue:         1,
		GoalValueProgress: 0,
		UserID:            user_id,
		Deadline:          time.Now().AddDate(0, 0, -1),
	}

	_ = db.Create(&testGoal)

	updatedGoal, err := repository.UpdateBySession(ctx, trainingSession)
	assert.NoError(t, err)
	assert.Equal(t, len(updatedGoal), 0)

	gettedGoal, err := repository.GetByID(ctx, testGoal.ID)
	assert.NoError(t, err)
	assert.Equal(t, gettedGoal.GoalValueProgress, uint(0))
}
