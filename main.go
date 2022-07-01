package main

import (
	"database/sql/driver"
	"fmt"
	"reflect"
	"regexp"
	"time"
)

type Argument interface {
	Match(driver.Value) bool
}

type AnyTime struct{}

// Match satisfies sqlmock.Argument interface
func (a AnyTime) Match(v driver.Value) bool {
	_, ok := v.(time.Time)
	return ok
}

type AnyToken struct{}

func (AnyToken) Match(v driver.Value) bool {
	re := regexp.MustCompile(`(?m)[a-f0-9]{32}`)
	if code, ok := v.(string); ok {
		return re.MatchString(code)
	}
	return false
}

func main() {
  // As interface types are only used for static typing, a
	// common idiom to find the reflection Type for an interface
	// type Foo is to use a *Foo value.
	// writerType := reflect.TypeOf((*io.Writer)(nil)).Elem()

	// fileType := reflect.TypeOf((*os.File)(nil))
	// fmt.Println(fileType.Implements(writerType))

  var (
    v0           = AnyToken{}
    v1           = AnyTime{}
    v2           = ""
    argumentType = reflect.TypeOf((*Argument)(nil)).Elem()
  )
  
	fmt.Printf("v0 implements Argument: %t\n", reflect.TypeOf(v0).Implements(argumentType))
  fmt.Printf("v1 implements Argument: %t\n", reflect.TypeOf(v1).Implements(argumentType))
  fmt.Printf("v2 implements Argument: %t\n", reflect.TypeOf(v2).Implements(argumentType))
}
