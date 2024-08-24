package database

import (
	"database/sql/driver"
	"fmt"
	"strings"
	"time"
)

// TimeArray is a custom type for handling arrays of time.Time
type TimeArray []time.Time

// Scan implements the sql.Scanner interface
func (t *TimeArray) Scan(src interface{}) error {
	if src == nil {
		*t = nil
		return nil
	}

	switch src := src.(type) {
	case string:
		// Remove curly braces and split by comma
		elements := strings.Split(strings.Trim(src, "{}"), ",")
		times := make([]time.Time, len(elements))

		for i, elem := range elements {
			// Parse each time element
			parsedTime, err := time.Parse(time.RFC3339, elem)
			if err != nil {
				return fmt.Errorf("cannot parse time: %v", err)
			}
			times[i] = parsedTime
		}

		*t = times
		return nil

	default:
		return fmt.Errorf("cannot scan %T into TimeArray", src)
	}
}

// Value implements the driver.Valuer interface
func (t TimeArray) Value() (driver.Value, error) {
	if len(t) == 0 {
		return "{}", nil
	}

	times := make([]string, len(t))
	for i, timeVal := range t {
		times[i] = timeVal.Format(time.RFC3339)
	}

	return "{" + strings.Join(times, ",") + "}", nil
}
