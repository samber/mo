package either3

import (
	"testing"

	"github.com/samber/mo"
	"github.com/stretchr/testify/assert"
)

func TestPipe2(t *testing.T) {
	is := assert.New(t)

	src := mo.NewEither3Arg2[int, string, bool]("hi")
	id := func(e mo.Either3[int, string, bool]) mo.Either3[int, string, bool] { return e }

	out := Pipe2(src, id, id)
	s, ok := out.Arg2()
	is.True(ok)
	is.Equal("hi", s)
}

func TestPipe1(t *testing.T) {
	is := assert.New(t)
	src := mo.NewEither3Arg2[int, string, bool]("hi")
	id := func(e mo.Either3[int, string, bool]) mo.Either3[int, string, bool] { return e }
	out := Pipe1(src, id)
	s, ok := out.Arg2()
	is.True(ok)
	is.Equal("hi", s)
}

func TestPipe3(t *testing.T) {
	is := assert.New(t)
	src := mo.NewEither3Arg2[int, string, bool]("hi")
	id := func(e mo.Either3[int, string, bool]) mo.Either3[int, string, bool] { return e }
	out := Pipe3(src, id, id, id)
	s, ok := out.Arg2()
	is.True(ok)
	is.Equal("hi", s)
}

func TestPipe4(t *testing.T) {
	is := assert.New(t)

	src := mo.NewEither3Arg2[int, string, bool]("hi")
	id := func(e mo.Either3[int, string, bool]) mo.Either3[int, string, bool] { return e }

	out4 := Pipe4(src, id, id, id, id)
	s4, ok := out4.Arg2()
	is.True(ok)
	is.Equal("hi", s4)

	out5 := Pipe5(src, id, id, id, id, id)
	s5, ok := out5.Arg2()
	is.True(ok)
	is.Equal("hi", s5)

	out6 := Pipe6(src, id, id, id, id, id, id)
	s6, ok := out6.Arg2()
	is.True(ok)
	is.Equal("hi", s6)

	out7 := Pipe7(src, id, id, id, id, id, id, id)
	s7, ok := out7.Arg2()
	is.True(ok)
	is.Equal("hi", s7)

	out8 := Pipe8(src, id, id, id, id, id, id, id, id)
	s8, ok := out8.Arg2()
	is.True(ok)
	is.Equal("hi", s8)

	out9 := Pipe9(src, id, id, id, id, id, id, id, id, id)
	s9, ok := out9.Arg2()
	is.True(ok)
	is.Equal("hi", s9)

	out10 := Pipe10(src, id, id, id, id, id, id, id, id, id, id)
	s10, ok := out10.Arg2()
	is.True(ok)
	is.Equal("hi", s10)
}
