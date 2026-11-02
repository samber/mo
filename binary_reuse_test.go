package mo

import (
	"bytes"
	"encoding"
	"encoding/gob"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type binaryReusePayload struct {
	Count  int
	Label  string
	Values map[string]int
}

func assertBinaryReuse[T any](t *testing.T, wrap func(binaryReusePayload) T) {
	t.Helper()
	for _, method := range []string{"binary", "gob"} {
		t.Run(method, func(t *testing.T) {
			previous := wrap(binaryReusePayload{Count: 7, Label: "old", Values: map[string]int{"old": 1}})
			expected := wrap(binaryReusePayload{Values: map[string]int{"new": 2}})
			if method == "binary" {
				data, err := any(expected).(encoding.BinaryMarshaler).MarshalBinary()
				require.NoError(t, err)
				require.NoError(t, any(&previous).(encoding.BinaryUnmarshaler).UnmarshalBinary(data))
			} else {
				var data bytes.Buffer
				require.NoError(t, gob.NewEncoder(&data).Encode(expected))
				require.NoError(t, gob.NewDecoder(&data).Decode(&previous))
			}
			assert.Equal(t, expected, previous)
		})
	}
}

func TestBinaryDecodeReusedValues(t *testing.T) {
	t.Run("Option", func(t *testing.T) {
		assertBinaryReuse(t, Some[binaryReusePayload])
	})
	t.Run("Either/left", func(t *testing.T) {
		assertBinaryReuse(t, Left[binaryReusePayload, int])
	})
	t.Run("Either/right", func(t *testing.T) {
		assertBinaryReuse(t, Right[int, binaryReusePayload])
	})
	t.Run("Either3/arg1", func(t *testing.T) {
		assertBinaryReuse(t, NewEither3Arg1[binaryReusePayload, int, int])
	})
	t.Run("Either3/arg2", func(t *testing.T) {
		assertBinaryReuse(t, NewEither3Arg2[int, binaryReusePayload, int])
	})
	t.Run("Either3/arg3", func(t *testing.T) {
		assertBinaryReuse(t, NewEither3Arg3[int, int, binaryReusePayload])
	})
	t.Run("Either4/arg1", func(t *testing.T) {
		assertBinaryReuse(t, NewEither4Arg1[binaryReusePayload, int, int, int])
	})
	t.Run("Either4/arg2", func(t *testing.T) {
		assertBinaryReuse(t, NewEither4Arg2[int, binaryReusePayload, int, int])
	})
	t.Run("Either4/arg3", func(t *testing.T) {
		assertBinaryReuse(t, NewEither4Arg3[int, int, binaryReusePayload, int])
	})
	t.Run("Either4/arg4", func(t *testing.T) {
		assertBinaryReuse(t, NewEither4Arg4[int, int, int, binaryReusePayload])
	})
	t.Run("Either5/arg1", func(t *testing.T) {
		assertBinaryReuse(t, NewEither5Arg1[binaryReusePayload, int, int, int, int])
	})
	t.Run("Either5/arg2", func(t *testing.T) {
		assertBinaryReuse(t, NewEither5Arg2[int, binaryReusePayload, int, int, int])
	})
	t.Run("Either5/arg3", func(t *testing.T) {
		assertBinaryReuse(t, NewEither5Arg3[int, int, binaryReusePayload, int, int])
	})
	t.Run("Either5/arg4", func(t *testing.T) {
		assertBinaryReuse(t, NewEither5Arg4[int, int, int, binaryReusePayload, int])
	})
	t.Run("Either5/arg5", func(t *testing.T) {
		assertBinaryReuse(t, NewEither5Arg5[int, int, int, int, binaryReusePayload])
	})
}
