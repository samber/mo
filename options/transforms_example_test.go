package options_test

import (
	"fmt"
	"strings"

	"github.com/samber/mo"
	"github.com/samber/mo/options"
)

func ExampleMap_some() {
	some := mo.Some("hello world")
	result := options.Map(some, func(v string) int { return len(v) })

	fmt.Printf("%t -> %d", result.IsPresent(), result.OrEmpty())
	// Output: true -> 11
}

func ExampleMap_none() {
	none := mo.None[string]()
	result := options.Map(none, func(v string) int { return len(v) })

	fmt.Printf("%t -> %d", result.IsPresent(), result.OrEmpty())
	// Output: false -> 0
}

func ExampleFlatMap_some() {
	some := mo.Some("hello world")
	result := options.FlatMap(some, func(v string) mo.Option[[]string] {
		if len(v) > 0 {
			return mo.Some(strings.Fields(v))
		}

		return mo.None[[]string]()
	})

	fmt.Printf("%t -> %q", result.IsPresent(), result.OrEmpty())
	// Output: true -> ["hello" "world"]
}

func ExampleFlatMap_none() {
	none := mo.None[string]()
	result := options.FlatMap(none, func(v string) mo.Option[[]string] {
		if len(v) > 0 {
			return mo.Some(strings.Fields(v))
		}

		return mo.None[[]string]()
	})

	fmt.Printf("%t -> %q", result.IsPresent(), result.OrEmpty())
	// Output: false -> []
}

func ExampleMatch_some() {
	some := mo.Some(42)
	result := options.Match(
		some,
		func(i int) (string, bool) {
			return fmt.Sprintf("The number is %d", i), true
		},
		func() (string, bool) {
			return "none", false
		},
	)

	fmt.Printf("%t -> %q", result.IsPresent(), result.OrEmpty())
	// Output: true -> "The number is 42"
}

func ExampleMatch_none() {
	none := mo.None[int]()
	result := options.Match(
		none,
		func(i int) (string, bool) {
			return fmt.Sprintf("The number is %d", i), false
		},
		func() (string, bool) {
			return "No value", true
		},
	)

	fmt.Printf("%t -> %q", result.IsPresent(), result.OrEmpty())
	// Output: true -> "No value"
}

func ExampleFlatMatch_some() {
	some := mo.Some(42)
	result := options.FlatMatch(
		some,
		func(i int) string {
			return fmt.Sprintf("The number is %d", i)
		},
		func() string {
			return "none"
		},
	)

	fmt.Printf("%q", result)
	// Output: "The number is 42"
}

func ExampleFlatMatch_none() {
	none := mo.None[int]()
	result := options.FlatMatch(
		none,
		func(i int) string {
			return fmt.Sprintf("The number is %d", i)
		},
		func() string {
			return "No value"
		},
	)

	fmt.Printf("%q", result)
	// Output: "No value"
}
