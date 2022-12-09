package mo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewEither5(t *testing.T) {
	is := assert.New(t)

	either5Arg1 := NewEither5Arg1[int, bool, float64, string, byte](42)
	is.Equal(Either5[int, bool, float64, string, byte]{argId: either5ArgId1, arg1: 42}, either5Arg1)

	either5Arg2 := NewEither5Arg2[int, bool, float64, string, byte](true)
	is.Equal(Either5[int, bool, float64, string, byte]{argId: either5ArgId2, arg2: true}, either5Arg2)

	either5Arg3 := NewEither5Arg3[int, bool, float64, string, byte](1.2)
	is.Equal(Either5[int, bool, float64, string, byte]{argId: either5ArgId3, arg3: 1.2}, either5Arg3)

	either5Arg4 := NewEither5Arg4[int, bool, float64, string, byte]("Hello")
	is.Equal(Either5[int, bool, float64, string, byte]{argId: either5ArgId4, arg4: "Hello"}, either5Arg4)

	either5Arg5 := NewEither5Arg5[int, bool, float64, string, byte](5)
	is.Equal(Either5[int, bool, float64, string, byte]{argId: either5ArgId5, arg5: 5}, either5Arg5)
}

func TestEither5IsArg(t *testing.T) {
	is := assert.New(t)

	either5Arg1 := NewEither5Arg1[int, bool, float64, string, byte](42)
	either5Arg2 := NewEither5Arg2[int, bool, float64, string, byte](true)
	either5Arg3 := NewEither5Arg3[int, bool, float64, string, byte](1.2)
	either5Arg4 := NewEither5Arg4[int, bool, float64, string, byte]("Hello")
	either5Arg5 := NewEither5Arg5[int, bool, float64, string, byte](5)

	is.True(either5Arg1.IsArg1())
	is.False(either5Arg1.IsArg2())
	is.False(either5Arg1.IsArg3())
	is.False(either5Arg1.IsArg4())
	is.False(either5Arg1.IsArg5())

	is.False(either5Arg2.IsArg1())
	is.True(either5Arg2.IsArg2())
	is.False(either5Arg2.IsArg3())
	is.False(either5Arg2.IsArg4())
	is.False(either5Arg2.IsArg5())

	is.False(either5Arg3.IsArg1())
	is.False(either5Arg3.IsArg2())
	is.True(either5Arg3.IsArg3())
	is.False(either5Arg3.IsArg4())
	is.False(either5Arg3.IsArg5())

	is.False(either5Arg4.IsArg1())
	is.False(either5Arg4.IsArg2())
	is.False(either5Arg4.IsArg3())
	is.True(either5Arg4.IsArg4())
	is.False(either5Arg4.IsArg5())

	is.False(either5Arg5.IsArg1())
	is.False(either5Arg5.IsArg2())
	is.False(either5Arg5.IsArg3())
	is.False(either5Arg5.IsArg4())
	is.True(either5Arg5.IsArg5())
}

