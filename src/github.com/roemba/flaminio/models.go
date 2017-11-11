package flaminio

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	ID int `json:"id"`
	FirstName string ""
	Name string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}
