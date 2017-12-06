package database

import (
	"log"
	"github.com/roemba/flaminio/models"
	"github.com/roemba/flaminio/utility"
)

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

	err = createLogAndOperationsTable()
	if err != nil {
		return err
	}

	db.AutoMigrate(&models.Location{}, &models.Metadata{})

	createSequenceTable()
	createReservationsTable()

	log.Println("Finished migration")
	return
}

func createLogAndOperationsTable() (err error) {
	db.AutoMigrate(&models.LogOperationType{}, &models.Log{})
	db.Model(&models.Log{}).AddForeignKey("operation_type_id", "log_operation_types(uuid)", "RESTRICT", "CASCADE")
	db.Model(&models.Log{}).AddForeignKey("user_id", "users(uuid)", "CASCADE", "CASCADE")

	var operationTypesArray = []models.LogOperationType {
		{
			Name: "Changed",
		},
		{
			Name: "Deleted",
		},
		{
			Name: "Added",
		},
		{
			Name: "Moved",
		},
		{
			Name: "Authentication",
		},
	}

	for i, e := range operationTypesArray {
		var operationType models.LogOperationType
		db.FirstOrCreate(&operationType, e)
		operationTypesArray[i] = operationType
	}

	return
}

func createUsersTable() (err error) {
	db.AutoMigrate(&models.User{})

	if query := db.First(&models.User{}, models.User{Email:"admin@admin.com"}); query.RecordNotFound() {
		hashedPassword, err := utility.HashPassword("admin")
		if err != nil {
			return err
		}
		db.Create(&models.User{FirstName:"admin", LastName:"admin", Email:"admin@admin.com", Password:hashedPassword})
	}
	return
}

func createPermissionsTable() (err error) {
	db.AutoMigrate(&models.Permission{})
	db.Exec("ALTER TABLE user_permissions ADD CONSTRAINT user_foreign_key FOREIGN KEY (user_uuid) " +
		"REFERENCES users (uuid) ON DELETE CASCADE ON UPDATE CASCADE;")
	db.Exec("ALTER TABLE user_permissions ADD CONSTRAINT permission_foreign_key FOREIGN KEY (permission_uuid) " +
		"REFERENCES permissions (uuid) ON DELETE CASCADE ON UPDATE CASCADE;")

	var permissionArray = []models.Permission {
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
		{
			Name: "canEditLocations",
		},
	}

	for i, e := range permissionArray {
		var permission models.Permission
		db.FirstOrCreate(&permission, e)
		permissionArray[i] = permission
	}

	var user models.User
	db.First(&user, models.User{Email:"admin@admin.com"})
	user.Permissions = permissionArray
	db.Save(user)
	return
}

func createSequenceTable() {
	db.AutoMigrate(&models.Sequence{})
	db.Model(&models.Sequence{}).AddForeignKey("meta_id", "metadata(uuid)", "RESTRICT", "CASCADE")
}

func createReservationsTable() {
	db.AutoMigrate(&models.Reservation{})
	db.Model(&models.Reservation{}).AddForeignKey("meta_id", "metadata(uuid)", "RESTRICT", "CASCADE")
	db.Model(&models.Reservation{}).AddForeignKey("sequence_id", "sequences(uuid)", "CASCADE", "CASCADE")
	db.Model(&models.Reservation{}).AddForeignKey("location_id", "locations(uuid)", "CASCADE", "CASCADE")
	db.Model(&models.Reservation{}).AddForeignKey("creator_id", "users(uuid)", "RESTRICT", "CASCADE")
}