func TestEither5GetArg(t *testing.T) {
	is := assert.New(t)

	either5Arg1 := NewEither5Arg1[int, bool, float64, string, byte](42)
	either5Arg2 := NewEither5Arg2[int, bool, float64, string, byte](true)
	either5Arg3 := NewEither5Arg3[int, bool, float64, string, byte](1.2)
	either5Arg4 := NewEither5Arg4[int, bool, float64, string, byte]("Hello")
	either5Arg5 := NewEither5Arg5[int, bool, float64, string, byte](5)

	result1_1, ok1_1 := either5Arg1.Arg1()
	_, ok1_2 := either5Arg1.Arg2()
	_, ok1_3 := either5Arg1.Arg3()
	_, ok1_4 := either5Arg1.Arg4()
	_, ok1_5 := either5Arg1.Arg5()

	is.Equal(42, result1_1)
	is.True(ok1_1)
	is.False(ok1_2)
	is.False(ok1_3)
	is.False(ok1_4)
	is.False(ok1_5)

	_, ok2_1 := either5Arg2.Arg1()
	result2, ok2_2 := either5Arg2.Arg2()
	_, ok2_3 := either5Arg2.Arg3()
	_, ok2_4 := either5Arg2.Arg4()
	_, ok2_5 := either5Arg2.Arg5()

	is.Equal(true, result2)
	is.False(ok2_1)
	is.True(ok2_2)
	is.False(ok2_3)
	is.False(ok2_4)
	is.False(ok2_5)

	_, ok3_1 := either5Arg3.Arg1()
	_, ok3_2 := either5Arg3.Arg2()
	result3, ok3_3 := either5Arg3.Arg3()
	_, ok3_4 := either5Arg3.Arg4()
	_, ok3_5 := either5Arg3.Arg5()

	is.Equal(1.2, result3)
	is.False(ok3_1)
	is.False(ok3_2)
	is.True(ok3_3)
	is.False(ok3_4)
	is.False(ok3_5)

	_, ok4_1 := either5Arg4.Arg1()
	_, ok4_2 := either5Arg4.Arg2()
	_, ok4_3 := either5Arg4.Arg3()
	result4, ok4_4 := either5Arg4.Arg4()
	_, ok4_5 := either5Arg4.Arg5()

	is.Equal("Hello", result4)
	is.False(ok4_1)
	is.False(ok4_2)
	is.False(ok4_3)
	is.True(ok4_4)
	is.False(ok4_5)

	_, ok5_1 := either5Arg5.Arg1()
	_, ok5_2 := either5Arg5.Arg2()
	_, ok5_3 := either5Arg5.Arg3()
	_, ok5_4 := either5Arg5.Arg4()
	result5, ok5_5 := either5Arg5.Arg5()

	is.Equal(byte(5), result5)
	is.False(ok5_1)
	is.False(ok5_2)
	is.False(ok5_3)
	is.False(ok5_4)
	is.True(ok5_5)
}

func TestEither5MustArg(t *testing.T) {
	is := assert.New(t)

	either5Arg1 := NewEither5Arg1[int, bool, float64, string, byte](42)
	either5Arg2 := NewEither5Arg2[int, bool, float64, string, byte](true)
	either5Arg3 := NewEither5Arg3[int, bool, float64, string, byte](1.2)
	either5Arg4 := NewEither5Arg4[int, bool, float64, string, byte]("Hello")
	either5Arg5 := NewEither5Arg5[int, bool, float64, string, byte](5)

	is.NotPanics(func() {
		is.Equal(42, either5Arg1.MustArg1())
	})
	is.Panics(func() {
		either5Arg1.MustArg2()
	})
	is.Panics(func() {
		either5Arg1.MustArg3()
	})
	is.Panics(func() {
		either5Arg1.MustArg4()
	})
	is.Panics(func() {
		either5Arg1.MustArg5()
	})

	is.Panics(func() {
		either5Arg2.MustArg1()
	})
	is.NotPanics(func() {
		is.Equal(true, either5Arg2.MustArg2())
	})
	is.Panics(func() {
		either5Arg2.MustArg3()
	})
	is.Panics(func() {
		either5Arg2.MustArg4()
	})
	is.Panics(func() {
		either5Arg2.MustArg5()
	})

	is.Panics(func() {
		either5Arg3.MustArg1()
	})
	is.Panics(func() {
		either5Arg3.MustArg2()
	})
	is.NotPanics(func() {
		is.Equal(1.2, either5Arg3.MustArg3())
	})
	is.Panics(func() {
		either5Arg3.MustArg4()
	})
	is.Panics(func() {
		either5Arg3.MustArg5()
	})

	is.Panics(func() {
		either5Arg4.MustArg1()
	})
	is.Panics(func() {
		either5Arg4.MustArg2()
	})
	is.Panics(func() {
		either5Arg4.MustArg3()
	})
	is.NotPanics(func() {
		is.Equal("Hello", either5Arg4.MustArg4())
	})
	is.Panics(func() {
		either5Arg4.MustArg5()
	})

	is.Panics(func() {
		either5Arg5.MustArg1()
	})
	is.Panics(func() {
		either5Arg5.MustArg2()
	})
	is.Panics(func() {
		either5Arg5.MustArg3()
	})
	is.Panics(func() {
		either5Arg5.MustArg4()
	})
	is.NotPanics(func() {
		is.Equal(byte(5), either5Arg5.MustArg5())
	})
}

