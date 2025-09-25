package result

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/samber/mo"
	"github.com/stretchr/testify/assert"
)

func Pipe1[A any, B any](
	source mo.Result[A],
	operator1 func(mo.Result[A]) mo.Result[B],
) mo.Result[B] {
	return operator1(
		source,
	)
}

func Pipe2[A any, B any, C any](
	source mo.Result[A],
	operator1 func(mo.Result[A]) mo.Result[B],
	operator2 func(mo.Result[B]) mo.Result[C],
) mo.Result[C] {
	return operator2(
		operator1(
			source,
		),
	)
}

func Pipe3[A any, B any, C any, D any](
	source mo.Result[A],
	operator1 func(mo.Result[A]) mo.Result[B],
	operator2 func(mo.Result[B]) mo.Result[C],
	operator3 func(mo.Result[C]) mo.Result[D],
) mo.Result[D] {
	return operator3(
		operator2(
			operator1(source),
		),
	)
}

func Pipe4[A any, B any, C any, D any, E any](
	source mo.Result[A],
	operator1 func(mo.Result[A]) mo.Result[B],
	operator2 func(mo.Result[B]) mo.Result[C],
	operator3 func(mo.Result[C]) mo.Result[D],
	operator4 func(mo.Result[D]) mo.Result[E],
) mo.Result[E] {
	return operator4(
		operator3(
			operator2(
				operator1(source),
			),
		),
	)
}

func Pipe5[A any, B any, C any, D any, E any, F any](
	source mo.Result[A],
	operator1 func(mo.Result[A]) mo.Result[B],
	operator2 func(mo.Result[B]) mo.Result[C],
	operator3 func(mo.Result[C]) mo.Result[D],
	operator4 func(mo.Result[D]) mo.Result[E],
	operator5 func(mo.Result[E]) mo.Result[F],
) mo.Result[F] {
	return operator5(
		operator4(
			operator3(
				operator2(
					operator1(source),
				),
			),
		),
	)
}

func Pipe6[A any, B any, C any, D any, E any, F any, G any](
	source mo.Result[A],
	operator1 func(mo.Result[A]) mo.Result[B],
	operator2 func(mo.Result[B]) mo.Result[C],
	operator3 func(mo.Result[C]) mo.Result[D],
	operator4 func(mo.Result[D]) mo.Result[E],
	operator5 func(mo.Result[E]) mo.Result[F],
	operator6 func(mo.Result[F]) mo.Result[G],
) mo.Result[G] {
	return operator6(
		operator5(
			operator4(
				operator3(
					operator2(
						operator1(source),
					),
				),
			),
		),
	)
}

func Pipe7[A any, B any, C any, D any, E any, F any, G any, H any](
	source mo.Result[A],
	operator1 func(mo.Result[A]) mo.Result[B],
	operator2 func(mo.Result[B]) mo.Result[C],
	operator3 func(mo.Result[C]) mo.Result[D],
	operator4 func(mo.Result[D]) mo.Result[E],
	operator5 func(mo.Result[E]) mo.Result[F],
	operator6 func(mo.Result[F]) mo.Result[G],
	operator7 func(mo.Result[G]) mo.Result[H],
) mo.Result[H] {
	return operator7(
		operator6(
			operator5(
				operator4(
					operator3(
						operator2(
							operator1(source),
						),
					),
				),
			),
		),
	)
}

func Pipe8[A any, B any, C any, D any, E any, F any, G any, H any, I any](
	source mo.Result[A],
	operator1 func(mo.Result[A]) mo.Result[B],
	operator2 func(mo.Result[B]) mo.Result[C],
	operator3 func(mo.Result[C]) mo.Result[D],
	operator4 func(mo.Result[D]) mo.Result[E],
	operator5 func(mo.Result[E]) mo.Result[F],
	operator6 func(mo.Result[F]) mo.Result[G],
	operator7 func(mo.Result[G]) mo.Result[H],
	operator8 func(mo.Result[H]) mo.Result[I],
) mo.Result[I] {
	return operator8(
		operator7(
			operator6(
				operator5(
					operator4(
						operator3(
							operator2(
								operator1(source),
							),
						),
					),
				),
			),
		),
	)
}

func Pipe9[A any, B any, C any, D any, E any, F any, G any, H any, I any, J any](
	source mo.Result[A],
	operator1 func(mo.Result[A]) mo.Result[B],
	operator2 func(mo.Result[B]) mo.Result[C],
	operator3 func(mo.Result[C]) mo.Result[D],
	operator4 func(mo.Result[D]) mo.Result[E],
	operator5 func(mo.Result[E]) mo.Result[F],
	operator6 func(mo.Result[F]) mo.Result[G],
	operator7 func(mo.Result[G]) mo.Result[H],
	operator8 func(mo.Result[H]) mo.Result[I],
	operator9 func(mo.Result[I]) mo.Result[J],
) mo.Result[J] {
	return operator9(
		operator8(
			operator7(
				operator6(
					operator5(
						operator4(
							operator3(
								operator2(
									operator1(source),
								),
							),
						),
					),
				),
			),
		),
	)
}

func Pipe10[A any, B any, C any, D any, E any, F any, G any, H any, I any, J any, K any](
	source mo.Result[A],
	operator1 func(mo.Result[A]) mo.Result[B],
	operator2 func(mo.Result[B]) mo.Result[C],
	operator3 func(mo.Result[C]) mo.Result[D],
	operator4 func(mo.Result[D]) mo.Result[E],
	operator5 func(mo.Result[E]) mo.Result[F],
	operator6 func(mo.Result[F]) mo.Result[G],
	operator7 func(mo.Result[G]) mo.Result[H],
	operator8 func(mo.Result[H]) mo.Result[I],
	operator9 func(mo.Result[I]) mo.Result[J],
	operator10 func(mo.Result[J]) mo.Result[K],
) mo.Result[K] {
	return operator10(
		operator9(
			operator8(
				operator7(
					operator6(
						operator5(
							operator4(
								operator3(
									operator2(
										operator1(source),
									),
								),
							),
						),
					),
				),
			),
		),
	)
}

func TestPipeTypeTransformations(t *testing.T) {
	is := assert.New(t)

	out := Pipe3(
		mo.Ok("42"),
		FlatMap(func(str string) mo.Result[int] {
			return mo.TupleToResult(strconv.Atoi(str))
		}),
		Map(func(n int) float64 {
			return float64(n)
		}),
		Map(func(n float64) string {
			return fmt.Sprintf("%.2f", n)
		}),
	)
	is.Equal(mo.Ok("42.00"), out)
}
