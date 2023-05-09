package training

type UpdateExerciseRequest struct {
	Title          string `json:"title"`
	Description    string `json:"description"`
	TrainerID      string `json:"trainer_id" binding:"required"`
	TrainingPlanID string
	ExerciseID     string
}
