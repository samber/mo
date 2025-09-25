package mo

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type atomicInt32 struct {
	v  int32
	mu sync.Mutex
}

func (a *atomicInt32) Add(n int32) int32 {
	a.mu.Lock()
	defer a.mu.Unlock()
	return atomic.AddInt32(&a.v, n)
}

func assertAndIncrement(t *testing.T, is *assert.Assertions, expected int, i *atomicInt32) {
	got := i.Add(1)
	is.Equal(int32(expected), got-1)
}

func TestFuture(t *testing.T) {
	is := assert.New(t)

	result, err := NewFuture[int](func(resolve func(int), reject func(error)) {
		resolve(42)
	}).Then(func(value int) (int, error) {
		is.Equal(42, value)
		return 21, assert.AnError
	}).Catch(func(err error) (int, error) {
		is.Equal(assert.AnError, err)
		return 0, nil
	}).Then(func(value int) (int, error) {
		is.Equal(0, value)
		return 84, nil
	}).Collect()

	is.Equal(84, result)
	is.Nil(err)
}

func TestFutureSimpleResolve(t *testing.T) {
	is := assert.New(t)

	result, err := NewFuture[int](func(resolve func(int), reject func(error)) {
		resolve(42)
	}).Collect()

	is.Equal(42, result)
	is.Nil(err)
}

func TestFutureSimpleReject(t *testing.T) {
	is := assert.New(t)

	result, err := NewFuture[int](func(resolve func(int), reject func(error)) {
		reject(assert.AnError)
	}).Collect()

	is.Equal(0, result)
	is.NotNil(err)
	is.Equal(assert.AnError, err)
}

func TestFutureMultipleResolve(t *testing.T) {
	is := assert.New(t)

	result, err := NewFuture[int](func(resolve func(int), reject func(error)) {
		resolve(42)
	}).Then(func(value int) (int, error) {
		is.Equal(42, value)
		return 84, nil
	}).Then(func(value int) (int, error) {
		is.Equal(84, value)
		return 21, nil
	}).Collect()

	is.Equal(21, result)
	is.Nil(err)
}

func TestFutureMultipleReject(t *testing.T) {
	is := assert.New(t)

	result, err := NewFuture[int](func(resolve func(int), reject func(error)) {
		resolve(42)
	}).Catch(func(err error) (int, error) {
		is.Fail("should not enter here")
		return 84, assert.AnError
	}).Then(func(value int) (int, error) {
		is.Equal(42, value)
		return 21, assert.AnError
	}).Catch(func(err error) (int, error) {
		is.Equal(assert.AnError, err)
		return 1, nil
	}).Collect()

	is.Equal(1, result)
	is.Nil(err)
}

func TestFutureSingleReject(t *testing.T) {
	is := assert.New(t)

	result, err := NewFuture[int](func(resolve func(int), reject func(error)) {
		reject(assert.AnError)
	}).Catch(func(err error) (int, error) {
		is.Equal(assert.AnError, err)
		return 84, nil
	}).Collect()

	is.Equal(84, result)
	is.Nil(err)
}

func TestFutureErrorResult(t *testing.T) {
	is := assert.New(t)

	result, err := NewFuture[int](func(resolve func(int), reject func(error)) {
		reject(assert.AnError)
	}).Collect()

	is.Equal(0, result)
	is.NotNil(err)
	is.Equal(assert.AnError, err)
}

func TestFutureFinallyResolve(t *testing.T) {
	is := assert.New(t)

	result, err := NewFuture[int](func(resolve func(int), reject func(error)) {
		resolve(21)
	}).Finally(func(value int, err error) (int, error) {
		is.Equal(21, value)
		is.Nil(err)

		return 42, nil
	}).Collect()

	is.Equal(42, result)
	is.Nil(err)
}

func TestFutureFinallyReject(t *testing.T) {
	is := assert.New(t)

	result, err := NewFuture[int](func(resolve func(int), reject func(error)) {
		reject(assert.AnError)
	}).Finally(func(value int, err error) (int, error) {
		is.Equal(0, value)
		is.NotNil(err)
		is.Equal(assert.AnError, err)

		return 42, nil
	}).Collect()

	is.Equal(42, result)
	is.Nil(err)
}

