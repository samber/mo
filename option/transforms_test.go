package option

import (
    "testing"

    "github.com/samber/mo"
    "github.com/stretchr/testify/assert"
)

func TestMap(t *testing.T) {
    is := assert.New(t)

    out := Map[int, string](func(i int) string { return "n=" + string(rune('0'+i)) })(mo.Some(3))
    v, ok := out.Get()
    is.True(ok)
    is.Equal("n=3", v)

    none := Map[int, string](func(i int) string { return "x" })(mo.None[int]())
    _, ok = none.Get()
    is.False(ok)
}

func TestFlatMap(t *testing.T) {
    is := assert.New(t)

    out := FlatMap[int, string](func(i int) mo.Option[string] { return mo.Some("k") })(mo.Some(1))
    v, ok := out.Get()
    is.True(ok)
    is.Equal("k", v)

    none := FlatMap[int, string](func(i int) mo.Option[string] { return mo.Some("k") })(mo.None[int]())
    _, ok = none.Get()
    is.False(ok)
}

func TestMatch(t *testing.T) {
    is := assert.New(t)

    m := Match[int, string](func(i int) (string, bool) { return "i", true }, func() (string, bool) { return "n", true })
    v1, ok1 := m(mo.Some(1)).Get()
    is.True(ok1)
    is.Equal("i", v1)

    v2, ok2 := m(mo.None[int]()).Get()
    is.True(ok2)
    is.Equal("n", v2)
}

func TestFlatMatch(t *testing.T) {
    is := assert.New(t)

    fm := FlatMatch[int, string](func(i int) mo.Option[string] { return mo.Some("i") }, func() mo.Option[string] { return mo.Some("n") })
    v1, ok1 := fm(mo.Some(1)).Get()
    is.True(ok1)
    is.Equal("i", v1)

    v2, ok2 := fm(mo.None[int]()).Get()
    is.True(ok2)
    is.Equal("n", v2)
}
