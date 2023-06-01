package training_sessions

type UpdateExerciseSessionRequest struct {
	ID   uint `json:"id" binding:"required"`
	Done bool `json:"done" binding:"required"`
}

type UpdateTrainingSessionRequest struct {
	RequesterID      string                         `json:"requester_id" binding:"required"`
	ID               uint                           `json:"id" binding:"required"`
	ExerciseSessions []UpdateExerciseSessionRequest `json:"exercise_sessions" binding:"required"`
	Done             bool                           `json:"done" binding:"required"`
	StepCount        uint                           `json:"step_count" binding:"required"`
	SecondsCount     uint                           `json:"seconds_count" binding:"required"`
}

type UpdateTrainingSessionResponse CreateTrainingSessionResponse
