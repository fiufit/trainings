package training

type CreateExerciseRequest struct {
	Title          string `json:"title" binding:"required"`
	Description    string `json:"description" binding:"required"`
	TrainerID      string `json:"trainer_id" binding:"required"`
	TrainingPlanID uint
}
