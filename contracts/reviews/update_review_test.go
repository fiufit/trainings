package reviews

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateValidUpateReviewOk(t *testing.T) {
	req := &UpdateReviewRequest{
		Score:          5,
		Comment:        "test",
		UserID:         "test",
		TrainingPlanID: 1,
		ReviewID:       1,
	}
	err := req.Validate()
	assert.NoError(t, err)
}

func TestValidateInvalidUpdateReviewWithScoreGreatenThanFiveErr(t *testing.T) {
	req := &UpdateReviewRequest{
		Score:          6,
		Comment:        "test",
		UserID:         "test",
		TrainingPlanID: 1,
		ReviewID:       1,
	}
	err := req.Validate()
	assert.Error(t, err)
}

func TestValidateInvalidUpdateReviewWithScoreLessThanOneErr(t *testing.T) {
	req := &UpdateReviewRequest{
		Score:          0,
		Comment:        "test",
		UserID:         "test",
		TrainingPlanID: 1,
		ReviewID:       1,
	}
	err := req.Validate()
	assert.Error(t, err)
}
