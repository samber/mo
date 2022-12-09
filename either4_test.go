package mo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewEither4(t *testing.T) {
	is := assert.New(t)

	either4Arg1 := NewEither4Arg1[int, bool, float64, string](42)
	is.Equal(Either4[int, bool, float64, string]{argId: either4ArgId1, arg1: 42}, either4Arg1)

	either4Arg2 := NewEither4Arg2[int, bool, float64, string](true)
	is.Equal(Either4[int, bool, float64, string]{argId: either4ArgId2, arg2: true}, either4Arg2)

	either4Arg3 := NewEither4Arg3[int, bool, float64, string](1.2)
	is.Equal(Either4[int, bool, float64, string]{argId: either4ArgId3, arg3: 1.2}, either4Arg3)

	either4Arg4 := NewEither4Arg4[int, bool, float64, string]("Hello")
	is.Equal(Either4[int, bool, float64, string]{argId: either4ArgId4, arg4: "Hello"}, either4Arg4)
}

func TestEither4IsArg(t *testing.T) {
	is := assert.New(t)

	either4Arg1 := NewEither4Arg1[int, bool, float64, string](42)
	either4Arg2 := NewEither4Arg2[int, bool, float64, string](true)
	either4Arg3 := NewEither4Arg3[int, bool, float64, string](1.2)
	either4Arg4 := NewEither4Arg4[int, bool, float64, string]("Hello")

	is.True(either4Arg1.IsArg1())
	is.False(either4Arg1.IsArg2())
	is.False(either4Arg1.IsArg3())
	is.False(either4Arg1.IsArg4())

	is.False(either4Arg2.IsArg1())
	is.True(either4Arg2.IsArg2())
	is.False(either4Arg2.IsArg3())
	is.False(either4Arg2.IsArg4())

	is.False(either4Arg3.IsArg1())
	is.False(either4Arg3.IsArg2())
	is.True(either4Arg3.IsArg3())
	is.False(either4Arg3.IsArg4())

	is.False(either4Arg4.IsArg1())
	is.False(either4Arg4.IsArg2())
	is.False(either4Arg4.IsArg3())
	is.True(either4Arg4.IsArg4())
}

func TestEither4GetArg(t *testing.T) {
	is := assert.New(t)

	either4Arg1 := NewEither4Arg1[int, bool, float64, string](42)
	either4Arg2 := NewEither4Arg2[int, bool, float64, string](true)
	either4Arg3 := NewEither4Arg3[int, bool, float64, string](1.2)
	either4Arg4 := NewEither4Arg4[int, bool, float64, string]("Hello")

	result1_1, ok1_1 := either4Arg1.Arg1()
	_, ok1_2 := either4Arg1.Arg2()
	_, ok1_3 := either4Arg1.Arg3()
	_, ok1_4 := either4Arg1.Arg4()

	is.Equal(42, result1_1)
	is.True(ok1_1)
	is.False(ok1_2)
	is.False(ok1_3)
	is.False(ok1_4)

	_, ok2_1 := either4Arg2.Arg1()
	result2, ok2_2 := either4Arg2.Arg2()
	_, ok2_3 := either4Arg2.Arg3()
	_, ok2_4 := either4Arg2.Arg4()

	is.Equal(true, result2)
	is.False(ok2_1)
	is.True(ok2_2)
	is.False(ok2_3)
	is.False(ok2_4)

	_, ok3_1 := either4Arg3.Arg1()
	_, ok3_2 := either4Arg3.Arg2()
	result3, ok3_3 := either4Arg3.Arg3()
	_, ok3_4 := either4Arg3.Arg4()

	is.Equal(1.2, result3)
	is.False(ok3_1)
	is.False(ok3_2)
	is.True(ok3_3)
	is.False(ok3_4)

	_, ok4_1 := either4Arg4.Arg1()
	_, ok4_2 := either4Arg4.Arg2()
	_, ok4_3 := either4Arg4.Arg3()
	result4, ok4_4 := either4Arg4.Arg4()

	is.Equal("Hello", result4)
	is.False(ok4_1)
	is.False(ok4_2)
	is.False(ok4_3)
	is.True(ok4_4)
}

