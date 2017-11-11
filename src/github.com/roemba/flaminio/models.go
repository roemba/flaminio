package flaminio

import (
	"time"
)

type StandardModel struct {
	UUID string `json:"uuid" gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
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
	Permissions []Permission `gorm:"many2many:user_permissions"`
}

type Permission struct {
	StandardModel
	Name string `gorm:"size:255;not null"`
	Users []User `gorm:"many2many:user_permissions"`
}
