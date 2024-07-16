package mo

import (
	"database/sql"
	"encoding/json"
	"fmt"
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

func TestTupleToOption(t *testing.T) {
	is := assert.New(t)

	cb := func(v int, ok bool) func() (int, bool) {
		return func() (int, bool) {
			return v, ok
		}
	}

	is.Equal(Option[int]{isPresent: false}, TupleToOption(cb(42, false)()))
	is.Equal(Option[int]{isPresent: true, value: 42}, TupleToOption(cb(42, true)()))
}

func TestOptionEmptyableToOption(t *testing.T) {
	is := assert.New(t)

	is.Equal(Option[error]{isPresent: false}, EmptyableToOption[error](nil))
	is.Equal(Option[error]{isPresent: true, value: assert.AnError}, EmptyableToOption(assert.AnError))

	is.Equal(Option[int]{isPresent: false}, EmptyableToOption(0))
	is.Equal(Option[int]{isPresent: true, value: 42}, EmptyableToOption(42))
}

func TestOptionPointerToOption(t *testing.T) {
	is := assert.New(t)

	is.Equal(Option[error]{isPresent: false}, PointerToOption[error](nil))
	is.Equal(Option[error]{isPresent: true, value: assert.AnError}, PointerToOption(&assert.AnError))

	zero := 0
	fortyTwo := 42
	is.Equal(Option[int]{isPresent: true, value: 0}, PointerToOption(&zero))
	is.Equal(Option[int]{isPresent: true, value: 42}, PointerToOption(&fortyTwo))
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

func TestOptionToPointer(t *testing.T) {
	is := assert.New(t)

	p := Some(42).ToPointer()
	is.NotNil(p)
	is.Equal(42, *p)

	is.Nil(None[int]().ToPointer())
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
	option3 := Some("")

	value, err := option1.MarshalJSON()
	is.NoError(err)
	is.Equal(`"foo"`, string(value))

	value, err = option2.MarshalJSON()
	is.NoError(err)
	is.Equal(`null`, string(value))

	value, err = option3.MarshalJSON()
	is.NoError(err)
	is.Equal(`""`, string(value))

	type testStruct struct {
		Field Option[string]
	}

	optionInStruct := testStruct{
		Field: option1,
	}
	var marshalled []byte
	marshalled, err = json.Marshal(optionInStruct)
	is.NoError(err)
	is.Equal(`{"Field":"foo"}`, string(marshalled))
}

func TestOptionUnmarshalJSON(t *testing.T) {
	is := assert.New(t)

	option1 := Some("foo")
	option2 := None[string]()

	err := option1.UnmarshalJSON([]byte(`"foo"`))
	is.NoError(err)
	is.Equal(Some("foo"), option1)

	var res Option[string]
	err = json.Unmarshal([]byte(`"foo"`), &res)
	is.NoError(err)
	is.Equal(res, option1)

	err = option2.UnmarshalJSON([]byte(`null`))
	is.NoError(err)
	is.Equal(None[string](), option2)

	type testStruct struct {
		Field Option[string]
	}

	unmarshal := testStruct{}
	err = json.Unmarshal([]byte(`{"Field": "foo"}`), &unmarshal)
	is.NoError(err)
	is.Equal(testStruct{
		Field: Some("foo"),
	}, unmarshal)

	unmarshal = testStruct{}
	err = json.Unmarshal([]byte(`{"Field": null}`), &unmarshal)
	is.NoError(err)
	is.Equal(testStruct{Field: None[string]()}, unmarshal)

	unmarshal = testStruct{}
	err = json.Unmarshal([]byte(`{}`), &unmarshal)
	is.NoError(err)
	is.Equal(testStruct{Field: None[string]()}, unmarshal)

	unmarshal = testStruct{}
	err = json.Unmarshal([]byte(`{"Field": ""}`), &unmarshal)
	is.NoError(err)
	is.Equal(testStruct{Field: Some("")}, unmarshal)

	unmarshal = testStruct{}
	err = json.Unmarshal([]byte(`{"Field": "}`), &unmarshal)
	is.Error(err)
}

func TestOptionMarshalText(t *testing.T) {
	is := assert.New(t)

	bytes1, err1 := Some(42).MarshalText()
	bytes2, err2 := None[int]().MarshalText()
	bytes3, err3 := Some("42").MarshalText()

	is.Equal([]byte("42"), bytes1)
	is.Nil(err1)
	is.Equal([]byte("null"), bytes2)
	is.Nil(err2)
	is.Equal([]byte("\"42\""), bytes3)
	is.Nil(err3)
}

func TestOptionUnmarshalText(t *testing.T) {
	is := assert.New(t)

	option1 := Option[int]{}
	option2 := Option[int]{}
	option3 := Option[string]{}

	err1 := option1.UnmarshalText([]byte("null"))
	err2 := option2.UnmarshalText([]byte("42"))
	err3 := option3.UnmarshalText([]byte("\"42\""))

	is.Equal(None[int](), option1)
	is.Nil(err1)
	is.Equal(Some[int](42), option2)
	is.Nil(err2)
	is.Equal(Some[string]("42"), option3)
	is.Nil(err3)
}

func TestOptionMarshalBinary(t *testing.T) {
	is := assert.New(t)

	binary1, err1 := Some(42).MarshalBinary()
	binary2, err2 := None[int]().MarshalBinary()
	binary3, err3 := Some("42").MarshalBinary()

	is.Equal([]byte{1, 0x3, 0x4, 0x0, 0x54}, binary1)
	is.Nil(err1)
	is.Equal([]byte{0}, binary2)
	is.Nil(err2)
	is.Equal([]byte{1, 0x5, 0xc, 0x0, 0x2, 0x34, 0x32}, binary3)
	is.Nil(err3)
}

func TestOptionUnmarshalBinary(t *testing.T) {
	is := assert.New(t)

	option1 := Option[int]{}
	option2 := Option[int]{}
	option3 := Option[string]{}

	err1 := option1.UnmarshalBinary([]byte{0})
	err2 := option2.UnmarshalBinary([]byte{1, 0x3, 0x4, 0x0, 0x54})
	err3 := option3.UnmarshalBinary([]byte{1, 0x5, 0xc, 0x0, 0x2, 0x34, 0x32})

	is.Equal(None[int](), option1)
	is.Nil(err1)
	is.Equal(Some[int](42), option2)
	is.Nil(err2)
	is.Equal(Some[string]("42"), option3)
	is.Nil(err3)
}

func TestOptionGobEncode(t *testing.T) {
	is := assert.New(t)

	binary1, err1 := Some(42).GobEncode()
	binary2, err2 := None[int]().GobEncode()
	binary3, err3 := Some("42").GobEncode()

	is.Equal([]byte{1, 0x3, 0x4, 0x0, 0x54}, binary1)
	is.Nil(err1)
	is.Equal([]byte{0}, binary2)
	is.Nil(err2)
	is.Equal([]byte{1, 0x5, 0xc, 0x0, 0x2, 0x34, 0x32}, binary3)
	is.Nil(err3)
}

func TestOptionGobDecode(t *testing.T) {
	is := assert.New(t)

	option1 := Option[int]{}
	option2 := Option[int]{}
	option3 := Option[string]{}

	err1 := option1.GobDecode([]byte{0})
	err2 := option2.GobDecode([]byte{1, 0x3, 0x4, 0x0, 0x54})
	err3 := option3.GobDecode([]byte{1, 0x5, 0xc, 0x0, 0x2, 0x34, 0x32})

	is.Equal(None[int](), option1)
	is.Nil(err1)
	is.Equal(Some[int](42), option2)
	is.Nil(err2)
	is.Equal(Some[string]("42"), option3)
	is.Nil(err3)
}

func TestOptionScan(t *testing.T) {
	is := assert.New(t)

	option1 := Some("foo")
	option2 := None[string]()

	nullString1 := sql.NullString{String: "foo", Valid: true}
	nullString2 := sql.NullString{String: "", Valid: false}

	res1Exp, err1Exp := nullString1.Value()
	res1, err1 := option1.Value()

	res2Exp, err2Exp := nullString2.Value()
	res2, err2 := option2.Value()

	is.Equal(res1Exp, res1)
	is.Equal(err1Exp, err1)
	is.Equal(res2Exp, res2)
	is.Equal(err2Exp, err2)
}

func TestOptionScanWithPossibleConvert(t *testing.T) {
	is := assert.New(t)

	// As passed by the sql package in some cases, src is a []byte.
	// https://github.com/golang/go/blob/071b8d51c1a70fa6b12f0bed2e93370e193333fd/src/database/sql/convert.go#L396
	src1 := []byte{65, 66, 67}
	dest1 := None[string]()
	src2 := int32(32)
	dest2 := None[int]()

	err1 := dest1.Scan(src1)
	err2 := dest2.Scan(src2)

	is.Nil(err1)
	is.Equal(Some("ABC"), dest1)
	is.Nil(err2)
	is.Equal(Some(32), dest2)
}

func TestOptionValue(t *testing.T) {
	is := assert.New(t)

	option1 := Option[string]{}
	option2 := Option[string]{}

	nullString1, _ := sql.NullString{String: "foo", Valid: true}.Value()
	nullString2, _ := sql.NullString{String: "", Valid: false}.Value()

	err1 := option1.Scan(nullString1)
	err2 := option2.Scan(nullString2)

	is.EqualValues(Some("foo"), option1)
	is.Nil(err1)
	is.EqualValues(None[string](), option2)
	is.Nil(err2)
}

func TestOptionValueWithPossibleConvert(t *testing.T) {
	is := assert.New(t)

	opt := Some(uint32(42))
	expected := int64(42)

	value, err := opt.Value()
	is.Nil(err)
	is.Equal(expected, value)
}

type SomeScanner struct {
	Cool bool
	Some int
}

func (ss *SomeScanner) Scan(src any) error {
	val, ok := src.(string)
	if !ok {
		return fmt.Errorf("cannot scan - src is not a string")
	}

	var unmarshalled SomeScanner
	if err := json.Unmarshal([]byte(val), &unmarshalled); err != nil {
		return fmt.Errorf("failed to unmarshal json: %w", err)
	}

	*ss = unmarshalled
	return nil
}

// If T is a sql.Scanner, make use of that.
func TestOptionScanner(t *testing.T) {
	is := assert.New(t)

	jsonString := `{"cool": true, "some": 123}`
	nullString, _ := sql.NullString{}.Value()

	var someScanner Option[SomeScanner]
	var noneScanner Option[SomeScanner]

	err1 := someScanner.Scan(jsonString)
	err2 := noneScanner.Scan(nullString)

	is.NoError(err1)
	is.EqualValues(Some(SomeScanner{Cool: true, Some: 123}), someScanner)
	is.NoError(err2)
	is.EqualValues(None[SomeScanner](), noneScanner)
}

// TestOptionFoldSuccess tests the Fold method with a successful result.
func TestOptionFoldSuccess(t *testing.T) {
	is := assert.New(t)
	option := Option[int]{isPresent: true, value: 10}

	successFunc := func(value int) string {
		return fmt.Sprintf("Success: %v", value)
	}
	failureFunc := func(err error) string {
		return fmt.Sprintf("Failure: %v", err)
	}

	folded := Fold[error, int, string](option, successFunc, failureFunc)
	expected := "Success: 10"

	is.Equal(expected, folded)
}

// TestOptionFoldFailure tests the Fold method with a failure result.
func TestOptionFoldFailure(t *testing.T) {
	is := assert.New(t)
	option := Option[int]{isPresent: false}

	successFunc := func(value int) string {
		return fmt.Sprintf("Success: %v", value)
	}
	failureFunc := func(err error) string {
		return fmt.Sprintf("Failure: %v", err)
	}

	folded := Fold[error, int, string](option, successFunc, failureFunc)
	expected := fmt.Sprintf("Failure: %v", optionNoSuchElement)

	is.Equal(expected, folded)
}
