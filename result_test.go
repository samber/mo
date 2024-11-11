package mo

import (
	"encoding/json"
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResultOk(t *testing.T) {
	is := assert.New(t)

	is.Equal(Result[int]{value: 42, isErr: false, err: nil}, Ok(42))
}

func TestResultErr(t *testing.T) {
	is := assert.New(t)

	is.Equal(Result[int]{value: 0, isErr: true, err: assert.AnError}, Err[int](assert.AnError))
}

func TestResultErrf(t *testing.T) {
	is := assert.New(t)

	is.Equal(Result[int]{value: 0, isErr: true, err: assert.AnError}, Errf[int](assert.AnError.Error())) //nolint:govet
}

func TestResultTupleToResult(t *testing.T) {
	is := assert.New(t)

	is.Equal(Result[int]{value: 0, isErr: true, err: assert.AnError}, TupleToResult(42, assert.AnError))
}

func TestResultTry(t *testing.T) {
	is := assert.New(t)

	is.Equal(Result[int]{value: 42, isErr: false, err: nil}, Try(func() (int, error) {
		return 42, nil
	}))
	is.Equal(Result[int]{value: 0, isErr: true, err: assert.AnError}, Try(func() (int, error) {
		return 42, assert.AnError
	}))
}

func TestResultIsOk(t *testing.T) {
	is := assert.New(t)

	is.True(Ok(42).IsOk())
	is.False(Err[int](assert.AnError).IsOk())
}

func TestResultIsError(t *testing.T) {
	is := assert.New(t)

	is.False(Ok(42).IsError())
	is.True(Err[int](assert.AnError).IsError())
}

func TestResultError(t *testing.T) {
	is := assert.New(t)

	is.Nil(Ok(42).Error())
	is.NotNil(Err[int](assert.AnError).Error())
	is.Equal(assert.AnError, Err[int](assert.AnError).Error())
}

func TestResultGet(t *testing.T) {
	is := assert.New(t)

	v1, err1 := Ok(42).Get()
	v2, err2 := Err[int](assert.AnError).Get()

	is.Equal(42, v1)
	is.Nil(err1)
	is.Error(assert.AnError, err1)

	is.Equal(0, v2)
	is.NotNil(err2)
	is.Error(assert.AnError, err2)
}

func TestResultMustGet(t *testing.T) {
	is := assert.New(t)

	is.NotPanics(func() {
		Ok(42).MustGet()
	})
	is.Panics(func() {
		Err[int](assert.AnError).MustGet()
	})

	is.Equal(42, Ok(42).MustGet())
}

func TestResultOrElse(t *testing.T) {
	is := assert.New(t)

	is.Equal(42, Ok(42).OrElse(21))
	is.Equal(21, Err[int](assert.AnError).OrElse(21))
}

func TestResultOrEmpty(t *testing.T) {
	is := assert.New(t)

	is.Equal(42, Ok(42).OrEmpty())
	is.Equal(0, Err[int](assert.AnError).OrEmpty())
}

func TestResultToEither(t *testing.T) {
	is := assert.New(t)

	right, ok1 := Ok(42).ToEither().Right()
	left, ok2 := Err[int](assert.AnError).ToEither().Left()

	is.Equal(42, right)
	is.True(ok1)
	is.Equal(assert.AnError, left)
	is.True(ok2)
}

func TestResultForEach(t *testing.T) {
	is := assert.New(t)

	Err[int](assert.AnError).ForEach(func(i int) {
		is.Fail("should not enter here")
	})

	Ok(42).ForEach(func(i int) {
		is.Equal(42, i)
	})
}

func TestResultMatch(t *testing.T) {
	is := assert.New(t)

	opt1 := Ok(21).Match(
		func(i int) (int, error) {
			is.Equal(21, i)
			return i * 2, nil
		},
		func(err error) (int, error) {
			is.Fail("should not enter here")
			return 0, err
		},
	)
	opt2 := Err[int](assert.AnError).Match(
		func(i int) (int, error) {
			is.Fail("should not enter here")
			return i * 2, nil
		},
		func(err error) (int, error) {
			is.Equal(assert.AnError, err)
			return 0, err
		},
	)

	is.Equal(Result[int]{value: 42, isErr: false, err: nil}, opt1)
	is.Equal(Result[int]{value: 0, isErr: true, err: assert.AnError}, opt2)
}

func TestResultMap(t *testing.T) {
	is := assert.New(t)

	opt1 := Ok(21).Map(func(i int) (int, error) {
		return i * 2, nil
	})
	opt2 := Err[int](assert.AnError).Map(func(i int) (int, error) {
		is.Fail("should not be called")
		return 42, nil
	})

	is.Equal(Result[int]{value: 42, isErr: false, err: nil}, opt1)
	is.Equal(Result[int]{value: 0, isErr: true, err: assert.AnError}, opt2)
}

func TestResultMapErr(t *testing.T) {
	is := assert.New(t)

	opt1 := Ok(21).MapErr(func(err error) (int, error) {
		is.Fail("should not be called")
		return 42, nil
	})
	opt2 := Err[int](assert.AnError).MapErr(func(err error) (int, error) {
		return 42, nil
	})

	is.Equal(Result[int]{value: 21, isErr: false, err: nil}, opt1)
	is.Equal(Result[int]{value: 42, isErr: false, err: nil}, opt2)
}

func TestResultFlatMap(t *testing.T) {
	is := assert.New(t)

	opt1 := Ok(21).FlatMap(func(i int) Result[int] {
		return Ok(42)
	})
	opt2 := Err[int](assert.AnError).FlatMap(func(i int) Result[int] {
		is.Fail("should not be called")
		return Ok(42)
	})

	is.Equal(Result[int]{value: 42, isErr: false, err: nil}, opt1)
	is.Equal(Result[int]{value: 0, isErr: true, err: assert.AnError}, opt2)
}

func TestResultMarshalJSON(t *testing.T) {
	is := assert.New(t)

	result1 := Ok("foo")
	result2 := Err[string](fmt.Errorf("an error"))
	result3 := Ok("")

	value, err := result1.MarshalJSON()
	is.NoError(err)
	is.Equal(`{"result":"foo"}`, string(value))

	value, err = result2.MarshalJSON()
	is.NoError(err)
	is.Equal(`{"error":{"message":"an error"}}`, string(value))

	value, err = result3.MarshalJSON()
	is.NoError(err)
	is.Equal(`{"result":""}`, string(value))

	type testStruct struct {
		Field Result[string]
	}

	resultInStruct := testStruct{
		Field: result1,
	}
	var marshalled []byte
	marshalled, err = json.Marshal(resultInStruct)
	is.NoError(err)
	is.Equal(`{"Field":{"result":"foo"}}`, string(marshalled))
}

func TestResultUnmarshalJSON(t *testing.T) {
	is := assert.New(t)

	result1 := Ok("foo")
	result2 := Err[string](fmt.Errorf("an error"))
	result3 := Ok("")

	err := result1.UnmarshalJSON([]byte(`{"result":"foo"}`))
	is.NoError(err)
	is.Equal(Ok("foo"), result1)

	var res Result[string]
	err = json.Unmarshal([]byte(`{"result":"foo"}`), &res)
	is.NoError(err)
	is.Equal(res, result1)

	err = result2.UnmarshalJSON([]byte(`{"error":{"message":"an error"}}`))
	is.NoError(err)
	is.Equal(Err[string](fmt.Errorf("an error")), result2)

	err = result3.UnmarshalJSON([]byte(`{"result":""}`))
	is.NoError(err)
	is.Equal(Ok(""), result3)

	type testStruct struct {
		Field Result[string]
	}

	unmarshal := testStruct{}
	err = json.Unmarshal([]byte(`{"Field":{"result":"foo"}}`), &unmarshal)
	is.NoError(err)
	is.Equal(testStruct{Field: Ok("foo")}, unmarshal)

	unmarshal = testStruct{}
	err = json.Unmarshal([]byte(`{"Field":{"error":{"message":"an error"}}}`), &unmarshal)
	is.NoError(err)
	is.Equal(testStruct{Field: Err[string](fmt.Errorf("an error"))}, unmarshal)

	unmarshal = testStruct{}
	err = json.Unmarshal([]byte(`{}`), &unmarshal)
	is.NoError(err)
	is.Equal(testStruct{Field: Ok("")}, unmarshal)

	// Both result and error are set; unmarshal to Err
	unmarshal = testStruct{}
	err = json.Unmarshal([]byte(`{"Field":{"result":"foo","error":{"message":"an error"}}}`), &unmarshal)
	is.NoError(err)
	is.Equal(testStruct{Field: Err[string](fmt.Errorf("an error"))}, unmarshal)

	// Bad structure for error; cannot unmarshal
	unmarshal = testStruct{}
	err = json.Unmarshal([]byte(`{"Field":{"result":"foo","error":true}}`), &unmarshal)
	is.Error(err)

	unmarshal = testStruct{}
	err = json.Unmarshal([]byte(`{"Field": "}`), &unmarshal)
	is.Error(err)
}

// TestResultFoldSuccess tests the Fold method with a successful result.
func TestResultFoldSuccess(t *testing.T) {
	is := assert.New(t)
	result := Result[int]{value: 42, isErr: false, err: nil}

	successFunc := func(value int) string {
		return fmt.Sprintf("Success: %v", value)
	}
	failureFunc := func(err error) string {
		return fmt.Sprintf("Failure: %v", err)
	}

	folded := Fold[error, int, string](result, successFunc, failureFunc)
	expected := "Success: 42"

	is.Equal(expected, folded)
}

// TestResultFoldFailure tests the Fold method with a failure result.
func TestResultFoldFailure(t *testing.T) {
	err := errors.New("result error")
	is := assert.New(t)

	result := Result[int]{value: 0, isErr: true, err: err}

	successFunc := func(value int) string {
		return fmt.Sprintf("Success: %v", value)
	}
	failureFunc := func(err error) string {
		return fmt.Sprintf("Failure: %v", err)
	}

	folded := Fold[error, int, string](result, successFunc, failureFunc)
	expected := fmt.Sprintf("Failure: %v", err)

	is.Equal(expected, folded)
}
