package trainings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateValidUpdateTrainingRequestOk(t *testing.T) {
	req := UpdateTrainingRequest{
		BaseTrainingRequest: BaseTrainingRequest{
			Name:        "Test training",
			Description: "Test description",
			TagStrings:  []string{"strength", "speed"},
			Exercises: []ExerciseRequest{
				{
					Title:       "Test exercise",
					Description: "Test description",
				},
			},
		},
	}
	err := req.Validate()
	assert.NoError(t, err)
}

func TestValidateInvalidUpdateTrainingRequestErr(t *testing.T) {
	req := UpdateTrainingRequest{
		BaseTrainingRequest: BaseTrainingRequest{
			Name:        "Test training",
			Description: "Test description",
			TagStrings:  []string{"invalid", "speed"},
			Exercises: []ExerciseRequest{
				{
					Title:       "Test exercise",
					Description: "Test description",
				},
			},
		},
	}
	err := req.Validate()
	assert.Error(t, err)
}
