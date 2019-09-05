package validator

import (
	"strconv"
)

// ChecksT contains defines all checks
type ChecksT struct {
	// Strings contains checks realted to strings
	Strings StringChecksT

	// Int contains checks related to intages
	Int IntChecksT
}

// StringChecksT contains all checks related to strings
type StringChecksT struct {
	// MinLength checks the minimal length of a string
	MinLength Check

	// Maxlength checks the minimal length of a string
	Maxlength Check

	// Password checks a password
	Password Check
}

// IntChecksT contains all checks related to strings
type IntChecksT struct {
	HigherThan Check
	LowerThan  Check
}

// Checks is a list of pre defined checks that
// can be added by the user to the *Checker
var Checks = ChecksT{
	Strings: StringChecksT{
		Maxlength: func(c *Context) error {
			val, ok := c.Val.(string)
			if !ok {
				return ErrValIsNotString
			}
			arg, err := strconv.Atoi(c.CheckArg)
			if err != nil {
				return err
			}
			if len(val) > arg {
				return ErrValToLong
			}
			return nil
		},
		MinLength: func(c *Context) error {
			val, ok := c.Val.(string)
			if !ok {
				return ErrValIsNotString
			}
			arg, err := strconv.Atoi(c.CheckArg)
			if err != nil {
				return err
			}
			if len(val) < arg {
				return ErrValToShort
			}
			return nil
		},
		Password: func(c *Context) error {
			val, ok := c.Val.(string)
			if !ok {
				return ErrValIsNotString
			}
			if len(val) > 8 {
				return nil
			}
			return ErrValToShort
		},
	},
	Int: IntChecksT{
		HigherThan: func(c *Context) error {
			arg, err := strconv.Atoi(c.CheckArg)
			if err != nil {
				return err
			}

			toCheck := 0
			switch c.Val.(type) {
			case int:
				toCheck = c.Val.(int)
			case int8:
				toCheck = int(c.Val.(int8))
			case int16:
				toCheck = int(c.Val.(int16))
			case int32:
				toCheck = int(c.Val.(int32))
			case int64:
				toCheck = int(c.Val.(int64))
			case uint:
				toCheck = int(c.Val.(uint))
			case uint8:
				toCheck = int(c.Val.(uint8))
			case uint16:
				toCheck = int(c.Val.(uint16))
			case uint32:
				toCheck = int(c.Val.(uint32))
			case uint64:
				toCheck = int(c.Val.(uint64))
			case float32:
				toCheck = int(c.Val.(float32))
			case float64:
				toCheck = int(c.Val.(float32))
			default:
				return ErrValIsNotInt
			}

			if arg > toCheck {
				return ErrValToSmall
			}
			return nil
		},
		LowerThan: func(c *Context) error {
			arg, err := strconv.Atoi(c.CheckArg)
			if err != nil {
				return err
			}

			toCheck := 0
			switch c.Val.(type) {
			case int:
				toCheck = c.Val.(int)
			case int8:
				toCheck = int(c.Val.(int8))
			case int16:
				toCheck = int(c.Val.(int16))
			case int32:
				toCheck = int(c.Val.(int32))
			case int64:
				toCheck = int(c.Val.(int64))
			case uint:
				toCheck = int(c.Val.(uint))
			case uint8:
				toCheck = int(c.Val.(uint8))
			case uint16:
				toCheck = int(c.Val.(uint16))
			case uint32:
				toCheck = int(c.Val.(uint32))
			case uint64:
				toCheck = int(c.Val.(uint64))
			case float32:
				toCheck = int(c.Val.(float32))
			case float64:
				toCheck = int(c.Val.(float32))
			default:
				return ErrValIsNotInt
			}

			if arg < toCheck {
				return ErrValToBig
			}
			return nil
		},
	},
}
