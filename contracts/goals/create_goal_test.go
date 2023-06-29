package goals

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/undefinedlabs/go-mpatch"
)

func TestValidateValidCreateGoalRequestOk(t *testing.T) {
	creationDate := time.Now()
	_, err := mpatch.PatchMethod(time.Now, func() time.Time {
		return creationDate
	})
	if err != nil {
		t.Fatal(err)
	}
	req := CreateGoalRequest{
		Title:       "test",
		GoalType:    "sessions count",
		GoalSubtype: "beginner",
		GoalValue:   10,
		Deadline:    creationDate.Add(24 * time.Hour),
		UserID:      "test",
	}
	err = req.Validate()
	assert.NoError(t, err)
}

func TestValidateInvalidCreateGoalRequestErr(t *testing.T) {
	creationDate := time.Now()
	req := CreateGoalRequest{
		Title:       "test",
		GoalType:    "sessions count",
		GoalSubtype: "beginner",
		GoalValue:   10,
		Deadline:    creationDate.Add(-24 * time.Hour),
		UserID:      "test",
	}
	err := req.Validate()
	assert.Error(t, err)
}
