package mo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOptionSome(t *testing.T) {
	is := assert.New(t)

	is.Equal(Option[int]{value: 42, isPresent: true}, Some(42))
}

func TestOptionNone(t *testing.T) {
	is := assert.New(t)

	is.Equal(Option[int]{isPresent: false}, None[int]())
}

func TestOptionIsPresent(t *testing.T) {
	is := assert.New(t)

	is.True(Some(42).IsPresent())
	is.False(None[int]().IsPresent())
}

func TestOptionIsAbsent(t *testing.T) {
	is := assert.New(t)

	is.False(Some(42).IsAbsent())
	is.True(None[int]().IsAbsent())
}

func TestOptionSize(t *testing.T) {
	is := assert.New(t)

	is.Equal(1, Some(42).Size())
	is.Equal(0, None[int]().Size())
}

func TestOptionGet(t *testing.T) {
	is := assert.New(t)

	v1, ok1 := Some(42).Get()
	v2, ok2 := None[int]().Get()

	is.Equal(42, v1)
	is.Equal(true, ok1)
	is.Equal(0, v2)
	is.Equal(false, ok2)
}

func TestOptionMustGet(t *testing.T) {
	is := assert.New(t)

	is.NotPanics(func() {
		Some(42).MustGet()
	})
	is.Panics(func() {
		None[int]().MustGet()
	})

	is.Equal(42, Some(42).MustGet())
}

func TestOptionOrElse(t *testing.T) {
	is := assert.New(t)

	is.Equal(42, Some(42).OrElse(21))
	is.Equal(21, None[int]().OrElse(21))
}

func TestOptionOrEmpty(t *testing.T) {
	is := assert.New(t)

	is.Equal(42, Some(42).OrEmpty())
	is.Equal(0, None[int]().OrEmpty())
}

func TestOptionForEach(t *testing.T) {
	is := assert.New(t)

	tmp := 0
	f := func(x int) {
		tmp = x
	}

	None[int]().ForEach(f)
	is.Equal(0, tmp)

	Some(42).ForEach(f)
	is.Equal(42, tmp)
}

func TestOptionMatch(t *testing.T) {
	is := assert.New(t)

	onValue := func(i int) (int, bool) {
		return i * 2, true
	}
	onNone := func() (int, bool) {
		return 0, false
	}

	opt1 := Some(21).Match(onValue, onNone)
	opt2 := None[int]().Match(onValue, onNone)

	is.Equal(Option[int]{value: 42, isPresent: true}, opt1)
	is.Equal(Option[int]{value: 0, isPresent: false}, opt2)
}

func TestOptionMap(t *testing.T) {
	is := assert.New(t)

	opt1 := Some(21).Map(func(i int) (int, bool) {
		return i * 2, true
	})
	opt2 := None[int]().Map(func(i int) (int, bool) {
		is.Fail("should not be called")
		return 42, true
	})

	is.Equal(Option[int]{value: 42, isPresent: true}, opt1)
	is.Equal(Option[int]{value: 0, isPresent: false}, opt2)
}

func TestOptionMapNone(t *testing.T) {
	is := assert.New(t)

	opt1 := Some(21).MapNone(func() (int, bool) {
		is.Fail("should not be called")
		return 42, true
	})
	opt2 := None[int]().MapNone(func() (int, bool) {
		return 42, true
	})

	is.Equal(Option[int]{value: 21, isPresent: true}, opt1)
	is.Equal(Option[int]{value: 42, isPresent: true}, opt2)
}

func TestOptionFlatMap(t *testing.T) {
	is := assert.New(t)

	opt1 := Some(21).FlatMap(func(i int) Option[int] {
		return Some(42)
	})
	opt2 := None[int]().FlatMap(func(i int) Option[int] {
		return Some(42)
	})

	is.Equal(Option[int]{value: 42, isPresent: true}, opt1)
	is.Equal(Option[int]{value: 0, isPresent: false}, opt2)
}

func TestOptionMarshalJSON(t *testing.T) {
	is := assert.New(t)

	option1 := Some("foo")
	option2 := None[string]()

	value, err := option1.MarshalJSON()
	is.NoError(err)
	is.Equal(`"foo"`, string(value))

	value, err = option2.MarshalJSON()
	is.NoError(err)
	is.Equal(`null`, string(value))
}

func TestOptionUnmarshalJSON(t *testing.T) {
	is := assert.New(t)

	option1 := Some("foo")
	option2 := None[string]()

	err := option1.UnmarshalJSON([]byte(`"foo"`))
	is.NoError(err)
	is.Equal(option1, Some("foo"))

	err = option2.UnmarshalJSON([]byte(`null`))
	is.NoError(err)
	is.Equal(option2, None[string]())
}
