package mo

import (
	"sync"
)

// NewFuture instanciate a new future.
func NewFuture[T any](cb func(resolve func(T), reject func(error))) *Future[T] {
	future := Future[T]{
		mu:       sync.RWMutex{},
		next:     nil,
		cancelCb: func() {},
	}

	go func() {
		cb(future.resolve, future.reject)
	}()

	return &future
}

// Future represents a value which may or may not currently be available, but will be
// available at some point, or an exception if that value could not be made available.
type Future[T any] struct {
	mu sync.RWMutex

	next     func(T, error)
	cancelCb func()
}

func (f *Future[T]) resolve(value T) {
	f.mu.RLock()
	defer f.mu.RUnlock()

	if f.next != nil {
		f.next(value, nil)
	}
}

func (f *Future[T]) reject(err error) {
	f.mu.RLock()
	defer f.mu.RUnlock()

	if f.next != nil {
		f.next(empty[T](), err)
	}
}

// Then is called when Future is resolved. It returns a new Future.
func (f *Future[T]) Then(cb func(T) (T, error)) *Future[T] {
	f.mu.Lock()
	defer f.mu.Unlock()

	future := &Future[T]{
		mu:   sync.RWMutex{},
		next: nil,
		cancelCb: func() {
			f.Cancel()
		},
	}

	f.next = func(value T, err error) {
		if err != nil {
			future.reject(err)
			return
		}

		newValue, err := cb(value)
		if err != nil {
			future.reject(err)
			return
		}

		future.resolve(newValue)
	}

	return future
}

// Catch is called when Future is rejected. It returns a new Future.
func (f *Future[T]) Catch(cb func(error) (T, error)) *Future[T] {
	f.mu.Lock()
	defer f.mu.Unlock()

	future := &Future[T]{
		mu:   sync.RWMutex{},
		next: nil,
		cancelCb: func() {
			f.Cancel()
		},
	}

	f.next = func(value T, err error) {
		if err == nil {
			future.resolve(value)
			return
		}

		newValue, err := cb(err)
		if err != nil {
			future.reject(err)
			return
		}

		future.resolve(newValue)
	}

	return future
}

// Finally is called when Future is processed either resolved or rejected. It returns a new Future.
func (f *Future[T]) Finally(cb func(T, error) (T, error)) *Future[T] {
	f.mu.Lock()
	defer f.mu.Unlock()

	future := &Future[T]{
		mu:   sync.RWMutex{},
		next: nil,
		cancelCb: func() {
			f.Cancel()
		},
	}

	f.next = func(value T, err error) {
		newValue, err := cb(value, err)
		if err != nil {
			future.reject(err)
			return
		}

		future.resolve(newValue)
	}

	return future
}

// Cancel cancels the Future chain.
func (f *Future[T]) Cancel() {
	f.mu.Lock()
	defer f.mu.Unlock()

	f.next = nil
	if f.cancelCb != nil {
		f.cancelCb()
	}
}

// Collect awaits and return result of the Future.
func (f *Future[T]) Collect() (T, error) {
	done := make(chan struct{})

	var a T
	var b error

	f.mu.Lock()
	f.next = func(value T, err error) {
		a = value
		b = err

		done <- struct{}{}
	}
	f.mu.Unlock()

	<-done

	return a, b
}

// Result wraps Collect and returns a Result.
func (f *Future[T]) Result() Result[T] {
	return TupleToResult(f.Collect())
}

// Either wraps Collect and returns a Either.
func (f *Future[T]) Either() Either[error, T] {
	v, err := f.Collect()
	if err != nil {
		return Left[error, T](err)
	}
	return Right[error, T](v)
}
