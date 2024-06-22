package mo

import (
	"errors"
	"fmt"
)

func ExampleFold() {
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
