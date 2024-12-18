package mo

import (
	"bytes"
	"database/sql"
	"database/sql/driver"
	"encoding/gob"
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
)

var optionNoSuchElement = fmt.Errorf("no such element")

// Some builds an Option when value is present.
// Play: https://go.dev/play/p/iqz2n9n0tDM
func Some[T any](value T) Option[T] {
	return Option[T]{value}
}

// None builds an Option when value is absent.
// Play: https://go.dev/play/p/yYQPsYCSYlD
func None[T any]() Option[T] {
	return make(Option[T], 0)
}

// TupleToOption builds a Some Option when second argument is true, or None.
// Play: https://go.dev/play/p/gkrg2pZwOty
func TupleToOption[T any](value T, ok bool) Option[T] {
	if ok {
		return Some(value)
	}
	return None[T]()
}

// EmptyableToOption builds a Some Option when value is not empty, or None.
// Play: https://go.dev/play/p/GSpQQ-q-UES
func EmptyableToOption[T any](value T) Option[T] {
	// 🤮
	isZero := reflect.ValueOf(&value).Elem().IsZero()
	if isZero {
		return None[T]()
	}

	return Some(value)
}

// PointerToOption builds a Some Option when value is not nil, or None.
// Play: https://go.dev/play/p/yPVMj4DUb-I
func PointerToOption[T any](value *T) Option[T] {
	if value == nil {
		return None[T]()
	}

	return Some(*value)
}

// Option is a container for an optional value of type T. If value exists, Option is
// of type Some. If the value is absent, Option is of type None.
type Option[T any] []T

// IsPresent returns false when value is absent.
// Play: https://go.dev/play/p/nDqIaiihyCA
func (o Option[T]) IsPresent() bool {
	return len(o) > 0
}

// IsAbsent returns false when value is present.
// Play: https://go.dev/play/p/23e2zqyVOQm
func (o Option[T]) IsAbsent() bool {
	return len(o) == 0
}

// Size returns 1 when value is present or 0 instead.
// Play: https://go.dev/play/p/7ixCNG1E9l7
func (o Option[T]) Size() int {
	return len(o)
}

// Get returns value and presence.
// Play: https://go.dev/play/p/0-JBa1usZRT
func (o Option[T]) Get() (T, bool) {
	if len(o) == 0 {
		return empty[T](), false
	}

	return o[0], true
}

// MustGet returns value if present or panics instead.
// Play: https://go.dev/play/p/RVBckjdi5WR
func (o Option[T]) MustGet() T {
	if len(o) == 0 {
		panic(optionNoSuchElement)
	}

	return o[0]
}

// OrElse returns value if present or default value.
// Play: https://go.dev/play/p/TrGByFWCzXS
func (o Option[T]) OrElse(fallback T) T {
	if len(o) == 0 {
		return fallback
	}

	return o[0]
}

// OrEmpty returns value if present or empty value.
// Play: https://go.dev/play/p/SpSUJcE-tQm
func (o Option[T]) OrEmpty() T {
	if len(o) == 0 {
		return empty[T]()
	}
	return o[0]
}

// ForEach executes the given side-effecting function of value is present.
func (o Option[T]) ForEach(onValue func(value T)) {
	if len(o) > 0 {
		onValue(o[0])
	}
}

// Match executes the first function if value is present and second function if absent.
// It returns a new Option.
// Play: https://go.dev/play/p/1V6st3LDJsM
func (o Option[T]) Match(onValue func(value T) (T, bool), onNone func() (T, bool)) Option[T] {
	if len(o) == 0 {
		return TupleToOption(onNone())
	}
	return TupleToOption(onValue(o[0]))
}

// Map executes the mapper function if value is present or returns None if absent.
// Play: https://go.dev/play/p/mvfP3pcP_eJ
func (o Option[T]) Map(mapper func(value T) (T, bool)) Option[T] {
	if len(o) == 0 {
		return None[T]()
	}
	return TupleToOption(mapper(o[0]))
}

// MapNone executes the mapper function if value is absent or returns Option.
// Play: https://go.dev/play/p/_KaHWZ6Q17b
func (o Option[T]) MapNone(mapper func() (T, bool)) Option[T] {
	if len(o) == 0 {
		return TupleToOption(mapper())
	}

	return Some(o[0])
}

