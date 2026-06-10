package mo

import (
	"errors"
	"testing"
)

var (
	sinkResultInt    Result[int]
	sinkResultString Result[string]
)

func BenchmarkResultConstructors(b *testing.B) {
	err := errors.New("an error")

	b.Run("Ok", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			sinkResultInt = Ok(i)
		}
	})

	b.Run("Err", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			sinkResultInt = Err[int](err)
		}
	})

	b.Run("TupleToResult", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			sinkResultInt = TupleToResult(i, nil)
		}
	})
}

func BenchmarkResultAccessors(b *testing.B) {
	ok := Ok(42)
	ko := Err[int](errors.New("an error"))

	b.Run("Get", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			sinkInt, sinkErr = ok.Get()
		}
	})

	b.Run("Match", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			sinkResultInt = ok.Match(
				func(v int) (int, error) { return v * 2, nil },
				func(err error) (int, error) { return 0, err },
			)
		}
	})

	b.Run("Map", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			sinkResultInt = ok.Map(func(v int) (int, error) { return v * 2, nil })
		}
	})

	b.Run("MapErr", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			sinkResultInt = ko.MapErr(func(err error) (int, error) { return 0, err })
		}
	})

	b.Run("FlatMap", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			sinkResultInt = ok.FlatMap(func(v int) Result[int] { return Ok(v * 2) })
		}
	})
}

func BenchmarkResultMarshalJSON(b *testing.B) {
	ok := Ok("hello world")
	ko := Err[string](errors.New("an error"))

	b.Run("Ok", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			sinkBytes, sinkErr = ok.MarshalJSON()
		}
	})

	b.Run("Err", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			sinkBytes, sinkErr = ko.MarshalJSON()
		}
	})
}

func BenchmarkResultUnmarshalJSON(b *testing.B) {
	okPayload := []byte(`{"result":"hello world"}`)
	koPayload := []byte(`{"error":{"message":"an error"}}`)

	b.Run("Ok", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			sinkErr = sinkResultString.UnmarshalJSON(okPayload)
		}
	})

	b.Run("Err", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			sinkErr = sinkResultString.UnmarshalJSON(koPayload)
		}
	})
}
