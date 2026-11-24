package options_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/samber/mo"
	"github.com/samber/mo/options"
)

func nonEmptyFields(s string) mo.Option[[]string] {
	fields := strings.Fields(s)
	if len(fields) > 0 {
		return mo.Some(fields)
	}
	return mo.None[[]string]()
}

func TestMap(t *testing.T) {
	t.Run("Some", func(t *testing.T) {
		original := mo.Some("hello world")

		mapped := options.Map(original, strings.Fields)

		assert.Equal(t, mo.Some([]string{"hello", "world"}), mapped)
	})
	t.Run("None", func(t *testing.T) {
		original := mo.None[string]()

		mapped := options.Map(original, strings.Fields)

		assert.Equal(t, mo.None[[]string](), mapped)
	})
}

func TestFlatMap(t *testing.T) {
	t.Run("Some-to-Some", func(t *testing.T) {
		original := mo.Some("hello world")

		mapped := options.FlatMap(original, nonEmptyFields)

		assert.Equal(t, mo.Some([]string{"hello", "world"}), mapped)
	})
	t.Run("Some-to-None", func(t *testing.T) {
		original := mo.Some("")

		mapped := options.FlatMap(original, nonEmptyFields)

		assert.Equal(t, mo.None[[]string](), mapped)
	})
	t.Run("None", func(t *testing.T) {
		original := mo.None[string]()

		mapped := options.FlatMap(original, nonEmptyFields)

		assert.Equal(t, mo.None[[]string](), mapped)
	})
}

func TestMatch(t *testing.T) {
	t.Run("Some-to-Present", func(t *testing.T) {
		original := mo.Some("hello world")

		result := options.Match(
			original,
			func(val string) (int, bool) {
				assert.Equal(t, "hello world", val)
				return 1234, true
			},
			func() (int, bool) {
				require.Fail(t, "should not be called")
				return 0, false
			},
		)

		assert.Equal(t, mo.Some(1234), result)
	})
	t.Run("Some-to-Absent", func(t *testing.T) {
		original := mo.Some("hello world")

		result := options.Match(
			original,
			func(val string) (int, bool) {
				assert.Equal(t, "hello world", val)
				return 0, false
			},
			func() (int, bool) {
				require.Fail(t, "should not be called")
				return 0, false
			},
		)

		assert.Equal(t, mo.None[int](), result)
	})
	t.Run("None-to-Present", func(t *testing.T) {
		original := mo.None[string]()

		result := options.Match(
			original,
			func(val string) (int, bool) {
				require.Fail(t, "should not be called")
				return 0, false
			},
			func() (int, bool) {
				return 1234, true
			},
		)

		assert.Equal(t, mo.Some(1234), result)
	})
	t.Run("None-to-Absent", func(t *testing.T) {
		original := mo.None[string]()

		result := options.Match(
			original,
			func(val string) (int, bool) {
				require.Fail(t, "should not be called")
				return 0, false
			},
			func() (int, bool) {
				return 0, false
			},
		)

		assert.Equal(t, mo.None[int](), result)
	})
}

func TestFlatMatch(t *testing.T) {
	t.Run("Some", func(t *testing.T) {
		original := mo.Some("hello world")

		result := options.FlatMatch(
			original,
			func(val string) int {
				assert.Equal(t, "hello world", val)
				return 1234
			},
			func() int {
				require.Fail(t, "should not be called")
				return 0
			},
		)

		assert.Equal(t, 1234, result)
	})
	t.Run("None", func(t *testing.T) {
		original := mo.None[string]()

		result := options.FlatMatch(
			original,
			func(val string) int {
				require.Fail(t, "should not be called")
				return 0
			},
			func() int {
				return 1234
			},
		)

		assert.Equal(t, 1234, result)
	})
}