// FlatMap executes the mapper function if value is present or returns None if absent.
// Play: https://go.dev/play/p/OXO-zJx6n5r
func (o Option[T]) FlatMap(mapper func(value T) Option[T]) Option[T] {
	if len(o) == 0 {
		return None[T]()
	}
	return mapper(o[0])
}

// ToPointer returns value if present or a nil pointer.
// Play: https://go.dev/play/p/N43w92SM-Bs
func (o Option[T]) ToPointer() *T {
	if len(o) == 0 {
		return nil
	}

	return &(o[0])
}

// MarshalJSON encodes Option into json.
func (o Option[T]) MarshalJSON() ([]byte, error) {
	if len(o) == 0 {
		// if anybody find a way to support `omitempty` param, please contribute!
		return json.Marshal(nil)
	}

	return json.Marshal(o[0])
}

// UnmarshalJSON decodes Option from json.
func (o *Option[T]) UnmarshalJSON(b []byte) error {
	if bytes.Equal(b, []byte("null")) {
		*o = make(Option[T], 0)
		return nil
	}

	var v T
	err := json.Unmarshal(b, &v)
	if err != nil {
		return err
	}

	*o = Option[T]{v}
	return nil
}

// MarshalText implements the encoding.TextMarshaler interface.
func (o Option[T]) MarshalText() ([]byte, error) {
	return json.Marshal(o)
}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (o *Option[T]) UnmarshalText(data []byte) error {
	return json.Unmarshal(data, o)
}

// MarshalBinary is the interface implemented by an object that can marshal itself into a binary form.
func (o Option[T]) MarshalBinary() ([]byte, error) {
	if len(o) == 0 {
		return []byte{0}, nil
	}

	var buf bytes.Buffer

	enc := gob.NewEncoder(&buf)
	if err := enc.Encode(o[0]); err != nil {
		return []byte{}, err
	}

	return append([]byte{1}, buf.Bytes()...), nil
}

// UnmarshalBinary is the interface implemented by an object that can unmarshal a binary representation of itself.
func (o *Option[T]) UnmarshalBinary(data []byte) error {
	if len(data) == 0 {
		return errors.New("Option[T].UnmarshalBinary: no data")
	}

	if data[0] == 0 {
		*o = make(Option[T], 0)
		return nil
	}

	buf := bytes.NewBuffer(data[1:])
	dec := gob.NewDecoder(buf)
	var v T
	err := dec.Decode(&v)
	if err != nil {
		return err
	}

	*o = Option[T]{v}
	return nil
}

// GobEncode implements the gob.GobEncoder interface.
func (o Option[T]) GobEncode() ([]byte, error) {
	return o.MarshalBinary()
}

// GobDecode implements the gob.GobDecoder interface.
func (o *Option[T]) GobDecode(data []byte) error {
	return o.UnmarshalBinary(data)
}

// Scan implements the SQL sql.Scanner interface.
func (o *Option[T]) Scan(src any) error {
	if src == nil {
		*o = make(Option[T], 0)
		return nil
	}

	// is is only possible to assert interfaces, so convert first
	// https://go.googlesource.com/proposal/+/refs/heads/master/design/43651-type-parameters.md#why-not-permit-type-assertions-on-values-whose-type-is-a-type-parameter
	var t T
	if tScanner, ok := interface{}(&t).(sql.Scanner); ok {
		if err := tScanner.Scan(src); err != nil {
			return fmt.Errorf("failed to scan: %w", err)
		}

		*o = Option[T]{t}
		return nil
	}

	if av, err := driver.DefaultParameterConverter.ConvertValue(src); err == nil {
		if v, ok := av.(T); ok {
			*o = Option[T]{v}
			return nil
		}
	}

	return o.scanConvertValue(src)
}

// Value implements the driver Valuer interface.
func (o Option[T]) Value() (driver.Value, error) {
	if len(o) == 0 {
		return nil, nil
	}

	return driver.DefaultParameterConverter.ConvertValue(o[0])
}

// leftValue returns an error if the Option is None, otherwise nil
//
//nolint:unused
func (o Option[T]) leftValue() error {
	if len(o) == 0 {
		return optionNoSuchElement
	}
	return nil
}

// rightValue returns the value if the Option is Some, otherwise the zero value of T
//
//nolint:unused
func (o Option[T]) rightValue() T {
	if len(o) == 0 {
		var zero T
		return zero
	}
	return o[0]
}

// hasLeftValue returns true if the Option represents a None state
//
//nolint:unused
func (o Option[T]) hasLeftValue() bool {
	return len(o) == 0
}
