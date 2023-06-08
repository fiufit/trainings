package goals

type DeleteGoalRequest struct {
	UserID string `json:"user_id" binding:"required"`
	GoalID uint
}
