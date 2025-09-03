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
func Match[I1 any, I2 any, I3 any, O1 any, O2 any, O3 any](onArg1 func(I1) O1, onArg2 func(I2) O2, onArg3 func(I3) O3) func(either mo.Either3[I1, I2, I3]) mo.Either3[O1, O2, O3] {
	return func(either mo.Either3[I1, I2, I3]) mo.Either3[O1, O2, O3] {
		if either.IsArg1() {
			return mo.NewEither3Arg1[O1, O2, O3](onArg1(either.MustArg1()))
		}

		if either.IsArg2() {
			return mo.NewEither3Arg2[O1, O2, O3](onArg2(either.MustArg2()))
		}

		return mo.NewEither3Arg3[O1, O2, O3](onArg3(either.MustArg3()))
	}
}

// MapArg1 executes the given function, if Either3 use the first argument, and returns result.
func MapArg1[I1 any, I2 any, I3 any, O1 any](f func(I1) O1) func(either mo.Either3[I1, I2, I3]) mo.Either3[O1, I2, I3] {
	return func(either mo.Either3[I1, I2, I3]) mo.Either3[O1, I2, I3] {
		return Match(
			func(arg1 I1) O1 { return f(arg1) },
			func(arg2 I2) I2 { return arg2 },
			func(arg3 I3) I3 { return arg3 },
		)(either)
	}
}

// MapArg2 executes the given function, if Either3 use the second argument, and returns result.
func MapArg2[I1 any, I2 any, I3 any, O2 any](f func(I2) O2) func(either mo.Either3[I1, I2, I3]) mo.Either3[I1, O2, I3] {
	return func(either mo.Either3[I1, I2, I3]) mo.Either3[I1, O2, I3] {
		return Match(
			func(arg1 I1) I1 { return arg1 },
			func(arg2 I2) O2 { return f(arg2) },
			func(arg3 I3) I3 { return arg3 },
		)(either)
	}
}

// MapArg3 executes the given function, if Either3 use the third argument, and returns result.
func MapArg3[I1 any, I2 any, I3 any, O3 any](f func(I3) O3) func(either mo.Either3[I1, I2, I3]) mo.Either3[I1, I2, O3] {
	return func(either mo.Either3[I1, I2, I3]) mo.Either3[I1, I2, O3] {
		return Match(
			func(arg1 I1) I1 { return arg1 },
			func(arg2 I2) I2 { return arg2 },
			func(arg3 I3) O3 { return f(arg3) },
		)(either)
	}
}
