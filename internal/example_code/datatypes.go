package example

import "fmt"

func DemoDatatypes() {

	// Variables & Zero init
	var b bool    // = false
	var i int     // = 0
	var f float32 // = 0.0
	var s string  // = ""

	// bool string
	// int int8 int16 int32 int64
	// uint uint8 uint16 uint32 uint64 uintptr
	// float32 float64
	// complex64 complex128

	// Short declaration
	var someString string = "hello"
	sameAsAbove := "hello"

	// Pointers
	var stringLiteral string = "hello"
	var stringPointer *string = &stringLiteral
	var pointersCanBeNil *string = nil

	stringLiteral = "hello2"
	var stringLiteralCopy string = *stringPointer
	fmt.Printf("%v", stringLiteralCopy) // "hello2"

	// Arrays & Slices
	var arr [3]int                  // = [0,0,0]
	var slc []int                   // = nil
	var slc2 []int = make([]int, 0) // = []
	slc3 := []int{1, 2, 3, 4, 5}    // = [1,2,3,4,5]
	slc4 := slc3[1:4]               // = [2,3,4]

	// Maps
	m := make(map[string]int)
	m["one"] = 1
	m2 := map[string]int{"two": 2, "three": 3}
	delete(m2, "two")

	// Functions as values
	var square func(n int) int = func(n int) int {
		return n * n
	}
	fmt.Printf("%v", square(5)) // 25

	// Print all to avoid unused variable errors for the screenshots
	fmt.Printf("%v %v %v %v %v %v %v %v %v %v %v %v %v", b, i, f, s, someString, sameAsAbove, stringLiteralCopy, pointersCanBeNil, arr, slc, slc2, slc4, m2)
}
