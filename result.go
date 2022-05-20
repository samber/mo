package mo

// Ok builds a Result when value is valid.
func Ok[T any](value T) Result[T] {
	return Result[T]{
		value: value,
		isErr: false,
	}
}

// Err builds a Result when value is invalid.
func Err[T any](err error) Result[T] {
	return Result[T]{
		err:   err,
		isErr: true,
	}
}

// TupleToResult convert a pair of T and error into a Result.
func TupleToResult[T any](value T, err error) Result[T] {
	if err != nil {
		return Err[T](err)
	}
	return Ok(value)
}

// Result respresent a result of an action having one
// of the following output: success or failure.
// An instance of Result is an instance of either Ok or Err.
// It could be compared to `Either[error, T]`.
type Result[T any] struct {
	isErr bool
	value T
	err   error
}

// IsOk returns true when value is valid.
func (r Result[T]) IsOk() bool {
	return !r.isErr
}

// IsError returns true when value is invalid.
func (r Result[T]) IsError() bool {
	return r.isErr
}

// Error returns error when value is invalid or nil.
func (r Result[T]) Error() error {
	return r.err
}

// MustGet returns value and error.
func (r Result[T]) Get() (T, error) {
	if r.isErr {
		return empty[T](), r.err
	}

	return r.value, nil
}

// MustGet returns value when Result is valid or panics.
func (r Result[T]) MustGet() T {
	if r.isErr {
		panic(r.err)
	}

	return r.value
}

// OrElse returns value when Result is valid or default value.
func (r Result[T]) OrElse(fallback T) T {
	if r.isErr {
		return fallback
	}

	return r.value
}

// OrEmpty returns value when Result is valid or empty value.
func (r Result[T]) OrEmpty() T {
	return r.value
}

// ToEither transforms a Result into an Either type.
func (r Result[T]) ToEither() Either[error, T] {
	if r.isErr {
		return Left[error, T](r.err)
	}

	return Right[error, T](r.value)
}

// ForEach executes the given side-effecting function if Result is valid.
func (r Result[T]) ForEach(mapper func(value T)) {
	if !r.isErr {
		mapper(r.value)
	}
}

// Match executes the first function if Result is valid and second function if invalid.
// It returns a new Result.
func (r Result[T]) Match(onSuccess func(value T) (T, error), onError func(err error) (T, error)) Result[T] {
	if r.isErr {
		return TupleToResult(onError(r.err))
	}
	return TupleToResult(onSuccess(r.value))
}

// Map executes the mapper function if Result is valid. It returns a new Result.
func (r Result[T]) Map(mapper func(value T) (T, error)) Result[T] {
	if !r.isErr {
		return TupleToResult(mapper(r.value))
	}

	return Err[T](r.err)
}

// MapErr executes the mapper function if Result is invalid. It returns a new Result.
func (r Result[T]) MapErr(mapper func(error) (T, error)) Result[T] {
	if r.isErr {
		return TupleToResult(mapper(r.err))
	}

	return Ok(r.value)
}

// FlatMap executes the mapper function if Result is valid. It returns a new Result.
func (r Result[T]) FlatMap(mapper func(value T) Result[T]) Result[T] {
	if !r.isErr {
		return mapper(r.value)
	}

	return Err[T](r.err)
}
