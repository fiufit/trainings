package repositories

import (
	"context"
	"errors"
	"testing"

	rcontracts "github.com/fiufit/trainings/contracts/reviews"
	"github.com/fiufit/trainings/models"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap/zaptest"
)

func TestReviewRepository_Create_OK(t *testing.T) {
	defer testSuite.TruncateModels()
	ctx := context.Background()
	db := testSuite.DB
	repository := NewReviewRepository(db, zaptest.NewLogger(t))

	testTrainingPlan := models.TrainingPlan{
		Name: "test",
	}
	_ = db.Create(&testTrainingPlan)
	testReview := models.Review{
		TrainingPlanID:      testTrainingPlan.ID,
		TrainingPlanVersion: testTrainingPlan.Version,
		Score:               5,
		Comment:             "test",
	}

	createdReview, err := repository.CreateReview(ctx, testReview)
	assert.NoError(t, err)
	assert.Equal(t, createdReview.ID, uint(1))

}

func TestReviewRepository_Create_DBError(t *testing.T) {
	defer testSuite.TruncateModels()
	ctx := context.Background()
	db := testSuite.DB
	repository := NewReviewRepository(db, zaptest.NewLogger(t))

	testTrainingPlan := models.TrainingPlan{
		Name: "test",
	}
	_ = db.Create(&testTrainingPlan)
	testReview := models.Review{
		TrainingPlanID:      testTrainingPlan.ID,
		TrainingPlanVersion: testTrainingPlan.Version,
		Score:               5,
		Comment:             "test",
	}

	_ = repository.db.AddError(errors.New("test error"))
	_, err := repository.CreateReview(ctx, testReview)
	assert.Error(t, err)
	db.Error = nil //overwrite the db Error so that TruncateModels() doesn't panic
}

func TestReviewRepository_CreateWithTrainingIdNotFound_DBError(t *testing.T) {
	defer testSuite.TruncateModels()
	ctx := context.Background()
	db := testSuite.DB
	repository := NewReviewRepository(db, zaptest.NewLogger(t))

	testReview := models.Review{
		TrainingPlanID:      1,
		TrainingPlanVersion: 1,
		Score:               5,
		Comment:             "test",
	}

	_, err := repository.CreateReview(ctx, testReview)
	assert.Error(t, err)
}

func TestReviewRepository_CreateTwiceWithSameUserId_Error(t *testing.T) {
	defer testSuite.TruncateModels()
	ctx := context.Background()
	db := testSuite.DB
	repository := NewReviewRepository(db, zaptest.NewLogger(t))

	testTrainingPlan := models.TrainingPlan{
		Name: "test",
	}

	_ = db.Create(&testTrainingPlan)
	testReview := models.Review{
		TrainingPlanID:      testTrainingPlan.ID,
		TrainingPlanVersion: testTrainingPlan.Version,
		Score:               5,
		Comment:             "test",
	}

	_, err := repository.CreateReview(ctx, testReview)
	assert.NoError(t, err)

	_, err = repository.CreateReview(ctx, testReview)
	assert.Error(t, err)
}

func TestReviewRepository_Update_OK(t *testing.T) {
	defer testSuite.TruncateModels()
	ctx := context.Background()
	db := testSuite.DB
	repository := NewReviewRepository(db, zaptest.NewLogger(t))

	testTrainingPlan := models.TrainingPlan{
		Name: "test",
	}
	_ = db.Create(&testTrainingPlan)
	testReview := models.Review{
		TrainingPlanID:      testTrainingPlan.ID,
		TrainingPlanVersion: testTrainingPlan.Version,
		Score:               5,
		Comment:             "test",
	}

	createdReview, err := repository.CreateReview(ctx, testReview)
	assert.NoError(t, err)

	createdReview.Comment = "updated"
	updatedReview, err := repository.UpdateReview(ctx, createdReview)
	assert.NoError(t, err)
	assert.Equal(t, updatedReview.ID, createdReview.ID)
	assert.Equal(t, updatedReview.Comment, createdReview.Comment)
}

func TestReviewRepository_Update_DBError(t *testing.T) {
	defer testSuite.TruncateModels()
	ctx := context.Background()
	db := testSuite.DB
	repository := NewReviewRepository(db, zaptest.NewLogger(t))

	testTrainingPlan := models.TrainingPlan{
		Name: "test",
	}

	_ = db.Create(&testTrainingPlan)
	testReview := models.Review{
		TrainingPlanID:      testTrainingPlan.ID,
		TrainingPlanVersion: testTrainingPlan.Version,
		Score:               5,
		Comment:             "test",
	}

	createdReview, err := repository.CreateReview(ctx, testReview)
	assert.NoError(t, err)

	createdReview.Comment = "updated"
	_ = repository.db.AddError(errors.New("test error"))
	_, err = repository.UpdateReview(ctx, createdReview)
	assert.Error(t, err)
	db.Error = nil //overwrite the db Error so that TruncateModels() doesn't panic
}

