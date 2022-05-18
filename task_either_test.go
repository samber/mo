package mo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTaskEither(t *testing.T) {
	is := assert.New(t)

	taskEither := NewTaskEither(func() *Future[int] {
		return NewFuture(func(resolve Resolver[int], reject Rejection) {
			resolve(42)
		})
	})

	result := taskEither.Run().Result().MustGet()

	is.Equal(42, result)
}

func TestTaskEitherOrElse(t *testing.T) {
	is := assert.New(t)

	taskEither1 := NewTaskEither(func() *Future[int] {
		return NewFuture(func(resolve Resolver[int], reject Rejection) {
			resolve(42)
		})
	})
	taskEither2 := NewTaskEither(func() *Future[int] {
		return NewFuture(func(resolve Resolver[int], reject Rejection) {
			reject(assert.AnError)
		})
	})

	result1 := taskEither1.OrElse(1234)
	result2 := taskEither2.OrElse(1234)

	is.Equal(42, result1)
	is.Equal(1234, result2)
}

func TestTaskEitherMatch(t *testing.T) {
	is := assert.New(t)

	taskEither := NewTaskEither(func() *Future[int] {
		return NewFuture(func(resolve Resolver[int], reject Rejection) {
			resolve(42)
		})
	})

	mapped := taskEither.Match(
		func(err error) Either[error, int] {
			return Right[error, int](1234)
		},
		func(i int) Either[error, int] {
			return Right[error, int](i)
		},
	)

	v, ok := mapped.Right()

	is.Equal(42, v)
	is.True(ok)
}

func TestTaskEitherTryCatch(t *testing.T) {
	is := assert.New(t)

	taskEither := NewTaskEither(func() *Future[int] {
		return NewFuture(func(resolve Resolver[int], reject Rejection) {
			resolve(42)
		})
	})

	mapped := taskEither.TryCatch(
		func(err error) Either[error, int] {
			return Right[error, int](1234)
		},
		func(i int) Either[error, int] {
			return Right[error, int](i)
		},
	)

	v, ok := mapped.Right()

	is.Equal(42, v)
	is.True(ok)
}

func TestTaskEitherToTask(t *testing.T) {
	is := assert.New(t)

	taskEither := NewTaskEither(func() *Future[int] {
		return NewFuture(func(resolve Resolver[int], reject Rejection) {
			reject(assert.AnError)
		})
	})

	task := taskEither.ToTask(1234)

	result := task.Run().Result().MustGet()

	is.Equal(1234, result)
}

func TestTaskEitherToEither(t *testing.T) {
	is := assert.New(t)

	taskEither := NewTaskEither(func() *Future[int] {
		return NewFuture(func(resolve Resolver[int], reject Rejection) {
			reject(assert.AnError)
		})
	})

	either := taskEither.ToEither()
	err, isError := either.Left()

	is.True(isError)
	is.NotNil(err)
	is.Equal(assert.AnError, err)
}
