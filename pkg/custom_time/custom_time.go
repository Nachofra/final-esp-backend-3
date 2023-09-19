package custom_time

import (
	"fmt"
	"strings"
	"time"
)

// Time wraps time native library from Go.
type Time struct {
	time.Time
}

// UnmarshalJSON customizes date parsing while unmarshalling to custom_time.DateTime format.
func (t *Time) UnmarshalJSON(b []byte) (err error) {
	dateStr := string(b)

	// Remove extra quotes from string, remove T and Z when date comes from database.
	dateStr = strings.Trim(dateStr, "\"")
	dateStr = strings.Replace(dateStr, "T", "", -1)
	dateStr = strings.Replace(dateStr, "Z", "", -1)

	date, err := time.Parse(time.DateTime, dateStr)
	if err != nil {
		return fmt.Errorf("the field must be in datetime format (YYYY-MM-DD HH:mm:ss): %s", dateStr)
	}

	t.Time = date
	return
}
