package typeclass

type Monadic[T any] interface {
	ForEach(func(T))
	Map(func(T) T) T
	FlatMap(func(T) Monadic[T]) T
}
