package mo

import (
	"testing"
)

var (
	sinkEither  Either[error, int]
	sinkEither2 Either[string, int]
)

func BenchmarkEitherConstructors(b *testing.B) {
	b.Run("Left", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			sinkEither2 = Left[string, int]("hello")
		}
	})

	b.Run("Right", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			sinkEither2 = Right[string, int](i)
		}
	})
}

func BenchmarkEitherAccessors(b *testing.B) {
	right := Right[string, int](42)

	b.Run("Match", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			sinkEither2 = right.Match(
				func(l string) Either[string, int] { return Left[string, int](l) },
				func(r int) Either[string, int] { return Right[string, int](r * 2) },
			)
		}
	})

	b.Run("MapRight", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			sinkEither2 = right.MapRight(func(r int) Either[string, int] { return Right[string, int](r * 2) })
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
	right := Right[string, int](42)
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
		var e Either[string, int]
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			sinkErr = e.UnmarshalBinary(encoded)
		}
	})
}

func BenchmarkEitherNBinary(b *testing.B) {
	b.Run("Either3/MarshalBinary", func(b *testing.B) {
		e := NewEither3Arg2[string, int, bool](42)
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			sinkBytes, sinkErr = e.MarshalBinary()
		}
	})

	b.Run("Either4/MarshalBinary", func(b *testing.B) {
		e := NewEither4Arg2[string, int, bool, float64](42)
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			sinkBytes, sinkErr = e.MarshalBinary()
		}
	})

	b.Run("Either5/MarshalBinary", func(b *testing.B) {
		e := NewEither5Arg2[string, int, bool, float64, uint](42)
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			sinkBytes, sinkErr = e.MarshalBinary()
		}
	})
}
