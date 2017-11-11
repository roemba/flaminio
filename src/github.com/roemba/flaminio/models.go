package flaminio

import (
	"time"
)

type StandardModel struct {
	ID string `json:"id" gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type User struct {
	StandardModel
	FirstName string `json:"firstname" gorm:"size:255;not null"`
	MiddleName string `json:"middlename" gorm:"size:255"`
	LastName string `json:"lastname" gorm:"size:255;not null"`
	Username string `json:"username" gorm:"size:255;not null;unique"`
	Password string `json:"password" gorm:"type:bytea;not null"`
}
