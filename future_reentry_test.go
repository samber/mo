package mo

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFutureCallbackCanCollectParent(t *testing.T) {
	start := make(chan struct{})
	parent := NewFuture(func(resolve func(int), reject func(error)) {
		<-start
		resolve(42)
	})
	child := parent.Then(func(value int) (int, error) {
		previous, err := parent.Collect()
		return value + previous, err
	})
	close(start)
	value, err := child.Collect()
	assert.Equal(t, 84, value)
	assert.NoError(t, err)
}

func TestFutureCatchCanCollectParent(t *testing.T) {
	start := make(chan struct{})
	parent := NewFuture(func(resolve func(int), reject func(error)) {
		<-start
		reject(assert.AnError)
	})
	child := parent.Catch(func(error) (int, error) {
		return parent.Collect()
	})
	close(start)
	value, err := child.Collect()
	assert.Zero(t, value)
	assert.ErrorIs(t, err, assert.AnError)
}

func TestFutureCallbackCanCancelParent(t *testing.T) {
	for _, rejectParent := range []bool{false, true} {
		t.Run(fmt.Sprint(rejectParent), func(t *testing.T) {
			start := make(chan struct{})
			parent := NewFuture(func(resolve func(int), reject func(error)) {
				<-start
				if rejectParent {
					reject(assert.AnError)
				} else {
					resolve(42)
				}
			})
			child := parent.Finally(func(value int, err error) (int, error) {
				parent.Cancel()
				return value, err
			})
			close(start)
			value, err := child.Collect()
			if rejectParent {
				assert.Zero(t, value)
				assert.ErrorIs(t, err, assert.AnError)
			} else {
				assert.Equal(t, 42, value)
				assert.NoError(t, err)
			}
		})
	}
}

func TestFutureCallbackCanChainParent(t *testing.T) {
	start := make(chan struct{})
	parent := NewFuture(func(resolve func(int), reject func(error)) {
		<-start
		resolve(42)
	})
	child := parent.Then(func(int) (int, error) {
		return parent.Then(func(value int) (int, error) {
			return value * 2, nil
		}).Collect()
	})
	close(start)
	value, err := child.Collect()
	assert.Equal(t, 84, value)
	assert.NoError(t, err)
}

func TestFutureCallbackCanSettleParentAgain(t *testing.T) {
	start := make(chan struct{})
	var settle func(int)
	parent := NewFuture(func(resolve func(int), reject func(error)) {
		settle = resolve
		<-start
		resolve(42)
	})
	child := parent.Then(func(value int) (int, error) {
		settle(84)
		return value, nil
	})
	close(start)
	value, err := child.Collect()
	assert.Equal(t, 42, value)
	assert.NoError(t, err)
	value, err = parent.Collect()
	assert.Equal(t, 42, value)
	assert.NoError(t, err)
}
