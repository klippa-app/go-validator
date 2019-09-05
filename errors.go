package validator

import "errors"

// ErrorWithKey is a wrapper around error with an extra Key value
type ErrorWithKey struct {
	ErrorKey  error
	FullError error
}

func (e ErrorWithKey) String() string {
	return e.FullError.Error()
}

func (e ErrorWithKey) Error() string {
	return e.FullError.Error()
}

var (
	ErrCheckNotDefined error = ErrorWithKey{errors.New("NOT_DEFINED"), errors.New("Check is not defined")}
	ErrValIsNotString  error = ErrorWithKey{errors.New("NOT_STRING"), errors.New("Value is not a string")}
	ErrValIsNotInt     error = ErrorWithKey{errors.New("NOT_INT"), errors.New("value is not a int")}
	ErrValToLong       error = ErrorWithKey{errors.New("VAL_TO_LONG"), errors.New("Value is to long")}
	ErrValToShort      error = ErrorWithKey{errors.New("VAL_TO_SHORT"), errors.New("Value is to short")}
	ErrValToBig        error = ErrorWithKey{errors.New("VAL_TO_BIG"), errors.New("Value is to big")}
	ErrValToSmall      error = ErrorWithKey{errors.New("VAL_TO_SMALL"), errors.New("Value is to small")}
)