func TestEither4MustArg(t *testing.T) {
	is := assert.New(t)

	either4Arg1 := NewEither4Arg1[int, bool, float64, string](42)
	either4Arg2 := NewEither4Arg2[int, bool, float64, string](true)
	either4Arg3 := NewEither4Arg3[int, bool, float64, string](1.2)
	either4Arg4 := NewEither4Arg4[int, bool, float64, string]("Hello")

	is.NotPanics(func() {
		is.Equal(42, either4Arg1.MustArg1())
	})
	is.Panics(func() {
		either4Arg1.MustArg2()
	})
	is.Panics(func() {
		either4Arg1.MustArg3()
	})
	is.Panics(func() {
		either4Arg1.MustArg4()
	})

	is.Panics(func() {
		either4Arg2.MustArg1()
	})
	is.NotPanics(func() {
		is.Equal(true, either4Arg2.MustArg2())
	})
	is.Panics(func() {
		either4Arg2.MustArg3()
	})
	is.Panics(func() {
		either4Arg2.MustArg4()
	})

	is.Panics(func() {
		either4Arg3.MustArg1()
	})
	is.Panics(func() {
		either4Arg3.MustArg2()
	})
	is.NotPanics(func() {
		is.Equal(1.2, either4Arg3.MustArg3())
	})
	is.Panics(func() {
		either4Arg3.MustArg4()
	})

	is.Panics(func() {
		either4Arg4.MustArg1()
	})
	is.Panics(func() {
		either4Arg4.MustArg2()
	})
	is.Panics(func() {
		either4Arg4.MustArg3()
	})
	is.NotPanics(func() {
		is.Equal("Hello", either4Arg4.MustArg4())
	})
}

func TestEither4Unpack(t *testing.T) {
	is := assert.New(t)

	either := NewEither4Arg1[int, bool, float64, string](42)
	either4Arg1, either4Arg2, either4Arg3, either4Arg4 := either.Unpack()

	is.Equal(42, either4Arg1)
	is.Equal(false, either4Arg2)
	is.Equal(float64(0), either4Arg3)
	is.Equal("", either4Arg4)
}

func TestEither4GetOrElse(t *testing.T) {
	is := assert.New(t)

	either4Arg1 := NewEither4Arg1[int, bool, float64, string](42)
	either4Arg2 := NewEither4Arg2[int, bool, float64, string](true)
	either4Arg3 := NewEither4Arg3[int, bool, float64, string](1.2)
	either4Arg4 := NewEither4Arg4[int, bool, float64, string]("Hello")

	is.Equal(42, either4Arg1.Arg1OrElse(21))
	is.Equal(false, either4Arg1.Arg2OrElse(false))
	is.Equal(2.1, either4Arg1.Arg3OrElse(2.1))
	is.Equal("Bye", either4Arg1.Arg4OrElse("Bye"))

	is.Equal(21, either4Arg2.Arg1OrElse(21))
	is.Equal(true, either4Arg2.Arg2OrElse(false))
	is.Equal(2.1, either4Arg2.Arg3OrElse(2.1))
	is.Equal("Bye", either4Arg2.Arg4OrElse("Bye"))

	is.Equal(21, either4Arg3.Arg1OrElse(21))
	is.Equal(false, either4Arg3.Arg2OrElse(false))
	is.Equal(1.2, either4Arg3.Arg3OrElse(2.1))
	is.Equal("Bye", either4Arg3.Arg4OrElse("Bye"))

	is.Equal(21, either4Arg4.Arg1OrElse(21))
	is.Equal(false, either4Arg4.Arg2OrElse(false))
	is.Equal(2.1, either4Arg4.Arg3OrElse(2.1))
	is.Equal("Hello", either4Arg4.Arg4OrElse("Bye"))
}

func TestEither4GetOrEmpty(t *testing.T) {
	is := assert.New(t)

	either4Arg1 := NewEither4Arg1[int, bool, float64, string](42)
	either4Arg2 := NewEither4Arg2[int, bool, float64, string](true)
	either4Arg3 := NewEither4Arg3[int, bool, float64, string](1.2)
	either4Arg4 := NewEither4Arg4[int, bool, float64, string]("Hello")

	is.Equal(42, either4Arg1.Arg1OrEmpty())
	is.Equal(false, either4Arg1.Arg2OrEmpty())
	is.Equal(0.0, either4Arg1.Arg3OrEmpty())
	is.Equal("", either4Arg1.Arg4OrEmpty())

	is.Equal(0, either4Arg2.Arg1OrEmpty())
	is.Equal(true, either4Arg2.Arg2OrEmpty())
	is.Equal(0.0, either4Arg2.Arg3OrEmpty())
	is.Equal("", either4Arg2.Arg4OrEmpty())

	is.Equal(0, either4Arg3.Arg1OrEmpty())
	is.Equal(false, either4Arg3.Arg2OrEmpty())
	is.Equal(1.2, either4Arg3.Arg3OrEmpty())
	is.Equal("", either4Arg3.Arg4OrEmpty())

	is.Equal(0, either4Arg4.Arg1OrEmpty())
	is.Equal(false, either4Arg4.Arg2OrEmpty())
	is.Equal(0.0, either4Arg4.Arg3OrEmpty())
	is.Equal("Hello", either4Arg4.Arg4OrEmpty())
}

