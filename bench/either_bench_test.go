package bench

import (
	"testing"

	mo "github.com/samber/mo"
)

var (
	sinkEither  mo.Either[error, int]
	sinkEither2 mo.Either[string, int]
)

func BenchmarkEitherConstructors(b *testing.B) {
	b.Run("Left", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			sinkEither2 = mo.Left[string, int]("hello")
		}
	})

	b.Run("Right", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			sinkEither2 = mo.Right[string, int](i)
		}
	})
}

func BenchmarkEitherAccessors(b *testing.B) {
	right := mo.Right[string, int](42)

	b.Run("Match", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			sinkEither2 = right.Match(
				func(l string) mo.Either[string, int] { return mo.Left[string, int](l) },
				func(r int) mo.Either[string, int] { return mo.Right[string, int](r * 2) },
			)
		}
	})

	b.Run("MapRight", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			sinkEither2 = right.MapRight(func(r int) mo.Either[string, int] { return mo.Right[string, int](r * 2) })
		}
	})

	b.Run("Swap", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = right.Swap()
		}
	})
}

func BenchmarkEitherBinary(b *testing.B) {
	right := mo.Right[string, int](42)
	encoded, err := right.MarshalBinary()
	if err != nil {
		b.Fatal(err)
	}

	b.Run("MarshalBinary", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			sinkBytes, sinkErr = right.MarshalBinary()
		}
	})

	b.Run("UnmarshalBinary", func(b *testing.B) {
		var e mo.Either[string, int]
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			sinkErr = e.UnmarshalBinary(encoded)
		}
	})
}

func BenchmarkEitherNBinary(b *testing.B) {
	b.Run("Either3/MarshalBinary", func(b *testing.B) {
		e := mo.NewEither3Arg2[string, int, bool](42)
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			sinkBytes, sinkErr = e.MarshalBinary()
		}
	})

	b.Run("Either4/MarshalBinary", func(b *testing.B) {
		e := mo.NewEither4Arg2[string, int, bool, float64](42)
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			sinkBytes, sinkErr = e.MarshalBinary()
		}
	})

	b.Run("Either5/MarshalBinary", func(b *testing.B) {
		e := mo.NewEither5Arg2[string, int, bool, float64, uint](42)
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			sinkBytes, sinkErr = e.MarshalBinary()
		}
	})
}
