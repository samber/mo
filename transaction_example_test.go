package mo

import (
	"fmt"
)

func ExampleTransaction() {
	transaction := NewTransaction[int]().
		Then(
			func(state int) (int, error) {
				fmt.Println("step 1")
				return state + 10, nil
			},
			func(state int) int {
				fmt.Println("rollback 1")
				return state - 10
			},
		).
		Then(
			func(state int) (int, error) {
				fmt.Println("step 2")
				return state + 15, nil
			},
			func(state int) int {
				fmt.Println("rollback 2")
				return state - 15
			},
		).
		Then(
			func(state int) (int, error) {
				fmt.Println("step 3")

				if true {
					return state, fmt.Errorf("error")
				}

				return state + 42, nil
			},
			func(state int) int {
				fmt.Println("rollback 3")
				return state - 42
			},
		)

	_, _ = transaction.Process(-5)

	// Output:
	// step 1
	// step 2
	// step 3
	// rollback 2
	// rollback 1
}

func ExampleTransaction_ok() {
	transaction := NewTransaction[int]().
		Then(
			func(state int) (int, error) {
				return state + 10, nil
			},
			func(state int) int {
				return state - 10
			},
		).
		Then(
			func(state int) (int, error) {
				return state + 15, nil
			},
			func(state int) int {
				return state - 15
			},
		).
		Then(
			func(state int) (int, error) {
				return state + 42, nil
			},
			func(state int) int {
				return state - 42
			},
		)

	state, err := transaction.Process(-5)

	fmt.Println(state)
	fmt.Println(err)
	// Output:
	// 62
	// <nil>
}

func ExampleTransaction_error() {
	transaction := NewTransaction[int]().
		Then(
			func(state int) (int, error) {
				return state + 10, nil
			},
			func(state int) int {
				return state - 10
			},
		).
		Then(
			func(state int) (int, error) {
				return state, fmt.Errorf("error")
			},
			func(state int) int {
				return state - 15
			},
		).
		Then(
			func(state int) (int, error) {
				return state + 42, nil
			},
			func(state int) int {
				return state - 42
			},
		)

	state, err := transaction.Process(-5)

	fmt.Println(state)
	fmt.Println(err)
	// Output:
	// -5
	// error
}
