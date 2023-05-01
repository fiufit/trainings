package contracts

import "errors"

var (
	ErrInternal   = errors.New("something went wrong")
	ErrBadRequest = errors.New("unable to parse request")
)
