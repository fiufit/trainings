package training

type DeleteTrainingRequest struct {
	TrainerID      string `json:"trainer_id" binding:"required"`
	TrainingPlanID string
}
