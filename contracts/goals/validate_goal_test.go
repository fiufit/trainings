package goals

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateValidTypeWithNoSubtypeOk(t *testing.T) {
	err := ValidateGoalType("step_count", "")
	assert.NoError(t, err)
}

func TestValidateValidTypeWithNoSubtypeErr(t *testing.T) {
	err := ValidateGoalType("step_count", "subtipo que no existe")
	assert.Error(t, err)

}

func TestValidateInvalidTypeWithNoSubtypeErr(t *testing.T) {
	err := ValidateGoalType("tipo que no existe", "")
	assert.Error(t, err)
}

func TestValidateInvalidTypeWithNoSubtype2Err(t *testing.T) {
	err := ValidateGoalType("tipo que no existe", "subtipo que no existe")
	assert.Error(t, err)
}

func TestValidateValidTypeWithSubtypeOk(t *testing.T) {
	err := ValidateGoalType("sessions_count", "beginner")
	assert.NoError(t, err)
}

func TestValidateValidTypeWithSubtypeErr(t *testing.T) {
	err := ValidateGoalType("sessions_count", "")
	assert.Error(t, err)

}

func TestValidateValidTypeWithSubtypeErr2(t *testing.T) {
	err := ValidateGoalType("sessions_count", "subtipo que no existe")
	assert.Error(t, err)
}
