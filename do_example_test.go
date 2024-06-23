package mo

import (
	"errors"
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

func ExampleDo_panic() {
	a := Ok("Hello, World!")
	b := Some("42")
	c := Err[string](errors.New("result error"))

	result := Do(func() []string {
		return []string{
			a.MustGet(),
			b.MustGet(),
			c.MustGet(), // would panic without Do-notation
		}
	})

	fmt.Println(result.IsError())
	fmt.Println(result.Error().Error())
	// Output:
	// true
	// result error
}
