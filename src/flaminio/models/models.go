package models

import (
	"time"
	"github.com/satori/go.uuid"
	"flaminio/utility"
	"encoding/json"
	"database/sql"
)

//Basic models
type StandardModel struct {
	UUID uuid.UUID `json:"uuid"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

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

//Functional models
type User struct {
	StandardModel
	FirstName string `json:"firstname"`
	MiddleName sql.NullString `json:"middlename"`
	LastName string `json:"lastname"`
	Password string `json:"-"`
	Email string `json:"email"`
	Permissions []Permission `json:"permissions"`
}

type Permission struct {
	StandardModel
	Name string `json:"name"`
	Users []User `json:"users"`
}

type Location struct {
	StandardModel
	Name string `json:"name"`
	Description string `json:"description"`
}

type Sequence struct {
	StandardModel
	Name string `json:"name"`
	Description string `json:"description"`
}

//TODO: Fix CustomDateAndTime not working in Database
type Reservation struct {
	StandardModel
	Name string `json:"name"`
	Description string `json:"description"`
	Creator     User              `json:"-"`
	CreatorID   uuid.UUID         `json:"creator-id"`
	Location    Location          `json:"-"`
	LocationID  uuid.UUID         `json:"location-id"`
	Sequence    Sequence          `json:"-"`
	SequenceID  uuid.NullUUID     `json:"sequence-id"`
	StartTimestamp CustomDateAndTime `json:"start"`
	EndTimestamp CustomDateAndTime `json:"end"`
}

type Log struct {
	StandardModel
	User User
	UserID uuid.UUID
	OperationType LogOperationType
	OperationTypeID uuid.UUID
	Message string `json:"message"`
}

type LogOperationType struct {
	StandardModel
	Name string `json:"name"`
}
