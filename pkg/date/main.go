package date

import "time"

const TimeFormat = "2006-01-02"

type Date time.Time

func (d *Date) UnmarshalJSON(data []byte) (err error) {
	if len(data) == 2 {
		*d = Date(time.Time{})
		return
	}

	now, err := time.Parse(`"`+TimeFormat+`"`, string(data))
	*d = Date(now)
	return
}

func (d Date) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(TimeFormat)+2)
	b = append(b, '"')
	b = time.Time(d).AppendFormat(b, TimeFormat)
	b = append(b, '"')
	return b, nil
}

func (d Date) String() string {
	return time.Time(d).Format(TimeFormat)
}
