package exercises

type UpdateExerciseRequest struct {
	Title          string `json:"title"`
	Description    string `json:"description"`
	TrainerID      string `json:"trainer_id" binding:"required"`
	TrainingPlanID uint
	ExerciseID     uint
}
