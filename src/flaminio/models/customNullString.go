package models

import (
	"encoding/json"
	"database/sql/driver"
)

//Behaves exactly like sql.NullString, but behaves better with JSON, as it outputs either 'null' if there is no data,
//or the string itself if it is not 'null'.
type CustomNullString struct {
	String string
	Valid bool
}

func (c *CustomNullString) UnmarshalJSON(j []byte) (err error) {
	s := string(j)
	s = s[1:len(s)-1]

	if s != "" {
		c.String = s
		c.Valid = true
		return nil
	}
	c.String = ""
	c.Valid = false
	return nil
}

func (c CustomNullString) MarshalJSON() ([]byte, error) {
	if c.Valid {
		return json.Marshal(c.String)
	}
	return json.Marshal(nil)
}

func (c *CustomNullString) Scan(src interface{}) error {
	if src == nil {
		c.String, c.Valid = "", false
		return nil
	}
	if convertedString, ok := src.(string); ok {
		c.String = convertedString
		c.Valid = true
		return nil
	}
	return nil
}

func (c CustomNullString) Value() (driver.Value, error) {
	if !c.Valid {
		return nil, nil
	}
	return c.String, nil
}

func ToNullString(s string) CustomNullString {
	if s == "" {
		return CustomNullString{
			String: "",
			Valid: false,
		}
	}
	return CustomNullString{
		String: s,
		Valid: true,
	}
}
