# `go-validator` A nice little validator for go
*Based on the idea of [asaskevich/govalidator](https://github.com/asaskevich/govalidator) but with more checking data and with more extensive responses*  
[![GoDoc](https://godoc.org/github.com/klippa-app/go-validator?status.svg)](https://godoc.org/github.com/klippa-app/go-validator)

### Example
```go
package main

import (
	"fmt"

	"github.com/klippa-app/go-validator"
)

type userInput struct {
	Name     string `valid:"minLength 1, maxLength 200"`
	Password string `valid:"password"`
}

func main() {
	checker := validator.NewChecker()
	checker.AddCheck("minLength", validator.Checks.Strings.MinLength)
	checker.AddCheck("maxLength", validator.Checks.Strings.Maxlength)
	checker.AddCheck("password", validator.Checks.Strings.Password)

	output := checker.Check(userInput{
		Name:     "Mario",
		Password: "Jump",
	})

  fmt.Println(output)
  // > map[Password:[Value is to short]]
}
```

### Docs
For all documentation look on: [godoc.org/github.com/klippa-app/go-validator](https://godoc.org/github.com/klippa-app/go-validator)  

#### Add a custom check
The package also has some build in check that can be found in: `validator.Checks`  
```go
checker := validator.NewChecker()
checker.AddCheck("notNill", func(c *validator.Context) error {
  slice, ok := c.Val.([]string)
  if !ok || slice == nil {
    return errors.New("Value is not a valid slice")
  }
  return nil
})

type Test struct {
  List []string `valid:"notNill"`
}

output := checker.Check(Test{
  List: nil,
})

// output = map[List:[Value is not a valid slice]]
```

#### Use json tag as error key
```go
checker := validator.NewChecker(validator.Options{
  JSONTag: &validator.JSONTag{} // Add the JSONTag here
})
checker.AddCheck("password", validator.Checks.Strings.Password)

type userInput struct {
	Password string `json:"pass" valid:"password"`
}
output := checker.Check(userInput{
  Password: "abcd",
})

// The output Password field now has "pass" as key
// output = map[pass:[Value is to short]]
```

### Q & A

#### Can i use nested values?
Yes you can use pointers, structs in structs and slices.  
The responses will also change if you use this:  
- Struct with Struct: `structField1.structField2`
- Struct with Array with Structs: `structField.3.structField`

#### Can i validate private fields?
Yes you can
