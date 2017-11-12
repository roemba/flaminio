package flaminio

import "github.com/satori/go.uuid"

/*
All database calls (except migrations) are done from this file
 */
func getUser(email string) (user User, err bool) {
	if db.First(&user, User{Email: email}).RecordNotFound() {
		return user, true
	}
	return user, false
}

func getUserWithPermissions(uuid uuid.UUID) (user User) {
	var permissions []Permission
	db.First(&user, User{StandardModel:StandardModel{UUID: uuid}})
	db.Model(&user).Related(&permissions, "Permissions")
	user.Permissions = permissions
	return user
}