func TestEither4ForEach(t *testing.T) {
	is := assert.New(t)

	NewEither4Arg1[int, bool, float64, string](42).ForEach(func(v1 int) {
		is.Equal(42, v1)
	}, func(v2 bool) {
		is.Fail("should not enter here")
	}, func(v3 float64) {
		is.Fail("should not enter here")
	}, func(v4 string) {
		is.Fail("should not enter here")
	})

	NewEither4Arg2[int, bool, float64, string](true).ForEach(func(v1 int) {
		is.Fail("should not enter here")
	}, func(v2 bool) {
		is.Equal(true, v2)
	}, func(v3 float64) {
		is.Fail("should not enter here")
	}, func(v4 string) {
		is.Fail("should not enter here")
	})

	NewEither4Arg3[int, bool, float64, string](1.2).ForEach(func(v1 int) {
		is.Fail("should not enter here")
	}, func(v2 bool) {
		is.Fail("should not enter here")
	}, func(v3 float64) {
		is.Equal(1.2, v3)
	}, func(v4 string) {
		is.Fail("should not enter here")
	})

	NewEither4Arg4[int, bool, float64, string]("Hello").ForEach(func(v1 int) {
		is.Fail("should not enter here")
	}, func(v2 bool) {
		is.Fail("should not enter here")
	}, func(v3 float64) {
		is.Fail("should not enter here")
	}, func(v4 string) {
		is.Equal("Hello", v4)
	})
}

func TestEither4Match(t *testing.T) {
	is := assert.New(t)

	e1 := NewEither4Arg1[int, bool, float64, string](42).Match(func(v1 int) Either4[int, bool, float64, string] {
		is.Equal(42, v1)
		return NewEither4Arg1[int, bool, float64, string](21)
	}, func(v2 bool) Either4[int, bool, float64, string] {
		is.Fail("should not enter here")
		return NewEither4Arg2[int, bool, float64, string](false)
	}, func(v3 float64) Either4[int, bool, float64, string] {
		is.Fail("should not enter here")
		return NewEither4Arg3[int, bool, float64, string](2.1)
	}, func(v4 string) Either4[int, bool, float64, string] {
		is.Fail("should not enter here")
		return NewEither4Arg4[int, bool, float64, string]("Bye")
	})

	is.Equal(NewEither4Arg1[int, bool, float64, string](21), e1)

	e2 := NewEither4Arg2[int, bool, float64, string](true).Match(func(v1 int) Either4[int, bool, float64, string] {
		is.Fail("should not enter here")
		return NewEither4Arg1[int, bool, float64, string](21)
	}, func(v2 bool) Either4[int, bool, float64, string] {
		is.Equal(true, v2)
		return NewEither4Arg2[int, bool, float64, string](false)
	}, func(v3 float64) Either4[int, bool, float64, string] {
		is.Fail("should not enter here")
		return NewEither4Arg3[int, bool, float64, string](2.1)
	}, func(v4 string) Either4[int, bool, float64, string] {
		is.Fail("should not enter here")
		return NewEither4Arg4[int, bool, float64, string]("Bye")
	})

	is.Equal(NewEither4Arg2[int, bool, float64, string](false), e2)

	e3 := NewEither4Arg3[int, bool, float64, string](1.2).Match(func(v1 int) Either4[int, bool, float64, string] {
		is.Fail("should not enter here")
		return NewEither4Arg1[int, bool, float64, string](21)
	}, func(v2 bool) Either4[int, bool, float64, string] {
		is.Fail("should not enter here")
		return NewEither4Arg2[int, bool, float64, string](false)
	}, func(v3 float64) Either4[int, bool, float64, string] {
		is.Equal(1.2, v3)
		return NewEither4Arg3[int, bool, float64, string](2.1)
	}, func(v4 string) Either4[int, bool, float64, string] {
		is.Fail("should not enter here")
		return NewEither4Arg4[int, bool, float64, string]("Bye")
	})

	is.Equal(NewEither4Arg3[int, bool, float64, string](2.1), e3)

	e4 := NewEither4Arg4[int, bool, float64, string]("Hello").Match(func(v1 int) Either4[int, bool, float64, string] {
		is.Fail("should not enter here")
		return NewEither4Arg1[int, bool, float64, string](21)
	}, func(v2 bool) Either4[int, bool, float64, string] {
		is.Fail("should not enter here")
		return NewEither4Arg2[int, bool, float64, string](false)
	}, func(v3 float64) Either4[int, bool, float64, string] {
		is.Fail("should not enter here")
		return NewEither4Arg3[int, bool, float64, string](2.1)
	}, func(v4 string) Either4[int, bool, float64, string] {
		is.Equal("Hello", v4)
		return NewEither4Arg4[int, bool, float64, string]("Bye")
	})

	is.Equal(NewEither4Arg4[int, bool, float64, string]("Bye"), e4)
}

