package database

import (
	"database/sql"
	"flaminio/models"
	"flaminio/utility"
	"log"

	"github.com/satori/go.uuid"
)

const standardModel = `uuid uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
					createdAt timestamp NOT NULL DEFAULT (NOW() AT TIME ZONE 'utc'),
					updatedAt timestamp NOT NULL DEFAULT (NOW() AT TIME ZONE 'utc'),`

func fatal(err error, tx *sql.Tx) {
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}
}

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

	err = createLocationsTable()
	if err != nil {
		return err
	}

	err = createSequenceTable()
	if err != nil {
		return err
	}

	err = createReservationsTable()
	if err != nil {
		return err
	}

	err = createTokenBlacklistTable()
	if err != nil {
		return err
	}

	log.Println("Finished migration")
	return
}

func createLogAndOperationsTable() (err error) {
	tx, err := db.Begin()
	utility.Fatal(err)

	_, err = tx.Exec(`CREATE TABLE IF NOT EXISTS flaminio.log_operation_types (
					` + standardModel + `
					name character varying(255) NOT NULL UNIQUE
				)`)
	fatal(err, tx)

	_, err = tx.Exec(`CREATE TABLE IF NOT EXISTS flaminio.logs (
					` + standardModel + `
					userId uuid NOT NULL REFERENCES flaminio.users ON DELETE CASCADE ON UPDATE CASCADE,
					operationTypeId uuid NOT NULL REFERENCES flaminio.log_operation_types ON DELETE RESTRICT ON UPDATE CASCADE,
					message text
				)`)
	fatal(err, tx)

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

	stmt, err := tx.Prepare(`INSERT INTO flaminio.log_operation_types (name)
					VALUES ($1)
					ON CONFLICT (name) DO NOTHING
					RETURNING uuid`)
	fatal(err, tx)

	for _, e := range operationTypesArray {
		_, err = stmt.Exec(e.Name)
		fatal(err, tx)
	}

	err = tx.Commit()
	return err
}

func createUsersTable() (err error) {
	tx, err := db.Begin()
	utility.Fatal(err)

	_, err = tx.Exec(`CREATE TABLE IF NOT EXISTS flaminio.users (
					` + standardModel + `
					firstName character varying(255) NOT NULL,
					middleName character varying(255),
					lastName character varying(255) NOT NULL,
					password bytea NOT NULL,
					email citext NOT NULL UNIQUE,
					preferredLocale character varying(10) NOT NULL DEFAULT 'en'
				)`)
	fatal(err, tx)

	hashedPassword, err := utility.HashPassword("admin")
	utility.Fatal(err)

	user := models.User{FirstName:"admin", LastName:"admin", Email:"admin@admin.com", Password:hashedPassword}
	_, err = tx.Exec(`INSERT INTO flaminio.users (firstName, lastName, password, email)
					VALUES ($1, $2, $3, $4)
					ON CONFLICT (email) DO NOTHING`, user.FirstName, user.LastName,
					user.Password, user.Email)
	fatal(err, tx)

	err = tx.Commit()
	utility.Fatal(err)
	return
}

func createPermissionsTable() (err error) {
	tx, err := db.Begin()
	utility.Fatal(err)

	_, err = tx.Exec(`CREATE TABLE IF NOT EXISTS flaminio.permissions (
					` + standardModel + `
					name character varying(255) NOT NULL UNIQUE
				)`)
	fatal(err, tx)

	_, err = tx.Exec(`CREATE TABLE IF NOT EXISTS flaminio.user_permissions (
					userId uuid REFERENCES flaminio.users ON DELETE CASCADE ON UPDATE CASCADE,
					permissionId uuid REFERENCES flaminio.permissions ON DELETE CASCADE ON UPDATE CASCADE,
					PRIMARY KEY (userId, permissionId)
				)`)
	fatal(err, tx)

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

	//Giving the admin user all permissions
	var user = models.User{Email: "admin@admin.com"}
	err = tx.QueryRow("SELECT uuid FROM flaminio.users WHERE email = $1", user.Email).Scan(&user.UUID)
	fatal(err, tx)

	stmt, err := tx.Prepare(`INSERT INTO flaminio.permissions AS p (name)
					VALUES ($1)
					ON CONFLICT (name) DO NOTHING
					RETURNING uuid`)
	fatal(err, tx)
	stmt2, err := tx.Prepare(`INSERT INTO flaminio.user_permissions VALUES ($1, $2)`)
	fatal(err, tx)

	for _, e := range permissionArray {
		var permissionUUID uuid.UUID
		err = stmt.QueryRow(e.Name).Scan(&permissionUUID)
		if err != nil {
			continue
		}

		_, err = stmt2.Exec(user.UUID, permissionUUID)
		fatal(err, tx)
	}

	err = tx.Commit()
	utility.Fatal(err)
	return
}

func createLocationsTable() (err error) {
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS flaminio.locations (
					` + standardModel + `
					name character varying(255) NOT NULL UNIQUE,
					description text
				)`)
	return err
}

func createSequenceTable() (err error) {
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS flaminio.sequences (
					` + standardModel + `
					name character varying(255) NOT NULL UNIQUE,
					description text
				)`)
	return err
}

func createReservationsTable() (err error) {
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS flaminio.reservations (
					` + standardModel + `
					name character varying(255) NOT NULL UNIQUE,
					description text,
					creatorId uuid NOT NULL REFERENCES flaminio.users ON DELETE RESTRICT ON UPDATE CASCADE,
					locationId uuid NOT NULL REFERENCES flaminio.locations ON DELETE CASCADE ON UPDATE CASCADE,
					sequenceId uuid REFERENCES flaminio.sequences ON DELETE CASCADE ON UPDATE CASCADE,
					duration tsrange NOT NULL,
					EXCLUDE USING gist (locationid WITH =, duration WITH &&)
				)`)
	return err
}

func createTokenBlacklistTable() (err error) {
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS flaminio.token_blacklist (
					jwtTokenDigest bytea PRIMARY KEY,
					revocationDate timestamp NOT NULL DEFAULT (NOW() AT TIME ZONE 'utc')
					)`)
	return err
}