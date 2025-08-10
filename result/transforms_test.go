package result

import (
    "errors"
    "testing"

    "github.com/samber/mo"
    "github.com/stretchr/testify/assert"
)

func TestMap(t *testing.T) {
    is := assert.New(t)

    out := Map[int, string](func(i int) string { return "n=" + string(rune('0'+i)) })(mo.Ok(3))
    v, err := out.Get()
    is.NoError(err)
    is.Equal("n=3", v)

    errRes := Map[int, string](func(i int) string { return "x" })(mo.Err[int](errors.New("boom")))
    _, err = errRes.Get()
    is.Error(err)
}

func TestFlatMap(t *testing.T) {
    is := assert.New(t)

    out := FlatMap[int, string](func(i int) mo.Result[string] { return mo.Ok("k") })(mo.Ok(1))
    v, err := out.Get()
    is.NoError(err)
    is.Equal("k", v)

    errRes := FlatMap[int, string](func(i int) mo.Result[string] { return mo.Ok("k") })(mo.Err[int](errors.New("boom")))
    _, err = errRes.Get()
    is.Error(err)
}

func TestMatch(t *testing.T) {
    is := assert.New(t)

    m := Match[int, string](func(i int) (string, error) { return "i", nil }, func() (string, error) { return "n", nil })
    v1, err1 := m(mo.Ok(1)).Get()
    is.NoError(err1)
    is.Equal("i", v1)

    v2, err2 := m(mo.Err[int](errors.New("boom"))).Get()
    is.NoError(err2)
    is.Equal("n", v2)
}

func TestFlatMatch(t *testing.T) {
    is := assert.New(t)

    fm := FlatMatch[int, string](func(i int) mo.Result[string] { return mo.Ok("i") }, func() mo.Result[string] { return mo.Ok("n") })
    v1, err1 := fm(mo.Ok(1)).Get()
    is.NoError(err1)
    is.Equal("i", v1)

    v2, err2 := fm(mo.Err[int](errors.New("boom"))).Get()
    is.NoError(err2)
    is.Equal("n", v2)
}
