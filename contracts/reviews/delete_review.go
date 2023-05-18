package reviews

type DeleteReviewRequest struct {
	UserID         string `json:"user_id" binding:"required"`
	TrainingPlanID uint
	ReviewID       uint
}
