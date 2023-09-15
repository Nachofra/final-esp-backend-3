package time

import "time"

type Time struct {
	time.Time
}

func (t *Time) UnmarshalJSON(b []byte) (err error) {
	date, err := time.Parse(time.DateTime, string(b))
	if err != nil {
		return err
	}

	t.Time = date
	return
}
