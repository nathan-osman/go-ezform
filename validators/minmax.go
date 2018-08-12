package validators

import (
	"fmt"
)

// MinMax ensures that an integer falls within the specified range.
type MinMax struct {
	// Min is the lowest value that will be accepted.
	Min int64
	// Max is the highest value that will be accepted.
	Max int64
}

// Validate ensures the integer is valid.
func (m MinMax) Validate(v int64) error {
	if v < m.Min {
		return fmt.Errorf("value cannot be less than %d", m.Min)
	}
	if v > m.Max {
		return fmt.Errorf("value cannot be greater than %d", m.Max)
	}
	return nil
}
