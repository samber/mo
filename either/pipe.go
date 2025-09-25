package either

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/samber/mo"
	"github.com/stretchr/testify/assert"
)

func Pipe1[A1 any, A2 any, B1 any, B2 any](
	source mo.Either[A1, A2],
	operator1 func(mo.Either[A1, A2]) mo.Either[B1, B2],
) mo.Either[B1, B2] {
	return operator1(
		source,
	)
}

func Pipe2[A1 any, A2 any, B1 any, B2 any, C1 any, C2 any](
	source mo.Either[A1, A2],
	operator1 func(mo.Either[A1, A2]) mo.Either[B1, B2],
	operator2 func(mo.Either[B1, B2]) mo.Either[C1, C2],
) mo.Either[C1, C2] {
	return operator2(
		operator1(
			source,
		),
	)
}

func Pipe3[A1 any, A2 any, B1 any, B2 any, C1 any, C2 any, D1 any, D2 any](
	source mo.Either[A1, A2],
	operator1 func(mo.Either[A1, A2]) mo.Either[B1, B2],
	operator2 func(mo.Either[B1, B2]) mo.Either[C1, C2],
	operator3 func(mo.Either[C1, C2]) mo.Either[D1, D2],
) mo.Either[D1, D2] {
	return operator3(
		operator2(
			operator1(
				source,
			),
		),
	)
}

func Pipe4[A1 any, A2 any, B1 any, B2 any, C1 any, C2 any, D1 any, D2 any, E1 any, E2 any](
	source mo.Either[A1, A2],
	operator1 func(mo.Either[A1, A2]) mo.Either[B1, B2],
	operator2 func(mo.Either[B1, B2]) mo.Either[C1, C2],
	operator3 func(mo.Either[C1, C2]) mo.Either[D1, D2],
	operator4 func(mo.Either[D1, D2]) mo.Either[E1, E2],
) mo.Either[E1, E2] {
	return operator4(
		operator3(
			operator2(
				operator1(
					source,
				),
			),
		),
	)
}

func Pipe5[A1 any, A2 any, B1 any, B2 any, C1 any, C2 any, D1 any, D2 any, E1 any, E2 any, F1 any, F2 any](
	source mo.Either[A1, A2],
	operator1 func(mo.Either[A1, A2]) mo.Either[B1, B2],
	operator2 func(mo.Either[B1, B2]) mo.Either[C1, C2],
	operator3 func(mo.Either[C1, C2]) mo.Either[D1, D2],
	operator4 func(mo.Either[D1, D2]) mo.Either[E1, E2],
	operator5 func(mo.Either[E1, E2]) mo.Either[F1, F2],
) mo.Either[F1, F2] {
	return operator5(
		operator4(
			operator3(
				operator2(
					operator1(
						source,
					),
				),
			),
		),
	)
}
func Pipe6[A1 any, A2 any, B1 any, B2 any, C1 any, C2 any, D1 any, D2 any, E1 any, E2 any, F1 any, F2 any, G1 any, G2 any](
	source mo.Either[A1, A2],
	operator1 func(mo.Either[A1, A2]) mo.Either[B1, B2],
	operator2 func(mo.Either[B1, B2]) mo.Either[C1, C2],
	operator3 func(mo.Either[C1, C2]) mo.Either[D1, D2],
	operator4 func(mo.Either[D1, D2]) mo.Either[E1, E2],
	operator5 func(mo.Either[E1, E2]) mo.Either[F1, F2],
	operator6 func(mo.Either[F1, F2]) mo.Either[G1, G2],
) mo.Either[G1, G2] {
	return operator6(
		operator5(
			operator4(
				operator3(
					operator2(
						operator1(
							source,
						),
					),
				),
			),
		),
	)
}

func Pipe7[A1 any, A2 any, B1 any, B2 any, C1 any, C2 any, D1 any, D2 any, E1 any, E2 any, F1 any, F2 any, G1 any, G2 any, H1 any, H2 any](
	source mo.Either[A1, A2],
	operator1 func(mo.Either[A1, A2]) mo.Either[B1, B2],
	operator2 func(mo.Either[B1, B2]) mo.Either[C1, C2],
	operator3 func(mo.Either[C1, C2]) mo.Either[D1, D2],
	operator4 func(mo.Either[D1, D2]) mo.Either[E1, E2],
	operator5 func(mo.Either[E1, E2]) mo.Either[F1, F2],
	operator6 func(mo.Either[F1, F2]) mo.Either[G1, G2],
	operator7 func(mo.Either[G1, G2]) mo.Either[H1, H2],
) mo.Either[H1, H2] {
	return operator7(
		operator6(
			operator5(
				operator4(
					operator3(
						operator2(
							operator1(
								source,
							),
						),
					),
				),
			),
		),
	)
}

