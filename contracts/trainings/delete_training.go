package trainings

type DeleteTrainingRequest struct {
	TrainerID      string `json:"trainer_id" binding:"required"`
	TrainingPlanID uint
}
