package models

import (
	"encoding/json"
	"errors"
	"flaminio/utility"
	"time"
)

//TODO Change to customTimeStamp, and use timestamp type
type CustomDateAndTime struct {
	time.Time
}

func (c *CustomDateAndTime) UnmarshalJSON(j []byte) (err error) {
	s := string(j)
	s = s[1:len(s)-1]
	t, err := time.Parse(utility.ISO8601DATE_TIME, s)
	if err != nil {
		return err
	}
	c.Time = t
	return
}

func (c CustomDateAndTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.Time.Format(utility.ISO8601DATE_TIME))
}

func (c *CustomDateAndTime) Scan(src interface{}) error {
	if convertedTime, ok := src.(time.Time); ok {
		c.Time = convertedTime
		return nil
	}
	return errors.New("Could not convert input to Time")
}