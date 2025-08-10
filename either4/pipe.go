package either4

import "github.com/samber/mo"

func Pipe1[A1 any, A2 any, A3 any, A4 any, B1 any, B2 any, B3 any, B4 any](
	source mo.Either4[A1, A2, A3, A4],
	operator1 func(mo.Either4[A1, A2, A3, A4]) mo.Either4[B1, B2, B3, B4],
) mo.Either4[B1, B2, B3, B4] {
	return operator1(
		source,
	)
}

func Pipe2[A1 any, A2 any, A3 any, A4 any, B1 any, B2 any, B3 any, B4 any, C1 any, C2 any, C3 any, C4 any](
	source mo.Either4[A1, A2, A3, A4],
	operator1 func(mo.Either4[A1, A2, A3, A4]) mo.Either4[B1, B2, B3, B4],
	operator2 func(mo.Either4[B1, B2, B3, B4]) mo.Either4[C1, C2, C3, C4],
) mo.Either4[C1, C2, C3, C4] {
	return operator2(
		operator1(
			source,
		),
	)
}

func Pipe3[A1 any, A2 any, A3 any, A4 any, B1 any, B2 any, B3 any, B4 any, C1 any, C2 any, C3 any, C4 any, D1 any, D2 any, D3 any, D4 any](
	source mo.Either4[A1, A2, A3, A4],
	operator1 func(mo.Either4[A1, A2, A3, A4]) mo.Either4[B1, B2, B3, B4],
	operator2 func(mo.Either4[B1, B2, B3, B4]) mo.Either4[C1, C2, C3, C4],
	operator3 func(mo.Either4[C1, C2, C3, C4]) mo.Either4[D1, D2, D3, D4],
) mo.Either4[D1, D2, D3, D4] {
	return operator3(
		operator2(
			operator1(
				source,
			),
		),
	)
}