func TestEither5Unpack(t *testing.T) {
	is := assert.New(t)

	either := NewEither5Arg1[int, bool, float64, string, string](42)
	either5Arg1, either5Arg2, either5Arg3, either5Arg4, either5Arg5 := either.Unpack()

	is.Equal(42, either5Arg1)
	is.Equal(false, either5Arg2)
	is.Equal(float64(0), either5Arg3)
	is.Equal("", either5Arg4)
	is.Equal("", either5Arg5)
}

func TestEither5GetOrElse(t *testing.T) {
	is := assert.New(t)

	either5Arg1 := NewEither5Arg1[int, bool, float64, string, byte](42)
	either5Arg2 := NewEither5Arg2[int, bool, float64, string, byte](true)
	either5Arg3 := NewEither5Arg3[int, bool, float64, string, byte](1.2)
	either5Arg4 := NewEither5Arg4[int, bool, float64, string, byte]("Hello")
	either5Arg5 := NewEither5Arg5[int, bool, float64, string, byte](5)

	is.Equal(42, either5Arg1.Arg1OrElse(21))
	is.Equal(false, either5Arg1.Arg2OrElse(false))
	is.Equal(2.1, either5Arg1.Arg3OrElse(2.1))
	is.Equal("Bye", either5Arg1.Arg4OrElse("Bye"))
	is.Equal(byte(10), either5Arg1.Arg5OrElse(10))

	is.Equal(21, either5Arg2.Arg1OrElse(21))
	is.Equal(true, either5Arg2.Arg2OrElse(false))
	is.Equal(2.1, either5Arg2.Arg3OrElse(2.1))
	is.Equal("Bye", either5Arg2.Arg4OrElse("Bye"))
	is.Equal(byte(10), either5Arg2.Arg5OrElse(10))

	is.Equal(21, either5Arg3.Arg1OrElse(21))
	is.Equal(false, either5Arg3.Arg2OrElse(false))
	is.Equal(1.2, either5Arg3.Arg3OrElse(2.1))
	is.Equal("Bye", either5Arg3.Arg4OrElse("Bye"))
	is.Equal(byte(10), either5Arg3.Arg5OrElse(10))

	is.Equal(21, either5Arg4.Arg1OrElse(21))
	is.Equal(false, either5Arg4.Arg2OrElse(false))
	is.Equal(2.1, either5Arg4.Arg3OrElse(2.1))
	is.Equal("Hello", either5Arg4.Arg4OrElse("Bye"))
	is.Equal(byte(10), either5Arg4.Arg5OrElse(10))

	is.Equal(21, either5Arg5.Arg1OrElse(21))
	is.Equal(false, either5Arg5.Arg2OrElse(false))
	is.Equal(2.1, either5Arg5.Arg3OrElse(2.1))
	is.Equal("Bye", either5Arg5.Arg4OrElse("Bye"))
	is.Equal(byte(5), either5Arg5.Arg5OrElse(10))
}

