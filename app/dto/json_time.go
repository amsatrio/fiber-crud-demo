package dto

import (
	"database/sql/driver"
	"errors"
	"strings"
	"time"
)

const timeLayout = "2006-01-02 15:04:05"

type JSONTime struct {
	time.Time
}

// MarshalJSON writes the time in "YYYY-MM-DD HH:mm:ss" format
func (jt JSONTime) MarshalJSON() ([]byte, error) {
	formatted := jt.Format(timeLayout)
	return []byte(`"` + formatted + `"`), nil
}

// UnmarshalJSON parses the time from "YYYY-MM-DD HH:mm:ss" format
func (jt *JSONTime) UnmarshalJSON(data []byte) error {
	str := strings.Trim(string(data), `"`)
	if str == "null" || str == "" {
		return nil
	}

	parsedTime, err := time.Parse(timeLayout, str)
	if err != nil {
		return err
	}

	jt.Time = parsedTime
	return nil
}

// Scan implements the sql.Scanner interface for database retrieval
func (jt *JSONTime) Scan(value interface{}) error {
	if value == nil {
		jt.Time = time.Time{}
		return nil
	}

	switch v := value.(type) {
	case time.Time:
		jt.Time = v
	case []byte:
		return jt.UnmarshalJSON(v)
	case string:
		return jt.UnmarshalJSON([]byte(v))
	default:
		return errors.New("unsupported type for JSONTime")
	}
	return nil
}

// Value implements the driver.Valuer interface for database insertion
func (jt JSONTime) Value() (driver.Value, error) {
	if jt.IsZero() {
		return nil, nil
	}
	return jt.Time, nil
}