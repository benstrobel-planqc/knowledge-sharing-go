package example

import "fmt"

// No native enums
type Status string

const (
	Scheduled Status = "SCHEDULED"
	Running   Status = "RUNNING"
	Completed Status = "COMPLETED"
)

func DemoEnums(s Status) {
	switch s {
	case Scheduled:
		// Handle scheduled status
	case Running:
		// Handle running status
	case Completed:
		// Handle completed status
	}
}

// Functional patterns not first class citizens
func ListComprehension() {
	// Python: squares = [x*x for x in range(10) if x%2 == 0]
	squares := []int{}
	for x := range 10 {
		if x%2 == 0 {
			squares = append(squares, x*x)
		}
	}
	fmt.Println(squares)
}

// Mixed concerns: Nullability and pointers
type User struct {
	Id            string
	FavoriteColor *string
}

func SomeMethod(u *User) {
	// Is u nullable or do we just want to pass it by reference to avoid copying?
	if u == nil {
		fmt.Println("User is nil")
		return
	}
}

// Often used: Custom type
type Optional[T any] struct {
	Present bool
	Value   T
}

func SomeMethod2(u Optional[User], u2 *Optional[User]) {
	// Clearly separates nullability from reference semantics
}