func TestEither5GetOrEmpty(t *testing.T) {
	is := assert.New(t)

	either5Arg1 := NewEither5Arg1[int, bool, float64, string, byte](42)
	either5Arg2 := NewEither5Arg2[int, bool, float64, string, byte](true)
	either5Arg3 := NewEither5Arg3[int, bool, float64, string, byte](1.2)
	either5Arg4 := NewEither5Arg4[int, bool, float64, string, byte]("Hello")
	either5Arg5 := NewEither5Arg5[int, bool, float64, string, byte](5)

	is.Equal(42, either5Arg1.Arg1OrEmpty())
	is.Equal(false, either5Arg1.Arg2OrEmpty())
	is.Equal(0.0, either5Arg1.Arg3OrEmpty())
	is.Equal("", either5Arg1.Arg4OrEmpty())
	is.Equal(byte(0), either5Arg1.Arg5OrEmpty())

	is.Equal(0, either5Arg2.Arg1OrEmpty())
	is.Equal(true, either5Arg2.Arg2OrEmpty())
	is.Equal(0.0, either5Arg2.Arg3OrEmpty())
	is.Equal("", either5Arg2.Arg4OrEmpty())
	is.Equal(byte(0), either5Arg2.Arg5OrEmpty())

	is.Equal(0, either5Arg3.Arg1OrEmpty())
	is.Equal(false, either5Arg3.Arg2OrEmpty())
	is.Equal(1.2, either5Arg3.Arg3OrEmpty())
	is.Equal("", either5Arg3.Arg4OrEmpty())
	is.Equal(byte(0), either5Arg3.Arg5OrEmpty())

	is.Equal(0, either5Arg4.Arg1OrEmpty())
	is.Equal(false, either5Arg4.Arg2OrEmpty())
	is.Equal(0.0, either5Arg4.Arg3OrEmpty())
	is.Equal("Hello", either5Arg4.Arg4OrEmpty())
	is.Equal(byte(0), either5Arg4.Arg5OrEmpty())

	is.Equal(0, either5Arg5.Arg1OrEmpty())
	is.Equal(false, either5Arg5.Arg2OrEmpty())
	is.Equal(0.0, either5Arg5.Arg3OrEmpty())
	is.Equal("", either5Arg5.Arg4OrEmpty())
	is.Equal(byte(5), either5Arg5.Arg5OrEmpty())
}

func TestEither5ForEach(t *testing.T) {
	is := assert.New(t)

	NewEither5Arg1[int, bool, float64, string, byte](42).ForEach(func(v1 int) {
		is.Equal(42, v1)
	}, func(v2 bool) {
		is.Fail("should not enter here")
	}, func(v3 float64) {
		is.Fail("should not enter here")
	}, func(v4 string) {
		is.Fail("should not enter here")
	}, func(v5 byte) {
		is.Fail("should not enter here")
	})

	NewEither5Arg2[int, bool, float64, string, byte](true).ForEach(func(v1 int) {
		is.Fail("should not enter here")
	}, func(v2 bool) {
		is.Equal(true, v2)
	}, func(v3 float64) {
		is.Fail("should not enter here")
	}, func(v4 string) {
		is.Fail("should not enter here")
	}, func(v5 byte) {
		is.Fail("should not enter here")
	})

	NewEither5Arg3[int, bool, float64, string, byte](1.2).ForEach(func(v1 int) {
		is.Fail("should not enter here")
	}, func(v2 bool) {
		is.Fail("should not enter here")
	}, func(v3 float64) {
		is.Equal(1.2, v3)
	}, func(v4 string) {
		is.Fail("should not enter here")
	}, func(v5 byte) {
		is.Fail("should not enter here")
	})

	NewEither5Arg4[int, bool, float64, string, byte]("Hello").ForEach(func(v1 int) {
		is.Fail("should not enter here")
	}, func(v2 bool) {
		is.Fail("should not enter here")
	}, func(v3 float64) {
		is.Fail("should not enter here")
	}, func(v4 string) {
		is.Equal("Hello", v4)
	}, func(v5 byte) {
		is.Fail("should not enter here")
	})

	NewEither5Arg5[int, bool, float64, string, byte](5).ForEach(func(v1 int) {
		is.Fail("should not enter here")
	}, func(v2 bool) {
		is.Fail("should not enter here")
	}, func(v3 float64) {
		is.Fail("should not enter here")
	}, func(v4 string) {
		is.Fail("should not enter here")
	}, func(v5 byte) {
		is.Equal(byte(5), v5)
	})
}

