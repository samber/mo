package mo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTask(t *testing.T) {
	is := assert.New(t)

	task := NewTask(func() *Future[int] {
		return NewFuture(func(resolve func(int), reject func(error)) {
			resolve(42)
		})
	})
	result := task.Run().Result().MustGet()

	is.Equal(42, result)
}

func TestTask1(t *testing.T) {
	is := assert.New(t)

	task := NewTask1(func(a string) *Future[int] {
		return NewFuture(func(resolve func(int), reject func(error)) {
			is.Equal("foo", a)
			resolve(42)
		})
	})
	result := task.Run("foo").Result().MustGet()

	is.Equal(42, result)
}

func TestTask2(t *testing.T) {
	is := assert.New(t)

	task := NewTask2(func(a string, b string) *Future[int] {
		return NewFuture(func(resolve func(int), reject func(error)) {
			is.Equal("foo", a)
			is.Equal("bar", b)
			resolve(42)
		})
	})
	result := task.Run("foo", "bar").Result().MustGet()

	is.Equal(42, result)
}

func TestTask3(t *testing.T) {
	is := assert.New(t)

	task := NewTask3(func(a string, b string, c string) *Future[int] {
		return NewFuture(func(resolve func(int), reject func(error)) {
			is.Equal("foo", a)
			is.Equal("bar", b)
			is.Equal("hello", c)
			resolve(42)
		})
	})
	result := task.Run("foo", "bar", "hello").Result().MustGet()

	is.Equal(42, result)
}

func TestTask4(t *testing.T) {
	is := assert.New(t)

	task := NewTask4(func(a string, b string, c string, d string) *Future[int] {
		return NewFuture(func(resolve func(int), reject func(error)) {
			is.Equal("foo", a)
			is.Equal("bar", b)
			is.Equal("hello", c)
			is.Equal("world", d)
			resolve(42)
		})
	})
	result := task.Run("foo", "bar", "hello", "world").Result().MustGet()

	is.Equal(42, result)
}

func TestTask5(t *testing.T) {
	is := assert.New(t)

	task := NewTask5(func(a string, b string, c string, d string, e bool) *Future[int] {
		return NewFuture(func(resolve func(int), reject func(error)) {
			is.Equal("foo", a)
			is.Equal("bar", b)
			is.Equal("hello", c)
			is.Equal("world", d)
			is.True(e)
			resolve(42)
		})
	})
	result := task.Run("foo", "bar", "hello", "world", true).Result().MustGet()

	is.Equal(42, result)
}
