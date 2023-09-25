package mo

import (
	"fmt"
)

var err = fmt.Errorf("error")

func ExampleOk() {
	ok := Ok(42)
	result := ok.OrElse(1234)
	_err := ok.Error()

	fmt.Println(result, _err)
	// Output: 42 <nil>
}

func ExampleErr() {
	ko := Err[int](err)
	result := ko.OrElse(1234)
	_err := ko.Error()

	fmt.Println(result, _err)
	// Output: 1234 error
}

func ExampleErrf() {
	ko := Errf[int]("error")
	result := ko.OrElse(1234)
	_err := ko.Error()

	fmt.Println(result, _err)
	// Output: 1234 error
}

func ExampleTupleToResult() {
	randomFunc := func() (int, error) {
		return 42, err
	}

	value, _err := randomFunc()

	none := TupleToResult(value, _err)
	result := none.OrElse(1234)

	fmt.Println(result)
	// Output: 1234
}

func ExampleTry_ok() {
	randomFunc := func() (int, error) {
		return 42, nil
	}

	result := Try(randomFunc)
	value, err := result.Get()

	fmt.Println(value)
	fmt.Println(err)
	// Output:
	// 42
	// <nil>
}

func ExampleTry_err() {
	randomFunc := func() (int, error) {
		return 42, err
	}

	result := Try(randomFunc)
	value, err := result.Get()

	fmt.Println(value)
	fmt.Println(err)
	// Output:
	// 0
	// error
}

func ExampleResult_ok() {
	ok := Ok(42)
	result := ok.OrElse(1234)
	_err := ok.Error()

	fmt.Println(result, _err)
	// Output: 42 <nil>
}

func ExampleResult_err() {
	ko := Err[int](err)
	result := ko.OrElse(1234)
	_err := ko.Error()

	fmt.Println(result, _err)
	// Output: 1234 error
}

func ExampleResult_IsOk_ok() {
	ok := Ok(42)
	result := ok.IsOk()

	fmt.Println(result)
	// Output: true
}

func ExampleResult_IsOk_err() {
	ko := Err[int](err)
	result := ko.IsOk()

	fmt.Println(result)
	// Output: false
}

func ExampleResult_IsError_ok() {
	ok := Ok(42)
	result := ok.IsError()

	fmt.Println(result)
	// Output: false
}

func ExampleResult_IsError_err() {
	ko := Err[int](err)
	result := ko.IsError()

	fmt.Println(result)
	// Output: true
}

func ExampleResult_Error_ok() {
	ok := Ok(42)
	result := ok.Error()

	fmt.Println(result)
	// Output: <nil>
}

func ExampleResult_Error_err() {
	ko := Err[int](err)
	result := ko.Error()

	fmt.Println(result)
	// Output: error
}

func ExampleResult_Get_ok() {
	ok := Ok(42)
	result, err := ok.Get()

	fmt.Println(result)
	fmt.Println(err)
	// Output:
	// 42
	// <nil>
}

func ExampleResult_Get_err() {
	ko := Err[int](err)
	result, err := ko.Get()

	fmt.Println(result)
	fmt.Println(err)
	// Output:
	// 0
	// error
}

func ExampleResult_MustGet_ok() {
	ok := Ok(42)
	result := ok.MustGet()

	fmt.Println(result)
	// Output: 42
}

// func ExampleResult_MustGet_err() {
// 	ko := Err[int](err)
// 	result := ko.MustGet()

// 	fmt.Println(result)
// 	// Output: panics
// }

func ExampleResult_OrElse_ok() {
	ok := Ok(42)
	result := ok.OrElse(1234)

	fmt.Println(result)
	// Output: 42
}

func ExampleResult_OrElse_err() {
	ko := Err[int](err)
	result := ko.OrElse(1234)

	fmt.Println(result)
	// Output: 1234
}

func ExampleResult_OrEmpty_ok() {
	ok := Ok(42)
	result := ok.OrEmpty()

	fmt.Println(result)
	// Output: 42
}

func ExampleResult_OrEmpty_err() {
	ko := Err[int](err)
	result := ko.OrEmpty()

	fmt.Println(result)
	// Output: 0
}

func ExampleResult_ToEither_ok() {
	ok := Ok(42)
	either := ok.ToEither()

	err, isLeft := either.Left()
	value, isRight := either.Right()

	fmt.Println(isLeft, isRight)
	fmt.Println(err)
	fmt.Println(value)
	// Output:
	// false true
	// <nil>
	// 42
}

func ExampleResult_ToEither_err() {
	ko := Err[int](err)
	either := ko.ToEither()

	err, isLeft := either.Left()
	value, isRight := either.Right()

	fmt.Println(isLeft, isRight)
	fmt.Println(err)
	fmt.Println(value)
	// Output:
	// true false
	// error
	// 0
}

func ExampleResult_Match_ok() {
	ok := Ok(42)
	result := ok.Match(
		func(i int) (int, error) {
			return i * 2, nil
		},
		func(err error) (int, error) {
			return 21, nil
		},
	)

	fmt.Println(result.IsError(), result.OrEmpty(), result.Error())
	// Output: false 84 <nil>
}

func ExampleResult_Match_err() {
	ko := Err[int](err)
	result := ko.Match(
		func(i int) (int, error) {
			return i * 2, nil
		},
		func(err error) (int, error) {
			return 21, nil
		},
	)

	fmt.Println(result.IsError(), result.OrEmpty(), result.Error())
	// Output: false 21 <nil>
}

func ExampleResult_Map_ok() {
	ok := Ok(42)
	result := ok.Map(
		func(i int) (int, error) {
			return i * 2, nil
		},
	)

	fmt.Println(result.IsError(), result.OrEmpty(), result.Error())
	// Output: false 84 <nil>
}

func ExampleResult_Map_err() {
	ko := Err[int](err)
	result := ko.Map(
		func(i int) (int, error) {
			return i * 2, nil
		},
	)

	fmt.Println(result.IsError(), result.OrEmpty(), result.Error())
	// Output: true 0 error
}

func ExampleResult_MapErr_ok() {
	ok := Ok(42)
	result := ok.MapErr(
		func(_err error) (int, error) {
			return 1234, nil
		},
	)

	fmt.Println(result.IsError(), result.OrEmpty(), result.Error())
	// Output: false 42 <nil>
}

func ExampleResult_MapErr_err() {
	ko := Err[int](err)
	result := ko.MapErr(
		func(_err error) (int, error) {
			return 1234, nil
		},
	)

	fmt.Println(result.IsError(), result.OrEmpty(), result.Error())
	// Output: false 1234 <nil>
}

func ExampleResult_FlatMap_ok() {
	ok := Ok(42)
	result := ok.FlatMap(
		func(i int) Result[int] {
			return Ok(1234)
		},
	)

	fmt.Println(result.IsError(), result.OrEmpty(), result.Error())
	// Output: false 1234 <nil>
}

func ExampleResult_FlatMap_err() {
	ko := Err[int](err)
	result := ko.FlatMap(
		func(i int) Result[int] {
			return Ok(1234)
		},
	)

	fmt.Println(result.IsError(), result.OrEmpty(), result.Error())
	// Output: true 0 error
}
