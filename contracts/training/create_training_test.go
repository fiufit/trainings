package training

import (
	"testing"
	"time"

	"github.com/fiufit/trainings/models"
	"github.com/go-playground/assert/v2"
	"github.com/undefinedlabs/go-mpatch"
)

func TestConvertToExerciseOk(t *testing.T) {

	req := ExerciseRequest{
		Title:       "Test exercise",
		Description: "Test description",
	}
	exercise := ConvertToExercise(req)

	assert.Equal(t, exercise.Title, "Test exercise")
	assert.Equal(t, exercise.Description, "Test description")
	assert.Equal(t, exercise.ID, int8(0))
	assert.Equal(t, exercise.TrainingPlanID, int8(0))
}

func TestConvertToExercisesWithEmptySliceOk(t *testing.T) {
	reqs := []ExerciseRequest{}
	exercises := ConvertToExercises(reqs)

	assert.Equal(t, []models.Exercise{}, exercises)
}

func TestConvertToExercisesOk(t *testing.T) {
	req1 := ExerciseRequest{
		Title:       "Test exercise 1",
		Description: "Test description 1",
	}

	req2 := ExerciseRequest{
		Title:       "Test exercise 2",
		Description: "Test description 2",
	}
	reqs := []ExerciseRequest{req1, req2}

	exercise1 := models.Exercise{
		ID:             int8(0),
		TrainingPlanID: int8(0),
		Title:          "Test exercise 1",
		Description:    "Test description 1",
	}

	exercise2 := models.Exercise{
		ID:             int8(0),
		TrainingPlanID: int8(0),
		Title:          "Test exercise 2",
		Description:    "Test description 2",
	}

	exercises := ConvertToExercises(reqs)

	assert.Equal(t, []models.Exercise{exercise1, exercise2}, exercises)
}

func TestConverToTrainingPlanOk(t *testing.T) {
	creationDate := time.Now()
	_, err := mpatch.PatchMethod(time.Now, func() time.Time {
		return creationDate
	})
	if err != nil {
		t.Fatal(err)
	}

	req1 := ExerciseRequest{
		Title:       "Test exercise 1",
		Description: "Test description 1",
	}

	req2 := ExerciseRequest{
		Title:       "Test exercise 2",
		Description: "Test description 2",
	}
	reqs := []ExerciseRequest{req1, req2}
	req := CreateTrainingRequest{
		Name:        "Test Training Name",
		Description: "Test Training Description",
		TrainerID:   "Test Trainer ID",
		Exercises:   reqs,
	}

	exercise1 := models.Exercise{
		ID:             int8(0),
		TrainingPlanID: int8(0),
		Title:          "Test exercise 1",
		Description:    "Test description 1",
	}

	exercise2 := models.Exercise{
		ID:             int8(0),
		TrainingPlanID: int8(0),
		Title:          "Test exercise 2",
		Description:    "Test description 2",
	}

	training := models.TrainingPlan{
		ID:          0,
		Name:        "Test Training Name",
		Description: "Test Training Description",
		TrainerID:   "Test Trainer ID",
		CreatedAt:   creationDate,
		Exercises:   []models.Exercise{exercise1, exercise2},
	}

	res := ConverToTrainingPlan(req)

	assert.Equal(t, res, training)
}
