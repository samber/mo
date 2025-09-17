package either3

import (
	"testing"

	"github.com/samber/mo"
	"github.com/stretchr/testify/assert"
)

func TestMatch_AllBranches(t *testing.T) {
	is := assert.New(t)

	f := Match[int, string, bool, string, int, int](
		func(i int) string { return "i" },
		func(s string) int { return len(s) },
		func(b bool) int {
			if b {
				return 1
			}
			return 0
		},
	)

	o1 := f(mo.NewEither3Arg1[int, string, bool](3))
	a1, ok1 := o1.Arg1()
	is.True(ok1)
	is.Equal("i", a1)

	o2 := f(mo.NewEither3Arg2[int, string, bool]("abcd"))
	a2, ok2 := o2.Arg2()
	is.True(ok2)
	is.Equal(4, a2)

	o3 := f(mo.NewEither3Arg3[int, string, bool](true))
	a3, ok3 := o3.Arg3()
	is.True(ok3)
	is.Equal(1, a3)
}

func TestMapArgN(t *testing.T) {
	is := assert.New(t)

	out1 := MapArg1[int, string, bool, string](func(i int) string { return "x" })(mo.NewEither3Arg1[int, string, bool](1))
	v1, ok1 := out1.Arg1()
	is.True(ok1)
	is.Equal("x", v1)

	out2 := MapArg2[int, string, bool, int](func(s string) int { return len(s) })(mo.NewEither3Arg2[int, string, bool]("ab"))
	v2, ok2 := out2.Arg2()
	is.True(ok2)
	is.Equal(2, v2)

	out3 := MapArg3[int, string, bool, int](func(b bool) int {
		if b {
			return 7
		}
		return 8
	})(mo.NewEither3Arg3[int, string, bool](true))
	v3, ok3 := out3.Arg3()
	is.True(ok3)
	is.Equal(7, v3)
}