func TestReviewRepository_UpdateWithTrainingIdNotFound_DBError(t *testing.T) {
	defer testSuite.TruncateModels()
	ctx := context.Background()
	db := testSuite.DB
	repository := NewReviewRepository(db, zaptest.NewLogger(t))

	testReview := models.Review{
		TrainingPlanID:      1,
		TrainingPlanVersion: 1,
		Score:               5,
		Comment:             "test",
	}

	_, err := repository.UpdateReview(ctx, testReview)
	assert.Error(t, err)
}

func TestReviewRepository_Delete_OK(t *testing.T) {
	defer testSuite.TruncateModels()
	ctx := context.Background()
	db := testSuite.DB
	repository := NewReviewRepository(db, zaptest.NewLogger(t))

	testTrainingPlan := models.TrainingPlan{
		Name: "test",
	}
	_ = db.Create(&testTrainingPlan)
	testReview := models.Review{
		TrainingPlanID:      testTrainingPlan.ID,
		TrainingPlanVersion: testTrainingPlan.Version,
		Score:               5,
		Comment:             "test",
	}

	createdReview, err := repository.CreateReview(ctx, testReview)
	assert.NoError(t, err)

	err = repository.DeleteReview(ctx, createdReview.ID)
	assert.NoError(t, err)
}

func TestReviewRepository_Delete_DBError(t *testing.T) {
	defer testSuite.TruncateModels()
	ctx := context.Background()
	db := testSuite.DB
	repository := NewReviewRepository(db, zaptest.NewLogger(t))

	testTrainingPlan := models.TrainingPlan{
		Name: "test",
	}
	_ = db.Create(&testTrainingPlan)
	testReview := models.Review{
		TrainingPlanID:      testTrainingPlan.ID,
		TrainingPlanVersion: testTrainingPlan.Version,
		Score:               5,
		Comment:             "test",
	}

	createdReview, err := repository.CreateReview(ctx, testReview)
	assert.NoError(t, err)

	_ = repository.db.AddError(errors.New("test error"))
	err = repository.DeleteReview(ctx, createdReview.ID)
	assert.Error(t, err)
	db.Error = nil //overwrite the db Error so that TruncateModels() doesn't panic
}

func TestReviewRepository_DeleteWithTrainingIdNotFound_DBError(t *testing.T) {
	defer testSuite.TruncateModels()
	ctx := context.Background()
	db := testSuite.DB
	repository := NewReviewRepository(db, zaptest.NewLogger(t))

	err := repository.DeleteReview(ctx, 1)
	assert.Error(t, err)
}

func TestReviewRepository_GetById_OK(t *testing.T) {
	defer testSuite.TruncateModels()
	ctx := context.Background()
	db := testSuite.DB
	repository := NewReviewRepository(db, zaptest.NewLogger(t))

	testTrainingPlan := models.TrainingPlan{
		Name: "test",
	}
	_ = db.Create(&testTrainingPlan)
	testReview := models.Review{
		TrainingPlanID:      testTrainingPlan.ID,
		TrainingPlanVersion: testTrainingPlan.Version,
		Score:               5,
		Comment:             "test",
	}

	createdReview, err := repository.CreateReview(ctx, testReview)
	assert.NoError(t, err)

	review, err := repository.GetReviewByID(ctx, createdReview.ID)
	assert.NoError(t, err)
	assert.Equal(t, review.ID, createdReview.ID)
	assert.Equal(t, review.Comment, createdReview.Comment)
}

func TestReviewRepository_GetById_DBError(t *testing.T) {
	defer testSuite.TruncateModels()
	ctx := context.Background()
	db := testSuite.DB
	repository := NewReviewRepository(db, zaptest.NewLogger(t))

	testTrainingPlan := models.TrainingPlan{
		Name: "test",
	}
	_ = db.Create(&testTrainingPlan)
	testReview := models.Review{
		TrainingPlanID:      testTrainingPlan.ID,
		TrainingPlanVersion: testTrainingPlan.Version,
		Score:               5,
		Comment:             "test",
	}

	createdReview, err := repository.CreateReview(ctx, testReview)
	assert.NoError(t, err)

	_ = repository.db.AddError(errors.New("test error"))
	_, err = repository.GetReviewByID(ctx, createdReview.ID)
	assert.Error(t, err)
	db.Error = nil //overwrite the db Error so that TruncateModels() doesn't panic
}

