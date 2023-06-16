package users

type UserID struct {
	UserID string `json:"user_id" binding:"required"`
}
