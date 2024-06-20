package mo

import (
	"errors"
	"fmt"
)

// Do executes a function within a monadic context, capturing any errors that occur.
// If the function executes successfully, its result is wrapped in a successful Result.
// If the function panics (indicating a failure), the panic is caught and converted into an error Result.
func Do[T any](fn func() T) (result Result[T]) {
	defer func() {
		if r := recover(); r != nil {
			if err, ok := r.(error); ok {
				result = Err[T](err)
			} else {
				result = Err[T](errors.New(fmt.Sprint(r)))
			}
		}
	}()
	return Ok(fn())
}
