package typeclass

type Filterable[T any] interface {
	Filter(func(T) bool) T
}
