// Package options provides cross type transformations for `mo.Option`.
//
// The functions provided by this package are not methods of `mo.Option` due to the lack of method type parameters
// on methods. This is part of the design decision of the Go's generics as explained here:
// https://go.googlesource.com/proposal/+/refs/heads/master/design/43651-type-parameters.md#No-parameterized-methods
//
// Providing these methods as a separate package also matches Go's primitives and standard library:
// - The `string` type don't have methods, but there we have the `strings` package.
// - The `[]byte` type don't have methods, but there we have the `bytes` package.
// - The `io.Reader` defines a single method, and all manipulations of a reader is done on packages `io` and `ioutil`.
package options

// import (
// 	"github.com/samber/mo"
// )

// // Map returns a new `mo.Option` wrapping the result of applying `f` to the value of opt, if present, and None otherwise.
// func Map[I any, O any](opt mo.Option[I], f func(I) O) mo.Option[O] {
// 	if val, ok := opt.Get(); ok {
// 		return mo.Some(f(val))
// 	}

// 	return mo.None[O]()
// }

// // FlatMap returns the result of applying `f` to the value of opt, if present, and None otherwise.
// func FlatMap[I any, O any](opt mo.Option[I], f func(I) mo.Option[O]) mo.Option[O] {
// 	if val, ok := opt.Get(); ok {
// 		return f(val)
// 	}

// 	return mo.None[O]()
// }

// // Match returns a new `mo.Option` from the result of applying `onValue` to the value of opt, if present,
// // or from the result of calling `onNone` if absent.
// func Match[I any, O any](opt mo.Option[I], onValue func(I) (O, bool), onNone func() (O, bool)) mo.Option[O] {
// 	if val, ok := opt.Get(); ok {
// 		return mo.TupleToOption(onValue(val))
// 	}

// 	return mo.TupleToOption(onNone())
// }

// // FlatMatch returns the result of applying `onValue` to the value of opt, if present,
// // or the result of `onNone` if absent.
// func FlatMatch[I any, O any](opt mo.Option[I], onValue func(I) O, onNone func() O) O {
// 	if val, ok := opt.Get(); ok {
// 		return onValue(val)
// 	}

// 	return onNone()
// }
