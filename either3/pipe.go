package either3

import "github.com/samber/mo"

func Pipe1[A1 any, A2 any, A3 any, B1 any, B2 any, B3 any](
	source mo.Either3[A1, A2, A3],
	operator1 func(mo.Either3[A1, A2, A3]) mo.Either3[B1, B2, B3],
) mo.Either3[B1, B2, B3] {
	return operator1(
		source,
	)
}

func Pipe2[A1 any, A2 any, A3 any, B1 any, B2 any, B3 any, C1 any, C2 any, C3 any](
	source mo.Either3[A1, A2, A3],
	operator1 func(mo.Either3[A1, A2, A3]) mo.Either3[B1, B2, B3],
	operator2 func(mo.Either3[B1, B2, B3]) mo.Either3[C1, C2, C3],
) mo.Either3[C1, C2, C3] {
	return operator2(
		operator1(
			source,
		),
	)
}

func Pipe3[A1 any, A2 any, A3 any, B1 any, B2 any, B3 any, C1 any, C2 any, C3 any, D1 any, D2 any, D3 any](
	source mo.Either3[A1, A2, A3],
	operator1 func(mo.Either3[A1, A2, A3]) mo.Either3[B1, B2, B3],
	operator2 func(mo.Either3[B1, B2, B3]) mo.Either3[C1, C2, C3],
	operator3 func(mo.Either3[C1, C2, C3]) mo.Either3[D1, D2, D3],
) mo.Either3[D1, D2, D3] {
	return operator3(
		operator2(
			operator1(
				source,
			),
		),
	)
}

