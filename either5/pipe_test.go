package either5

import (
	"testing"

	"github.com/samber/mo"
	"github.com/stretchr/testify/assert"
)

func TestPipe3(t *testing.T) {
	is := assert.New(t)

	src := mo.NewEither5Arg5[int, string, bool, float64, rune]('a')
	id := func(e mo.Either5[int, string, bool, float64, rune]) mo.Either5[int, string, bool, float64, rune] {
		return e
	}

	out := Pipe3[int, string, bool, float64, rune, int, string, bool, float64, rune, int, string, bool, float64, rune](src, id, id, id)
	v, ok := out.Arg5()
	is.True(ok)
	is.Equal(rune('a'), v)
}

func TestPipe1(t *testing.T) {
	is := assert.New(t)

	src := mo.NewEither5Arg5[int, string, bool, float64, rune]('a')
	id := func(e mo.Either5[int, string, bool, float64, rune]) mo.Either5[int, string, bool, float64, rune] {
		return e
	}

	out := Pipe1(src, id)
	v, ok := out.Arg5()
	is.True(ok)
	is.Equal(rune('a'), v)
}

func TestPipe2(t *testing.T) {
	is := assert.New(t)

	src := mo.NewEither5Arg5[int, string, bool, float64, rune]('a')
	id := func(e mo.Either5[int, string, bool, float64, rune]) mo.Either5[int, string, bool, float64, rune] {
		return e
	}

	out := Pipe2(src, id, id)
	v, ok := out.Arg5()
	is.True(ok)
	is.Equal(rune('a'), v)
}

func TestPipe4(t *testing.T) {
	is := assert.New(t)

	src := mo.NewEither5Arg5[int, string, bool, float64, rune]('a')
	id := func(e mo.Either5[int, string, bool, float64, rune]) mo.Either5[int, string, bool, float64, rune] {
		return e
	}

	out4 := Pipe4(src, id, id, id, id)
	v4, ok := out4.Arg5()
	is.True(ok)
	is.Equal(rune('a'), v4)

	out5 := Pipe5(src, id, id, id, id, id)
	v5, ok := out5.Arg5()
	is.True(ok)
	is.Equal(rune('a'), v5)

	out6 := Pipe6(src, id, id, id, id, id, id)
	v6, ok := out6.Arg5()
	is.True(ok)
	is.Equal(rune('a'), v6)

}

func TestPipe5(t *testing.T) {
	is := assert.New(t)
	src := mo.NewEither5Arg5[int, string, bool, float64, rune]('a')
	id := func(e mo.Either5[int, string, bool, float64, rune]) mo.Either5[int, string, bool, float64, rune] {
		return e
	}
	out := Pipe5(src, id, id, id, id, id)
	v, ok := out.Arg5()
	is.True(ok)
	is.Equal(rune('a'), v)
}

func TestPipe6(t *testing.T) {
	is := assert.New(t)
	src := mo.NewEither5Arg5[int, string, bool, float64, rune]('a')
	id := func(e mo.Either5[int, string, bool, float64, rune]) mo.Either5[int, string, bool, float64, rune] {
		return e
	}
	out := Pipe6(src, id, id, id, id, id, id)
	v, ok := out.Arg5()
	is.True(ok)
	is.Equal(rune('a'), v)
}

func TestPipe7(t *testing.T) {
	is := assert.New(t)
	src := mo.NewEither5Arg5[int, string, bool, float64, rune]('a')
	id := func(e mo.Either5[int, string, bool, float64, rune]) mo.Either5[int, string, bool, float64, rune] {
		return e
	}
	out := Pipe7(src, id, id, id, id, id, id, id)
	v, ok := out.Arg5()
	is.True(ok)
	is.Equal(rune('a'), v)
}
