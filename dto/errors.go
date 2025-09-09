package dto

import "errors"

var (
	ErrInternal      = errors.New("Internal error")
	ErrBadRequest    = errors.New("Bad request")
	ErrNotFound      = errors.New("Not found")
	ErrAlreadyExists = errors.New("Already exists")
	ErrBadData       = errors.New("Bad data")
)