func Pipe8[A1 any, A2 any, B1 any, B2 any, C1 any, C2 any, D1 any, D2 any, E1 any, E2 any, F1 any, F2 any, G1 any, G2 any, H1 any, H2 any, I1 any, I2 any](
	source mo.Either[A1, A2],
	operator1 func(mo.Either[A1, A2]) mo.Either[B1, B2],
	operator2 func(mo.Either[B1, B2]) mo.Either[C1, C2],
	operator3 func(mo.Either[C1, C2]) mo.Either[D1, D2],
	operator4 func(mo.Either[D1, D2]) mo.Either[E1, E2],
	operator5 func(mo.Either[E1, E2]) mo.Either[F1, F2],
	operator6 func(mo.Either[F1, F2]) mo.Either[G1, G2],
	operator7 func(mo.Either[G1, G2]) mo.Either[H1, H2],
	operator8 func(mo.Either[H1, H2]) mo.Either[I1, I2],
) mo.Either[I1, I2] {
	return operator8(
		operator7(
			operator6(
				operator5(
					operator4(
						operator3(
							operator2(
								operator1(
									source,
								),
							),
						),
					),
				),
			),
		),
	)
}

func Pipe9[A1 any, A2 any, B1 any, B2 any, C1 any, C2 any, D1 any, D2 any, E1 any, E2 any, F1 any, F2 any, G1 any, G2 any, H1 any, H2 any, I1 any, I2 any, J1 any, J2 any](
	source mo.Either[A1, A2],
	operator1 func(mo.Either[A1, A2]) mo.Either[B1, B2],
	operator2 func(mo.Either[B1, B2]) mo.Either[C1, C2],
	operator3 func(mo.Either[C1, C2]) mo.Either[D1, D2],
	operator4 func(mo.Either[D1, D2]) mo.Either[E1, E2],
	operator5 func(mo.Either[E1, E2]) mo.Either[F1, F2],
	operator6 func(mo.Either[F1, F2]) mo.Either[G1, G2],
	operator7 func(mo.Either[G1, G2]) mo.Either[H1, H2],
	operator8 func(mo.Either[H1, H2]) mo.Either[I1, I2],
	operator9 func(mo.Either[I1, I2]) mo.Either[J1, J2],
) mo.Either[J1, J2] {
	return operator9(
		operator8(
			operator7(
				operator6(
					operator5(
						operator4(
							operator3(
								operator2(
									operator1(
										source,
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

func Pipe10[A1 any, A2 any, B1 any, B2 any, C1 any, C2 any, D1 any, D2 any, E1 any, E2 any, F1 any, F2 any, G1 any, G2 any, H1 any, H2 any, I1 any, I2 any, J1 any, J2 any, K1 any, K2 any](
	source mo.Either[A1, A2],
	operator1 func(mo.Either[A1, A2]) mo.Either[B1, B2],
	operator2 func(mo.Either[B1, B2]) mo.Either[C1, C2],
	operator3 func(mo.Either[C1, C2]) mo.Either[D1, D2],
	operator4 func(mo.Either[D1, D2]) mo.Either[E1, E2],
	operator5 func(mo.Either[E1, E2]) mo.Either[F1, F2],
	operator6 func(mo.Either[F1, F2]) mo.Either[G1, G2],
	operator7 func(mo.Either[G1, G2]) mo.Either[H1, H2],
	operator8 func(mo.Either[H1, H2]) mo.Either[I1, I2],
	operator9 func(mo.Either[I1, I2]) mo.Either[J1, J2],
	operator10 func(mo.Either[J1, J2]) mo.Either[K1, K2],
) mo.Either[K1, K2] {
	return operator10(
		operator9(
			operator8(
				operator7(
					operator6(
						operator5(
							operator4(
								operator3(
									operator2(
										operator1(
											source,
										),
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
		mo.Left[string, error]("42"),
		FlatMapLeft(func(str string) mo.Either[int, error] {
			v, err := strconv.Atoi(str)
			if err != nil {
				mo.Right[int](err)
			}
			return mo.Left[int, error](v)
		}),
		MapLeft[int, error](func(n int) float64 {
			return float64(n)
		}),
		MapLeft[float64, error](func(n float64) string {
			return fmt.Sprintf("%.2f", n)
		}),
	)
	is.Equal(mo.Left[string, error]("42.00"), out)
}
