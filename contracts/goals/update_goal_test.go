package goals

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestValidateValidUpdateGoalRequestOk(t *testing.T) {
	req := &UpdateGoalRequest{
		Title:     "titulo",
		GoalValue: 10,
		Deadline:  time.Now().Add(24 * time.Hour),
		UserID:    "1",
		GoalID:    1,
	}
	err := req.Validate()
	assert.NoError(t, err)
}

func TestValidateInvalidUpdateGoalRequestErr(t *testing.T) {
	req := &UpdateGoalRequest{
		Title:     "titulo",
		GoalValue: 10,
		Deadline:  time.Now().Add(-24 * time.Hour),
		UserID:    "1",
		GoalID:    1,
	}
	err := req.Validate()
	assert.Error(t, err)
}
