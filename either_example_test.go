package mo

import "fmt"

func ExampleLeft() {
	left := Left[string, int]("hello")
	result1 := left.LeftOrElse("world")
	result2 := left.RightOrElse(1234)

	fmt.Println(result1, result2)
	// Output: hello 1234
}

func ExampleRight() {
	right := Right[string, int](42)
	result1 := right.LeftOrElse("world")
	result2 := right.RightOrElse(1234)

	fmt.Println(result1, result2)
	// Output: world 42
}

func ExampleEither_Unpack_left() {
	either := Left[string, int]("42")
	left, right := either.Unpack()

	fmt.Println(left, right)
	// Output: 42 0
}

func ExampleEither_Unpack_right() {
	either := Right[string, int](42)
	left, right := either.Unpack()

	fmt.Println(left, right)
	// Output: 42
}

func ExampleEither_IsLeft_left() {
	left := Left[string, int]("hello")
	result := left.IsLeft()

	fmt.Println(result)
	// Output: true
}

func ExampleEither_IsLeft_right() {
	right := Right[string, int](42)
	result := right.IsLeft()

	fmt.Println(result)
	// Output: false
}

func ExampleEither_IsRight_left() {
	left := Left[string, int]("hello")
	result := left.IsRight()

	fmt.Println(result)
	// Output: false
}

func ExampleEither_IsRight_right() {
	right := Right[string, int](42)
	result := right.IsRight()

	fmt.Println(result)
	// Output: true
}

func ExampleEither_Left_left() {
	left := Left[string, int]("hello")
	result, ok := left.Left()

	fmt.Println(result)
	fmt.Println(ok)
	// Output:
	// hello
	// true
}

func ExampleEither_Left_right() {
	right := Right[string, int](42)
	result, ok := right.Left()

	fmt.Println(result)
	fmt.Println(ok)
	// Output:
	// false
}

func ExampleEither_Right_left() {
	left := Left[string, int]("hello")
	result, ok := left.Right()

	fmt.Println(result)
	fmt.Println(ok)
	// Output:
	// 0
	// false
}

func ExampleEither_Right_right() {
	right := Right[string, int](42)
	result, ok := right.Right()

	fmt.Println(result)
	fmt.Println(ok)
	// Output:
	// 42
	// true
}

func ExampleEither_MustLeft_left() {
	left := Left[string, int]("hello")
	result := left.MustLeft()

	fmt.Println(result)
	// Output: hello
}

// func ExampleEither_MustLeft_right() {
// 	right := Right[string, int](42)
// 	result := right.MustLeft()

// 	fmt.Println(result)
// 	// Output: panics
// }

// func ExampleEither_MustRight_left() {
// 	left := Left[string, int]("hello")
// 	result := left.MustRight()

// 	fmt.Println(result)
// 	// Output: panics
// }

func ExampleEither_MustRight_right() {
	right := Right[string, int](42)
	result := right.MustRight()

	fmt.Println(result)
	// Output: 42
}

func ExampleEither_LeftOrElse_left() {
	left := Left[string, int]("hello")
	result := left.LeftOrElse("world")

	fmt.Println(result)
	// Output: hello
}

func ExampleEither_LeftOrElse_right() {
	right := Right[string, int](42)
	result := right.LeftOrElse("world")

	fmt.Println(result)
	// Output: world
}

func ExampleEither_RightOrElse_left() {
	left := Left[string, int]("hello")
	result := left.RightOrElse(1234)

	fmt.Println(result)
	// Output: 1234
}

func ExampleEither_RightOrElse_right() {
	right := Right[string, int](42)
	result := right.RightOrElse(1234)

	fmt.Println(result)
	// Output: 42
}

func ExampleEither_LeftOrEmpty_left() {
	left := Left[string, int]("hello")
	result := left.LeftOrEmpty()

	fmt.Println(result)
	// Output: hello
}

func ExampleEither_LeftOrEmpty_right() {
	right := Right[string, int](42)
	result := right.LeftOrEmpty()

	fmt.Println(result)
	// Output:
}

