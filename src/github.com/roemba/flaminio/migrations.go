package flaminio

import "log"

/*
All migrations will run only once on startup so performance impact will be minimal when the application is running
 */
func migrate() (err error) {
	log.Println("Starting migration...")

	err = createUsersTable()
	if err != nil {
		return err
	}

	err = createPermissionsTable()
	if err != nil {
		return err
	}
	log.Println("Finished migration")
	return
}

func createUsersTable() (err error) {
	db.AutoMigrate(&User{})

	if query := db.First(&User{}, User{Email:"admin@admin.com"}); query.RecordNotFound() {
		hashedPassword, err := hashPassword("admin")
		if err != nil {
			return err
		}
		db.Create(&User{FirstName:"admin", LastName:"admin", Email:"admin@admin.com", Password:hashedPassword})
	}
	return
}

func createPermissionsTable() (err error) {
	db.AutoMigrate(&Permission{})

	var permissionArray = []Permission {
		{
			Name: "canViewSchedule",
		},
		{
			Name: "canEditUsers",
		},
		{
			Name: "canViewUsers",
		},
		{
			Name: "canEditSchedule",
		},
	}

	for i, e := range permissionArray {
		var permission Permission
		db.FirstOrCreate(&permission, e)
		permissionArray[i] = permission
	}

	var user User
	db.First(&user, User{Email:"admin@admin.com"})
	user.Permissions = permissionArray
	db.Save(user)
	return
}
