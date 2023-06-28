package repositories

import (
	"context"
	"errors"
	"testing"

	"github.com/fiufit/trainings/contracts"
	tcontracts "github.com/fiufit/trainings/contracts/trainings"
	"github.com/fiufit/trainings/models"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap/zaptest"
	"gorm.io/gorm"
)

func TestTrainingRepository_Create_DBError(t *testing.T) {
	defer testSuite.TruncateModels()
	ctx := context.Background()

	db := testSuite.DB
	repository := NewTrainingRepository(db, zaptest.NewLogger(t))

	_ = repository.db.AddError(errors.New("test error"))
	testTrainingPlan := models.TrainingPlan{}

	_, err := repository.CreateTrainingPlan(ctx, testTrainingPlan)
	assert.Error(t, err)
	db.Error = nil //overwrite the db Error so that TruncateModels() doesn't panic
}

func TestTrainingRepository_Create_OK(t *testing.T) {
	ctx := context.Background()

	defer testSuite.TruncateModels()

	db := testSuite.DB
	repository := NewTrainingRepository(db, zaptest.NewLogger(t))

	testTrainingPlan := models.TrainingPlan{
		Name: "test",
	}
	createdTrainingPlan, err := repository.CreateTrainingPlan(ctx, testTrainingPlan)

	assert.NoError(t, err)

	var dbTrainingPlan models.TrainingPlan
	_ = db.First(&dbTrainingPlan)

	assert.Equal(t, createdTrainingPlan.ID, dbTrainingPlan.ID)
}

func TestTrainingRepository_GetByID_DBError(t *testing.T) {
	defer testSuite.TruncateModels()
	ctx := context.Background()

	db := testSuite.DB
	repository := NewTrainingRepository(db, zaptest.NewLogger(t))

	_ = repository.db.AddError(errors.New("test error"))

	_, err := repository.GetTrainingByID(ctx, 1)
	assert.Error(t, err)
	db.Error = nil //overwrite the db Error so that TruncateModels() doesn't panic
}

func TestTrainingRepository_GetByID_NotFound(t *testing.T) {
	defer testSuite.TruncateModels()
	ctx := context.Background()

	db := testSuite.DB
	repository := NewTrainingRepository(db, zaptest.NewLogger(t))

	_, err := repository.GetTrainingByID(ctx, 1)
	assert.Error(t, err)
	assert.ErrorIs(t, err, contracts.ErrTrainingPlanNotFound)
}

func TestTrainingRepository_GetByID_OK(t *testing.T) {
	defer testSuite.TruncateModels()
	ctx := context.Background()

	db := testSuite.DB
	repository := NewTrainingRepository(db, zaptest.NewLogger(t))

	testTrainingPlan := models.TrainingPlan{
		Name: "test",
	}
	_ = db.Create(&testTrainingPlan)

	resultTrainingPlan, err := repository.GetTrainingByID(ctx, testTrainingPlan.ID)
	assert.NoError(t, err)
	assert.Equal(t, resultTrainingPlan.ID, testTrainingPlan.ID)
	assert.Equal(t, resultTrainingPlan.Name, testTrainingPlan.Name)
}

func TestTrainingRepository_Get_OK(t *testing.T) {
	defer testSuite.TruncateModels()
	ctx := context.Background()

	db := testSuite.DB
	repository := NewTrainingRepository(db, zaptest.NewLogger(t))

	req := tcontracts.GetTrainingsRequest{
		Difficulty: "beginner",
	}
	testTrainingPlan1 := models.TrainingPlan{
		Name:       "test1",
		Difficulty: "beginner",
	}
	testTrainingPlan2 := models.TrainingPlan{
		Name:       "test2",
		Difficulty: "expert",
	}
	_ = db.Create(&testTrainingPlan1)
	_ = db.Create(&testTrainingPlan2)

	resultTrainingPlan, err := repository.GetTrainingPlans(ctx, req)
	assert.NoError(t, err)
	for _, trainingPlan := range resultTrainingPlan.TrainingPlans {
		assert.Equal(t, trainingPlan.Difficulty, req.Difficulty)
	}
	assert.Equal(t, len(resultTrainingPlan.TrainingPlans), 1)
	assert.Equal(t, resultTrainingPlan.TrainingPlans[0].ID, testTrainingPlan1.ID)
}

