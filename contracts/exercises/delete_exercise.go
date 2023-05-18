package exercises

type DeleteExerciseRequest struct {
	TrainerID      string `json:"trainer_id" binding:"required"`
	TrainingPlanID uint
	ExerciseID     uint
}
