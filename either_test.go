package mo

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEitherLeft(t *testing.T) {
	is := assert.New(t)

	left := Left[int, bool](42)
	is.Equal(Either[int, bool]{left: 42, right: false, isLeft: true}, left)
}

func TestEitherRight(t *testing.T) {
	is := assert.New(t)

	right := Right[int, bool](true)
	is.Equal(Either[int, bool]{left: 0, right: true, isLeft: false}, right)
}

func TestEitherUnpack(t *testing.T) {
	is := assert.New(t)

	left1, right1 := Left[int, bool](42).Unpack()
	left2, right2 := Right[int, bool](true).Unpack()

	is.Equal(42, left1)
	is.Equal(false, right1)
	is.Equal(0, left2)
	is.Equal(true, right2)
}

func TestEitherIsLeftOrRight(t *testing.T) {
	is := assert.New(t)

	left := Left[int, bool](42)
	right := Right[int, bool](true)

	is.True(left.IsLeft())
	is.False(left.IsRight())
	is.False(right.IsLeft())
	is.True(right.IsRight())
}

func TestEitherLeftOrRight(t *testing.T) {
	is := assert.New(t)

	left := Left[int, bool](42)
	right := Right[int, bool](true)

	result1, ok1 := left.Left()
	result2, ok2 := left.Right()
	result3, ok3 := right.Left()
	result4, ok4 := right.Right()

	is.Equal(42, result1)
	is.True(ok1)
	is.Equal(false, result2)
	is.False(ok2)
	is.Equal(0, result3)
	is.False(ok3)
	is.Equal(true, result4)
	is.True(ok4)
}

func TestEitherMustLeftOrRight(t *testing.T) {
	is := assert.New(t)

	left := Left[int, bool](42)
	right := Right[int, bool](true)

	is.NotPanics(func() {
		is.Equal(42, left.MustLeft())
	})
	is.Panics(func() {
		left.MustRight()
	})
	is.Panics(func() {
		right.MustLeft()
	})
	is.NotPanics(func() {
		is.Equal(true, right.MustRight())
	})
}

func TestEitherGetOrElse(t *testing.T) {
	is := assert.New(t)

	left := Left[int, string](42)
	right := Right[int, string]("foobar")

	is.Equal(42, left.LeftOrElse(21))
	is.Equal(21, right.LeftOrElse(21))
	is.Equal("baz", left.RightOrElse("baz"))
	is.Equal("foobar", right.RightOrElse("baz"))
}

func TestEitherGetOrEmpty(t *testing.T) {
	is := assert.New(t)

	left := Left[int, string](42)
	right := Right[int, string]("foobar")

	is.Equal(42, left.LeftOrEmpty())
	is.Equal(0, right.LeftOrEmpty())
	is.Equal("", left.RightOrEmpty())
	is.Equal("foobar", right.RightOrEmpty())
}

func TestEitherSwap(t *testing.T) {
	is := assert.New(t)

	left := Left[int, string](42)
	right := Right[int, string]("foobar")

	is.Equal(Either[string, int]{left: "", right: 42, isLeft: false}, left.Swap())
	is.Equal(Either[string, int]{left: "foobar", right: 0, isLeft: true}, right.Swap())
}

func TestEitherForEach(t *testing.T) {
	is := assert.New(t)

	Left[int, string](42).ForEach(
		func(a int) {
			is.Equal(42, a)
		},
		func(b string) {
			is.Fail("should not enter here")
		},
	)

	Right[int, string]("foobar").ForEach(
		func(a int) {
			is.Fail("should not enter here")
		},
		func(b string) {
			is.Equal("foobar", b)
		},
	)
}

func TestEitherMatch(t *testing.T) {
	is := assert.New(t)

	e1 := Left[int, string](42).Match(
		func(a int) Either[int, string] {
			is.Equal(42, a)
			return Left[int, string](21)
		},
		func(b string) Either[int, string] {
			is.Fail("should not enter here")
			return Left[int, string](1)
		},
	)

	e2 := Right[int, string]("foobar").Match(
		func(a int) Either[int, string] {
			is.Fail("should not enter here")
			return Right[int, string]("baz")
		},
		func(b string) Either[int, string] {
			is.Equal("foobar", b)
			return Right[int, string]("plop")
		},
	)

	is.Equal(Either[int, string]{left: 21, right: "", isLeft: true}, e1)
	is.Equal(Either[int, string]{left: 0, right: "plop", isLeft: false}, e2)
}

func TestEitherMapLeft(t *testing.T) {
	is := assert.New(t)

	e1 := Left[int, string](42).MapLeft(
		func(a int) Either[int, string] {
			is.Equal(42, a)
			return Left[int, string](21)
		},
	)

	e2 := Right[int, string]("foobar").MapLeft(
		func(a int) Either[int, string] {
			is.Fail("should not enter here")
			return Right[int, string]("plop")
		},
	)

	is.Equal(Either[int, string]{left: 21, right: "", isLeft: true}, e1)
	is.Equal(Either[int, string]{left: 0, right: "foobar", isLeft: false}, e2)
}

func TestEitherMapRight(t *testing.T) {
	is := assert.New(t)

	e1 := Left[int, string](42).MapRight(
		func(b string) Either[int, string] {
			is.Fail("should not enter here")
			return Left[int, string](21)
		},
	)

	e2 := Right[int, string]("foobar").MapRight(
		func(b string) Either[int, string] {
			is.Equal("foobar", b)
			return Right[int, string]("plop")
		},
	)

	is.Equal(Either[int, string]{left: 42, right: "", isLeft: true}, e1)
	is.Equal(Either[int, string]{left: 0, right: "plop", isLeft: false}, e2)
}

// TestEitherFoldSuccess tests the Fold method with a successful result.
func TestEitherFoldSuccess(t *testing.T) {
	is := assert.New(t)
	either := Either[error, int]{left: nil, right: 10, isLeft: false}

	successFunc := func(value int) string {
		return fmt.Sprintf("Success: %v", value)
	}
	failureFunc := func(err error) string {
		return fmt.Sprintf("Failure: %v", err)
	}

	folded := Fold[error, int, string](either, successFunc, failureFunc)
	expected := "Success: 10"

	is.Equal(expected, folded)
}

// TestEitherFoldFailure tests the Fold method with a failure result.
func TestEitherFoldFailure(t *testing.T) {
	err := errors.New("either error")
	is := assert.New(t)

	either := Either[error, int]{left: err, right: 0, isLeft: true}

	successFunc := func(value int) string {
		return fmt.Sprintf("Success: %v", value)
	}
	failureFunc := func(err error) string {
		return fmt.Sprintf("Failure: %v", err)
	}

	folded := Fold[error, int, string](either, successFunc, failureFunc)
	expected := fmt.Sprintf("Failure: %v", err)

	is.Equal(expected, folded)
}
