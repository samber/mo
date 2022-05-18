package typeclass

type Foldable[T any, R any] interface {
	Map(func(T) R) R
	FoldLeft(func(T) R) R
	FoldRight(func(T) R) R
}
