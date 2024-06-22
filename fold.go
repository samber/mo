package mo

// Foldable represents a type that can be folded into a single value
// based on its state.
//
// - T: the type of the value in the failure state (e.g., an error type).
// - U: the type of the value in the success state.
type Foldable[T any, U any] interface {
	leftValue() T
	rightValue() U
	hasLeftValue() bool
}

// Fold applies one of the two functions based on the state of the Foldable type,
// and it returns the result of applying either successFunc or failureFunc.
//
// - T: the type of the failure value (e.g., an error type)
// - U: the type of the success value
// - R: the type of the return value from the folding functions
//
// successFunc is applied when the Foldable is in the success state (i.e., isLeft() is false).
// failureFunc is applied when the Foldable is in the failure state (i.e., isLeft() is true).
func Fold[T, U, R any](f Foldable[T, U], successFunc func(U) R, failureFunc func(T) R) R {
	if f.hasLeftValue() {
		return failureFunc(f.leftValue())
	}
	return successFunc(f.rightValue())
}
