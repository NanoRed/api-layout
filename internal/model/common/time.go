package common

import (
	"database/sql/driver"
	"time"
)

type Time struct {
	time.Time
}

func NewTime(t time.Time) Time {
	return Time{t}
}

func (t *Time) Scan(value any) error {
	orig := value.(time.Time)
	t.Time = time.Date(
		orig.Year(), orig.Month(), orig.Day(),
		orig.Hour(), orig.Minute(), orig.Second(),
		orig.Nanosecond(), time.Local,
	)
	return nil
}

func (t Time) Value() (driver.Value, error) {
	if t.IsZero() {
		return nil, nil
	}
	return t.Format(time.RFC3339Nano), nil
}
