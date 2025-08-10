package result

import (
	"errors"
	"testing"

	"github.com/samber/mo"
	"github.com/stretchr/testify/assert"
)

func TestPipe4Transforms(t *testing.T) {
	is := assert.New(t)

	src := mo.Ok("a")

	op := func(suffix string) func(mo.Result[string]) mo.Result[string] {
		return func(r mo.Result[string]) mo.Result[string] {
			if v, err := r.Get(); err == nil {
				return mo.Ok(v + suffix)
			}
			return r
		}
	}

	out := Pipe4[string, string, string, string, string](src, op("1"), op("2"), op("3"), op("4"))
	v, err := out.Get()
	is.NoError(err)
	is.Equal("a1234", v)
}

func TestPipePreservesError(t *testing.T) {
	is := assert.New(t)

	src := mo.Err[string](errors.New("boom"))
	id := func(r mo.Result[string]) mo.Result[string] { return r }

	out := Pipe5[string, string, string, string, string, string](src, id, id, id, id, id)
	_, err := out.Get()
	is.Error(err)
}

func TestPipe1To10(t *testing.T) {
	is := assert.New(t)

	id := func(r mo.Result[string]) mo.Result[string] { return r }

	// Ok scenario
	okR := mo.Ok("v")
	v1, err := Pipe1(okR, id).Get()
	is.NoError(err)
	is.Equal("v", v1)
	v2, err := Pipe2(okR, id, id).Get()
	is.NoError(err)
	is.Equal("v", v2)
	v4, err := Pipe4(okR, id, id, id, id).Get()
	is.NoError(err)
	is.Equal("v", v4)
	v5, err := Pipe5(okR, id, id, id, id, id).Get()
	is.NoError(err)
	is.Equal("v", v5)
	v6, err := Pipe6(okR, id, id, id, id, id, id).Get()
	is.NoError(err)
	is.Equal("v", v6)
	v7, err := Pipe7(okR, id, id, id, id, id, id, id).Get()
	is.NoError(err)
	is.Equal("v", v7)
	v8, err := Pipe8(okR, id, id, id, id, id, id, id, id).Get()
	is.NoError(err)
	is.Equal("v", v8)
	v9, err := Pipe9(okR, id, id, id, id, id, id, id, id, id).Get()
	is.NoError(err)
	is.Equal("v", v9)
	v10, err := Pipe10(okR, id, id, id, id, id, id, id, id, id, id).Get()
	is.NoError(err)
	is.Equal("v", v10)

	// Err scenario
	boom := mo.Err[string](errors.New("boom"))
	_, err = Pipe1(boom, id).Get()
	is.Error(err)
	_, err = Pipe2(boom, id, id).Get()
	is.Error(err)
	_, err = Pipe4(boom, id, id, id, id).Get()
	is.Error(err)
	_, err = Pipe5(boom, id, id, id, id, id).Get()
	is.Error(err)
	_, err = Pipe6(boom, id, id, id, id, id, id).Get()
	is.Error(err)
	_, err = Pipe7(boom, id, id, id, id, id, id, id).Get()
	is.Error(err)
	_, err = Pipe8(boom, id, id, id, id, id, id, id, id).Get()
	is.Error(err)
	_, err = Pipe9(boom, id, id, id, id, id, id, id, id, id).Get()
	is.Error(err)
	_, err = Pipe10(boom, id, id, id, id, id, id, id, id, id, id).Get()
	is.Error(err)
}

func TestPipe3Identity(t *testing.T) {
	is := assert.New(t)

	id := func(r mo.Result[string]) mo.Result[string] { return r }
	out := Pipe3[string, string, string, string](mo.Ok("v"), id, id, id)
	v, err := out.Get()
	is.NoError(err)
	is.Equal("v", v)
}
