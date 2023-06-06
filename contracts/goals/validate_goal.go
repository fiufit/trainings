package goals

import (
	"github.com/fiufit/trainings/contracts"
)

var validGoalTypes = map[string]struct{}{
	"step count":     {},
	"minutes count":  {},
	"sessions count": {},
}

var validGoalSubtypes = map[string]map[string]struct{}{
	"sessions count": {
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
