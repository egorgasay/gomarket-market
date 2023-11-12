package constants

import "errors"

var (
	ErrInvalidLogin   = errors.New("invalid data for login")
	ErrRecordNotFound = errors.New("record not found")
)
