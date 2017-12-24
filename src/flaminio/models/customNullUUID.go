package models

import (
	"encoding/json"
	"database/sql/driver"
	"github.com/satori/go.uuid"
)

//Behaves exactly like uuid.NullUUID, but behaves better with JSON, as it outputs either 'null' if there is no data,
//or the uuid itself if it is not 'null'.
type CustomNullUUID struct {
	UUID uuid.UUID
	Valid bool
}

func (c *CustomNullUUID) UnmarshalJSON(j []byte) (err error) {
	u, err := uuid.FromBytes(j)
	if err != nil {
		return err
	}

	if u != uuid.Nil {
		c.UUID = u
		c.Valid = true
		return nil
	}
	c.UUID = uuid.Nil
	c.Valid = false
	return nil
}

func (c CustomNullUUID) MarshalJSON() ([]byte, error) {
	if c.Valid {
		return json.Marshal(c.UUID)
	}
	return json.Marshal(nil)
}

func (c *CustomNullUUID) Scan(src interface{}) error {
	if src == nil {
		c.UUID, c.Valid = uuid.Nil, false
		return nil
	}
	c.Valid = true
	return c.UUID.Scan(src)
}

func (c CustomNullUUID) Value() (driver.Value, error) {
	if !c.Valid {
		return nil, nil
	}
	return c.UUID.Value()
}

func ToNullUUID(u uuid.UUID) CustomNullUUID {
	if u == uuid.Nil {
		return CustomNullUUID{
			UUID: uuid.Nil,
			Valid: false,
		}
	}
	return CustomNullUUID{
		UUID: u,
		Valid: true,
	}
}
