package models

import (
	"database/sql/driver"
	"encoding/json"

	"github.com/jackc/pgx/pgtype"
)

type CustomTsrange struct {
	pgtype.Tsrange
}

type ConversionStruct struct {
	Start CustomDateAndTime `json:"start"`
	End CustomDateAndTime `json:"end"`
}

func (c *CustomTsrange) UnmarshalJSON(j []byte) (err error) {
	var inputStruct ConversionStruct
	err = json.Unmarshal(j, &inputStruct)
	if err != nil {
		return err
	}

	var startTimeStamp pgtype.Timestamp
	err = startTimeStamp.Set(inputStruct.Start.Time)
	if err != nil {
		return err
	}

	var endTimeStamp pgtype.Timestamp
	err = endTimeStamp.Set(inputStruct.End.Time)
	if err != nil {
		return err
	}

	c.Tsrange = pgtype.Tsrange{
		Lower: startTimeStamp,
		Upper: endTimeStamp,
		LowerType: pgtype.Inclusive,
		UpperType: pgtype.Exclusive,
		Status: pgtype.Present,
	}

	return err
}

func (c CustomTsrange) MarshalJSON() ([]byte, error) {
	return json.Marshal(ConversionStruct{
		Start: CustomDateAndTime{Time: c.Tsrange.Lower.Time},
		End: CustomDateAndTime{c.Tsrange.Upper.Time},
	})
}

func (c *CustomTsrange) Scan(src interface{}) error {
	return c.Tsrange.Scan(src)
}

func (c CustomTsrange) Value() (driver.Value, error) {
	return c.Tsrange.Value()
}
