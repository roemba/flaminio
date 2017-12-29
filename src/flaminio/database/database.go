package database

import (
	"crypto/sha256"
	"database/sql"
	"flaminio/models"
	"flaminio/utility"
	"time"

	"github.com/dgrijalva/jwt-go"
	_ "github.com/jackc/pgx"
	"github.com/jmoiron/sqlx"
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

func GetReservationsByDate(date time.Time) (reservationsArray []models.Reservation, err error) {
	rows, err := db.Queryx(`SELECT * FROM flaminio.reservations AS r WHERE lower(r.duration)::date <= $1::date
									AND upper(r.duration)::date >= $1::date`, date.Format(utility.ISO8601DATE))
	if err != sql.ErrNoRows {
		utility.Fatal(err)
	} else {
		rows.Close()
		return nil, err
	}
	defer rows.Close()


	for rows.Next() {
		var reservation models.Reservation
		err := rows.StructScan(&reservation)
		utility.Fatal(err)

		reservationsArray = append(reservationsArray, reservation)
	}
	err = rows.Err()
	return reservationsArray, err
}

func GetReservationsByDateAndLocation(date time.Time, locationStringArray []string, v *[]models.Reservation) (err error) {
	//err = db.Where("date_and_time::date = ? AND location_id in (?)", date.Format(utility.ISO8601DATE), locationStringArray).Find(&v).Error
	return
}

func CreateReservation(r *models.Reservation) (reservationUUID uuid.UUID, err error) {
	err = db.QueryRow(`INSERT INTO flaminio.reservations (name, description, creatorId, locationId, sequenceId,
 								duration) VALUES ($1, $2, $3, $4, $5, $6) RETURNING uuid`, r.Name, r.Description,
 									r.CreatorID, r.LocationID, r.SequenceID, r.Duration).Scan(&reservationUUID)
	return reservationUUID, err
}

func CreateLocation(l *models.Location) (locationUUID uuid.UUID, err error) {
	err = db.QueryRow("INSERT INTO flaminio.locations (name, description) VALUES ($1, $2) RETURNING uuid", l.Name, l.Description).Scan(&locationUUID)
	return locationUUID, err
}

func GetLocations() (locationsArray []models.Location, err error) {
	rows, err := db.Queryx(`SELECT * FROM flaminio.locations ORDER BY name ASC`)
	if err != sql.ErrNoRows {
		utility.Fatal(err)
	} else {
		rows.Close()
		return nil, err
	}
	defer rows.Close()


	for rows.Next() {
		var location models.Location
		err := rows.StructScan(&location)
		utility.Fatal(err)

		locationsArray = append(locationsArray, location)
	}
	err = rows.Err()
	return locationsArray, err
}

func GetLocationByUUID(locationUUID uuid.UUID) (location models.Location, err error) {
	err = db.QueryRowx(`SELECT * FROM flaminio.locations WHERE uuid = $1`, locationUUID).StructScan(&location)
	return location, err
}

func DeleteLocation(locationUUID uuid.UUID) (err error) {
	_, err = db.Exec(`DELETE FROM flaminio.locations WHERE uuid = $1`, locationUUID)
	return err
}

func UpdateLocation(l *models.Location) (err error) {
	_, err = db.Exec(`UPDATE flaminio.locations SET (updatedat, name, description) = ((NOW() AT TIME ZONE 'utc'),
 								$2, $3) WHERE uuid = $1`, l.UUID, l.Name, l.Description)
	return err
}

func AddTokenToBlacklist(token *jwt.Token) (err error) {
	_, err = db.Exec(`DELETE FROM flaminio.token_blacklist WHERE revocationDate <= ((NOW() AT TIME ZONE 'utc') - INTERVAL '12 hours')`)
	utility.Fatal(err)

	hasher := sha256.New()
	hasher.Write([]byte(token.Raw))
	_, err = db.Exec(`INSERT INTO flaminio.token_blacklist(jwtTokenDigest) VALUES ($1)`, string(hasher.Sum(nil)))
	return err
}

func IsTokenBlacklisted(token *jwt.Token) (revoked bool) {
	var blacklistedToken models.BlacklistedToken
	hasher := sha256.New()
	hasher.Write([]byte(token.Raw))
	err := db.QueryRowx(`SELECT * FROM flaminio.token_blacklist WHERE jwtTokenDigest = $1`, string(hasher.Sum(nil))).StructScan(&blacklistedToken)
	if err == sql.ErrNoRows {
		return false
	} else if err != nil {
		utility.Fatal(err)
	}
	return true
}