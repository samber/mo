package either

import (
	"testing"

	"github.com/samber/mo"
	"github.com/stretchr/testify/assert"
)

func TestPipe1(t *testing.T) {
	is := assert.New(t)

	src := mo.Right[int]("x")
	id := func(e mo.Either[int, string]) mo.Either[int, string] { return e }

	out := Pipe1(src, id)
	r, ok := out.Right()
	is.True(ok)
	is.Equal("x", r)
}

func TestPipe2(t *testing.T) {
	is := assert.New(t)

	src := mo.Left[int, string](42)
	id1 := func(e mo.Either[int, string]) mo.Either[int, string] { return e }
	id2 := func(e mo.Either[int, string]) mo.Either[int, string] { return e }

	out := Pipe2(src, id1, id2)
	l, ok := out.Left()
	is.True(ok)
	is.Equal(42, l)
}

func TestPipe3(t *testing.T) {
	is := assert.New(t)

	src := mo.Right[int, string]("a")

	op1 := func(e mo.Either[int, string]) mo.Either[int, string] {
		if e.IsRight() {
			return mo.Right[int, string](e.MustRight() + "1")
		}
		return mo.Left[int, string](e.MustLeft())
	}
	op2 := func(e mo.Either[int, string]) mo.Either[int, string] {
		if e.IsRight() {
			return mo.Right[int, string](e.MustRight() + "2")
		}
		return mo.Left[int, string](e.MustLeft())
	}
	op3 := func(e mo.Either[int, string]) mo.Either[int, string] {
		if e.IsRight() {
			return mo.Right[int, string](e.MustRight() + "3")
		}
		return mo.Left[int, string](e.MustLeft())
	}

	out := Pipe3[int, string, int, string, int, string, int, string](src, op1, op2, op3)
	r, ok := out.Right()
	is.True(ok)
	is.Equal("a123", r)
}

func TestPipe4(t *testing.T) {
	is := assert.New(t)
	src := mo.Right[int, string]("x")
	id := func(e mo.Either[int, string]) mo.Either[int, string] { return e }
	out := Pipe4(src, id, id, id, id)
	r, ok := out.Right()
	is.True(ok)
	is.Equal("x", r)
}

func TestPipe5(t *testing.T) {
	is := assert.New(t)
	src := mo.Right[int, string]("x")
	id := func(e mo.Either[int, string]) mo.Either[int, string] { return e }
	out := Pipe5(src, id, id, id, id, id)
	r, ok := out.Right()
	is.True(ok)
	is.Equal("x", r)
}

func TestPipe6(t *testing.T) {
	is := assert.New(t)
	src := mo.Right[int, string]("x")
	id := func(e mo.Either[int, string]) mo.Either[int, string] { return e }
	out := Pipe6(src, id, id, id, id, id, id)
	r, ok := out.Right()
	is.True(ok)
	is.Equal("x", r)
}

func TestPipe7(t *testing.T) {
	is := assert.New(t)
	src := mo.Right[int, string]("x")
	id := func(e mo.Either[int, string]) mo.Either[int, string] { return e }
	out := Pipe7(src, id, id, id, id, id, id, id)
	r, ok := out.Right()
	is.True(ok)
	is.Equal("x", r)
}

func TestPipe8(t *testing.T) {
	is := assert.New(t)
	src := mo.Right[int, string]("x")
	id := func(e mo.Either[int, string]) mo.Either[int, string] { return e }
	out := Pipe8(src, id, id, id, id, id, id, id, id)
	r, ok := out.Right()
	is.True(ok)
	is.Equal("x", r)
}

func TestPipe9(t *testing.T) {
	is := assert.New(t)
	src := mo.Right[int, string]("x")
	id := func(e mo.Either[int, string]) mo.Either[int, string] { return e }
	out := Pipe9(src, id, id, id, id, id, id, id, id, id)
	r, ok := out.Right()
	is.True(ok)
	is.Equal("x", r)
}

func TestPipe10(t *testing.T) {
	is := assert.New(t)
	src := mo.Right[int, string]("x")
	id := func(e mo.Either[int, string]) mo.Either[int, string] { return e }
	out := Pipe10(src, id, id, id, id, id, id, id, id, id, id)
	r, ok := out.Right()
	is.True(ok)
	is.Equal("x", r)
}
