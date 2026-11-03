package mo

import (
	"testing"
	"time"
)

func TestOptionIsZeroNilZeroer(t *testing.T) {
	var value *time.Time
	if !Some(value).IsZero() {
		t.Error("a nil time pointer should be zero")
	}
	if !Some[any](value).IsZero() {
		t.Error("an interface holding a nil time pointer should be zero")
	}
	if !Some(&time.Time{}).IsZero() {
		t.Error("a pointer to zero time should still use time.IsZero")
	}
	nonzero := time.Unix(1, 0)
	if Some(&nonzero).IsZero() {
		t.Error("a nonzero time pointer should not be zero")
	}
}
