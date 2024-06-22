package mo

import (
	"fmt"
)

func ExampleDo() {
	a := Ok("Hello, World!")
	b := Some("42")

	result := Do(func() []string {
		return []string{
			a.MustGet(),
			b.MustGet(),
		}
	})

	fmt.Println(result.IsError())
	fmt.Println(result.MustGet())
	// Output:
	// false
	// [Hello, World! 42]
}
