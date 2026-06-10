package mo

import (
	"testing"
)

func BenchmarkIO(b *testing.B) {
	io := NewIO(func() int { return 42 })

	b.Run("Run", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			sinkInt = io.Run()
		}
	})
}

func BenchmarkState(b *testing.B) {
	state := NewState(func(s int) (string, int) { return "ok", s + 1 })

	b.Run("Run", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_, sinkInt = state.Run(i)
		}
	})

	b.Run("Get", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_, sinkInt = state.Get().Run(i)
		}
	})
}

func BenchmarkDo(b *testing.B) {
	b.Run("Success", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			sinkResultInt = Do(func() int { return 42 })
		}
	})
}
