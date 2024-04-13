package date

import (
	"fmt"
	"time"
)

const DateFormat = "2006-01-02"

type Date time.Time

func (d *Date) UnmarshalJSON(data []byte) (err error) {
	if len(data) == 2 {
		*d = Date(time.Time{})
		return
	}

	now, err := time.Parse(`"`+DateFormat+`"`, string(data))
	*d = Date(now)
	return
}

func (d Date) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(DateFormat)+2)
	b = append(b, '"')
	b = time.Time(d).AppendFormat(b, DateFormat)
	b = append(b, '"')
	return b, nil
}

func (d *Date) Scan(value interface{}) error {
	if value == nil {
		*d = Date(time.Time{})
		return nil
	}

	date, ok := value.(time.Time)
	if !ok {
		return fmt.Errorf("failed to scan Date: %v", value)
	}
	*d = Date(date)
	return nil
}

func (d Date) String() string {
	return time.Time(d).Format(DateFormat)
}