func TestEither4MapArg(t *testing.T) {
	is := assert.New(t)

	either4Arg1 := NewEither4Arg1[int, bool, float64, string](42)
	either4Arg2 := NewEither4Arg2[int, bool, float64, string](true)
	either4Arg3 := NewEither4Arg3[int, bool, float64, string](1.2)
	either4Arg4 := NewEither4Arg4[int, bool, float64, string]("Hello")

	result1_1 := either4Arg1.MapArg1(func(v int) Either4[int, bool, float64, string] {
		is.Equal(42, v)
		return NewEither4Arg1[int, bool, float64, string](21)
	})
	is.Equal(NewEither4Arg1[int, bool, float64, string](21), result1_1)

	result1_2 := either4Arg1.MapArg2(func(v bool) Either4[int, bool, float64, string] {
		is.Fail("should not enter here")
		return NewEither4Arg2[int, bool, float64, string](false)
	})
	is.Equal(either4Arg1, result1_2)

	result1_3 := either4Arg1.MapArg3(func(v float64) Either4[int, bool, float64, string] {
		is.Fail("should not enter here")
		return NewEither4Arg3[int, bool, float64, string](2.1)
	})
	is.Equal(either4Arg1, result1_3)

	result1_4 := either4Arg1.MapArg4(func(v string) Either4[int, bool, float64, string] {
		is.Fail("should not enter here")
		return NewEither4Arg4[int, bool, float64, string]("Bye")
	})
	is.Equal(either4Arg1, result1_4)

	result2_1 := either4Arg2.MapArg1(func(v int) Either4[int, bool, float64, string] {
		is.Fail("should not enter here")
		return NewEither4Arg1[int, bool, float64, string](21)
	})
	is.Equal(either4Arg2, result2_1)

	result2_2 := either4Arg2.MapArg2(func(v bool) Either4[int, bool, float64, string] {
		is.Equal(true, v)
		return NewEither4Arg2[int, bool, float64, string](false)
	})
	is.Equal(NewEither4Arg2[int, bool, float64, string](false), result2_2)

	result2_3 := either4Arg2.MapArg3(func(v float64) Either4[int, bool, float64, string] {
		is.Fail("should not enter here")
		return NewEither4Arg3[int, bool, float64, string](2.1)
	})
	is.Equal(either4Arg2, result2_3)

	result2_4 := either4Arg2.MapArg4(func(v string) Either4[int, bool, float64, string] {
		is.Fail("should not enter here")
		return NewEither4Arg4[int, bool, float64, string]("Bye")
	})
	is.Equal(either4Arg2, result2_4)

	result3_1 := either4Arg3.MapArg1(func(v int) Either4[int, bool, float64, string] {
		is.Fail("should not enter here")
		return NewEither4Arg3[int, bool, float64, string](21)
	})
	is.Equal(either4Arg3, result3_1)

	result3_2 := either4Arg3.MapArg2(func(v bool) Either4[int, bool, float64, string] {
		is.Fail("should not enter here")
		return NewEither4Arg2[int, bool, float64, string](false)
	})
	is.Equal(either4Arg3, result3_2)

	result3_3 := either4Arg3.MapArg3(func(v float64) Either4[int, bool, float64, string] {
		is.Equal(1.2, v)
		return NewEither4Arg3[int, bool, float64, string](2.1)
	})
	is.Equal(NewEither4Arg3[int, bool, float64, string](2.1), result3_3)

	result3_4 := either4Arg3.MapArg4(func(v string) Either4[int, bool, float64, string] {
		is.Fail("should not enter here")
		return NewEither4Arg4[int, bool, float64, string]("Bye")
	})
	is.Equal(either4Arg3, result3_4)

	result4_1 := either4Arg4.MapArg1(func(v int) Either4[int, bool, float64, string] {
		is.Fail("should not enter here")
		return NewEither4Arg1[int, bool, float64, string](21)
	})
	is.Equal(either4Arg4, result4_1)

	result4_2 := either4Arg4.MapArg2(func(v bool) Either4[int, bool, float64, string] {
		is.Fail("should not enter here")
		return NewEither4Arg2[int, bool, float64, string](false)
	})
	is.Equal(either4Arg4, result4_2)

	result4_3 := either4Arg4.MapArg3(func(v float64) Either4[int, bool, float64, string] {
		is.Fail("should not enter here")
		return NewEither4Arg3[int, bool, float64, string](2.1)
	})
	is.Equal(either4Arg4, result4_3)

	result4_4 := either4Arg4.MapArg4(func(v string) Either4[int, bool, float64, string] {
		is.Equal("Hello", v)
		return NewEither4Arg4[int, bool, float64, string]("Bye")
	})
	is.Equal(NewEither4Arg4[int, bool, float64, string]("Bye"), result4_4)
}
