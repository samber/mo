package mo

import (
	"encoding/json"
	"fmt"
)

func ExampleSome() {
	some := Some(42)
	result := some.OrElse(1234)

	fmt.Println(result)
	// Output: 42
}

func ExampleNone() {
	none := None[int]()
	result := none.OrElse(1234)

	fmt.Println(result)
	// Output: 1234
}

func ExampleTupleToOption() {
	m := map[string]int{
		"foo": 21,
		"bar": 42,
		"baz": 84,
	}

	value, ok := m["hello world"]

	none := TupleToOption(value, ok)
	result := none.OrElse(1234)

	fmt.Println(result)
	// Output: 1234
}

func ExampleEmptyableToOption() {
	cb := func(ok bool) error {
		if ok {
			return nil
		}

		return fmt.Errorf("an error")
	}

	err := cb(false)

	none := EmptyableToOption(err)
	result, ok := none.Get()

	fmt.Println(result)
	fmt.Println(ok)
	// Output:
	// an error
	// true
}

func ExampleOption_some() {
	some := Some(42)
	result := some.OrElse(1234)

	fmt.Println(result)
	// Output: 42
}

func ExampleOption_none() {
	none := None[int]()
	result := none.OrElse(1234)

	fmt.Println(result)
	// Output: 1234
}

func ExampleOption_IsPresent_some() {
	some := Some(42)
	result := some.IsPresent()

	fmt.Println(result)
	// Output: true
}

func ExampleOption_IsPresent_none() {
	none := None[int]()
	result := none.IsPresent()

	fmt.Println(result)
	// Output: false
}

func ExampleOption_IsAbsent_some() {
	some := Some(42)
	result := some.IsAbsent()

	fmt.Println(result)
	// Output: false
}

func ExampleOption_IsAbsent_none() {
	none := None[int]()
	result := none.IsAbsent()

	fmt.Println(result)
	// Output: true
}

func ExampleOption_Size_some() {
	some := Some(42)
	result := some.Size()

	fmt.Println(result)
	// Output: 1
}

func ExampleOption_Size_none() {
	none := None[int]()
	result := none.Size()

	fmt.Println(result)
	// Output: 0
}

func ExampleOption_Get_some() {
	some := Some(42)
	result, ok := some.Get()

	fmt.Println(result)
	fmt.Println(ok)
	// Output:
	// 42
	// true
}

func ExampleOption_Get_none() {
	none := None[int]()
	result, ok := none.Get()

	fmt.Println(result)
	fmt.Println(ok)
	// Output:
	// 0
	// false
}

func ExampleOption_MustGet_some() {
	some := Some(42)
	result := some.MustGet()

	fmt.Println(result)
	// Output: 42
}

// func ExampleOption_MustGet_none() {
// 	none := None[int]()
// 	result := none.MustGet()

// 	fmt.Println(result)
// 	// Output: panics
// }

func ExampleOption_OrElse_some() {
	some := Some(42)
	result := some.OrElse(1234)

	fmt.Println(result)
	// Output: 42
}

func ExampleOption_OrElse_none() {
	none := None[int]()
	result := none.OrElse(1234)

	fmt.Println(result)
	// Output: 1234
}

func ExampleOption_OrEmpty_some() {
	some := Some(42)
	result := some.OrEmpty()

	fmt.Println(result)
	// Output: 42
}

func ExampleOption_OrEmpty_none() {
	none := None[int]()
	result := none.OrEmpty()

	fmt.Println(result)
	// Output: 0
}

func ExampleOption_ToPointer_some() {
	some := Some(42)
	result := some.ToPointer()

	fmt.Println(*result)
	// Output: 42
}

func ExampleOption_ToPointer_none() {
	none := None[int]()
	result := none.ToPointer()

	fmt.Println(result)
	// Output: <nil>
}

func ExampleOption_Match_some() {
	some := Some(42)
	result := some.Match(
		func(i int) (int, bool) {
			return 0, false
		},
		func() (int, bool) {
			return 2, true
		},
	)

	fmt.Println(result.IsPresent(), result.OrEmpty())
	// Output: false 0
}

func ExampleOption_Match_none() {
	none := None[int]()
	result := none.Match(
		func(i int) (int, bool) {
			return 0, false
		},
		func() (int, bool) {
			return 2, true
		},
	)

	fmt.Println(result.IsPresent(), result.OrEmpty())
	// Output: true 2
}

func ExampleOption_Map_some() {
	some := Some(42)
	result := some.Map(func(i int) (int, bool) {
		return 1234, true
	})

	fmt.Println(result.IsPresent(), result.OrEmpty())
	// Output: true 1234
}

func ExampleOption_Map_none() {
	none := None[int]()
	result := none.Map(func(i int) (int, bool) {
		return 1234, true
	})

	fmt.Println(result.IsPresent(), result.OrEmpty())
	// Output: false 0
}

func ExampleOption_MapNone_some() {
	some := Some(42)
	result := some.MapNone(func() (int, bool) {
		return 1234, true
	})

	fmt.Println(result.IsPresent(), result.OrEmpty())
	// Output: true 42
}

func ExampleOption_MapNone_none() {
	none := None[int]()
	result := none.MapNone(func() (int, bool) {
		return 1234, true
	})

	fmt.Println(result.IsPresent(), result.OrEmpty())
	// Output: true 1234
}

func ExampleOption_FlatMap_some() {
	some := Some(42)
	result := some.FlatMap(func(i int) Option[int] {
		return Some(21)
	})

	fmt.Println(result.IsPresent(), result.OrEmpty())
	// Output: true 21
}

func ExampleOption_FlatMap_none() {
	none := None[int]()
	result := none.FlatMap(func(i int) Option[int] {
		return Some(21)
	})

	fmt.Println(result.IsPresent(), result.OrEmpty())
	// Output: false 0
}

func ExampleOption_MarshalJSON_some() {
	type test struct {
		Email Option[string] `json:"email"`
	}

	value := test{Email: Some("samuel@example.com")}
	result, err := json.Marshal(value)

	fmt.Println(string(result))
	fmt.Println(err)
	// Output:
	// {"email":"samuel@example.com"}
	// <nil>
}

func ExampleOption_MarshalJSON_none() {
	type test struct {
		Email Option[string] `json:"email"`
	}

	value := test{Email: None[string]()}
	result, err := json.Marshal(value)

	fmt.Println(string(result))
	fmt.Println(err)
	// Output:
	// {"email":null}
	// <nil>
}

func ExampleOption_UnmarshalJSON_some() {
	type test struct {
		Email Option[string] `json:"email"`
	}

	value := []byte(`{"email":"samuel@example.com"}`)

	var result test
	err := json.Unmarshal(value, &result)

	fmt.Println(result.Email.Get())
	fmt.Println(err)
	// Output:
	// samuel@example.com true
	// <nil>
}

func ExampleOption_UnmarshalJSON_none() {
	type test struct {
		Email Option[string] `json:"email"`
	}

	value := []byte(`{"email":null}`)

	var result test
	err := json.Unmarshal(value, &result)

	fmt.Println(result.Email.Get())
	fmt.Println(err)
	// Output:
	// false
	// <nil>
}