func TestTrainingRepository_GetByIdAndVersion_OK(t *testing.T) {
	defer testSuite.TruncateModels()
	ctx := context.Background()

	db := testSuite.DB
	repository := NewTrainingRepository(db, zaptest.NewLogger(t))

	testTrainingPlanVersion1 := models.TrainingPlan{
		ID:      1,
		Version: 1,
		Name:    "test",
	}

	testTrainingPlanVersion2 := models.TrainingPlan{
		ID:      1,
		Version: 2,
		Name:    "test2",
	}
	_ = db.Create(&testTrainingPlanVersion1)
	_ = db.Create(&testTrainingPlanVersion2)

	resultTrainingPlan, err := repository.GetTrainingByIDAndVersion(ctx, testTrainingPlanVersion1.ID, testTrainingPlanVersion2.Version)
	assert.NoError(t, err)
	assert.Equal(t, resultTrainingPlan.ID, testTrainingPlanVersion1.ID)
	assert.Equal(t, resultTrainingPlan.Version, testTrainingPlanVersion2.Version)
}

func TestTrainingRepository_Get_DBError(t *testing.T) {
	defer testSuite.TruncateModels()
	ctx := context.Background()

	db := testSuite.DB
	repository := NewTrainingRepository(db, zaptest.NewLogger(t))

	_ = repository.db.AddError(errors.New("test error"))

	_, err := repository.GetTrainingByIDAndVersion(ctx, 1, 1)
	assert.Error(t, err)
	db.Error = nil //overwrite the db Error so that TruncateModels() doesn't panic
}

func TestTrainingRepository_GetByIdAndVersionOldVersion_DBError(t *testing.T) {
	defer testSuite.TruncateModels()
	ctx := context.Background()

	db := testSuite.DB
	repository := NewTrainingRepository(db, zaptest.NewLogger(t))

	testTrainingPlanVersion1 := models.TrainingPlan{
		ID:      1,
		Version: 1,
		Name:    "test",
	}

	testTrainingPlanVersion2 := models.TrainingPlan{
		ID:      1,
		Version: 2,
		Name:    "test2",
	}

	_ = repository.db.AddError(contracts.ErrTrainingPlanNotFound)

	_ = db.Create(&testTrainingPlanVersion1)
	_ = db.Create(&testTrainingPlanVersion2)

	_, err := repository.GetTrainingByIDAndVersion(ctx, testTrainingPlanVersion1.ID, testTrainingPlanVersion1.Version)
	assert.Error(t, err)
	assert.ErrorIs(t, err, contracts.ErrTrainingPlanNotFound)
	db.Error = nil //overwrite the db Error so that TruncateModels() doesn't panic
}

func TestTrainingRepository_GetFavorites_OK(t *testing.T) {
	defer testSuite.TruncateModels()
	ctx := context.Background()

	db := testSuite.DB
	repository := NewTrainingRepository(db, zaptest.NewLogger(t))

	req := tcontracts.GetFavoritesRequest{
		UserID: "test",
	}

	testTrainingPlan1 := models.TrainingPlan{
		Name:       "test1",
		Difficulty: "beginner",
	}
	testTrainingPlan2 := models.TrainingPlan{
		Name:       "test2",
		Difficulty: "expert",
	}

	_ = db.Create(&testTrainingPlan1)
	_ = db.Create(&testTrainingPlan2)

	testFavorite1 := models.Favorite{
		UserID:         req.UserID,
		TrainingPlan:   testTrainingPlan1,
		TrainingPlanID: testTrainingPlan1.ID,
	}

	testFavorite2 := models.Favorite{
		UserID:         req.UserID,
		TrainingPlan:   testTrainingPlan2,
		TrainingPlanID: testTrainingPlan2.ID,
	}

	_ = db.Create(&testFavorite1)
	_ = db.Create(&testFavorite2)

	resultTrainingPlan, err := repository.GetFavoriteTrainings(ctx, req)
	assert.NoError(t, err)
	assert.Equal(t, len(resultTrainingPlan.TrainingPlans), 2)
	assert.Equal(t, resultTrainingPlan.TrainingPlans[0].ID, testTrainingPlan1.ID)
	assert.Equal(t, resultTrainingPlan.TrainingPlans[1].ID, testTrainingPlan2.ID)

}

