package handlers

import (
	"github.com/satori/go.uuid"
	"net/http"
	"fmt"
	"flaminio/models"
	"flaminio/database"
	"github.com/gin-gonic/gin"
	"time"
	"encoding/json"
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

	jsonResponse(Response{STATUS_SUCCESS, reservations}, c.Writer)
}

func PUTReservationsHandler(c *gin.Context) {
	user := getUserFromContext(c)

	if !checkPermission(user, "canEditSchedule") {
		c.AbortWithError(http.StatusForbidden, errors.New("user has invalid permissions to access this resource"))
		fmt.Fprint(c.Writer, "Invalid permissions to access this resource")
		return
	}

	type putReservationsBody struct {
		Name        string                   `json:"name"`
		Description string                   `json:"description"`
		StartTimestamp models.CustomDateAndTime `json:"start"`
		EndTimestamp models.CustomDateAndTime `json:"end"`
		LocationID  uuid.UUID                `json:"location_id"`
		SequenceID  uuid.UUID                `json:"sequence_id"`
	}

	var userInput putReservationsBody
	//TODO Use disallowunknownfields once GO 1.10 is released
	err := json.NewDecoder(c.Request.Body).Decode(&userInput)

	if err != nil || userInput.Name == "" || userInput.Description == "" || userInput.LocationID == uuid.Nil {
		c.AbortWithError(http.StatusBadRequest, errors.New("error in request"))
		fmt.Fprint(c.Writer, "Error in request")
		return
	}

	reservation := models.Reservation {
		Name: userInput.Name,
		Description: userInput.Description,
		CreatorID:   user.UUID,
		LocationID:  userInput.LocationID,
		SequenceID:  toNullUUID(userInput.SequenceID),
		StartTimestamp: userInput.StartTimestamp,
		EndTimestamp: userInput.EndTimestamp,
	}

	reservationUUID, err := database.CreateReservation(&reservation)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, errors.New("error while creating a reservation: " + err.Error()))
		fmt.Fprint(c.Writer, "Error while creating reservation")
		return
	}

	jsonResponse(Response{STATUS_SUCCESS, struct {uuid uuid.UUID}{reservationUUID}}, c.Writer)
}
