package validator_test

import (
	"testing"

	"github.com/klippa-app/go-validator"
	"github.com/stretchr/testify/assert"
)

type testType struct {
	Name     string `json:"-" valid:"minLength 1, maxLength 200"`
	Password string `json:"pass" valid:"password"`
}

func TestNewChecker(t *testing.T) {
	assert.NotNil(t, validator.NewChecker())
}

func TestAddCheck(t *testing.T) {
	checker := validator.NewChecker()
	checker.AddCheck("password", validator.Checks.Strings.Password)
	check, ok := checker.DefinedChecks["password"]
	if !ok {
		assert.Fail(t, "password is not defined")
		return
	}
	assert.NotNil(t, check)
}

func TestCheck(t *testing.T) {
	checker := validator.NewChecker()
	checker.AddCheck("minLength", validator.Checks.Strings.MinLength)
	checker.AddCheck("maxLength", validator.Checks.Strings.Maxlength)
	checker.AddCheck("password", validator.Checks.Strings.Password)

	output := checker.Check(testType{
		Name:     "Mario",
		Password: "Jumping 1234",
	})
	assert.Equal(t, validator.ErrorsMap{}, output)

	output = checker.Check(testType{
		Name:     "Mario",
		Password: "Jump",
	})
	assert.Equal(t, validator.ErrorsMap{"Password": []error{validator.ErrValToShort}}, output)

	checker.JSONTag = &validator.JSONTag{}
	output = checker.Check(testType{
		Name:     "Mario",
		Password: "Jump",
	})
	assert.Equal(t, validator.ErrorsMap{"pass": []error{validator.ErrValToShort}}, output)

	output = checker.Check(testType{
		Name:     "",
		Password: "Jumping 1234",
	})
	assert.Equal(t, validator.ErrorsMap{"Name": []error{validator.ErrValToShort}}, output)

	checker.JSONTag = &validator.JSONTag{IgnoreDashFields: true}
	output = checker.Check(testType{
		Name:     "",
		Password: "Jumping 1234",
	})
	assert.Equal(t, validator.ErrorsMap{}, output)
}