func TestFutureOrder(t *testing.T) {
	is := assert.New(t)

	var i atomicInt32

	fut := NewFuture[int](func(resolve func(int), reject func(error)) {
		assertAndIncrement(t, is, 1, &i)

		resolve(42)
	}).Then(func(value int) (int, error) {
		assertAndIncrement(t, is, 2, &i)

		return 21, assert.AnError
	}).Catch(func(err error) (int, error) {
		assertAndIncrement(t, is, 3, &i)

		return 1, nil
	}).Finally(func(value int, err error) (int, error) {
		assertAndIncrement(t, is, 4, &i)

		return 21, nil
	})

	assertAndIncrement(t, is, 0, &i)

	_, _ = fut.Collect()

	assertAndIncrement(t, is, 5, &i)
}

func TestFutureOrderCollect(t *testing.T) {
	is := assert.New(t)

	var i atomicInt32

	_, _ = NewFuture[int](func(resolve func(int), reject func(error)) {
		assertAndIncrement(t, is, 0, &i)

		resolve(42)
	}).Then(func(value int) (int, error) {
		assertAndIncrement(t, is, 1, &i)

		return 21, assert.AnError
	}).Catch(func(err error) (int, error) {
		assertAndIncrement(t, is, 2, &i)

		return 1, nil
	}).Finally(func(value int, err error) (int, error) {
		assertAndIncrement(t, is, 3, &i)

		return 1, nil
	}).Collect()

	assertAndIncrement(t, is, 4, &i)
}

func TestFutureCancel(t *testing.T) {
	is := assert.New(t)

	var i atomicInt32

	future := NewFuture[int](func(resolve func(int), reject func(error)) {
		assertAndIncrement(t, is, 0, &i)

		time.Sleep(5 * time.Millisecond)

		resolve(42)
	}).Then(func(value int) (int, error) {
		assertAndIncrement(t, is, 3, &i)
		is.Fail("should not enter here")

		return 21, assert.AnError
	})

	time.Sleep(1 * time.Millisecond)
	assertAndIncrement(t, is, 1, &i)
	future.Cancel()

	time.Sleep(10 * time.Millisecond)
	assertAndIncrement(t, is, 2, &i)
}

func TestFutureCancelDelayed(t *testing.T) {
	is := assert.New(t)

	var i atomicInt32

	future := NewFuture[int](func(resolve func(int), reject func(error)) {
		time.Sleep(1 * time.Millisecond)
		assertAndIncrement(t, is, 1, &i)

		resolve(42)
	}).Then(func(value int) (int, error) {
		assertAndIncrement(t, is, 2, &i)

		return 21, assert.AnError
	})

	assertAndIncrement(t, is, 0, &i)

	time.Sleep(10 * time.Millisecond)

	future.Cancel()

	assertAndIncrement(t, is, 3, &i)
}

func TestFutureCancelTerminated(t *testing.T) {
	is := assert.New(t)

	var i atomicInt32

	future := NewFuture[int](func(resolve func(int), reject func(error)) {
		time.Sleep(1 * time.Millisecond)
		assertAndIncrement(t, is, 1, &i)

		resolve(42)
	}).Then(func(value int) (int, error) {
		assertAndIncrement(t, is, 2, &i)

		return 21, assert.AnError
	})

	assertAndIncrement(t, is, 0, &i)

	_, _ = future.Collect()

	assertAndIncrement(t, is, 3, &i)

	future.Cancel()

	assertAndIncrement(t, is, 4, &i)
}

func TestFutureResultResult(t *testing.T) {
	is := assert.New(t)

	result := NewFuture[int](func(resolve func(int), reject func(error)) {
		reject(assert.AnError)
	}).Result()

	is.Equal(Err[int](assert.AnError), result)
	is.NotNil(result.Error())
	is.Equal(assert.AnError, result.Error())
}

func TestFutureResultEither(t *testing.T) {
	is := assert.New(t)

	either := NewFuture[int](func(resolve func(int), reject func(error)) {
		reject(assert.AnError)
	}).Either()

	is.Equal(Left[error, int](assert.AnError), either)
	is.NotNil(either.Left())
	is.Equal(assert.AnError, either.MustLeft())
}

func TestFutureCompleteBeforeThen(t *testing.T) {
	completed := make(chan struct{})
	fut := NewFuture(func(resolve func(int), reject func(error)) {
		resolve(1)
		close(completed)
	})

	<-completed
	//nolint:errcheck
	fut.Then(func(in int) (int, error) {
		fmt.Println(in) // will never been print
		return in, nil
	}).Collect() // deadlock
}
