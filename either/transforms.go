// Package either provides cross type transformations for `mo.Either`.
//
// The functions provided by this package are not methods of `mo.Either` due to the lack of method type parameters
// on methods. This is part of the design decision of the Go's generics as explained here:
// https://go.googlesource.com/proposal/+/refs/heads/master/design/43651-type-parameters.md#No-parameterized-methods
package either

import (
	"github.com/samber/mo"
)

// MapLeft returns a new `mo.Either` wrapping the result of applying `f` to the left value of the either.
func MapLeft[I any, O any](f func(I) O) func(either mo.Either[I, O]) mo.Either[O, O] {
	return func(either mo.Either[I, O]) mo.Either[O, O] {
		if either.IsLeft() {
			return mo.Right[O](f(either.MustLeft()))
		}

		return mo.Left[O, O](either.MustRight())
	}
}

// MapRight returns a new `mo.Either` wrapping the result of applying `f` to the right value of the either.
func MapRight[I any, O any](f func(O) O) func(either mo.Either[I, O]) mo.Either[I, O] {
	return func(either mo.Either[I, O]) mo.Either[I, O] {
		if either.IsRight() {
			return mo.Right[I](f(either.MustRight()))
		}

		return mo.Left[I, O](either.MustLeft())
	}
}

// Match returns the result of applying `onLeft` to the left value of the either or `onRight` to the right value of the either.
func Match[ILeft any, IRight any, OLeft any, ORight any](onLeft func(ILeft) OLeft, onRight func(IRight) ORight) func(either mo.Either[ILeft, IRight]) mo.Either[OLeft, ORight] {
	return func(either mo.Either[ILeft, IRight]) mo.Either[OLeft, ORight] {
		if either.IsLeft() {
			return mo.Left[OLeft, ORight](onLeft(either.MustLeft()))
		}

		return mo.Right[OLeft](onRight(either.MustRight()))
	}
}

// Swap returns the left value in Right and vice versa.
func Swap[I any, O any]() func(either mo.Either[I, O]) mo.Either[O, I] {
	return func(either mo.Either[I, O]) mo.Either[O, I] {
		if either.IsLeft() {
			return mo.Right[O](either.MustLeft())
		}

		return mo.Left[O, I](either.MustRight())
	}
}
