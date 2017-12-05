package flaminio

import (
	"github.com/satori/go.uuid"
)

/*
All database calls (except migrations) are done from this file
 */
func getUser(email string) (user User, err bool) {
	if db.First(&user, User{Email: email}).RecordNotFound() {
		return user, true
	}
	return user, false
}

func getUserWithPermissions(userUUID uuid.UUID) (user User) {
	var permissions []Permission
	db.First(&user, User{StandardModel:StandardModel{UUID: userUUID}})
	db.Model(&user).Related(&permissions, "Permissions")
	user.Permissions = permissions
	return user
}

func addDatabaseLog(userUUID uuid.UUID, logTypeUUID uuid.UUID, message string) () {
	var log = Log{
		UserID:          userUUID,
		OperationTypeID: logTypeUUID,
		Message:         message,
	}
	db.Create(&log)
	return
}
