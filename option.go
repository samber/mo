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

var errOptionNoSuchElement = fmt.Errorf("no such element")

type zeroer interface {
	IsZero() bool
}

// Some builds an Option when value is present.
// Play: https://go.dev/play/p/iqz2n9n0tDM
func Some[T any](value T) Option[T] {
	return Option[T]{
		isPresent: true,
		value:     value,
	}
}

// None builds an Option when value is absent.
// Play: https://go.dev/play/p/yYQPsYCSYlD
func None[T any]() Option[T] {
	return Option[T]{
		isPresent: false,
	}
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
type Option[T any] struct {
	isPresent bool
	value     T
}

// IsPresent returns false when value is absent.
// Play: https://go.dev/play/p/nDqIaiihyCA
func (o Option[T]) IsPresent() bool {
	return o.isPresent
}

// IsSome is an alias to IsPresent.
// Play: https://go.dev/play/p/DyvGRy7fP9m
func (o Option[T]) IsSome() bool {
	return o.IsPresent()
}

// IsAbsent returns false when value is present.
// Play: https://go.dev/play/p/23e2zqyVOQm
func (o Option[T]) IsAbsent() bool {
	return !o.isPresent
}

// IsNone is an alias to IsAbsent.
// Play: https://go.dev/play/p/EdqxKhborIP
func (o Option[T]) IsNone() bool {
	return o.IsAbsent()
}

// Size returns 1 when value is present or 0 instead.
// Play: https://go.dev/play/p/7ixCNG1E9l7
func (o Option[T]) Size() int {
	if o.isPresent {
		return 1
	}

	return 0
}

// Get returns value and presence.
// Play: https://go.dev/play/p/0-JBa1usZRT
func (o Option[T]) Get() (T, bool) {
	if !o.isPresent {
		return empty[T](), false
	}

	return o.value, true
}

// MustGet returns value if present or panics instead.
// Play: https://go.dev/play/p/RVBckjdi5WR
func (o Option[T]) MustGet() T {
	if !o.isPresent {
		panic(errOptionNoSuchElement)
	}

	return o.value
}

// OrElse returns value if present or default value.
// Play: https://go.dev/play/p/TrGByFWCzXS
func (o Option[T]) OrElse(fallback T) T {
	if !o.isPresent {
		return fallback
	}

	return o.value
}

// OrEmpty returns value if present or empty value.
// Play: https://go.dev/play/p/SpSUJcE-tQm
func (o Option[T]) OrEmpty() T {
	return o.value
}

// ForEach executes the given side-effecting function of value is present.
func (o Option[T]) ForEach(onValue func(value T)) {
	if o.isPresent {
		onValue(o.value)
	}
}

// Match executes the first function if value is present and second function if absent.
// It returns a new Option.
// Play: https://go.dev/play/p/1V6st3LDJsM
func (o Option[T]) Match(onValue func(value T) (T, bool), onNone func() (T, bool)) Option[T] {
	if o.isPresent {
		return TupleToOption(onValue(o.value))
	}
	return TupleToOption(onNone())
}

// Map executes the mapper function if value is present or returns None if absent.
// Play: https://go.dev/play/p/mvfP3pcP_eJ
func (o Option[T]) Map(mapper func(value T) (T, bool)) Option[T] {
	if o.isPresent {
		return TupleToOption(mapper(o.value))
	}

	return None[T]()
}

// MapNone executes the mapper function if value is absent or returns Option.
// Play: https://go.dev/play/p/_KaHWZ6Q17b
func (o Option[T]) MapNone(mapper func() (T, bool)) Option[T] {
	if o.isPresent {
		return Some(o.value)
	}

	return TupleToOption(mapper())
}

// FlatMap executes the mapper function if value is present or returns None if absent.
// Play: https://go.dev/play/p/OXO-zJx6n5r
func (o Option[T]) FlatMap(mapper func(value T) Option[T]) Option[T] {
	if o.isPresent {
		return mapper(o.value)
	}

	return None[T]()
}

// MapValue executes the mapper function if value is present or returns None if absent.
func (o Option[T]) MapValue(mapper func(value T) T) Option[T] {
	if o.isPresent {
		return Some(mapper(o.value))
	}

	return None[T]()
}

// ToPointer returns value if present or a nil pointer.
// Play: https://go.dev/play/p/N43w92SM-Bs
func (o Option[T]) ToPointer() *T {
	if !o.isPresent {
		return nil
	}

	return &o.value
}

// MarshalJSON encodes Option into json.
// Go 1.20+ relies on the IsZero method when the `omitempty` tag is used
// unless a custom MarshalJSON method is defined.  Then the IsZero method is ignored.
// current best workaround is to instead use `omitzero` tag with Go 1.24+
func (o Option[T]) MarshalJSON() ([]byte, error) {
	if o.isPresent {
		return json.Marshal(o.value)
	}

	return json.Marshal(nil)
}

// UnmarshalJSON decodes Option from json.
func (o *Option[T]) UnmarshalJSON(b []byte) error {
	o.value = empty[T]() // reset the value if not set later.

	// If user manually set the field to be `null`, then it either means the option is absent or present with a zero value.
	if bytes.Equal([]byte("null"), bytes.ToLower(b)) {
		// // If the type is a pointer, then it means the option is present with a zero value.
		// o.isPresent = reflect.TypeOf(o.value).Kind() == reflect.Ptr
		// return nil

		o.isPresent = false
		return nil
	}

	err := json.Unmarshal(b, &o.value)
	if err != nil {
		return err
	}

	o.isPresent = true
	return nil
}

// IsZero assists `omitzero` tag introduced in Go 1.24
func (o Option[T]) IsZero() bool {
	if !o.isPresent {
		return true
	}

	var v any = o.value
	if v, ok := v.(zeroer); ok {
		return v.IsZero()
	}

	return reflect.ValueOf(o.value).IsZero()
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
	if !o.isPresent {
		return []byte{0}, nil
	}

	var buf bytes.Buffer

	enc := gob.NewEncoder(&buf)
	if err := enc.Encode(o.value); err != nil {
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
		o.isPresent = false
		o.value = empty[T]()
		return nil
	}

	buf := bytes.NewBuffer(data[1:])
	dec := gob.NewDecoder(buf)
	err := dec.Decode(&o.value)
	if err != nil {
		return err
	}

	o.isPresent = true
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
		o.isPresent = false
		o.value = empty[T]()
		return nil
	}

	// is is only possible to assert interfaces, so convert first
	// https://go.googlesource.com/proposal/+/refs/heads/master/design/43651-type-parameters.md#why-not-permit-type-assertions-on-values-whose-type-is-a-type-parameter
	var t T
	if tScanner, ok := interface{}(&t).(sql.Scanner); ok {
		if err := tScanner.Scan(src); err != nil {
			return fmt.Errorf("failed to scan: %w", err)
		}

		o.isPresent = true
		o.value = t
		return nil
	}

	if av, err := driver.DefaultParameterConverter.ConvertValue(src); err == nil {
		if v, ok := av.(T); ok {
			o.isPresent = true
			o.value = v
			return nil
		}
	}

	return o.scanConvertValue(src)
}

// Value implements the driver Valuer interface.
func (o Option[T]) Value() (driver.Value, error) {
	if !o.isPresent {
		return nil, nil
	}

	return driver.DefaultParameterConverter.ConvertValue(o.value)
}

// Equal compares two Option[T] instances for equality
func (o Option[T]) Equal(other Option[T]) bool {
	if !o.isPresent && !other.isPresent {
		return true
	}

	if o.isPresent != other.isPresent {
		return false
	}

	return reflect.DeepEqual(o.value, other.value)
}

// leftValue returns an error if the Option is None, otherwise nil
//
//nolint:unused
func (o Option[T]) leftValue() error {
	if !o.isPresent {
		return errOptionNoSuchElement
	}
	return nil
}

// rightValue returns the value if the Option is Some, otherwise the zero value of T
//
//nolint:unused
func (o Option[T]) rightValue() T {
	if !o.isPresent {
		var zero T
		return zero
	}
	return o.value
}

// hasLeftValue returns true if the Option represents a None state
//
//nolint:unused
func (o Option[T]) hasLeftValue() bool {
	return !o.isPresent
}
