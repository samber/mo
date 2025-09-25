package option

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/samber/mo"
	"github.com/stretchr/testify/assert"
)

func TestPipe3(t *testing.T) {
	is := assert.New(t)

	src := mo.Some("a")

	op1 := func(o mo.Option[string]) mo.Option[string] {
		if v, ok := o.Get(); ok {
			return mo.Some(v + "1")
		}
		return mo.None[string]()
	}
	op2 := func(o mo.Option[string]) mo.Option[string] {
		if v, ok := o.Get(); ok {
			return mo.Some(v + "2")
		}
		return mo.None[string]()
	}
	op3 := func(o mo.Option[string]) mo.Option[string] {
		if v, ok := o.Get(); ok {
			return mo.Some(v + "3")
		}
		return mo.None[string]()
	}

	out := Pipe3[string, string, string, string](src, op1, op2, op3)
	v, ok := out.Get()
	is.True(ok)
	is.Equal("a123", v)
}

func TestPipePreservesNone(t *testing.T) {
	is := assert.New(t)

	src := mo.None[string]()
	id := func(o mo.Option[string]) mo.Option[string] { return o }

	out := Pipe5[string, string, string, string, string, string](src, id, id, id, id, id)
	_, ok := out.Get()
	is.False(ok)
}

func TestPipe1(t *testing.T) {
	is := assert.New(t)
	id := func(o mo.Option[string]) mo.Option[string] { return o }
	some := mo.Some("v")
	r, ok := Pipe1(some, id).Get()
	is.True(ok)
	is.Equal("v", r)
}

func TestPipe2(t *testing.T) {
	is := assert.New(t)
	id := func(o mo.Option[string]) mo.Option[string] { return o }
	some := mo.Some("v")
	r, ok := Pipe2(some, id, id).Get()
	is.True(ok)
	is.Equal("v", r)
}

func TestPipe4(t *testing.T) {
	is := assert.New(t)
	id := func(o mo.Option[string]) mo.Option[string] { return o }
	some := mo.Some("v")
	r, ok := Pipe4(some, id, id, id, id).Get()
	is.True(ok)
	is.Equal("v", r)
}

func TestPipe5(t *testing.T) {
	is := assert.New(t)
	id := func(o mo.Option[string]) mo.Option[string] { return o }
	some := mo.Some("v")
	r, ok := Pipe5(some, id, id, id, id, id).Get()
	is.True(ok)
	is.Equal("v", r)
}

func TestPipe6(t *testing.T) {
	is := assert.New(t)
	id := func(o mo.Option[string]) mo.Option[string] { return o }
	some := mo.Some("v")
	r, ok := Pipe6(some, id, id, id, id, id, id).Get()
	is.True(ok)
	is.Equal("v", r)
}

func TestPipe7(t *testing.T) {
	is := assert.New(t)
	id := func(o mo.Option[string]) mo.Option[string] { return o }
	some := mo.Some("v")
	r, ok := Pipe7(some, id, id, id, id, id, id, id).Get()
	is.True(ok)
	is.Equal("v", r)
}

func TestPipe8(t *testing.T) {
	is := assert.New(t)
	id := func(o mo.Option[string]) mo.Option[string] { return o }
	some := mo.Some("v")
	r, ok := Pipe8(some, id, id, id, id, id, id, id, id).Get()
	is.True(ok)
	is.Equal("v", r)
}

func TestPipe9(t *testing.T) {
	is := assert.New(t)
	id := func(o mo.Option[string]) mo.Option[string] { return o }
	some := mo.Some("v")
	r, ok := Pipe9(some, id, id, id, id, id, id, id, id, id).Get()
	is.True(ok)
	is.Equal("v", r)
}

func TestPipe10(t *testing.T) {
	is := assert.New(t)

	id := func(o mo.Option[string]) mo.Option[string] { return o }
	some := mo.Some("v")
	r, ok := Pipe10(some, id, id, id, id, id, id, id, id, id, id).Get()
	is.True(ok)
	is.Equal("v", r)
}

func TestPipeTypeTransformations(t *testing.T) {
	is := assert.New(t)

	out := Pipe3(
		mo.Some("42"),
		FlatMap(func(str string) mo.Option[int] {
			v, err := strconv.Atoi(str)
			return mo.TupleToOption(v, err == nil)
		}),
		Map(func(n int) float64 {
			return float64(n)
		}),
		Map(func(n float64) string {
			return fmt.Sprintf("%.2f", n)
		}),
	)
	is.Equal(mo.Some("42.00"), out)
}
