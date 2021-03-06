package models

import (
	"time"

	"github.com/satori/go.uuid"
)

//Basic models
type StandardModel struct {
	UUID uuid.UUID `json:"uuid"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

//Functional models
type User struct {
	StandardModel
	FirstName string `json:"firstname"`
	MiddleName CustomNullString `json:"middlename"`
	LastName string `json:"lastname"`
	Password string `json:"-"`
	Email string `json:"email" binding:"required,email"`
	Permissions []Permission `json:"permissions"`
	PreferredLocale string `json:"preferred_locale"`
}

type Permission struct {
	StandardModel
	Name string `json:"name"`
	Users []User `json:"users"`
}

type Location struct {
	StandardModel
	Name string `json:"name"`
	Description CustomNullString `json:"description"`
}

type Sequence struct {
	StandardModel
	Name string `json:"name"`
	Description CustomNullString `json:"description"`
}

type Reservation struct {
	StandardModel
	Name string `json:"name"`
	Description CustomNullString `json:"description"`
	Creator     User              `json:"-"`
	CreatorID   uuid.UUID         `json:"creator_id"`
	Location    Location          `json:"-"`
	LocationID  uuid.UUID         `json:"location_id"`
	Sequence    Sequence          `json:"-"`
	SequenceID  CustomNullUUID    `json:"sequence_id"`
	Duration    CustomTsrange     `json:"duration"`
	Color       string            `json:"color"`
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

type BlacklistedToken struct {
	JwtTokenDigest []byte
	RevocationDate CustomDateAndTime
}
