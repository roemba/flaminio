package models

import (
	"time"
	"github.com/satori/go.uuid"
	"github.com/roemba/flaminio/utility"
	"encoding/json"
	"database/sql"
)

//Basic models
type StandardModel struct {
	UUID uuid.UUID `json:"uuid" gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
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
	FirstName string `json:"firstname" gorm:"size:255;not null"`
	MiddleName sql.NullString `json:"middlename" gorm:"size:255"`
	LastName string `json:"lastname" gorm:"size:255;not null"`
	Password string `json:"-" gorm:"type:bytea;not null"`
	Email string `json:"email" gorm:"type:citext;not null;unique_index"`
	Permissions []Permission `json:"permissions" gorm:"many2many:user_permissions"`
}

type Permission struct {
	StandardModel
	Name string `json:"name" gorm:"size:255;not null;unique"`
	Users []User `json:"users" gorm:"many2many:user_permissions"`
}

type Location struct {
	StandardModel
	Name string `json:"name" gorm:"size:255;not null;unique"`
	Description string `json:"description" gorm:"type:text;"`
}

type Sequence struct {
	StandardModel
	Meta Metadata
	MetaID uuid.UUID `gorm:"type:uuid;not null"`
}

type Metadata struct {
	StandardModel
	Name string `json:"name" gorm:"size:255;not null;"`
	Description string `json:"description" gorm:"type:text;"`
}

//TODO: Fix CustomDateAndTime not working in Database
type Reservation struct {
	StandardModel
	Creator     User              `json:"-"`
	CreatorID   uuid.UUID         `json:"creator-id" gorm:"type:uuid;not null"`
	Location    Location          `json:"-"`
	LocationID  uuid.UUID         `json:"location-id" gorm:"type:uuid;not null;"`
	Sequence    Sequence          `json:"-"`
	SequenceID  uuid.NullUUID     `json:"sequence-id" gorm:"type:uuid;"`
	Meta        Metadata          `json:"-"`
	MetaID      uuid.UUID         `json:"-" gorm:"type:uuid;not null;"`
	DateAndTime CustomDateAndTime `json:"date" gorm:"type:timestamp;not null;"`
}

type Log struct {
	StandardModel
	User User
	UserID uuid.UUID `gorm:"type:uuid;not null"`
	OperationType LogOperationType
	OperationTypeID uuid.UUID `gorm:"type:uuid;not null"`
	Message string `json:"message" gorm:"type:text;"`
}

type LogOperationType struct {
	StandardModel
	Name string `json:"name" gorm:"size:255;not null;unique"`
}
