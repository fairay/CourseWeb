package errors

import (
	"errors"
)

var (
	RecordNotFound = errors.New("Record is not found")
	UnknownError = errors.New("Unknown error")
	WrongPassword = errors.New("Wrong password")

	UnknownAccount = errors.New("Account is not found")
	UnknownRole = errors.New("Role is not found")
	AccountExists = errors.New("Account associated with such login is already exists")
	
	DBAdditionError = errors.New("Eror in addition new account to DB")
)
