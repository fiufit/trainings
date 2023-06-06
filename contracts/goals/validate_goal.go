package goals

import (
	"github.com/fiufit/trainings/contracts"
)

var validGoalTypes = map[string]struct{}{
	"step_count":     {},
	"minutes_count":  {},
	"sessions_count": {},
}

var validGoalSubtypes = map[string]map[string]struct{}{
	"sessions_count": {
		"beginner":     {},
		"intermediate": {},
		"expert":       {},
		"strength":     {},
		"speed":        {},
		"endurance":    {},
		"lose weight":  {},
		"gain weight":  {},
		"sports":       {},
	},
}

func ValidateGoalType(goalType string, goalSubtype string) error {
	if _, ok := validGoalTypes[goalType]; !ok {
		return contracts.ErrInvalidGoalType
	}

	if subTypes, ok := validGoalSubtypes[goalType]; ok {
		if _, ok := subTypes[goalSubtype]; !ok {
			return contracts.ErrInvalidGoalSubtype
		}
	} else {
		if goalSubtype != "" {
			return contracts.ErrInvalidGoalSubtype
		}
	}

	return nil
}