func Pipe4[A1 any, A2 any, A3 any, A4 any, B1 any, B2 any, B3 any, B4 any, C1 any, C2 any, C3 any, C4 any, D1 any, D2 any, D3 any, D4 any, E1 any, E2 any, E3 any, E4 any](
	source mo.Either4[A1, A2, A3, A4],
	operator1 func(mo.Either4[A1, A2, A3, A4]) mo.Either4[B1, B2, B3, B4],
	operator2 func(mo.Either4[B1, B2, B3, B4]) mo.Either4[C1, C2, C3, C4],
	operator3 func(mo.Either4[C1, C2, C3, C4]) mo.Either4[D1, D2, D3, D4],
	operator4 func(mo.Either4[D1, D2, D3, D4]) mo.Either4[E1, E2, E3, E4],
) mo.Either4[E1, E2, E3, E4] {
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

func Pipe5[A1 any, A2 any, A3 any, A4 any, B1 any, B2 any, B3 any, B4 any, C1 any, C2 any, C3 any, C4 any, D1 any, D2 any, D3 any, D4 any, E1 any, E2 any, E3 any, E4 any, F1 any, F2 any, F3 any, F4 any](
	source mo.Either4[A1, A2, A3, A4],
	operator1 func(mo.Either4[A1, A2, A3, A4]) mo.Either4[B1, B2, B3, B4],
	operator2 func(mo.Either4[B1, B2, B3, B4]) mo.Either4[C1, C2, C3, C4],
	operator3 func(mo.Either4[C1, C2, C3, C4]) mo.Either4[D1, D2, D3, D4],
	operator4 func(mo.Either4[D1, D2, D3, D4]) mo.Either4[E1, E2, E3, E4],
	operator5 func(mo.Either4[E1, E2, E3, E4]) mo.Either4[F1, F2, F3, F4],
) mo.Either4[F1, F2, F3, F4] {
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

func Pipe6[A1 any, A2 any, A3 any, A4 any, B1 any, B2 any, B3 any, B4 any, C1 any, C2 any, C3 any, C4 any, D1 any, D2 any, D3 any, D4 any, E1 any, E2 any, E3 any, E4 any, F1 any, F2 any, F3 any, F4 any, G1 any, G2 any, G3 any, G4 any](
	source mo.Either4[A1, A2, A3, A4],
	operator1 func(mo.Either4[A1, A2, A3, A4]) mo.Either4[B1, B2, B3, B4],
	operator2 func(mo.Either4[B1, B2, B3, B4]) mo.Either4[C1, C2, C3, C4],
	operator3 func(mo.Either4[C1, C2, C3, C4]) mo.Either4[D1, D2, D3, D4],
	operator4 func(mo.Either4[D1, D2, D3, D4]) mo.Either4[E1, E2, E3, E4],
	operator5 func(mo.Either4[E1, E2, E3, E4]) mo.Either4[F1, F2, F3, F4],
	operator6 func(mo.Either4[F1, F2, F3, F4]) mo.Either4[G1, G2, G3, G4],
) mo.Either4[G1, G2, G3, G4] {
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

func Pipe7[A1 any, A2 any, A3 any, A4 any, B1 any, B2 any, B3 any, B4 any, C1 any, C2 any, C3 any, C4 any, D1 any, D2 any, D3 any, D4 any, E1 any, E2 any, E3 any, E4 any, F1 any, F2 any, F3 any, F4 any, G1 any, G2 any, G3 any, G4 any, H1 any, H2 any, H3 any, H4 any](
	source mo.Either4[A1, A2, A3, A4],
	operator1 func(mo.Either4[A1, A2, A3, A4]) mo.Either4[B1, B2, B3, B4],
	operator2 func(mo.Either4[B1, B2, B3, B4]) mo.Either4[C1, C2, C3, C4],
	operator3 func(mo.Either4[C1, C2, C3, C4]) mo.Either4[D1, D2, D3, D4],
	operator4 func(mo.Either4[D1, D2, D3, D4]) mo.Either4[E1, E2, E3, E4],
	operator5 func(mo.Either4[E1, E2, E3, E4]) mo.Either4[F1, F2, F3, F4],
	operator6 func(mo.Either4[F1, F2, F3, F4]) mo.Either4[G1, G2, G3, G4],
	operator7 func(mo.Either4[G1, G2, G3, G4]) mo.Either4[H1, H2, H3, H4],
) mo.Either4[H1, H2, H3, H4] {
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

func Pipe8[A1 any, A2 any, A3 any, A4 any, B1 any, B2 any, B3 any, B4 any, C1 any, C2 any, C3 any, C4 any, D1 any, D2 any, D3 any, D4 any, E1 any, E2 any, E3 any, E4 any, F1 any, F2 any, F3 any, F4 any, G1 any, G2 any, G3 any, G4 any, H1 any, H2 any, H3 any, H4 any, I1 any, I2 any, I3 any, I4 any](
	source mo.Either4[A1, A2, A3, A4],
	operator1 func(mo.Either4[A1, A2, A3, A4]) mo.Either4[B1, B2, B3, B4],
	operator2 func(mo.Either4[B1, B2, B3, B4]) mo.Either4[C1, C2, C3, C4],
	operator3 func(mo.Either4[C1, C2, C3, C4]) mo.Either4[D1, D2, D3, D4],
	operator4 func(mo.Either4[D1, D2, D3, D4]) mo.Either4[E1, E2, E3, E4],
	operator5 func(mo.Either4[E1, E2, E3, E4]) mo.Either4[F1, F2, F3, F4],
	operator6 func(mo.Either4[F1, F2, F3, F4]) mo.Either4[G1, G2, G3, G4],
	operator7 func(mo.Either4[G1, G2, G3, G4]) mo.Either4[H1, H2, H3, H4],
	operator8 func(mo.Either4[H1, H2, H3, H4]) mo.Either4[I1, I2, I3, I4],
) mo.Either4[I1, I2, I3, I4] {
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

func Pipe9[A1 any, A2 any, A3 any, A4 any, B1 any, B2 any, B3 any, B4 any, C1 any, C2 any, C3 any, C4 any, D1 any, D2 any, D3 any, D4 any, E1 any, E2 any, E3 any, E4 any, F1 any, F2 any, F3 any, F4 any, G1 any, G2 any, G3 any, G4 any, H1 any, H2 any, H3 any, H4 any, I1 any, I2 any, I3 any, I4 any, J1 any, J2 any, J3 any, J4 any](
	source mo.Either4[A1, A2, A3, A4],
	operator1 func(mo.Either4[A1, A2, A3, A4]) mo.Either4[B1, B2, B3, B4],
	operator2 func(mo.Either4[B1, B2, B3, B4]) mo.Either4[C1, C2, C3, C4],
	operator3 func(mo.Either4[C1, C2, C3, C4]) mo.Either4[D1, D2, D3, D4],
	operator4 func(mo.Either4[D1, D2, D3, D4]) mo.Either4[E1, E2, E3, E4],
	operator5 func(mo.Either4[E1, E2, E3, E4]) mo.Either4[F1, F2, F3, F4],
	operator6 func(mo.Either4[F1, F2, F3, F4]) mo.Either4[G1, G2, G3, G4],
	operator7 func(mo.Either4[G1, G2, G3, G4]) mo.Either4[H1, H2, H3, H4],
	operator8 func(mo.Either4[H1, H2, H3, H4]) mo.Either4[I1, I2, I3, I4],
	operator9 func(mo.Either4[I1, I2, I3, I4]) mo.Either4[J1, J2, J3, J4],
) mo.Either4[J1, J2, J3, J4] {
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

func Pipe10[A1 any, A2 any, A3 any, A4 any, B1 any, B2 any, B3 any, B4 any, C1 any, C2 any, C3 any, C4 any, D1 any, D2 any, D3 any, D4 any, E1 any, E2 any, E3 any, E4 any, F1 any, F2 any, F3 any, F4 any, G1 any, G2 any, G3 any, G4 any, H1 any, H2 any, H3 any, H4 any, I1 any, I2 any, I3 any, I4 any, J1 any, J2 any, J3 any, J4 any, K1 any, K2 any, K3 any, K4 any](
	source mo.Either4[A1, A2, A3, A4],
	operator1 func(mo.Either4[A1, A2, A3, A4]) mo.Either4[B1, B2, B3, B4],
	operator2 func(mo.Either4[B1, B2, B3, B4]) mo.Either4[C1, C2, C3, C4],
	operator3 func(mo.Either4[C1, C2, C3, C4]) mo.Either4[D1, D2, D3, D4],
	operator4 func(mo.Either4[D1, D2, D3, D4]) mo.Either4[E1, E2, E3, E4],
	operator5 func(mo.Either4[E1, E2, E3, E4]) mo.Either4[F1, F2, F3, F4],
	operator6 func(mo.Either4[F1, F2, F3, F4]) mo.Either4[G1, G2, G3, G4],
	operator7 func(mo.Either4[G1, G2, G3, G4]) mo.Either4[H1, H2, H3, H4],
	operator8 func(mo.Either4[H1, H2, H3, H4]) mo.Either4[I1, I2, I3, I4],
	operator9 func(mo.Either4[I1, I2, I3, I4]) mo.Either4[J1, J2, J3, J4],
	operator10 func(mo.Either4[J1, J2, J3, J4]) mo.Either4[K1, K2, K3, K4],
) mo.Either4[K1, K2, K3, K4] {
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
