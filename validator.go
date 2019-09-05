package validator

import (
	"reflect"
	"strconv"
	"strings"
	"unsafe"
)

// Options are the options for running the ch
type Options struct {
	// JSONTagOptions uses the json tag (SomeItem string `json:"someItem"`) as
	// key in the Check return map
	// When this pointer is null this will be ignored
	JSONTag *JSONTag
}

// JSONTag are options for working with the json struct tag
type JSONTag struct {
	// IgnoreDashFields skips all fields with json:"-"
	IgnoreDashFields bool
}

// Checker contains the settings for the checker
type Checker struct {
	JSONTag       *JSONTag
	DefinedChecks map[string]Check
}

// Context contains the data for when running a check
type Context struct {
	Val interface{}

	CheckName string
	CheckArg  string

	FieldName string
	FieldPath string
}

// Check defines a check
type Check func(c *Context) error

// NewChecker creates a new pointer to a checker object
func NewChecker(options ...Options) *Checker {
	option := Options{}
	if len(options) > 0 {
		option = options[0]
	}
	return &Checker{
		JSONTag:       option.JSONTag,
		DefinedChecks: map[string]Check{},
	}
}

// AddCheck defines a new check
func (c *Checker) AddCheck(name string, check Check) {
	if check == nil {
		return
	}
	c.DefinedChecks[name] = check
}

// ErrorsMap is a map with errors for every field
type ErrorsMap map[string][]error

func (m ErrorsMap) Error() string {
	lines := []string{}
	for key, errors := range m {
		newLine := key + ":"
		for i, err := range errors {
			if i > 0 {
				newLine += ", "
			}
			newLine += err.Error()
		}
	}
	return strings.Join(lines, "\n")
}

// addError adds an error to the error map
func (m *ErrorsMap) addError(key string, err error) {
	if err == nil {
		return
	}

	field := (*m)[key]
	if field == nil {
		field = []error{}
	}
	field = append(field, err)
	(*m)[key] = field
}

// Check checks the input for problems
func (c *Checker) Check(input interface{}) ErrorsMap {
	errors := &ErrorsMap{}
	c.checkFieldType(errors, []string{}, reflect.ValueOf(input))
	return *errors
}

func (c *Checker) checkStruct(errors *ErrorsMap, path []string, input reflect.Value) {
	inputType := input.Type()
	for i := 0; i < input.NumField(); i++ {
		if !input.Field(i).IsValid() {
			return
		}

		copyOfInputType := reflect.New(inputType).Elem()
		copyOfInputType.Set(input)

		field := copyOfInputType.Field(i)
		field = reflect.NewAt(field.Type(), unsafe.Pointer(field.UnsafeAddr())).Elem()
		fieldType := inputType.Field(i)

		// Get the name prop
		name := fieldType.Name
		if c.JSONTag != nil {
			if val, ok := fieldType.Tag.Lookup("json"); ok {
				if val == "-" {
					if c.JSONTag.IgnoreDashFields {
						continue
					}
				} else {
					name = strings.Split(val, ",")[0]
				}
			}
		}

		path := append(path, name)

		// Check if the valid tag exsists and if so check the struct value
		validTag, found := fieldType.Tag.Lookup("valid")
		if found {
			checksToRun := strings.Split(validTag, ",")
			for _, check := range checksToRun {
				check = strings.Trim(check, " ")
				checkNameAndArgs := strings.Split(check, " ")

				checkFunction, ok := c.DefinedChecks[checkNameAndArgs[0]]
				if !ok || checkFunction == nil {
					errors.addError(strings.Join(path, "."), ErrCheckNotDefined)
					continue
				}

				c := &Context{
					CheckName: checkNameAndArgs[0],
					FieldName: name,
					FieldPath: strings.Join(path, "."),
					Val:       field.Interface(),
				}
				if len(checkNameAndArgs) > 1 {
					c.CheckArg = checkNameAndArgs[1]
				}

				err := checkFunction(c)
				if err != nil {
					errors.addError(strings.Join(path, "."), err)
				}
			}
		}

		c.checkFieldType(errors, path, field)
	}
}

func (c *Checker) checkFieldType(errors *ErrorsMap, path []string, input reflect.Value) {
	if !input.IsValid() {
		return
	}

	switch input.Kind() {
	case reflect.Struct:
		c.checkStruct(errors, path, input)
	case reflect.Ptr:
		c.checkFieldType(errors, path, input.Elem())
	case reflect.Slice:
		for i := 0; i < input.Len(); i++ {
			sliceItem := input.Index(i)
			c.checkFieldType(errors, append(path, strconv.Itoa(i)), sliceItem)
		}
	}
}