func TestEither5Match(t *testing.T) {
	is := assert.New(t)

	e1 := NewEither5Arg1[int, bool, float64, string, byte](42).Match(func(v1 int) Either5[int, bool, float64, string, byte] {
		is.Equal(42, v1)
		return NewEither5Arg1[int, bool, float64, string, byte](21)
	}, func(v2 bool) Either5[int, bool, float64, string, byte] {
		is.Fail("should not enter here")
		return NewEither5Arg2[int, bool, float64, string, byte](false)
	}, func(v3 float64) Either5[int, bool, float64, string, byte] {
		is.Fail("should not enter here")
		return NewEither5Arg3[int, bool, float64, string, byte](2.1)
	}, func(v4 string) Either5[int, bool, float64, string, byte] {
		is.Fail("should not enter here")
		return NewEither5Arg4[int, bool, float64, string, byte]("Bye")
	}, func(v5 byte) Either5[int, bool, float64, string, byte] {
		is.Fail("should not enter here")
		return NewEither5Arg5[int, bool, float64, string, byte](10)
	})

	is.Equal(NewEither5Arg1[int, bool, float64, string, byte](21), e1)

	e2 := NewEither5Arg2[int, bool, float64, string, byte](true).Match(func(v1 int) Either5[int, bool, float64, string, byte] {
		is.Fail("should not enter here")
		return NewEither5Arg1[int, bool, float64, string, byte](21)
	}, func(v2 bool) Either5[int, bool, float64, string, byte] {
		is.Equal(true, v2)
		return NewEither5Arg2[int, bool, float64, string, byte](false)
	}, func(v3 float64) Either5[int, bool, float64, string, byte] {
		is.Fail("should not enter here")
		return NewEither5Arg3[int, bool, float64, string, byte](2.1)
	}, func(v4 string) Either5[int, bool, float64, string, byte] {
		is.Fail("should not enter here")
		return NewEither5Arg4[int, bool, float64, string, byte]("Bye")
	}, func(v5 byte) Either5[int, bool, float64, string, byte] {
		is.Fail("should not enter here")
		return NewEither5Arg5[int, bool, float64, string, byte](10)
	})

	is.Equal(NewEither5Arg2[int, bool, float64, string, byte](false), e2)

	e3 := NewEither5Arg3[int, bool, float64, string, byte](1.2).Match(func(v1 int) Either5[int, bool, float64, string, byte] {
		is.Fail("should not enter here")
		return NewEither5Arg1[int, bool, float64, string, byte](21)
	}, func(v2 bool) Either5[int, bool, float64, string, byte] {
		is.Fail("should not enter here")
		return NewEither5Arg2[int, bool, float64, string, byte](false)
	}, func(v3 float64) Either5[int, bool, float64, string, byte] {
		is.Equal(1.2, v3)
		return NewEither5Arg3[int, bool, float64, string, byte](2.1)
	}, func(v4 string) Either5[int, bool, float64, string, byte] {
		is.Fail("should not enter here")
		return NewEither5Arg4[int, bool, float64, string, byte]("Bye")
	}, func(v5 byte) Either5[int, bool, float64, string, byte] {
		is.Fail("should not enter here")
		return NewEither5Arg5[int, bool, float64, string, byte](10)
	})

	is.Equal(NewEither5Arg3[int, bool, float64, string, byte](2.1), e3)

	e4 := NewEither5Arg4[int, bool, float64, string, byte]("Hello").Match(func(v1 int) Either5[int, bool, float64, string, byte] {
		is.Fail("should not enter here")
		return NewEither5Arg1[int, bool, float64, string, byte](21)
	}, func(v2 bool) Either5[int, bool, float64, string, byte] {
		is.Fail("should not enter here")
		return NewEither5Arg2[int, bool, float64, string, byte](false)
	}, func(v3 float64) Either5[int, bool, float64, string, byte] {
		is.Fail("should not enter here")
		return NewEither5Arg3[int, bool, float64, string, byte](2.1)
	}, func(v4 string) Either5[int, bool, float64, string, byte] {
		is.Equal("Hello", v4)
		return NewEither5Arg4[int, bool, float64, string, byte]("Bye")
	}, func(v5 byte) Either5[int, bool, float64, string, byte] {
		is.Fail("should not enter here")
		return NewEither5Arg5[int, bool, float64, string, byte](10)
	})

	is.Equal(NewEither5Arg4[int, bool, float64, string, byte]("Bye"), e4)

	e5 := NewEither5Arg5[int, bool, float64, string, byte](5).Match(func(v1 int) Either5[int, bool, float64, string, byte] {
		is.Fail("should not enter here")
		return NewEither5Arg1[int, bool, float64, string, byte](21)
	}, func(v2 bool) Either5[int, bool, float64, string, byte] {
		is.Fail("should not enter here")
		return NewEither5Arg2[int, bool, float64, string, byte](false)
	}, func(v3 float64) Either5[int, bool, float64, string, byte] {
		is.Fail("should not enter here")
		return NewEither5Arg3[int, bool, float64, string, byte](2.1)
	}, func(v4 string) Either5[int, bool, float64, string, byte] {
		is.Fail("should not enter here")
		return NewEither5Arg4[int, bool, float64, string, byte]("Bye")
	}, func(v5 byte) Either5[int, bool, float64, string, byte] {
		is.Equal(byte(5), v5)
		return NewEither5Arg5[int, bool, float64, string, byte](10)
	})

	is.Equal(NewEither5Arg5[int, bool, float64, string, byte](10), e5)
}

