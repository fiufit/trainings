package contracts

import "errors"

var (
	ErrInternal     = errors.New("something went wrong")
	ErrBadRequest   = errors.New("unable to parse request")
	ErrForeignKey   = errors.New("violates foreign key constraint")
	ErrUserNotFound = errors.New("user not found")
)
