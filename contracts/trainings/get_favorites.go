package trainings

import "github.com/fiufit/trainings/contracts"

type GetFavoritesRequest struct {
	UserID string `form:"user_id" binding:"required"`
	contracts.Pagination
}
