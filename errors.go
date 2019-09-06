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
	// ErrCheckNotDefined is used when a check in a tag is not found
	ErrCheckNotDefined error = ErrorWithKey{errors.New("NOT_DEFINED"), errors.New("Check is not defined")}

	// ErrValIsNotString is used when the error value is not a string
	ErrValIsNotString error = ErrorWithKey{errors.New("NOT_STRING"), errors.New("Value is not a string")}

	// ErrValIsNotInt is used when the value is not an int
	ErrValIsNotInt error = ErrorWithKey{errors.New("NOT_INT"), errors.New("value is not a int")}

	// ErrValToLong is used when the value is to long
	ErrValToLong error = ErrorWithKey{errors.New("VAL_TO_LONG"), errors.New("Value is to long")}

	// ErrValToShort is used when the value is to short
	ErrValToShort error = ErrorWithKey{errors.New("VAL_TO_SHORT"), errors.New("Value is to short")}

	// ErrValToBig is used when the value is to big
	ErrValToBig error = ErrorWithKey{errors.New("VAL_TO_BIG"), errors.New("Value is to big")}

	// ErrValToSmall is used when the value is to small
	ErrValToSmall error = ErrorWithKey{errors.New("VAL_TO_SMALL"), errors.New("Value is to small")}
)
