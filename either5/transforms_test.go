package either5

import (
	"testing"

	"github.com/samber/mo"
	"github.com/stretchr/testify/assert"
)

func TestMatch_AllBranches(t *testing.T) {
	is := assert.New(t)

	f := Match[int, string, bool, float64, rune, string, int, int, string, int](
		func(i int) string { return "i" },
		func(s string) int { return len(s) },
		func(b bool) int {
			if b {
				return 1
			}
			return 0
		},
		func(f float64) string { return "f" },
		func(r rune) int { return int(r) },
	)

	o1 := f(mo.NewEither5Arg1[int, string, bool, float64, rune](1))
	a1, ok1 := o1.Arg1()
	is.True(ok1)
	is.Equal("i", a1)

	o2 := f(mo.NewEither5Arg2[int, string, bool, float64, rune]("ab"))
	a2, ok2 := o2.Arg2()
	is.True(ok2)
	is.Equal(2, a2)

	o3 := f(mo.NewEither5Arg3[int, string, bool, float64, rune](true))
	a3, ok3 := o3.Arg3()
	is.True(ok3)
	is.Equal(1, a3)

	o4 := f(mo.NewEither5Arg4[int, string, bool, float64, rune](3.14))
	a4, ok4 := o4.Arg4()
	is.True(ok4)
	is.Equal("f", a4)

	o5 := f(mo.NewEither5Arg5[int, string, bool, float64, rune]('a'))
	a5, ok5 := o5.Arg5()
	is.True(ok5)
	is.Equal(int('a'), a5)
}

func TestMapArgN(t *testing.T) {
	is := assert.New(t)

	out1 := MapArg1[int, string, bool, float64, rune, string](func(i int) string { return "x" })(mo.NewEither5Arg1[int, string, bool, float64, rune](1))
	v1, ok1 := out1.Arg1()
	is.True(ok1)
	is.Equal("x", v1)

	out2 := MapArg2[int, string, bool, float64, rune, int](func(s string) int { return len(s) })(mo.NewEither5Arg2[int, string, bool, float64, rune]("ab"))
	v2, ok2 := out2.Arg2()
	is.True(ok2)
	is.Equal(2, v2)

	out3 := MapArg3[int, string, bool, float64, rune, int](func(b bool) int {
		if b {
			return 7
		}
		return 8
	})(mo.NewEither5Arg3[int, string, bool, float64, rune](true))
	v3, ok3 := out3.Arg3()
	is.True(ok3)
	is.Equal(7, v3)
}
