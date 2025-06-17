package errors

import (
	"errors"
	"fmt"
)

var (
	ErrNotFound     = errors.New("resource not found")
	ErrDB           = errors.New("error from the database")
	ErrMappingData  = errors.New("error mapping data")
	ErrInvalidInput = errors.New("invalid input")
)

func WrapError(baseError error, detail string) error {
	return fmt.Errorf("%w, %s", baseError, detail)
}