func TestTrainingRepository_UpdateTraining_OK(t *testing.T) {
	defer testSuite.TruncateModels()
	ctx := context.Background()

	db := testSuite.DB
	repository := NewTrainingRepository(db, zaptest.NewLogger(t))
	testTrainingPlan := models.TrainingPlan{
		Name: "test",
	}
	_ = db.Create(&testTrainingPlan)

	testTrainingPlan.Name = "test2"
	updatedTrainingPlan, err := repository.UpdateTrainingPlan(ctx, testTrainingPlan)
	assert.NoError(t, err)

	var dbTrainingPlan models.TrainingPlan
	_ = db.First(&dbTrainingPlan)

	assert.Equal(t, updatedTrainingPlan.ID, dbTrainingPlan.ID)
	assert.Equal(t, updatedTrainingPlan.Name, dbTrainingPlan.Name)
	assert.Equal(t, updatedTrainingPlan.Version, dbTrainingPlan.Version)
	assert.Equal(t, testTrainingPlan.Version+1, dbTrainingPlan.Version)
}

func TestTrainingRepository_UpdateTraining_DBError(t *testing.T) {
	defer testSuite.TruncateModels()
	ctx := context.Background()

	db := testSuite.DB
	repository := NewTrainingRepository(db, zaptest.NewLogger(t))
	testTrainingPlan := models.TrainingPlan{
		Name: "test",
	}
	_ = db.Create(&testTrainingPlan)

	_ = repository.db.AddError(errors.New("test error"))

	testTrainingPlan.Name = "test2"
	_, err := repository.UpdateTrainingPlan(ctx, testTrainingPlan)
	assert.Error(t, err)
	db.Error = nil //overwrite the db Error so that TruncateModels() doesn't panic
}

func TestTrainingRepository_DeleteTraining_OK(t *testing.T) {
	defer testSuite.TruncateModels()
	ctx := context.Background()

	db := testSuite.DB
	repository := NewTrainingRepository(db, zaptest.NewLogger(t))
	testTrainingPlan := models.TrainingPlan{
		Name: "test",
	}
	_ = db.Create(&testTrainingPlan)

	err := repository.DeleteTrainingPlan(ctx, testTrainingPlan.ID)
	assert.NoError(t, err)

	var dbTrainingPlan models.TrainingPlan
	err = db.First(&dbTrainingPlan).Error
	assert.ErrorIs(t, err, gorm.ErrRecordNotFound)
}

func TestTrainingRepository_DeleteTraining_DBError(t *testing.T) {
	defer testSuite.TruncateModels()
	ctx := context.Background()

	db := testSuite.DB
	repository := NewTrainingRepository(db, zaptest.NewLogger(t))
	testTrainingPlan := models.TrainingPlan{
		Name: "test",
	}
	_ = db.Create(&testTrainingPlan)

	_ = repository.db.AddError(errors.New("test error"))

	err := repository.DeleteTrainingPlan(ctx, testTrainingPlan.ID)
	assert.Error(t, err)
	db.Error = nil //overwrite the db Error so that TruncateModels() doesn't panic
}

func TestTrainingRepository_AddToFavorite_OK(t *testing.T) {
	defer testSuite.TruncateModels()
	ctx := context.Background()

	db := testSuite.DB
	repository := NewTrainingRepository(db, zaptest.NewLogger(t))
	testTrainingPlan := models.TrainingPlan{
		Name: "test",
	}
	_ = db.Create(&testTrainingPlan)

	err := repository.AddToFavorite(ctx, "test", testTrainingPlan.ID, testTrainingPlan.Version)
	assert.NoError(t, err)

	var dbFavorite models.Favorite
	err = db.First(&dbFavorite).Error
	assert.NoError(t, err)
	assert.Equal(t, dbFavorite.TrainingPlanID, testTrainingPlan.ID)
	assert.Equal(t, dbFavorite.TrainingPlanVersion, testTrainingPlan.Version)
	assert.Equal(t, dbFavorite.UserID, "test")
}

func TestTrainingRepository_AddToFavorite_DBError(t *testing.T) {
	defer testSuite.TruncateModels()
	ctx := context.Background()

	db := testSuite.DB
	repository := NewTrainingRepository(db, zaptest.NewLogger(t))
	testTrainingPlan := models.TrainingPlan{
		Name: "test",
	}
	_ = db.Create(&testTrainingPlan)

	_ = repository.db.AddError(errors.New("test error"))

	err := repository.AddToFavorite(ctx, "test", testTrainingPlan.ID, testTrainingPlan.Version)
	assert.Error(t, err)
	db.Error = nil //overwrite the db Error so that TruncateModels() doesn't panic
}

