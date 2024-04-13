package date

import (
	"fmt"
	"time"
)

const TimeFormat = "15:04"

type Time time.Time

func (t *Time) UnmarshalJSON(data []byte) (err error) {
	if len(data) == 2 {
		*t = Time(time.Time{})
		return
	}

	time, err := time.Parse(`"`+TimeFormat+`"`, string(data))
	*t = Time(time)
	return
}

func (t Time) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(TimeFormat)+2)
	b = append(b, '"')
	b = time.Time(t).AppendFormat(b, TimeFormat)
	b = append(b, '"')
	return b, nil
}

func (t *Time) Scan(value interface{}) error {
	if value == nil {
		*t = Time(time.Time{})
		return nil
	}

	time, ok := value.(time.Time)
	if !ok {
		return fmt.Errorf("failed to scan Time: %v", value)
	}
	*t = Time(time)
	return nil
}

func (t Time) String() string {
	return time.Time(t).Format(TimeFormat)
}