func ExampleEither_RightOrEmpty_left() {
	left := Left[string, int]("hello")
	result := left.RightOrEmpty()

	fmt.Println(result)
	// Output: 0
}

func ExampleEither_RightOrEmpty_right() {
	right := Right[string, int](42)
	result := right.RightOrEmpty()

	fmt.Println(result)
	// Output: 42
}

func ExampleEither_Swap_left() {
	left := Left[string, int]("hello")
	right := left.Swap()
	result1, ok1 := right.Left()
	result2, ok2 := right.Right()

	fmt.Println(result1)
	fmt.Println(ok1)
	fmt.Println(result2)
	fmt.Println(ok2)
	// Output:
	// 0
	// false
	// hello
	// true
}

func ExampleEither_Swap_right() {
	right := Right[string, int](42)
	left := right.Swap()
	result1, ok1 := left.Left()
	result2, ok2 := left.Right()

	fmt.Println(result1)
	fmt.Println(ok1)
	fmt.Println(result2)
	fmt.Println(ok2)
	// 42
	// true
	//
	// false
}

func ExampleEither_Match_left() {
	left := Left[string, int]("hello")
	result := left.Match(
		func(s string) Either[string, int] {
			return Right[string, int](1234)
		},
		func(i int) Either[string, int] {
			return Right[string, int](i * 42)
		},
	)
	result1, ok1 := result.Left()
	result2, ok2 := result.Right()

	fmt.Println(result1)
	fmt.Println(ok1)
	fmt.Println(result2)
	fmt.Println(ok2)
	// Output:
	// false
	// 1234
	// true
}
func ExampleEither_Match_right() {
	right := Right[string, int](42)
	result := right.Match(
		func(s string) Either[string, int] {
			return Left[string, int]("world")
		},
		func(i int) Either[string, int] {
			return Left[string, int]("foobar")
		},
	)
	result1, ok1 := result.Left()
	result2, ok2 := result.Right()

	fmt.Println(result1)
	fmt.Println(ok1)
	fmt.Println(result2)
	fmt.Println(ok2)
	// Output:
	// foobar
	// true
	// 0
	// false
}

func ExampleEither_MapLeft_left() {
	left := Left[string, int]("hello")
	result := left.MapLeft(
		func(s string) Either[string, int] {
			return Right[string, int](1234)
		},
	)
	result1, ok1 := result.Left()
	result2, ok2 := result.Right()

	fmt.Println(result1)
	fmt.Println(ok1)
	fmt.Println(result2)
	fmt.Println(ok2)
	// Output:
	// false
	// 1234
	// true
}
func ExampleEither_MapLeft_right() {
	right := Right[string, int](42)
	result := right.MapLeft(
		func(s string) Either[string, int] {
			return Left[string, int]("world")
		},
	)
	result1, ok1 := result.Left()
	result2, ok2 := result.Right()

	fmt.Println(result1)
	fmt.Println(ok1)
	fmt.Println(result2)
	fmt.Println(ok2)
	// Output:
	// false
	// 42
	// true
}

func ExampleEither_MapRight_left() {
	left := Left[string, int]("hello")
	result := left.MapRight(
		func(i int) Either[string, int] {
			return Right[string, int](1234)
		},
	)
	result1, ok1 := result.Left()
	result2, ok2 := result.Right()

	fmt.Println(result1)
	fmt.Println(ok1)
	fmt.Println(result2)
	fmt.Println(ok2)
	// Output:
	// hello
	// true
	// 0
	// false
}
func ExampleEither_MapRight_right() {
	right := Right[string, int](42)
	result := right.MapRight(
		func(i int) Either[string, int] {
			return Right[string, int](1234)
		},
	)
	result1, ok1 := result.Left()
	result2, ok2 := result.Right()

	fmt.Println(result1)
	fmt.Println(ok1)
	fmt.Println(result2)
	fmt.Println(ok2)
	// Output:
	// false
	// 1234
	// true
}
