package either5

import "github.com/samber/mo"

func Pipe1[A1 any, A2 any, A3 any, A4 any, A5 any, B1 any, B2 any, B3 any, B4 any, B5 any](
	source mo.Either5[A1, A2, A3, A4, A5],
	operator1 func(mo.Either5[A1, A2, A3, A4, A5]) mo.Either5[B1, B2, B3, B4, B5],
) mo.Either5[B1, B2, B3, B4, B5] {
	return operator1(
		source,
	)
}

func Pipe2[A1 any, A2 any, A3 any, A4 any, A5 any, B1 any, B2 any, B3 any, B4 any, B5 any, C1 any, C2 any, C3 any, C4 any, C5 any](
	source mo.Either5[A1, A2, A3, A4, A5],
	operator1 func(mo.Either5[A1, A2, A3, A4, A5]) mo.Either5[B1, B2, B3, B4, B5],
	operator2 func(mo.Either5[B1, B2, B3, B4, B5]) mo.Either5[C1, C2, C3, C4, C5],
) mo.Either5[C1, C2, C3, C4, C5] {
	return operator2(
		operator1(
			source,
		),
	)
}

func Pipe3[A1 any, A2 any, A3 any, A4 any, A5 any, B1 any, B2 any, B3 any, B4 any, B5 any, C1 any, C2 any, C3 any, C4 any, C5 any, D1 any, D2 any, D3 any, D4 any, D5 any](
	source mo.Either5[A1, A2, A3, A4, A5],
	operator1 func(mo.Either5[A1, A2, A3, A4, A5]) mo.Either5[B1, B2, B3, B4, B5],
	operator2 func(mo.Either5[B1, B2, B3, B4, B5]) mo.Either5[C1, C2, C3, C4, C5],
	operator3 func(mo.Either5[C1, C2, C3, C4, C5]) mo.Either5[D1, D2, D3, D4, D5],
) mo.Either5[D1, D2, D3, D4, D5] {
	return operator3(
		operator2(
			operator1(
				source,
			),
		),
	)
}

