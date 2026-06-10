package mo

import (
	"testing"
)

func BenchmarkFuture(b *testing.B) {
	b.Run("NewFutureCollect", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			f := NewFuture(func(resolve func(int), reject func(error)) {
				resolve(42)
			})
			sinkInt, sinkErr = f.Collect()
		}
	})

	b.Run("ThenChainCollect", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			f := NewFuture(func(resolve func(int), reject func(error)) {
				resolve(42)
			}).
				Then(func(v int) (int, error) { return v + 1, nil }).
				Then(func(v int) (int, error) { return v * 2, nil }).
				Then(func(v int) (int, error) { return v - 3, nil })
			sinkInt, sinkErr = f.Collect()
		}
	})
}

func BenchmarkTask(b *testing.B) {
	task := NewTask(func() *Future[int] {
		return NewFuture(func(resolve func(int), reject func(error)) {
			resolve(42)
		})
	})

	b.Run("RunCollect", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			sinkInt, sinkErr = task.Run().Collect()
		}
	})
}
