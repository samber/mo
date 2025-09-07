package either

import (
	"fmt"
	"testing"

	"github.com/samber/mo"
	"github.com/stretchr/testify/assert"
)

func TestMapLeft(t *testing.T) {
	is := assert.New(t)

	// Left should map to Left using the provided function
	out1 := MapLeft[int, string](func(i int) string { return fmt.Sprintf("n=%d", i) })(mo.Left[int, string](3))
	l1, ok1 := out1.Left()
	is.True(ok1)
	is.Equal("n=3", l1)

	// Right should stay Right carrying the right value
	out2 := MapLeft[int, string](func(i int) string { return "ignored" })(mo.Right[int, string]("x"))
	r2, ok2 := out2.Right()
	is.True(ok2)
	is.Equal("x", r2)
}

func TestMapRight(t *testing.T) {
	is := assert.New(t)

	// Right should map to Right using the provided function
	out1 := MapRight[int, string](func(s string) string { return s + "!" })(mo.Right[int, string]("a"))
	r1, ok1 := out1.Right()
	is.True(ok1)
	is.Equal("a!", r1)

	// Left should stay Left carrying the left value
	out2 := MapRight[int, string](func(s string) string { return s + "!" })(mo.Left[int, string](2))
	l2, ok2 := out2.Left()
	is.True(ok2)
	is.Equal(2, l2)
}

func TestMatch(t *testing.T) {
	is := assert.New(t)

	toLenOrTwice := Match[int, string, int, int](func(i int) int { return i * 2 }, func(s string) int { return len(s) })

	out1 := toLenOrTwice(mo.Left[int, string](3))
	l1, ok1 := out1.Left()
	is.True(ok1)
	is.Equal(6, l1)

	out2 := toLenOrTwice(mo.Right[int, string]("abcd"))
	r2, ok2 := out2.Right()
	is.True(ok2)
	is.Equal(4, r2)
}

func TestSwap(t *testing.T) {
	is := assert.New(t)

	out1 := Swap[int, string]()(mo.Left[int, string](2))
	r1, ok1 := out1.Right()
	is.True(ok1)
	is.Equal(2, r1)

	out2 := Swap[int, string]()(mo.Right[int, string]("x"))
	l2, ok2 := out2.Left()
	is.True(ok2)
	is.Equal("x", l2)
}