func TestTrainingRepository_RemoveFromFavorite_Ok(t *testing.T) {
	defer testSuite.TruncateModels()
	ctx := context.Background()

	db := testSuite.DB
	repository := NewTrainingRepository(db, zaptest.NewLogger(t))
	testTrainingPlan := models.TrainingPlan{
		Name: "test",
	}
	_ = db.Create(&testTrainingPlan)

	testFavorite := models.Favorite{
		UserID:         "test",
		TrainingPlan:   testTrainingPlan,
		TrainingPlanID: testTrainingPlan.ID,
	}

	_ = db.Create(&testFavorite)

	err := repository.RemoveFromFavorite(ctx, "test", testTrainingPlan.ID, testTrainingPlan.Version)
	assert.NoError(t, err)

	var dbFavorite models.Favorite
	err = db.First(&dbFavorite).Error
	assert.ErrorIs(t, err, gorm.ErrRecordNotFound)
}

func TestTrainingRepository_RemoveFromFavorite_DBError(t *testing.T) {
	defer testSuite.TruncateModels()
	ctx := context.Background()

	db := testSuite.DB
	repository := NewTrainingRepository(db, zaptest.NewLogger(t))
	testTrainingPlan := models.TrainingPlan{
		Name: "test",
	}
	_ = db.Create(&testTrainingPlan)

	testFavorite := models.Favorite{
		UserID:         "test",
		TrainingPlan:   testTrainingPlan,
		TrainingPlanID: testTrainingPlan.ID,
	}

	_ = db.Create(&testFavorite)

	_ = repository.db.AddError(errors.New("test error"))

	err := repository.RemoveFromFavorite(ctx, "test", testTrainingPlan.ID, testTrainingPlan.Version)
	assert.Error(t, err)
	db.Error = nil //overwrite the db Error so that TruncateModels() doesn't panic
}

func TestTrainingRepository_UpdateDisabledStatusToDisabled_Ok(t *testing.T) {
	defer testSuite.TruncateModels()
	ctx := context.Background()

	db := testSuite.DB
	repository := NewTrainingRepository(db, zaptest.NewLogger(t))

	testTrainingPlan := models.TrainingPlan{
		Name:     "test",
		Disabled: false,
	}
	_ = db.Create(&testTrainingPlan)

	err := repository.UpdateDisabledStatus(ctx, testTrainingPlan.ID, true)
	assert.NoError(t, err)

	var dbTrainingPlan models.TrainingPlan
	_ = db.First(&dbTrainingPlan)

	assert.Equal(t, dbTrainingPlan.Disabled, true)
}

func TestTrainingRepository_UpdateDisabledStatusToEnabled_Ok(t *testing.T) {
	defer testSuite.TruncateModels()
	ctx := context.Background()

	db := testSuite.DB
	repository := NewTrainingRepository(db, zaptest.NewLogger(t))

	testTrainingPlan := models.TrainingPlan{
		Name:     "test",
		Disabled: true,
	}
	_ = db.Create(&testTrainingPlan)

	err := repository.UpdateDisabledStatus(ctx, testTrainingPlan.ID, false)
	assert.NoError(t, err)

	var dbTrainingPlan models.TrainingPlan
	_ = db.First(&dbTrainingPlan)

	assert.Equal(t, dbTrainingPlan.Disabled, false)
}

func TestTrainingRepository_UpdateDisabledStatus_DBError(t *testing.T) {
	defer testSuite.TruncateModels()
	ctx := context.Background()

	db := testSuite.DB
	repository := NewTrainingRepository(db, zaptest.NewLogger(t))

	testTrainingPlan := models.TrainingPlan{
		Name:     "test",
		Disabled: true,
	}
	_ = db.Create(&testTrainingPlan)

	_ = repository.db.AddError(errors.New("test error"))

	err := repository.UpdateDisabledStatus(ctx, testTrainingPlan.ID, false)
	assert.Error(t, err)
	db.Error = nil //overwrite the db Error so that TruncateModels() doesn't panic
}
