package reviews

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateValidCreateReviewOk(t *testing.T) {
	req := &CreateReviewRequest{
		Score:          5,
		Comment:        "test",
		UserID:         "test",
		TrainingPlanID: 1,
	}
	err := req.Validate()
	assert.NoError(t, err)
}

func TestValidateInvalidCreateReviewWithScoreGreatenThanFiveErr(t *testing.T) {
	req := &CreateReviewRequest{
		Score:          6,
		Comment:        "test",
		UserID:         "test",
		TrainingPlanID: 1,
	}
	err := req.Validate()
	assert.Error(t, err)
}

func TestValidateInvalidCreateReviewWithScoreLessThanOneErr(t *testing.T) {
	req := &CreateReviewRequest{
		Score:          0,
		Comment:        "test",
		UserID:         "test",
		TrainingPlanID: 1,
	}
	err := req.Validate()
	assert.Error(t, err)
}
