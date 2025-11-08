package pkgerr

import (
	"errors"
	"fmt"
)

type Code int

const (
	CodeUnknownError Code = iota
	CodeValidationError
	CodeNotFoundError
	CodeUnauthorizedError
	CodeForbiddenError
	CodeDuplicateError
	CodeConflictError
	CodeTimeoutError
	CodeInternalError
	CodeTransactionError
	CodeCancelledError
	CodeAbortedError
	CodeRetryableError
)

type Error struct {
	Code    Code
	Key     string
	Message string
	Root    error
}

func (e Error) GetRoot() error     { return e.Root }
func (e Error) GetKey() string     { return e.Key }
func (e Error) GetCode() Code      { return e.Code }
func (e Error) GetMessage() string { return e.Message }
func (e Error) GetCodeMessage() string {
	switch e.Code {
	case CodeValidationError:
		return "Validation error"
	case CodeNotFoundError:
		return "Not found error"
	case CodeUnauthorizedError:
		return "Unauthorized error"
	case CodeForbiddenError:
		return "Forbidden error"
	case CodeDuplicateError:
		return "Duplicate error"
	case CodeConflictError:
		return "Conflict error"
	case CodeTimeoutError:
		return "Timeout error"
	case CodeInternalError:
		return "Internal error"
	case CodeTransactionError:
		return "Transaction error"
	case CodeCancelledError:
		return "Cancelled error"
	case CodeAbortedError:
		return "Aborted error"
	case CodeRetryableError:
		return "Retryable error"
	default:
		return "Unknown error"
	}
}

func (e Error) Error() string {
	return fmt.Sprintf("key: %s, message: %s, root: %v", e.Key, e.Message, e.Root)
}

func NewError(code Code, key, message string, root error) error {
	return &Error{Code: code, Key: key, Message: message, Root: root}
}

// Helper constructors
func NewValidationError(key, message string, root error) error {
	return NewError(CodeValidationError, key, message, root)
}
func NewNotFoundError(key, message string, root error) error {
	return NewError(CodeNotFoundError, key, message, root)
}
func NewUnauthorizedError(key, message string, root error) error {
	return NewError(CodeUnauthorizedError, key, message, root)
}
func NewForbiddenError(key, message string, root error) error {
	return NewError(CodeForbiddenError, key, message, root)
}
func NewDuplicateError(key, message string, root error) error {
	return NewError(CodeDuplicateError, key, message, root)
}
func NewConflictError(key, message string, root error) error {
	return NewError(CodeConflictError, key, message, root)
}
func NewTimeoutError(key, message string, root error) error {
	return NewError(CodeTimeoutError, key, message, root)
}
func NewInternalError(key, message string, root error) error {
	return NewError(CodeInternalError, key, message, root)
}
func NewTransactionError(key, message string, root error) error {
	return NewError(CodeTransactionError, key, message, root)
}
func NewCancelledError(key, message string, root error) error {
	return NewError(CodeCancelledError, key, message, root)
}
func NewAbortedError(key, message string, root error) error {
	return NewError(CodeAbortedError, key, message, root)
}
func NewRetryableError(key, message string, root error) error {
	return NewError(CodeRetryableError, key, message, root)
}

// Type checks
func Check(err error) (*Error, bool) {
	var appErr *Error
	if errors.As(err, &appErr) {
		return appErr, true
	}
	return nil, false
}

func IsValidationError(err error) bool   { return IsCode(err, CodeValidationError) }
func IsNotFoundError(err error) bool     { return IsCode(err, CodeNotFoundError) }
func IsUnauthorizedError(err error) bool { return IsCode(err, CodeUnauthorizedError) }
func IsForbiddenError(err error) bool    { return IsCode(err, CodeForbiddenError) }
func IsDuplicateError(err error) bool    { return IsCode(err, CodeDuplicateError) }
func IsConflictError(err error) bool     { return IsCode(err, CodeConflictError) }
func IsTimeoutError(err error) bool      { return IsCode(err, CodeTimeoutError) }
func IsInternalError(err error) bool     { return IsCode(err, CodeInternalError) }
func IsTransactionError(err error) bool  { return IsCode(err, CodeTransactionError) }
func IsCancelledError(err error) bool    { return IsCode(err, CodeCancelledError) }
func IsAbortedError(err error) bool      { return IsCode(err, CodeAbortedError) }
func IsRetryableError(err error) bool    { return IsCode(err, CodeRetryableError) }

func IsCode(err error, code Code) bool {
	var appErr *Error
	return errors.As(err, &appErr) && appErr.GetCode() == code
}

func IsKey(err error, key string) bool {
	var appErr *Error
	return errors.As(err, &appErr) && appErr.GetKey() == key
}
