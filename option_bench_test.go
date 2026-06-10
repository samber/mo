package mo

import (
	"strings"
	"testing"
)

var (
	sinkOptionInt    Option[int]
	sinkOptionString Option[string]
	sinkBytes        []byte
	sinkBool         bool
	sinkInt          int
	sinkErr          error
)

func BenchmarkOptionConstructors(b *testing.B) {
	b.Run("Some", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			sinkOptionInt = Some(i)
		}
	})

	b.Run("None", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			sinkOptionInt = None[int]()
		}
	})

	b.Run("TupleToOption", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			sinkOptionInt = TupleToOption(i, i%2 == 0)
		}
	})

	b.Run("EmptyableToOption", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			sinkOptionInt = EmptyableToOption(i)
		}
	})

	b.Run("PointerToOption", func(b *testing.B) {
		v := 42
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			sinkOptionInt = PointerToOption(&v)
		}
	})
}

func BenchmarkOptionAccessors(b *testing.B) {
	some := Some(42)
	none := None[int]()

	b.Run("Get", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			sinkInt, sinkBool = some.Get()
		}
	})

	b.Run("OrElse", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			sinkInt = none.OrElse(21)
		}
	})

	b.Run("Match", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			sinkOptionInt = some.Match(
				func(v int) (int, bool) { return v * 2, true },
				func() (int, bool) { return 0, false },
			)
		}
	})

	b.Run("Map", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			sinkOptionInt = some.Map(func(v int) (int, bool) { return v * 2, true })
		}
	})

	b.Run("FlatMap", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			sinkOptionInt = some.FlatMap(func(v int) Option[int] { return Some(v * 2) })
		}
	})
}

func BenchmarkOptionMarshalJSON(b *testing.B) {
	some := Some("hello world")
	none := None[string]()

	b.Run("Some", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			sinkBytes, sinkErr = some.MarshalJSON()
		}
	})

	b.Run("None", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			sinkBytes, sinkErr = none.MarshalJSON()
		}
	})
}

func BenchmarkOptionUnmarshalJSON(b *testing.B) {
	small := []byte(`"hello world"`)
	null := []byte(`null`)
	large := []byte(`"` + strings.Repeat("x", 1024) + `"`)

	b.Run("SmallValue", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			sinkErr = sinkOptionString.UnmarshalJSON(small)
		}
	})

	b.Run("Null", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			sinkErr = sinkOptionString.UnmarshalJSON(null)
		}
	})

	b.Run("LargeValue", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			sinkErr = sinkOptionString.UnmarshalJSON(large)
		}
	})
}

func BenchmarkOptionBinary(b *testing.B) {
	some := Some("hello world")
	encoded, err := some.MarshalBinary()
	if err != nil {
		b.Fatal(err)
	}

	b.Run("MarshalBinary", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			sinkBytes, sinkErr = some.MarshalBinary()
		}
	})

	b.Run("UnmarshalBinary", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			sinkErr = sinkOptionString.UnmarshalBinary(encoded)
		}
	})
}

func BenchmarkOptionIsZero(b *testing.B) {
	some := Some(42)
	none := None[int]()

	b.Run("Some", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			sinkBool = some.IsZero()
		}
	})

	b.Run("None", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			sinkBool = none.IsZero()
		}
	})
}

func BenchmarkOptionEqual(b *testing.B) {
	b.Run("Int", func(b *testing.B) {
		x := Some(42)
		y := Some(42)
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			sinkBool = x.Equal(y)
		}
	})

	b.Run("String", func(b *testing.B) {
		x := Some("hello world")
		y := Some("hello world")
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			sinkBool = x.Equal(y)
		}
	})

	b.Run("Struct", func(b *testing.B) {
		type payload struct {
			A int
			B string
		}
		x := Some(payload{A: 42, B: "hello"})
		y := Some(payload{A: 42, B: "hello"})
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			sinkBool = x.Equal(y)
		}
	})
}

func BenchmarkOptionScan(b *testing.B) {
	b.Run("Int64", func(b *testing.B) {
		var o Option[int64]
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			sinkErr = o.Scan(int64(42))
		}
	})

	b.Run("String", func(b *testing.B) {
		var o Option[string]
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			sinkErr = o.Scan("hello world")
		}
	})
}
