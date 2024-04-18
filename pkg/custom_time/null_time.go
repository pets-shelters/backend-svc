package custom_time

import (
	"database/sql/driver"
	"fmt"
	"log"
	"time"
)

const (
	TimeFormat = "15:04"
)

type NullTime struct {
	Time  time.Time
	Valid bool
}

func (nt NullTime) Value() (driver.Value, error) {
	if !nt.Valid {
		return nil, nil
	}
	return nt.Time, nil
}

func (nt *NullTime) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	stringTime, ok := value.(string)
	if !ok {
		return fmt.Errorf("failed to scan NullTime: %v", value)
	}

	parsedTime, err := time.Parse(TimeFormat, stringTime[0:5])
	if err != nil {
		return fmt.Errorf("failed to parse time: %w", err)
	}

	nt.Time, nt.Valid = parsedTime, true
	return nil
}

func (nt *NullTime) UnmarshalJSON(data []byte) (err error) {
	log.Printf("json time data %+v", data)
	if len(data) == 2 {
		log.Printf("json time is nil")
		return
	}

	time, err := time.Parse(`"`+TimeFormat+`"`, string(data))
	nt.Time, nt.Valid = time, true
	return
}

func (nt NullTime) MarshalJSON() ([]byte, error) {
	log.Printf("nt.Valid %+v", nt.Valid)
	log.Printf("nt %+v", nt)
	if !nt.Valid {
		return []byte(`""`), nil
	}

	b := make([]byte, 0, len(TimeFormat)+2)
	b = append(b, '"')
	b = nt.Time.AppendFormat(b, TimeFormat)
	b = append(b, '"')
	return b, nil
}

func (nt NullTime) String() string {
	if !nt.Valid {
		return ""
	}
	return nt.Time.Format(TimeFormat)
}
