// Package either3 provides cross type transformations for `mo.Either`.
//
// The functions provided by this package are not methods of `mo.Either` due to the lack of method type parameters
// on methods. This is part of the design decision of the Go's generics as explained here:
// https://go.googlesource.com/proposal/+/refs/heads/master/design/43651-type-parameters.md#No-parameterized-methods
package either3

import (
	"github.com/samber/mo"
)

// Match returns the result of applying `onLeft` to the left value of the either or `onRight` to the right value of the either.
func Match[In1 any, In2 any, In3 any, Out1 any, Out2 any, Out3 any](onArg1 func(In1) Out1, onArg2 func(In2) Out2, onArg3 func(In3) Out3) func(either mo.Either3[In1, In2, In3]) mo.Either3[Out1, Out2, Out3] {
	return func(either mo.Either3[In1, In2, In3]) mo.Either3[Out1, Out2, Out3] {
		if either.IsArg1() {
			return mo.NewEither3Arg1[Out1, Out2, Out3](onArg1(either.MustArg1()))
		}

		if either.IsArg2() {
			return mo.NewEither3Arg2[Out1, Out2, Out3](onArg2(either.MustArg2()))
		}

		return mo.NewEither3Arg3[Out1, Out2, Out3](onArg3(either.MustArg3()))
	}
}

// MapArg1 executes the given function, if Either3 use the first argument, and returns result.
func MapArg1[In1 any, In2 any, In3 any, Out1 any](f func(In1) Out1) func(either mo.Either3[In1, In2, In3]) mo.Either3[Out1, In2, In3] {
	return func(either mo.Either3[In1, In2, In3]) mo.Either3[Out1, In2, In3] {
		return Match(
			func(arg1 In1) Out1 { return f(arg1) },
			func(arg2 In2) In2 { return arg2 },
			func(arg3 In3) In3 { return arg3 },
		)(either)
	}
}

// MapArg2 executes the given function, if Either3 use the second argument, and returns result.
func MapArg2[In1 any, In2 any, In3 any, Out2 any](f func(In2) Out2) func(either mo.Either3[In1, In2, In3]) mo.Either3[In1, Out2, In3] {
	return func(either mo.Either3[In1, In2, In3]) mo.Either3[In1, Out2, In3] {
		return Match(
			func(arg1 In1) In1 { return arg1 },
			func(arg2 In2) Out2 { return f(arg2) },
			func(arg3 In3) In3 { return arg3 },
		)(either)
	}
}

// MapArg3 executes the given function, if Either3 use the third argument, and returns result.
func MapArg3[In1 any, In2 any, In3 any, Out3 any](f func(In3) Out3) func(either mo.Either3[In1, In2, In3]) mo.Either3[In1, In2, Out3] {
	return func(either mo.Either3[In1, In2, In3]) mo.Either3[In1, In2, Out3] {
		return Match(
			func(arg1 In1) In1 { return arg1 },
			func(arg2 In2) In2 { return arg2 },
			func(arg3 In3) Out3 { return f(arg3) },
		)(either)
	}
}
