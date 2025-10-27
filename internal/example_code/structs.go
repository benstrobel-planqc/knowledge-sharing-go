package example

import (
	"fmt"
)

// Structs and field visibility
type Vertex struct {
	X                float64
	Y                float64
	someInternalData string
}

func Demo() {
	// Initialize and modify structs
	var _ Vertex = Vertex{}      // = Vertex{0.0, 0.0, ""}
	v1 := Vertex{X: 1.0, Y: 2.0} // = Vertex{1.0, 2.0, ""}
	v1.X = 3.0

	// Pointers to structs
	v2 := &v1
	v2.Y = 4.0
	fmt.Printf("%v", v1.Y) // 4.0
}

func PassByValue(v Vertex) {
	// v is a copy of the argument
}

func PassByReference(v *Vertex) {
	// v is a pointer to the original argument
	// changes to *v affect the original
}

// Interfaces
type Stringer interface {
	String() string
}

func (v Vertex) String() string {
	return fmt.Sprintf("Vertex: %v %v %v", v.X, v.Y, v.someInternalData)
}

func Demo2() {
	var s Stringer = Vertex{X: 1.0, Y: 2.0, someInternalData: "secret"}
	fmt.Println(s.String())
}

// Generics
type Node[T Stringer] struct {
	Value    T
	Children []*Node[T]
}

func Demo3() {
	n0 := Node[Vertex]{Value: Vertex{X: 5.0, Y: 6.0}}
	n1 := Node[Vertex]{Value: Vertex{X: 2.0, Y: 3.0}, Children: []*Node[Vertex]{&n0}}
	fmt.Printf("%v %v", n1.Value, n1.Children[0].Value) // 2 5
}
