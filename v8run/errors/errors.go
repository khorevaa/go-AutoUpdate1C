package errors

import (
	"fmt"
	"github.com/pkg/errors"
)

const (
	NoType = ErrorType(iota)
	BadRequest
	NotFound
	//add any type you want
)

type ErrorType uint

type customError struct {
	errorType     ErrorType
	originalError error
	contextInfo   map[string]string
}

// Error returns the mssage of a customError
func (error customError) Error() string {
	return error.originalError.Error()
}

// New creates a new customError
func (e ErrorType) New(msg string) error {
	return customError{errorType: e, originalError: errors.New(msg)}
}

// New creates a new customError with formatted message
func (e ErrorType) Newf(msg string, args ...interface{}) error {
	err := fmt.Errorf(msg, args...)

	return customError{errorType: e, originalError: err}
}

// Wrap creates a new wrapped error
func (e ErrorType) Wrap(err error, msg string) error {
	return e.Wrapf(err, msg)
}

// Wrap creates a new wrapped error with formatted message
func (e ErrorType) Wrapf(err error, msg string, args ...interface{}) error {
	newErr := errors.Wrapf(err, msg, args...)

	return customError{errorType: e, originalError: newErr}
}
