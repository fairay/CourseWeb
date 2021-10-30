package errors

import (
	"errors"
)

var (
	AccessDeleteDenied = errors.New("Access to delete is denied")
)