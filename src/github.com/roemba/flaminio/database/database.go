package database

import (
	"github.com/satori/go.uuid"
	"github.com/jinzhu/gorm"
	"github.com/roemba/flaminio/models"
	"time"
	"github.com/roemba/flaminio/utility"
)

var db *gorm.DB

func ConnectToDatabase(){
	var err error
	db, err = gorm.Open("postgres", "host=localhost user=flaminio dbname=flaminio sslmode=disable password=ZzS08RNyosHD2xg49k9Z")
	utility.Fatal(err)

	err = migrate()
	utility.Fatal(err)
}

/*
All database calls (except migrations) are done from this file
 */
func GetUser(email string) (user models.User, err bool) {
	if db.First(&user, models.User{Email: email}).RecordNotFound() {
		return user, true
	}
	return user, false
}

func GetUserWithPermissions(userUUID uuid.UUID) (user models.User) {
	var permissions []models.Permission
	db.First(&user, models.User{
		StandardModel: models.StandardModel {
			UUID: userUUID,
			},
			})
	db.Model(&user).Related(&permissions, "Permissions")
	user.Permissions = permissions
	return user
}

func AddDatabaseLog(userUUID uuid.UUID, logTypeUUID uuid.UUID, message string) {
	var log = models.Log{
		UserID:          userUUID,
		OperationTypeID: logTypeUUID,
		Message:         message,
	}
	db.Create(&log)
	return
}

func GetPermissionArray(v *[]models.Permission) (err error) {
	err = db.Find(&v).Error
	return
}

func GetLogOperationsArray(v *[]models.LogOperationType) (err error) {
	err = db.Find(&v).Error
	return
}

func GetReservationsByDate(date time.Time, v *[]models.Reservation) (err error) {
	err = db.Where("date_and_time::date = ?", date.Format(utility.ISO8601DATE)).Find(&v).Error
	return
}

func GetReservationsByDateAndLocation(date time.Time, locationStringArray []string, v *[]models.Reservation) (err error) {
	err = db.Where("date_and_time::date = ? AND location_id in (?)", date.Format(utility.ISO8601DATE), locationStringArray).Find(&v).Error
	return
}

func CreateMetaData(m *models.Metadata) (err error) {
	err = db.Create(&m).Error
	return
}

func CreateReservation(r *models.Reservation) (err error) {
	err = db.Create(&r).Error
	return
}

func CreateLocation(l *models.Location) (err error) {
	err = db.Create(&l).Error
	return
}
