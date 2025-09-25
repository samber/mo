package either4

import (
	"testing"

	"github.com/samber/mo"
	"github.com/stretchr/testify/assert"
)

func TestPipe2(t *testing.T) {
	is := assert.New(t)

	src := mo.NewEither4Arg4[int, string, bool, float64](3.14)
	id := func(e mo.Either4[int, string, bool, float64]) mo.Either4[int, string, bool, float64] { return e }

	out := Pipe2[int, string, bool, float64, int, string, bool, float64, int, string, bool, float64](src, id, id)
	v, ok := out.Arg4()
	is.True(ok)
	is.Equal(3.14, v)
}

func TestPipe1(t *testing.T) {
	is := assert.New(t)
	src := mo.NewEither4Arg4[int, string, bool, float64](3.14)
	id := func(e mo.Either4[int, string, bool, float64]) mo.Either4[int, string, bool, float64] { return e }
	out := Pipe1(src, id)
	v, ok := out.Arg4()
	is.True(ok)
	is.Equal(3.14, v)
}

func TestPipe3(t *testing.T) {
	is := assert.New(t)
	src := mo.NewEither4Arg4[int, string, bool, float64](3.14)
	id := func(e mo.Either4[int, string, bool, float64]) mo.Either4[int, string, bool, float64] { return e }
	out := Pipe3(src, id, id, id)
	v, ok := out.Arg4()
	is.True(ok)
	is.Equal(3.14, v)
}

func TestPipe4(t *testing.T) {
	is := assert.New(t)

	src := mo.NewEither4Arg4[int, string, bool, float64](3.14)
	id := func(e mo.Either4[int, string, bool, float64]) mo.Either4[int, string, bool, float64] { return e }

	out4 := Pipe4(src, id, id, id, id)
	v4, ok := out4.Arg4()
	is.True(ok)
	is.Equal(3.14, v4)

	out5 := Pipe5(src, id, id, id, id, id)
	v5, ok := out5.Arg4()
	is.True(ok)
	is.Equal(3.14, v5)

	out6 := Pipe6(src, id, id, id, id, id, id)
	v6, ok := out6.Arg4()
	is.True(ok)
	is.Equal(3.14, v6)

	out7 := Pipe7(src, id, id, id, id, id, id, id)
	v7, ok := out7.Arg4()
	is.True(ok)
	is.Equal(3.14, v7)

	out8 := Pipe8(src, id, id, id, id, id, id, id, id)
	v8, ok := out8.Arg4()
	is.True(ok)
	is.Equal(3.14, v8)

	out9 := Pipe9(src, id, id, id, id, id, id, id, id, id)
	v9, ok := out9.Arg4()
	is.True(ok)
	is.Equal(3.14, v9)

	out10 := Pipe10(src, id, id, id, id, id, id, id, id, id, id)
	v10, ok := out10.Arg4()
	is.True(ok)
	is.Equal(3.14, v10)
}
