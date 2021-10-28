package errors

import (
	"errors"
)

var (
	RecordNotFound = errors.New("Account is not found")
	UnknownError = errors.New("Unknown error")
	WrongPassword = errors.New("Wrong password")
)