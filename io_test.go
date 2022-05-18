package mo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIO(t *testing.T) {
	is := assert.New(t)

	io := NewIO(func() int {
		return 42
	})
	result := io.Run()

	is.Equal(42, result)
}

func TestIO1(t *testing.T) {
	is := assert.New(t)

	io := NewIO1(func(a string) int {
		is.Equal("foo", a)
		return 42
	})
	result := io.Run("foo")

	is.Equal(42, result)
}

func TestIO2(t *testing.T) {
	is := assert.New(t)

	io := NewIO2(func(a string, b string) int {
		is.Equal("foo", a)
		is.Equal("bar", b)
		return 42
	})
	result := io.Run("foo", "bar")

	is.Equal(42, result)
}

func TestIO3(t *testing.T) {
	is := assert.New(t)

	io := NewIO3(func(a string, b string, c string) int {
		is.Equal("foo", a)
		is.Equal("bar", b)
		is.Equal("hello", c)
		return 42
	})
	result := io.Run("foo", "bar", "hello")

	is.Equal(42, result)
}

func TestIO4(t *testing.T) {
	is := assert.New(t)

	io := NewIO4(func(a string, b string, c string, d string) int {
		is.Equal("foo", a)
		is.Equal("bar", b)
		is.Equal("hello", c)
		is.Equal("world", d)
		return 42
	})
	result := io.Run("foo", "bar", "hello", "world")

	is.Equal(42, result)
}

func TestIO5(t *testing.T) {
	is := assert.New(t)

	io := NewIO5(func(a string, b string, c string, d string, e bool) int {
		is.Equal("foo", a)
		is.Equal("bar", b)
		is.Equal("hello", c)
		is.Equal("world", d)
		is.True(e)
		return 42
	})
	result := io.Run("foo", "bar", "hello", "world", true)

	is.Equal(42, result)
}
