package time

import "time"

// Time wraps time native library from Go.
type Time struct {
	time.Time
}

// UnmarshalJSON customizes date parsing while unmarshalling to time.DateTime format.
func (t *Time) UnmarshalJSON(b []byte) (err error) {
	date, err := time.Parse(time.DateTime, string(b))
	if err != nil {
		return err
	}

	t.Time = date
	return
}
