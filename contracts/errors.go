package contracts

import "errors"

var (
	ErrInternal             = errors.New("something went wrong")
	ErrBadRequest           = errors.New("unable to parse request")
	ErrForeignKey           = errors.New("violates foreign key constraint")
	ErrUserInternal         = errors.New("something went wrong")
	ErrUserBadRequest       = errors.New("unable to parse request")
	ErrUserNotFound         = errors.New("user not found")
	ErrTrainingPlanNotFound = errors.New("training plan not found")
	ErrExerciseNotFound     = errors.New("exercise not found")
	ErrUnauthorizedTrainer  = errors.New("user is not the training creator")
	ErrSelfReview           = errors.New("user is not allowed to review their own training")
	ErrReviewAlreadyExists  = errors.New("user already reviewed the training")
)
