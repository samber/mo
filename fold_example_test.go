package mo

import (
	"errors"
	"fmt"
)

func ExampleFold_result() {
	res1 := Result[int]{value: 42, isErr: false, err: nil}
	res2 := Result[int]{value: 0, isErr: true, err: errors.New("error")}

	successFunc := func(val int) string {
		return fmt.Sprintf("Success with value %d", val)
	}

	failureFunc := func(err error) string {
		return fmt.Sprintf("Failure with error %s", err)
	}

	fmt.Println(Fold[error, int, string](res1, successFunc, failureFunc))
	fmt.Println(Fold[error, int, string](res2, successFunc, failureFunc))
	// Output:
	// Success with value 42
	// Failure with error error
}

func ExampleFold_either() {
	either1 := Either[error, int]{isLeft: false, right: 42}
	either2 := Either[error, int]{isLeft: true, left: errors.New("either error")}

	successFunc := func(val int) string {
		return fmt.Sprintf("Success with value %d", val)
	}

	failureFunc := func(err error) string {
		return fmt.Sprintf("Failure with error %s", err)
	}

	fmt.Println(Fold[error, int, string](either1, successFunc, failureFunc))
	fmt.Println(Fold[error, int, string](either2, successFunc, failureFunc))
	// Output:
	// Success with value 42
	// Failure with error either error
}

func ExampleFold_option() {
	option1 := Option[int]{isPresent: true, value: 42}
	option2 := Option[int]{isPresent: false}

	successFunc := func(val int) string {
		return fmt.Sprintf("Success with value %d", val)
	}

	failureFunc := func(err error) string {
		return fmt.Sprintf("Failure with error %s", err)
	}

	fmt.Println(Fold[error, int, string](option1, successFunc, failureFunc))
	fmt.Println(Fold[error, int, string](option2, successFunc, failureFunc))
	// Output:
	// Success with value 42
	// Failure with error no such element
}