func TestReviewRepository_GetByIdWithIdNotFound_DBError(t *testing.T) {
	defer testSuite.TruncateModels()
	ctx := context.Background()
	db := testSuite.DB
	repository := NewReviewRepository(db, zaptest.NewLogger(t))

	_, err := repository.GetReviewByID(ctx, 1)
	assert.Error(t, err)
}

func TestReviewRepository_Get_OK(t *testing.T) {
	defer testSuite.TruncateModels()
	ctx := context.Background()
	db := testSuite.DB
	repository := NewReviewRepository(db, zaptest.NewLogger(t))

	testTrainingPlan := models.TrainingPlan{
		Name: "test",
	}
	_ = db.Create(&testTrainingPlan)

	req := rcontracts.GetReviewsRequest{
		MinScore:       2,
		MaxScore:       4,
		TrainingPlanID: testTrainingPlan.ID,
	}

	validTestReview := models.Review{
		TrainingPlanID:      testTrainingPlan.ID,
		TrainingPlanVersion: testTrainingPlan.Version,
		Score:               3,
		Comment:             "test_comment",
		UserID:              "reviewer_1",
	}
	invalidTestReview1 := models.Review{
		TrainingPlanID:      testTrainingPlan.ID,
		TrainingPlanVersion: testTrainingPlan.Version,
		Score:               1,
		Comment:             "test_comment_1",
		UserID:              "reviewer_2",
	}
	invalidTestReview2 := models.Review{
		TrainingPlanID:      testTrainingPlan.ID,
		TrainingPlanVersion: testTrainingPlan.Version,
		Score:               5,
		Comment:             "test_comment_2",
		UserID:              "reviewer_3",
	}

	_, err := repository.CreateReview(ctx, invalidTestReview1)
	assert.NoError(t, err)
	_, err = repository.CreateReview(ctx, invalidTestReview2)
	assert.NoError(t, err)
	createdReview, err := repository.CreateReview(ctx, validTestReview)
	assert.NoError(t, err)

	reviews, err := repository.GetReviews(ctx, req)
	assert.NoError(t, err)
	assert.Equal(t, len(reviews.Reviews), 1)
	assert.Equal(t, reviews.Reviews[0].ID, createdReview.ID)
	assert.Equal(t, reviews.Reviews[0].Comment, createdReview.Comment)
}

func TestReviewRepository_Get_DBError(t *testing.T) {
	defer testSuite.TruncateModels()
	ctx := context.Background()
	db := testSuite.DB
	repository := NewReviewRepository(db, zaptest.NewLogger(t))

	testTrainingPlan := models.TrainingPlan{
		Name: "test",
	}
	_ = db.Create(&testTrainingPlan)

	req := rcontracts.GetReviewsRequest{
		MinScore:       2,
		MaxScore:       4,
		TrainingPlanID: testTrainingPlan.ID,
	}

	validTestReview := models.Review{
		TrainingPlanID:      testTrainingPlan.ID,
		TrainingPlanVersion: testTrainingPlan.Version,
		Score:               3,
		Comment:             "test_comment",
		UserID:              "reviewer_1",
	}
	invalidTestReview1 := models.Review{
		TrainingPlanID:      testTrainingPlan.ID,
		TrainingPlanVersion: testTrainingPlan.Version,
		Score:               1,
		Comment:             "test_comment_1",
		UserID:              "reviewer_2",
	}
	invalidTestReview2 := models.Review{
		TrainingPlanID:      testTrainingPlan.ID,
		TrainingPlanVersion: testTrainingPlan.Version,
		Score:               5,
		Comment:             "test_comment_2",
		UserID:              "reviewer_3",
	}

	_, err := repository.CreateReview(ctx, invalidTestReview1)
	assert.NoError(t, err)
	_, err = repository.CreateReview(ctx, invalidTestReview2)
	assert.NoError(t, err)
	_, err = repository.CreateReview(ctx, validTestReview)
	assert.NoError(t, err)

	_ = repository.db.AddError(errors.New("test error"))
	_, err = repository.GetReviews(ctx, req)
	assert.Error(t, err)
	db.Error = nil //overwrite the db Error so that TruncateModels() doesn't panic
}

func TestReviewRepository_GetWithInvalidTrainingPlanIdReturnsEmptyArray_OK(t *testing.T) {
	defer testSuite.TruncateModels()
	ctx := context.Background()
	db := testSuite.DB
	repository := NewReviewRepository(db, zaptest.NewLogger(t))

	req := rcontracts.GetReviewsRequest{
		MinScore:       2,
		MaxScore:       4,
		TrainingPlanID: 1,
	}

	reviews, err := repository.GetReviews(ctx, req)
	assert.NoError(t, err)
	assert.Equal(t, len(reviews.Reviews), 0)
}
