package mo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTransation(t *testing.T) {
	is := assert.New(t)

	// no error
	{
		transaction := NewTransaction[int]().
			Then(
				func(state int) (int, error) {
					return state + 100, nil
				},
				func(state int) int {
					return state - 100
				},
			).
			Then(
				func(state int) (int, error) {
					return state + 21, nil
				},
				func(state int) int {
					return state - 21
				},
			)

		state, err := transaction.Process(21)
		is.Equal(142, state)
		is.Equal(nil, err)
	}

	// with error
	{
		transaction := NewTransaction[int]().
			Then(
				func(state int) (int, error) {
					return state + 100, nil
				},
				func(state int) int {
					return state - 100
				},
			).
			Then(
				func(state int) (int, error) {
					return state, assert.AnError
				},
				func(state int) int {
					return state - 21
				},
			).
			Then(
				func(state int) (int, error) {
					return state + 42, nil
				},
				func(state int) int {
					return state - 42
				},
			)

		state, err := transaction.Process(21)
		is.Equal(21, state)
		is.Equal(assert.AnError, err)
	}

	// with error + update value
	{
		transaction := NewTransaction[int]().
			Then(
				func(state int) (int, error) {
					return state + 100, nil
				},
				func(state int) int {
					return state - 100
				},
			).
			Then(
				func(state int) (int, error) {
					return state + 21, assert.AnError
				},
				func(state int) int {
					return state - 21
				},
			).
			Then(
				func(state int) (int, error) {
					return state + 42, nil
				},
				func(state int) int {
					return state - 42
				},
			)

		state, err := transaction.Process(21)
		is.Equal(42, state)
		is.Equal(assert.AnError, err)
	}
}
