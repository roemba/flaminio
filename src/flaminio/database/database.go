package database

import (
	"flaminio/utility"
	_ "github.com/lib/pq"
	//"github.com/roemba/flaminio/models"
	//"time"
	//"github.com/satori/go.uuid"
	"github.com/jmoiron/sqlx"
	"flaminio/models"
	"time"
	"github.com/satori/go.uuid"
)

var db *sqlx.DB

func ConnectToDatabase(){
	var err error
	db, err = sqlx.Open("postgres", "user=flaminio dbname=flaminio sslmode=disable password=ZzS08RNyosHD2xg49k9Z")
	utility.Fatal(err)

	err = migrate()
	utility.Fatal(err)
}

/*
All database calls (except migrations) are done from this file
 */
func GetUserByEmail(email string) (user models.User, err error) {
	err = db.QueryRowx("SELECT * FROM flaminio.users WHERE email = $1", email).StructScan(&user)
	return user, err
}

func GetUserByUUID(user models.User) (models.User, error) {
	err := db.QueryRowx("SELECT * FROM flaminio.users WHERE uuid = $1", user.UUID).StructScan(&user)
	return user, err
}

func GetPermissionsForUser(user models.User) (models.User, error) {
	rows, err := db.Queryx(`SELECT p.* FROM flaminio.user_permissions AS up
								INNER JOIN flaminio.permissions AS p ON up.permissionId = p.uuid
								WHERE up.userId = $1`, user.UUID)
	utility.Fatal(err)
	defer rows.Close()

	for rows.Next() {
		var permission models.Permission
		err := rows.StructScan(&permission)
		utility.Fatal(err)

		user.Permissions = append(user.Permissions, permission)
	}
	err = rows.Err()
	return user, err
}

func AddDatabaseLog(userUUID uuid.UUID, logTypeUUID uuid.UUID, message string) (err error) {
	_, err = db.Exec(`INSERT INTO flaminio.logs (userId, operationTypeId, message) VALUES ($1, $2, $3)`, userUUID, logTypeUUID, message)
	return
}

func GetPermissionArray() (permissionsArray []models.Permission, err error) {
	rows, err := db.Queryx("SELECT * FROM flaminio.permissions")
	defer rows.Close()

	for rows.Next() {
		var permission models.Permission
		err := rows.StructScan(&permission)
		utility.Fatal(err)

		permissionsArray = append(permissionsArray, permission)
	}
	err = rows.Err()
	return permissionsArray, err
}

func GetLogOperationsArray() (operationsArray []models.LogOperationType, err error) {
	rows, err := db.Queryx("SELECT * FROM flaminio.log_operation_types")
	utility.Fatal(err)
	defer rows.Close()

	for rows.Next() {
		var operationType models.LogOperationType
		err := rows.StructScan(&operationType)
		utility.Fatal(err)

		operationsArray = append(operationsArray, operationType)
	}
	err = rows.Err()
	return operationsArray, err
}

func GetReservationsByDate(date time.Time, v *[]models.Reservation) (err error) {
	//err = db.Where("date_and_time::date = ?", date.Format(utility.ISO8601DATE)).Find(&v).Error
	return
}

func GetReservationsByDateAndLocation(date time.Time, locationStringArray []string, v *[]models.Reservation) (err error) {
	//err = db.Where("date_and_time::date = ? AND location_id in (?)", date.Format(utility.ISO8601DATE), locationStringArray).Find(&v).Error
	return
}

func CreateMetaData(m *models.Metadata) (err error) {
	_, err = db.Exec("INSERT INTO flaminio.metadata (name, description) VALUES ($1, $2)", m.Name, m.Description)
	return err
}

func CreateReservation(r *models.Reservation) (err error) {
	_, err = db.Exec(`INSERT INTO flaminio.reservations (creatorId, locationId, sequenceId, metaId,
 								startTimestamp, endTimestamp) VALUES ($1, $2, $3, $4, $5, $6)`, r.CreatorID, r.LocationID,
 									r.SequenceID, r.MetaID) //Missing starttimestamp and end timestamp, so will always throw error
	return err
}

func CreateLocation(l *models.Location) (err error) {
	_, err = db.Exec("INSERT INTO flaminio.locations (name, description) VALUES ($1, $2)", l.Name, l.Description)
	return err
}
