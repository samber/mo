package mo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewState(t *testing.T) {
	is := assert.New(t)

	// A computation that returns the current state doubled as the result and
	// increments the state.
	s := NewState(func(state int) (int, int) {
		return state * 2, state + 1
	})

	result, newState := s.Run(10)
	is.Equal(20, result)
	is.Equal(11, newState)
}

func TestReturnState(t *testing.T) {
	is := assert.New(t)

	// ReturnState yields the given value and leaves the state untouched.
	s := ReturnState[string](42)

	result, newState := s.Run("unchanged")
	is.Equal(42, result)
	is.Equal("unchanged", newState)
}

func TestStateRun(t *testing.T) {
	is := assert.New(t)

	s := NewState(func(state int) (string, int) {
		return "value", state + 100
	})

	result, newState := s.Run(1)
	is.Equal("value", result)
	is.Equal(101, newState)
}

func TestStateGet(t *testing.T) {
	is := assert.New(t)

	s := NewState(func(state int) (string, int) {
		return "ignored", state
	})

	// Get returns the current state as both the result and the state.
	getter := s.Get()
	result, newState := getter.Run(7)
	is.Equal(7, result)
	is.Equal(7, newState)
}

func TestStateModify(t *testing.T) {
	is := assert.New(t)

	s := NewState(func(state int) (string, int) {
		return "ignored", state
	})

	// Modify applies f to the state and yields the zero value of A.
	modified := s.Modify(func(state int) int {
		return state * 10
	})

	result, newState := modified.Run(5)
	is.Equal("", result) // zero value of string
	is.Equal(50, newState)
}

func TestStatePut(t *testing.T) {
	is := assert.New(t)

	s := NewState(func(state int) (string, int) {
		return "ignored", state
	})

	// Put replaces the state and yields the zero value of A.
	put := s.Put(99)

	result, newState := put.Run(1)
	is.Equal("", result) // zero value of string
	is.Equal(99, newState)
}
