package validator

import "errors"

var (
	ErrCheckNotDefined = errors.New("Check is not defined")
	ErrValIsNotString  = errors.New("Value is not a string")
	ErrValIsNotInt     = errors.New("value is not a int")
	ErrValToLong       = errors.New("Value is to long")
	ErrValToShort      = errors.New("Value is to short")
	ErrValToBig        = errors.New("Value is to big")
	ErrValToSmall      = errors.New("Value is to small")
)