func TestEither5MapArg(t *testing.T) {
	is := assert.New(t)

	either5Arg1 := NewEither5Arg1[int, bool, float64, string, byte](42)
	either5Arg2 := NewEither5Arg2[int, bool, float64, string, byte](true)
	either5Arg3 := NewEither5Arg3[int, bool, float64, string, byte](1.2)
	either5Arg4 := NewEither5Arg4[int, bool, float64, string, byte]("Hello")
	either5Arg5 := NewEither5Arg5[int, bool, float64, string, byte](5)

	result1_1 := either5Arg1.MapArg1(func(v int) Either5[int, bool, float64, string, byte] {
		is.Equal(42, v)
		return NewEither5Arg1[int, bool, float64, string, byte](21)
	})
	is.Equal(NewEither5Arg1[int, bool, float64, string, byte](21), result1_1)

	result1_2 := either5Arg1.MapArg2(func(v bool) Either5[int, bool, float64, string, byte] {
		is.Fail("should not enter here")
		return NewEither5Arg2[int, bool, float64, string, byte](false)
	})
	is.Equal(either5Arg1, result1_2)

	result1_3 := either5Arg1.MapArg3(func(v float64) Either5[int, bool, float64, string, byte] {
		is.Fail("should not enter here")
		return NewEither5Arg3[int, bool, float64, string, byte](2.1)
	})
	is.Equal(either5Arg1, result1_3)

	result1_4 := either5Arg1.MapArg4(func(v string) Either5[int, bool, float64, string, byte] {
		is.Fail("should not enter here")
		return NewEither5Arg4[int, bool, float64, string, byte]("Bye")
	})
	is.Equal(either5Arg1, result1_4)

	result1_5 := either5Arg1.MapArg5(func(v byte) Either5[int, bool, float64, string, byte] {
		is.Fail("should not enter here")
		return NewEither5Arg5[int, bool, float64, string, byte](10)
	})
	is.Equal(either5Arg1, result1_5)

	result2_1 := either5Arg2.MapArg1(func(v int) Either5[int, bool, float64, string, byte] {
		is.Fail("should not enter here")
		return NewEither5Arg1[int, bool, float64, string, byte](21)
	})
	is.Equal(either5Arg2, result2_1)

	result2_2 := either5Arg2.MapArg2(func(v bool) Either5[int, bool, float64, string, byte] {
		is.Equal(true, v)
		return NewEither5Arg2[int, bool, float64, string, byte](false)
	})
	is.Equal(NewEither5Arg2[int, bool, float64, string, byte](false), result2_2)

	result2_3 := either5Arg2.MapArg3(func(v float64) Either5[int, bool, float64, string, byte] {
		is.Fail("should not enter here")
		return NewEither5Arg3[int, bool, float64, string, byte](2.1)
	})
	is.Equal(either5Arg2, result2_3)

	result2_4 := either5Arg2.MapArg4(func(v string) Either5[int, bool, float64, string, byte] {
		is.Fail("should not enter here")
		return NewEither5Arg4[int, bool, float64, string, byte]("Bye")
	})
	is.Equal(either5Arg2, result2_4)

	result2_5 := either5Arg2.MapArg5(func(v byte) Either5[int, bool, float64, string, byte] {
		is.Fail("should not enter here")
		return NewEither5Arg5[int, bool, float64, string, byte](10)
	})
	is.Equal(either5Arg2, result2_5)

	result3_1 := either5Arg3.MapArg1(func(v int) Either5[int, bool, float64, string, byte] {
		is.Fail("should not enter here")
		return NewEither5Arg3[int, bool, float64, string, byte](21)
	})
	is.Equal(either5Arg3, result3_1)

	result3_2 := either5Arg3.MapArg2(func(v bool) Either5[int, bool, float64, string, byte] {
		is.Fail("should not enter here")
		return NewEither5Arg2[int, bool, float64, string, byte](false)
	})
	is.Equal(either5Arg3, result3_2)

	result3_3 := either5Arg3.MapArg3(func(v float64) Either5[int, bool, float64, string, byte] {
		is.Equal(1.2, v)
		return NewEither5Arg3[int, bool, float64, string, byte](2.1)
	})
	is.Equal(NewEither5Arg3[int, bool, float64, string, byte](2.1), result3_3)

	result3_4 := either5Arg3.MapArg4(func(v string) Either5[int, bool, float64, string, byte] {
		is.Fail("should not enter here")
		return NewEither5Arg4[int, bool, float64, string, byte]("Bye")
	})
	is.Equal(either5Arg3, result3_4)

	result3_5 := either5Arg3.MapArg5(func(v byte) Either5[int, bool, float64, string, byte] {
		is.Fail("should not enter here")
		return NewEither5Arg5[int, bool, float64, string, byte](10)
	})
	is.Equal(either5Arg3, result3_5)

	result4_1 := either5Arg4.MapArg1(func(v int) Either5[int, bool, float64, string, byte] {
		is.Fail("should not enter here")
		return NewEither5Arg1[int, bool, float64, string, byte](21)
	})
	is.Equal(either5Arg4, result4_1)

	result4_2 := either5Arg4.MapArg2(func(v bool) Either5[int, bool, float64, string, byte] {
		is.Fail("should not enter here")
		return NewEither5Arg2[int, bool, float64, string, byte](false)
	})
	is.Equal(either5Arg4, result4_2)

	result4_3 := either5Arg4.MapArg3(func(v float64) Either5[int, bool, float64, string, byte] {
		is.Fail("should not enter here")
		return NewEither5Arg3[int, bool, float64, string, byte](2.1)
	})
	is.Equal(either5Arg4, result4_3)

	result4_4 := either5Arg4.MapArg4(func(v string) Either5[int, bool, float64, string, byte] {
		is.Equal("Hello", v)
		return NewEither5Arg4[int, bool, float64, string, byte]("Bye")
	})
	is.Equal(NewEither5Arg4[int, bool, float64, string, byte]("Bye"), result4_4)

	result4_5 := either5Arg4.MapArg5(func(v byte) Either5[int, bool, float64, string, byte] {
		is.Fail("should not enter here")
		return NewEither5Arg5[int, bool, float64, string, byte](10)
	})
	is.Equal(either5Arg4, result4_5)

	result5_1 := either5Arg5.MapArg1(func(v int) Either5[int, bool, float64, string, byte] {
		is.Fail("should not enter here")
		return NewEither5Arg1[int, bool, float64, string, byte](21)
	})
	is.Equal(either5Arg5, result5_1)

	result5_2 := either5Arg5.MapArg2(func(v bool) Either5[int, bool, float64, string, byte] {
		is.Fail("should not enter here")
		return NewEither5Arg2[int, bool, float64, string, byte](false)
	})
	is.Equal(either5Arg5, result5_2)

	result5_3 := either5Arg5.MapArg3(func(v float64) Either5[int, bool, float64, string, byte] {
		is.Fail("should not enter here")
		return NewEither5Arg3[int, bool, float64, string, byte](2.1)
	})
	is.Equal(either5Arg5, result5_3)

	result5_4 := either5Arg5.MapArg4(func(v string) Either5[int, bool, float64, string, byte] {
		is.Fail("should not enter here")
		return NewEither5Arg4[int, bool, float64, string, byte]("Bye")
	})
	is.Equal(either5Arg5, result5_4)

	result5_5 := either5Arg5.MapArg5(func(v byte) Either5[int, bool, float64, string, byte] {
		is.Equal(byte(5), v)
		return NewEither5Arg5[int, bool, float64, string, byte](10)
	})
	is.Equal(NewEither5Arg5[int, bool, float64, string, byte](10), result5_5)
}
