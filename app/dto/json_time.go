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

func (jt JSONTime) MarshalJSON() ([]byte, error) {
	formatted := jt.Format(timeLayout)
	return []byte(`"` + formatted + `"`), nil
}

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

func (jt JSONTime) Value() (driver.Value, error) {
	if jt.IsZero() {
		return nil, nil
	}
	return jt.Time, nil
}
