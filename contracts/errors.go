package contracts

import "errors"

var (
	ErrInternal       = errors.New("something went wrong")
	ErrBadRequest     = errors.New("unable to parse request")
	ErrForeignKey     = errors.New("violates foreign key constraint")
	ErrUserInternal   = errors.New("something went wrong")
	ErrUserBadRequest = errors.New("unable to parse request")
	ErrUserNotFound   = errors.New("user not found")
)
