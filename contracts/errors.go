package contracts

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	ErrInternal                       = errors.New("something went wrong")
	ErrBadRequest                     = errors.New("unable to parse request")
	ErrForeignKey                     = errors.New("violates foreign key constraint")
	ErrUserInternal                   = errors.New("something went wrong")
	ErrUserBadRequest                 = errors.New("unable to parse request")
	ErrUserNotFound                   = errors.New("user not found")
	ErrTrainingPlanNotFound           = errors.New("training plan not found")
	ErrExerciseNotFound               = errors.New("exercise not found")
	ErrUnauthorizedTrainer            = errors.New("user is not the training creator")
	ErrSelfReview                     = errors.New("user is not allowed to review their own training")
	ErrReviewAlreadyExists            = errors.New("user already reviewed the training")
	ErrReviewNotFound                 = errors.New("review not found")
	ErrUnauthorizedReviewer           = errors.New("user is not the review creator")
	ErrInvalidTag                     = errors.New("invalid tag")
	ErrTrainingSessionNotFound        = errors.New("training session not found")
	ErrTrainingSessionNotComplete     = errors.New("can't complete training session without completing exercises")
	ErrTrainingSessionAlreadyFinished = errors.New("can't modify an already finished training session")
	ErrUnauthorizedAthlete            = errors.New("user is not the training session creator")
	ErrInvalidGoalType                = errors.New("invalid goal type")
	ErrInvalidGoalSubtype             = errors.New("invalid goal subtype")
	ErrGoalNotFound                   = errors.New("goal not found")
	ErrAlreadyLiked                   = errors.New("user already added training plan to favorites")
	ErrNotLiked                       = errors.New("user didn't add training plan to favorites")
	ErrTrainingAlreadyDisabled        = errors.New("training plan already disabled")
	ErrTrainingNotDisabled            = errors.New("training plan not disabled")
)

func HandleErrorType(ctx *gin.Context, err error) {
	var status int

	switch {
	case errors.Is(err, ErrBadRequest):
		status = http.StatusBadRequest
	case errors.Is(err, ErrTrainingPlanNotFound):
		status = http.StatusNotFound
	case errors.Is(err, ErrUnauthorizedTrainer):
		status = http.StatusUnauthorized
	case errors.Is(err, ErrSelfReview):
		status = http.StatusForbidden
	case errors.Is(err, ErrReviewAlreadyExists):
		status = http.StatusConflict
	case errors.Is(err, ErrReviewNotFound):
		status = http.StatusNotFound
	case errors.Is(err, ErrUnauthorizedReviewer):
		status = http.StatusUnauthorized
	case errors.Is(err, ErrInvalidTag):
		status = http.StatusBadRequest
	case errors.Is(err, ErrTrainingSessionNotFound):
		status = http.StatusNotFound
	case errors.Is(err, ErrTrainingSessionNotComplete):
		status = http.StatusBadRequest
	case errors.Is(err, ErrTrainingSessionAlreadyFinished):
		status = http.StatusBadRequest
	case errors.Is(err, ErrUnauthorizedAthlete):
		status = http.StatusUnauthorized
	case errors.Is(err, ErrInvalidGoalType):
		status = http.StatusBadRequest
	case errors.Is(err, ErrInvalidGoalSubtype):
		status = http.StatusBadRequest
	case errors.Is(err, ErrGoalNotFound):
		status = http.StatusNotFound
	case errors.Is(err, ErrAlreadyLiked):
		status = http.StatusConflict
	case errors.Is(err, ErrNotLiked):
		status = http.StatusConflict
	case errors.Is(err, ErrTrainingAlreadyDisabled):
		status = http.StatusConflict
	case errors.Is(err, ErrTrainingNotDisabled):
		status = http.StatusConflict
	case errors.Is(err, ErrForeignKey):
		status = http.StatusConflict
	case errors.Is(err, ErrUserNotFound):
		status = http.StatusNotFound
	case errors.Is(err, ErrExerciseNotFound):
		status = http.StatusNotFound
	case errors.Is(err, ErrUserInternal):
		status = http.StatusInternalServerError
	case errors.Is(err, ErrUserBadRequest):
		status = http.StatusBadRequest
	default:
		status = http.StatusInternalServerError
		ctx.JSON(status, FormatErrResponse(ErrInternal))
		return
	}
	ctx.JSON(status, FormatErrResponse(err))
}
