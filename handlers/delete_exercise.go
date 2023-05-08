package handlers

import "go.uber.org/zap"

type DeleteExercise struct {
	//exercises usecases.ExerciseDeleter
	logger *zap.Logger
}

// func NewDeleteExercise(exercises usecases.ExerciseDeleter, logger *zap.Logger) DeleteExercise {
// 	return DeleteExercise{exercises: exercises, logger: logger}
// }
