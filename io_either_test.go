package mo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIOEither(t *testing.T) {
	is := assert.New(t)

	ioEither := NewIOEither(func() (int, error) {
		return 42, nil
	})
	result := ioEither.Run()

	is.False(result.isLeft)
	is.Nil(result.Left())
	is.True(result.IsRight())
	is.Equal(42, result.MustRight())
}

func TestIOEither1(t *testing.T) {
	is := assert.New(t)

	ioEither := NewIOEither1(func(a string) (int, error) {
		is.Equal("foo", a)
		return 42, nil
	})
	result := ioEither.Run("foo")

	is.False(result.isLeft)
	is.Nil(result.Left())
	is.True(result.IsRight())
	is.Equal(42, result.MustRight())
}

func TestIOEither2(t *testing.T) {
	is := assert.New(t)

	ioEither := NewIOEither2(func(a string, b string) (int, error) {
		is.Equal("foo", a)
		is.Equal("bar", b)
		return 42, nil
	})
	result := ioEither.Run("foo", "bar")

	is.False(result.isLeft)
	is.Nil(result.Left())
	is.True(result.IsRight())
	is.Equal(42, result.MustRight())
}

func TestIOEither3(t *testing.T) {
	is := assert.New(t)

	ioEither := NewIOEither3(func(a string, b string, c string) (int, error) {
		is.Equal("foo", a)
		is.Equal("bar", b)
		is.Equal("hello", c)
		return 42, nil
	})
	result := ioEither.Run("foo", "bar", "hello")

	is.False(result.isLeft)
	is.Nil(result.Left())
	is.True(result.IsRight())
	is.Equal(42, result.MustRight())
}

func TestIOEither4(t *testing.T) {
	is := assert.New(t)

	ioEither := NewIOEither4(func(a string, b string, c string, d string) (int, error) {
		is.Equal("foo", a)
		is.Equal("bar", b)
		is.Equal("hello", c)
		is.Equal("world", d)
		return 42, nil
	})
	result := ioEither.Run("foo", "bar", "hello", "world")

	is.False(result.isLeft)
	is.Nil(result.Left())
	is.True(result.IsRight())
	is.Equal(42, result.MustRight())
}

func TestIOEither5(t *testing.T) {
	is := assert.New(t)

	ioEither := NewIOEither5(func(a string, b string, c string, d string, e bool) (int, error) {
		is.Equal("foo", a)
		is.Equal("bar", b)
		is.Equal("hello", c)
		is.Equal("world", d)
		is.True(e)
		return 42, nil
	})
	result := ioEither.Run("foo", "bar", "hello", "world", true)

	is.False(result.isLeft)
	is.Nil(result.Left())
	is.True(result.IsRight())
	is.Equal(42, result.MustRight())
}
