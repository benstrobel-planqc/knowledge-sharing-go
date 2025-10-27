package example

import (
	"fmt"
	"time"
)

func Divide(a int, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}
	return a / b, nil
}

func ErrorDemo() {
	result, err := Divide(10, 0)
	if err != nil {
		fmt.Printf("Error occurred: %v\n", err)
		return
	}

	fmt.Printf("Result: %v\n", result)
}

// "error" is a built-in interface:
// type error interface {
//	  Error() string
// }

type MyCustomError struct {
	Detail string
	When   time.Time
}

func (e *MyCustomError) Error() string {
	return fmt.Sprintf("Error at %v: %s", e.When, e.Detail)
}

func Divide2(a int, b int) (int, error) {
	if b == 0 {
		return 0, &MyCustomError{
			Detail: "cannot divide by zero",
			When:   time.Now(),
		}
	}
	return a / b, nil
}
