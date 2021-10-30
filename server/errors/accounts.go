package errors

import (
	"errors"
)

var (
	RecordNotFound = errors.New("Record is not found")
	UnknownError = errors.New("Unknown error")
	WrongPassword = errors.New("Wrong password")

	UnknownAccount = errors.New("Account is not found")
)
