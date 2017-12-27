package handlers

import (
	"github.com/satori/go.uuid"
	"net/http"
	"fmt"
	"flaminio/models"
	"flaminio/database"
	"github.com/gin-gonic/gin"
	"time"
	"errors"
	"flaminio/utility"
	"database/sql"
)

func GETReservationsHandler(c *gin.Context) {
	user := getUserFromContext(c)

	if !checkPermission(user, "canViewSchedule") {
		c.AbortWithError(http.StatusForbidden, errors.New("user has invalid permissions to access this resource"))
		fmt.Fprint(c.Writer, "Invalid permissions to access this resource")
		return
	}

	var (
		date time.Time
		err error
		reservations []models.Reservation
	)
	if dateString, exists := c.GetQuery("date"); !exists {
		date = time.Now()
	} else {
		date, err = time.Parse(utility.ISO8601DATE, dateString)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, errors.New("error in request"))
			fmt.Fprint(c.Writer, "Error in request")
			return
		}
	}

	if locationStringArray, exists := c.GetQueryArray("location"); exists {
		database.GetReservationsByDateAndLocation(date, locationStringArray, &reservations)
	} else {
		reservations, err = database.GetReservationsByDate(date)
	}

	if err != nil && err != sql.ErrNoRows {
		c.AbortWithError(http.StatusInternalServerError, errors.New("error in getting reservations: " + err.Error()))
		fmt.Fprint(c.Writer, "Error in getting reservations")
		return
	}

	c.JSON(http.StatusOK, reservations)
}

func POSTReservationsHandler(c *gin.Context) {
	user := getUserFromContext(c)

	if !checkPermission(user, "canEditSchedule") {
		c.AbortWithError(http.StatusForbidden, errors.New("user has invalid permissions to access this resource"))
		fmt.Fprint(c.Writer, "Invalid permissions to access this resource")
		return
	}

	var userInput struct {
		Name        string                   `json:"name" binding:"required"`
		Description string                   `json:"description"`
		StartTimestamp models.CustomDateAndTime `json:"start" binding:"required"`
		EndTimestamp models.CustomDateAndTime `json:"end" binding:"required"`
		LocationID  uuid.UUID                `json:"location_id" binding:"required"`
		SequenceID  uuid.UUID                `json:"sequence_id"`
	}

	err := c.BindJSON(&userInput)

	if err != nil || userInput.StartTimestamp.Time.IsZero() || userInput.EndTimestamp.Time.IsZero() {
		c.AbortWithError(http.StatusBadRequest, errors.New("error in request"))
		fmt.Fprint(c.Writer, "Error in request")
		return
	}

	reservation := models.Reservation {
		Name: userInput.Name,
		Description: models.ToNullString(userInput.Description),
		CreatorID:   user.UUID,
		LocationID:  userInput.LocationID,
		SequenceID:  models.ToNullUUID(userInput.SequenceID),
		StartTimestamp: userInput.StartTimestamp,
		EndTimestamp: userInput.EndTimestamp,
	}

	reservationUUID, err := database.CreateReservation(&reservation)
	if err != nil {
		if isUniqueViolation(err) {
			c.AbortWithError(http.StatusConflict, errors.New("reservation name already exists"))
			fmt.Fprint(c.Writer, "Reservation name already exists")
			return
		}
		if isForeignKeyViolation(err) {
			c.AbortWithError(http.StatusConflict, errors.New("invalid location or sequence uuid given"))
			fmt.Fprint(c.Writer, "Invalid location or sequence UUID given")
			return
		}
		c.AbortWithError(http.StatusInternalServerError, errors.New("error while creating a reservation: " + err.Error()))
		fmt.Fprint(c.Writer, "Error while creating reservation")
		return
	}

	c.JSON(http.StatusCreated, struct {Uuid uuid.UUID `json:"uuid"`}{reservationUUID})
}
