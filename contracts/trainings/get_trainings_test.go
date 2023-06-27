package trainings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateValidGetTrainingsRequestOk(t *testing.T) {

	req := GetTrainingsRequest{
		UserID:     "test",
		TagStrings: []string{"strength", "speed"},
	}
	err := req.Validate()
	assert.NoError(t, err)
}

func TestValidateInvalidGetTrainingsRequestErr(t *testing.T) {

	req := GetTrainingsRequest{
		UserID:     "test",
		TagStrings: []string{"invalid", "speed"},
	}
	err := req.Validate()
	assert.Error(t, err)
}