func Pipe4[A1 any, A2 any, A3 any, A4 any, A5 any, B1 any, B2 any, B3 any, B4 any, B5 any, C1 any, C2 any, C3 any, C4 any, C5 any, D1 any, D2 any, D3 any, D4 any, D5 any, E1 any, E2 any, E3 any, E4 any, E5 any](
	source mo.Either5[A1, A2, A3, A4, A5],
	operator1 func(mo.Either5[A1, A2, A3, A4, A5]) mo.Either5[B1, B2, B3, B4, B5],
	operator2 func(mo.Either5[B1, B2, B3, B4, B5]) mo.Either5[C1, C2, C3, C4, C5],
	operator3 func(mo.Either5[C1, C2, C3, C4, C5]) mo.Either5[D1, D2, D3, D4, D5],
	operator4 func(mo.Either5[D1, D2, D3, D4, D5]) mo.Either5[E1, E2, E3, E4, E5],
) mo.Either5[E1, E2, E3, E4, E5] {
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

func Pipe5[A1 any, A2 any, A3 any, A4 any, A5 any, B1 any, B2 any, B3 any, B4 any, B5 any, C1 any, C2 any, C3 any, C4 any, C5 any, D1 any, D2 any, D3 any, D4 any, D5 any, E1 any, E2 any, E3 any, E4 any, E5 any, F1 any, F2 any, F3 any, F4 any, F5 any](
	source mo.Either5[A1, A2, A3, A4, A5],
	operator1 func(mo.Either5[A1, A2, A3, A4, A5]) mo.Either5[B1, B2, B3, B4, B5],
	operator2 func(mo.Either5[B1, B2, B3, B4, B5]) mo.Either5[C1, C2, C3, C4, C5],
	operator3 func(mo.Either5[C1, C2, C3, C4, C5]) mo.Either5[D1, D2, D3, D4, D5],
	operator4 func(mo.Either5[D1, D2, D3, D4, D5]) mo.Either5[E1, E2, E3, E4, E5],
	operator5 func(mo.Either5[E1, E2, E3, E4, E5]) mo.Either5[F1, F2, F3, F4, F5],
) mo.Either5[F1, F2, F3, F4, F5] {
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

func Pipe6[A1 any, A2 any, A3 any, A4 any, A5 any, B1 any, B2 any, B3 any, B4 any, B5 any, C1 any, C2 any, C3 any, C4 any, C5 any, D1 any, D2 any, D3 any, D4 any, D5 any, E1 any, E2 any, E3 any, E4 any, E5 any, F1 any, F2 any, F3 any, F4 any, F5 any, G1 any, G2 any, G3 any, G4 any, G5 any](
	source mo.Either5[A1, A2, A3, A4, A5],
	operator1 func(mo.Either5[A1, A2, A3, A4, A5]) mo.Either5[B1, B2, B3, B4, B5],
	operator2 func(mo.Either5[B1, B2, B3, B4, B5]) mo.Either5[C1, C2, C3, C4, C5],
	operator3 func(mo.Either5[C1, C2, C3, C4, C5]) mo.Either5[D1, D2, D3, D4, D5],
	operator4 func(mo.Either5[D1, D2, D3, D4, D5]) mo.Either5[E1, E2, E3, E4, E5],
	operator5 func(mo.Either5[E1, E2, E3, E4, E5]) mo.Either5[F1, F2, F3, F4, F5],
	operator6 func(mo.Either5[F1, F2, F3, F4, F5]) mo.Either5[G1, G2, G3, G4, G5],
) mo.Either5[G1, G2, G3, G4, G5] {
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

func Pipe7[A1 any, A2 any, A3 any, A4 any, A5 any, B1 any, B2 any, B3 any, B4 any, B5 any, C1 any, C2 any, C3 any, C4 any, C5 any, D1 any, D2 any, D3 any, D4 any, D5 any, E1 any, E2 any, E3 any, E4 any, E5 any, F1 any, F2 any, F3 any, F4 any, F5 any, G1 any, G2 any, G3 any, G4 any, G5 any, H1 any, H2 any, H3 any, H4 any, H5 any](
	source mo.Either5[A1, A2, A3, A4, A5],
	operator1 func(mo.Either5[A1, A2, A3, A4, A5]) mo.Either5[B1, B2, B3, B4, B5],
	operator2 func(mo.Either5[B1, B2, B3, B4, B5]) mo.Either5[C1, C2, C3, C4, C5],
	operator3 func(mo.Either5[C1, C2, C3, C4, C5]) mo.Either5[D1, D2, D3, D4, D5],
	operator4 func(mo.Either5[D1, D2, D3, D4, D5]) mo.Either5[E1, E2, E3, E4, E5],
	operator5 func(mo.Either5[E1, E2, E3, E4, E5]) mo.Either5[F1, F2, F3, F4, F5],
	operator6 func(mo.Either5[F1, F2, F3, F4, F5]) mo.Either5[G1, G2, G3, G4, G5],
	operator7 func(mo.Either5[G1, G2, G3, G4, G5]) mo.Either5[H1, H2, H3, H4, H5],
) mo.Either5[H1, H2, H3, H4, H5] {
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

func Pipe8[A1 any, A2 any, A3 any, A4 any, A5 any, B1 any, B2 any, B3 any, B4 any, B5 any, C1 any, C2 any, C3 any, C4 any, C5 any, D1 any, D2 any, D3 any, D4 any, D5 any, E1 any, E2 any, E3 any, E4 any, E5 any, F1 any, F2 any, F3 any, F4 any, F5 any, G1 any, G2 any, G3 any, G4 any, G5 any, H1 any, H2 any, H3 any, H4 any, H5 any, I1 any, I2 any, I3 any, I4 any, I5 any, J1 any, J2 any, J3 any, J4 any, J5 any](
	source mo.Either5[A1, A2, A3, A4, A5],
	operator1 func(mo.Either5[A1, A2, A3, A4, A5]) mo.Either5[B1, B2, B3, B4, B5],
	operator2 func(mo.Either5[B1, B2, B3, B4, B5]) mo.Either5[C1, C2, C3, C4, C5],
	operator3 func(mo.Either5[C1, C2, C3, C4, C5]) mo.Either5[D1, D2, D3, D4, D5],
	operator4 func(mo.Either5[D1, D2, D3, D4, D5]) mo.Either5[E1, E2, E3, E4, E5],
	operator5 func(mo.Either5[E1, E2, E3, E4, E5]) mo.Either5[F1, F2, F3, F4, F5],
	operator6 func(mo.Either5[F1, F2, F3, F4, F5]) mo.Either5[G1, G2, G3, G4, G5],
	operator7 func(mo.Either5[G1, G2, G3, G4, G5]) mo.Either5[H1, H2, H3, H4, H5],
	operator8 func(mo.Either5[H1, H2, H3, H4, H5]) mo.Either5[I1, I2, I3, I4, I5],
) mo.Either5[I1, I2, I3, I4, I5] {
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

func Pipe9[A1 any, A2 any, A3 any, A4 any, A5 any, B1 any, B2 any, B3 any, B4 any, B5 any, C1 any, C2 any, C3 any, C4 any, C5 any, D1 any, D2 any, D3 any, D4 any, D5 any, E1 any, E2 any, E3 any, E4 any, E5 any, F1 any, F2 any, F3 any, F4 any, F5 any, G1 any, G2 any, G3 any, G4 any, G5 any, H1 any, H2 any, H3 any, H4 any, H5 any, I1 any, I2 any, I3 any, I4 any, I5 any, J1 any, J2 any, J3 any, J4 any, J5 any, K1 any, K2 any, K3 any, K4 any, K5 any](
	source mo.Either5[A1, A2, A3, A4, A5],
	operator1 func(mo.Either5[A1, A2, A3, A4, A5]) mo.Either5[B1, B2, B3, B4, B5],
	operator2 func(mo.Either5[B1, B2, B3, B4, B5]) mo.Either5[C1, C2, C3, C4, C5],
	operator3 func(mo.Either5[C1, C2, C3, C4, C5]) mo.Either5[D1, D2, D3, D4, D5],
	operator4 func(mo.Either5[D1, D2, D3, D4, D5]) mo.Either5[E1, E2, E3, E4, E5],
	operator5 func(mo.Either5[E1, E2, E3, E4, E5]) mo.Either5[F1, F2, F3, F4, F5],
	operator6 func(mo.Either5[F1, F2, F3, F4, F5]) mo.Either5[G1, G2, G3, G4, G5],
	operator7 func(mo.Either5[G1, G2, G3, G4, G5]) mo.Either5[H1, H2, H3, H4, H5],
	operator8 func(mo.Either5[H1, H2, H3, H4, H5]) mo.Either5[I1, I2, I3, I4, I5],
	operator9 func(mo.Either5[I1, I2, I3, I4, I5]) mo.Either5[J1, J2, J3, J4, J5],
) mo.Either5[J1, J2, J3, J4, J5] {
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

func Pipe10[A1 any, A2 any, A3 any, A4 any, A5 any, B1 any, B2 any, B3 any, B4 any, B5 any, C1 any, C2 any, C3 any, C4 any, C5 any, D1 any, D2 any, D3 any, D4 any, D5 any, E1 any, E2 any, E3 any, E4 any, E5 any, F1 any, F2 any, F3 any, F4 any, F5 any, G1 any, G2 any, G3 any, G4 any, G5 any, H1 any, H2 any, H3 any, H4 any, H5 any, I1 any, I2 any, I3 any, I4 any, I5 any, J1 any, J2 any, J3 any, J4 any, J5 any, K1 any, K2 any, K3 any, K4 any, K5 any, L1 any, L2 any, L3 any, L4 any, L5 any](
	source mo.Either5[A1, A2, A3, A4, A5],
	operator1 func(mo.Either5[A1, A2, A3, A4, A5]) mo.Either5[B1, B2, B3, B4, B5],
	operator2 func(mo.Either5[B1, B2, B3, B4, B5]) mo.Either5[C1, C2, C3, C4, C5],
	operator3 func(mo.Either5[C1, C2, C3, C4, C5]) mo.Either5[D1, D2, D3, D4, D5],
	operator4 func(mo.Either5[D1, D2, D3, D4, D5]) mo.Either5[E1, E2, E3, E4, E5],
	operator5 func(mo.Either5[E1, E2, E3, E4, E5]) mo.Either5[F1, F2, F3, F4, F5],
	operator6 func(mo.Either5[F1, F2, F3, F4, F5]) mo.Either5[G1, G2, G3, G4, G5],
	operator7 func(mo.Either5[G1, G2, G3, G4, G5]) mo.Either5[H1, H2, H3, H4, H5],
	operator8 func(mo.Either5[H1, H2, H3, H4, H5]) mo.Either5[I1, I2, I3, I4, I5],
	operator9 func(mo.Either5[I1, I2, I3, I4, I5]) mo.Either5[J1, J2, J3, J4, J5],
	operator10 func(mo.Either5[J1, J2, J3, J4, J5]) mo.Either5[K1, K2, K3, K4, K5],
) mo.Either5[K1, K2, K3, K4, K5] {
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
