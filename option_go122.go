//go:build go1.22
// +build go1.22

package mo

import (
	"database/sql"
	"fmt"
)

func (o *Option[T]) scanConvertValue(src any) error {
	// we try to convertAssign values that we can't directly assign because ConvertValue
	// will return immediately for v that is already a Value, even if it is a different
	// Value type than the one we expect here.
	var st sql.Null[T]
	if err := st.Scan(src); err == nil {
		o.isPresent = true
		o.value = st.V
		return nil
	}
	return fmt.Errorf("failed to scan Option[T]")
}
