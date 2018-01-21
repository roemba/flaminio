package handlers

import (
	"database/sql"
	"errors"
	"flaminio/database"
	"flaminio/models"
	"flaminio/utility"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
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
		date = time.Now().UTC()
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
		Name        string                      `json:"name" binding:"required,max=255"`
		Description models.CustomNullString     `json:"description"`
		Duration    models.CustomTsrange        `json:"duration"`
		LocationID  uuid.UUID                   `json:"location_id" binding:"required"`
		SequenceID  models.CustomNullUUID       `json:"sequence_id"`
		Color       models.CustomNullString     `json:"color"`
	}
	err := c.BindJSON(&userInput)

	if err != nil || userInput.Duration.Tsrange.Lower.Time.IsZero() || userInput.Duration.Tsrange.Upper.Time.IsZero() ||
		userInput.Duration.Tsrange.Upper.Time.Before(userInput.Duration.Tsrange.Lower.Time) {
		c.AbortWithError(http.StatusBadRequest, errors.New("error in request"))
		fmt.Fprint(c.Writer, "Error in request")
		return
	}

	reservation := models.Reservation {
		Name:           userInput.Name,
		Description:    userInput.Description,
		CreatorID:      user.UUID,
		LocationID:     userInput.LocationID,
		SequenceID:     userInput.SequenceID,
		Duration:       userInput.Duration,
	}

	if !userInput.Color.Valid {
		reservation.Color = "006080"
	} else {
		reservation.Color = userInput.Color.String
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
		if isExcludeViolation(err) {
			c.AbortWithError(http.StatusConflict, errors.New("overlap in duration for the same location"))
			fmt.Fprint(c.Writer, "Overlap in duration for the same location")
			return
		}
		c.AbortWithError(http.StatusInternalServerError, errors.New("error while creating a reservation: " + err.Error()))
		fmt.Fprint(c.Writer, "Error while creating reservation")
		return
	}

	c.JSON(http.StatusCreated, struct {Uuid uuid.UUID `json:"uuid"`}{reservationUUID})
}
