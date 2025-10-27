package example

import (
	"fmt"
	"runtime"
)

// If
func DemoIf(x int) {
	if x%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}
}

func DemoFor(n int) {
	// Classic for
	for i := 0; i < n; i++ {
		fmt.Printf("%v\n", i)
	}

	// While
	sum := 0
	for sum < n {
		sum += sum + 1
	}

	// Forever
	for {
		break
	}

	// For each
	values := []int{1, 2, 3, 4, 5}
	for index, value := range values {
		fmt.Printf("%v: %v\n", index, value)
	}
}

func Helper() string {
	return "test"
}

// Switch
func DemoSwitch(s Stringer) {
	os := runtime.GOOS

	switch os {
	case "darwin":
		fmt.Println("Running on macOS.")
	case "linux":
		fmt.Println("Running on linux.")
	case Helper():
		fmt.Println("Running on test.")
	default:
		fmt.Printf("Running on %s.\n", os)
	}

	switch s.(type) {
	case Vertex:
		fmt.Println("It's a Vertex.")
	default:
		fmt.Println("Unknown type.")
	}
}
