package flaminio

import (
	"time"
	"github.com/satori/go.uuid"
)

type StandardModel struct {
	UUID uuid.UUID `json:"uuid" gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

type User struct {
	StandardModel
	FirstName string `json:"firstname" gorm:"size:255;not null"`
	MiddleName string `json:"middlename" gorm:"size:255"`
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