func Pipe4[A1 any, A2 any, A3 any, B1 any, B2 any, B3 any, C1 any, C2 any, C3 any, D1 any, D2 any, D3 any, E1 any, E2 any, E3 any](
	source mo.Either3[A1, A2, A3],
	operator1 func(mo.Either3[A1, A2, A3]) mo.Either3[B1, B2, B3],
	operator2 func(mo.Either3[B1, B2, B3]) mo.Either3[C1, C2, C3],
	operator3 func(mo.Either3[C1, C2, C3]) mo.Either3[D1, D2, D3],
	operator4 func(mo.Either3[D1, D2, D3]) mo.Either3[E1, E2, E3],
) mo.Either3[E1, E2, E3] {
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

func Pipe5[A1 any, A2 any, A3 any, B1 any, B2 any, B3 any, C1 any, C2 any, C3 any, D1 any, D2 any, D3 any, E1 any, E2 any, E3 any, F1 any, F2 any, F3 any](
	source mo.Either3[A1, A2, A3],
	operator1 func(mo.Either3[A1, A2, A3]) mo.Either3[B1, B2, B3],
	operator2 func(mo.Either3[B1, B2, B3]) mo.Either3[C1, C2, C3],
	operator3 func(mo.Either3[C1, C2, C3]) mo.Either3[D1, D2, D3],
	operator4 func(mo.Either3[D1, D2, D3]) mo.Either3[E1, E2, E3],
	operator5 func(mo.Either3[E1, E2, E3]) mo.Either3[F1, F2, F3],
) mo.Either3[F1, F2, F3] {
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

func Pipe6[A1 any, A2 any, A3 any, B1 any, B2 any, B3 any, C1 any, C2 any, C3 any, D1 any, D2 any, D3 any, E1 any, E2 any, E3 any, F1 any, F2 any, F3 any, G1 any, G2 any, G3 any](
	source mo.Either3[A1, A2, A3],
	operator1 func(mo.Either3[A1, A2, A3]) mo.Either3[B1, B2, B3],
	operator2 func(mo.Either3[B1, B2, B3]) mo.Either3[C1, C2, C3],
	operator3 func(mo.Either3[C1, C2, C3]) mo.Either3[D1, D2, D3],
	operator4 func(mo.Either3[D1, D2, D3]) mo.Either3[E1, E2, E3],
	operator5 func(mo.Either3[E1, E2, E3]) mo.Either3[F1, F2, F3],
	operator6 func(mo.Either3[F1, F2, F3]) mo.Either3[G1, G2, G3],
) mo.Either3[G1, G2, G3] {
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

func Pipe7[A1 any, A2 any, A3 any, B1 any, B2 any, B3 any, C1 any, C2 any, C3 any, D1 any, D2 any, D3 any, E1 any, E2 any, E3 any, F1 any, F2 any, F3 any, G1 any, G2 any, G3 any, H1 any, H2 any, H3 any](
	source mo.Either3[A1, A2, A3],
	operator1 func(mo.Either3[A1, A2, A3]) mo.Either3[B1, B2, B3],
	operator2 func(mo.Either3[B1, B2, B3]) mo.Either3[C1, C2, C3],
	operator3 func(mo.Either3[C1, C2, C3]) mo.Either3[D1, D2, D3],
	operator4 func(mo.Either3[D1, D2, D3]) mo.Either3[E1, E2, E3],
	operator5 func(mo.Either3[E1, E2, E3]) mo.Either3[F1, F2, F3],
	operator6 func(mo.Either3[F1, F2, F3]) mo.Either3[G1, G2, G3],
	operator7 func(mo.Either3[G1, G2, G3]) mo.Either3[H1, H2, H3],
) mo.Either3[H1, H2, H3] {
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

func Pipe8[A1 any, A2 any, A3 any, B1 any, B2 any, B3 any, C1 any, C2 any, C3 any, D1 any, D2 any, D3 any, E1 any, E2 any, E3 any, F1 any, F2 any, F3 any, G1 any, G2 any, G3 any, H1 any, H2 any, H3 any, I1 any, I2 any, I3 any](
	source mo.Either3[A1, A2, A3],
	operator1 func(mo.Either3[A1, A2, A3]) mo.Either3[B1, B2, B3],
	operator2 func(mo.Either3[B1, B2, B3]) mo.Either3[C1, C2, C3],
	operator3 func(mo.Either3[C1, C2, C3]) mo.Either3[D1, D2, D3],
	operator4 func(mo.Either3[D1, D2, D3]) mo.Either3[E1, E2, E3],
	operator5 func(mo.Either3[E1, E2, E3]) mo.Either3[F1, F2, F3],
	operator6 func(mo.Either3[F1, F2, F3]) mo.Either3[G1, G2, G3],
	operator7 func(mo.Either3[G1, G2, G3]) mo.Either3[H1, H2, H3],
	operator8 func(mo.Either3[H1, H2, H3]) mo.Either3[I1, I2, I3],
) mo.Either3[I1, I2, I3] {
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

func Pipe9[A1 any, A2 any, A3 any, B1 any, B2 any, B3 any, C1 any, C2 any, C3 any, D1 any, D2 any, D3 any, E1 any, E2 any, E3 any, F1 any, F2 any, F3 any, G1 any, G2 any, G3 any, H1 any, H2 any, H3 any, I1 any, I2 any, I3 any, J1 any, J2 any, J3 any](
	source mo.Either3[A1, A2, A3],
	operator1 func(mo.Either3[A1, A2, A3]) mo.Either3[B1, B2, B3],
	operator2 func(mo.Either3[B1, B2, B3]) mo.Either3[C1, C2, C3],
	operator3 func(mo.Either3[C1, C2, C3]) mo.Either3[D1, D2, D3],
	operator4 func(mo.Either3[D1, D2, D3]) mo.Either3[E1, E2, E3],
	operator5 func(mo.Either3[E1, E2, E3]) mo.Either3[F1, F2, F3],
	operator6 func(mo.Either3[F1, F2, F3]) mo.Either3[G1, G2, G3],
	operator7 func(mo.Either3[G1, G2, G3]) mo.Either3[H1, H2, H3],
	operator8 func(mo.Either3[H1, H2, H3]) mo.Either3[I1, I2, I3],
	operator9 func(mo.Either3[I1, I2, I3]) mo.Either3[J1, J2, J3],
) mo.Either3[J1, J2, J3] {
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

func Pipe10[A1 any, A2 any, A3 any, B1 any, B2 any, B3 any, C1 any, C2 any, C3 any, D1 any, D2 any, D3 any, E1 any, E2 any, E3 any, F1 any, F2 any, F3 any, G1 any, G2 any, G3 any, H1 any, H2 any, H3 any, I1 any, I2 any, I3 any, J1 any, J2 any, J3 any, K1 any, K2 any, K3 any](
	source mo.Either3[A1, A2, A3],
	operator1 func(mo.Either3[A1, A2, A3]) mo.Either3[B1, B2, B3],
	operator2 func(mo.Either3[B1, B2, B3]) mo.Either3[C1, C2, C3],
	operator3 func(mo.Either3[C1, C2, C3]) mo.Either3[D1, D2, D3],
	operator4 func(mo.Either3[D1, D2, D3]) mo.Either3[E1, E2, E3],
	operator5 func(mo.Either3[E1, E2, E3]) mo.Either3[F1, F2, F3],
	operator6 func(mo.Either3[F1, F2, F3]) mo.Either3[G1, G2, G3],
	operator7 func(mo.Either3[G1, G2, G3]) mo.Either3[H1, H2, H3],
	operator8 func(mo.Either3[H1, H2, H3]) mo.Either3[I1, I2, I3],
	operator9 func(mo.Either3[I1, I2, I3]) mo.Either3[J1, J2, J3],
	operator10 func(mo.Either3[J1, J2, J3]) mo.Either3[K1, K2, K3],
) mo.Either3[K1, K2, K3] {
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
