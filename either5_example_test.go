package mo

import "fmt"

func ExampleNewEither5Arg1() {
	either5Arg1 := NewEither5Arg1[int, bool, float64, string, byte](42)
	result1 := either5Arg1.Arg1OrElse(21)
	result2 := either5Arg1.Arg4OrElse("Bye")

	fmt.Println(result1, result2)
	// Output: 42 Bye
}

func ExampleEither5_IsArg1() {
	either5Arg1 := NewEither5Arg1[int, bool, float64, string, byte](42)
	result1 := either5Arg1.IsArg1()
	result2 := either5Arg1.IsArg4()

	fmt.Println(result1, result2)
	// Output: true false
}

func ExampleEither5_Arg1() {
	either5Arg1 := NewEither5Arg1[int, bool, float64, string, byte](42)
	result1, ok1 := either5Arg1.Arg1()
	result2, ok2 := either5Arg1.Arg3()

	fmt.Println(result1, ok1)
	fmt.Println(result2, ok2)
	// Output:
	// 42 true
	// 0 false
}

func ExampleEither5_MustArg1() {
	either5Arg1 := NewEither5Arg1[int, bool, float64, string, byte](42)

	// result = either5Arg1.MustArg4()
	// Panics

	result := either5Arg1.MustArg1()
	fmt.Println(result)
	// Output: 42
}

func ExampleEither5_Unpack() {
	either5 := NewEither5Arg1[int, bool, float64, string, byte](42)
	a, b, c, d, e := either5.Unpack()

	fmt.Println(a, b, c, d, e)
	// Output: 42 false 0  0
}

func ExampleEither5_Arg1OrElse() {
	either5Arg1 := NewEither5Arg1[int, bool, float64, string, byte](42)
	result1 := either5Arg1.Arg1OrElse(21)
	result2 := either5Arg1.Arg4OrElse("Bye")

	fmt.Println(result1, result2)
	// Output: 42 Bye
}

func ExampleEither5_Arg1OrEmpty() {
	either5Arg1 := NewEither5Arg1[int, bool, float64, string, byte](42)
	result1 := either5Arg1.Arg1OrEmpty()
	result2 := either5Arg1.Arg2OrEmpty()

	fmt.Println(result1, result2)
	// Output: 42 false
}

func ExampleEither5_Match() {
	either5Arg1 := NewEither5Arg1[int, bool, float64, string, byte](42)

	result1 := either5Arg1.Match(func(v int) Either5[int, bool, float64, string, byte] {
		return NewEither5Arg1[int, bool, float64, string, byte](21)
	}, func(v bool) Either5[int, bool, float64, string, byte] {
		return NewEither5Arg2[int, bool, float64, string, byte](false)
	}, func(v float64) Either5[int, bool, float64, string, byte] {
		return NewEither5Arg3[int, bool, float64, string, byte](2.1)
	}, func(v string) Either5[int, bool, float64, string, byte] {
		return NewEither5Arg4[int, bool, float64, string, byte]("Bye")
	}, func(v byte) Either5[int, bool, float64, string, byte] {
		return NewEither5Arg5[int, bool, float64, string, byte](10)
	})

	fmt.Println(result1.MustArg1())
	// Output: 21
}

func ExampleEither5_MapArg1() {
	either5Arg1 := NewEither5Arg1[int, bool, float64, string, byte](42)

	result1 := either5Arg1.MapArg1(
		func(v int) Either5[int, bool, float64, string, byte] {
			return NewEither5Arg1[int, bool, float64, string, byte](21)
		},
	)

	result2 := either5Arg1.MapArg4(
		func(v string) Either5[int, bool, float64, string, byte] {
			return NewEither5Arg4[int, bool, float64, string, byte]("Bye")
		},
	)

	fmt.Println(result1.MustArg1(), result2.MustArg1())
	// Output: 21 42
}
