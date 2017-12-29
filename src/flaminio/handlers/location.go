package handlers

import (
	"errors"
	"flaminio/database"
	"flaminio/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"database/sql"

	"github.com/satori/go.uuid"
)

func GETLocationsHandler(c *gin.Context) {
	var (
		err error
		locationArray []models.Location
	)
	if uuidString := removeSlash(c.Param("uuid")); uuidString != "" {
		if locationUUID, convertError := uuid.FromString(uuidString); convertError == nil {
			var location models.Location
			location, err = database.GetLocationByUUID(locationUUID)
			locationArray = append(locationArray, location)
		} else {
			c.AbortWithError(http.StatusBadRequest, errors.New("error in request, uuid not valid"))
			fmt.Fprint(c.Writer, "Error in request, uuid not valid")
			return
		}
	} else {
		locationArray, err = database.GetLocations()
	}

	if err == sql.ErrNoRows {
		c.AbortWithError(http.StatusNotFound, errors.New("could not find location(s)"))
		fmt.Fprint(c.Writer, "Could not find location(s)")
		return
	}
	if err != nil{
		c.AbortWithError(http.StatusInternalServerError, errors.New("error in getting locations: " + err.Error()))
		fmt.Fprint(c.Writer, "Error in getting locations")
		return
	}

	c.JSON(http.StatusOK, locationArray)
}

func POSTLocationsHandler(c *gin.Context) {
	user := getUserFromContext(c)

	if !checkPermission(user, "canEditLocations") {
		c.AbortWithError(http.StatusForbidden, errors.New("user has invalid permissions to access this resource"))
		fmt.Fprint(c.Writer, "Invalid permissions to access this resource")
		return
	}

	var userInput struct {
		Name string `json:"name" binding:"required,max=255"`
		Description models.CustomNullString `json:"description"`
	}
	err := c.BindJSON(&userInput)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, errors.New("error in request"))
		fmt.Fprint(c.Writer, "Error in request")
		return
	}

	location := models.Location {
		Name: userInput.Name,
		Description: userInput.Description,
	}

	locationUUID, err := database.CreateLocation(&location)
	if err != nil{
		if isUniqueViolation(err) {
			c.AbortWithError(http.StatusConflict, errors.New("location name already exists"))
			fmt.Fprint(c.Writer, "Location name already exists")
			return
		}
		c.AbortWithError(http.StatusInternalServerError, errors.New("error in creating locations: " + err.Error()))
		fmt.Fprint(c.Writer, "Error in creating locations")
		return
	}

	c.JSON(http.StatusCreated, struct {UUID uuid.UUID `json:"uuid"`}{UUID: locationUUID,})
}

func DELETELocationHandler(c *gin.Context) {
	user := getUserFromContext(c)

	if !checkPermission(user, "canEditLocations") {
		c.AbortWithError(http.StatusForbidden, errors.New("user has invalid permissions to access this resource"))
		fmt.Fprint(c.Writer, "Invalid permissions to access this resource")
		return
	}

	var	err error
	if locationUUID, convertError := uuid.FromString(removeSlash(c.Param("uuid"))); convertError == nil {
			err = database.DeleteLocation(locationUUID)
	} else {
			c.AbortWithError(http.StatusBadRequest, errors.New("error in request, uuid not valid"))
			fmt.Fprint(c.Writer, "Error in request, uuid not valid")
			return
	}

	if err != nil{
		c.AbortWithError(http.StatusInternalServerError, errors.New("error in deleting location: " + err.Error()))
		fmt.Fprint(c.Writer, "Error in deleting location")
		return
	}

	c.Status(http.StatusOK)
}

func PUTLocationsHandler(c *gin.Context) {
	user := getUserFromContext(c)

	if !checkPermission(user, "canEditLocations") {
		c.AbortWithError(http.StatusForbidden, errors.New("user has invalid permissions to access this resource"))
		fmt.Fprint(c.Writer, "Invalid permissions to access this resource")
		return
	}

	var userInput struct {
		UUID uuid.UUID `json:"uuid" binding:"required"`
		Name string `json:"name" binding:"required,max=255"`
		Description models.CustomNullString `json:"description"`
	}
	err := c.BindJSON(&userInput)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, errors.New("error in request"))
		fmt.Fprint(c.Writer, "Error in request")
		return
	}

	location := models.Location {
		StandardModel: models.StandardModel{UUID: userInput.UUID},
		Name: userInput.Name,
		Description: userInput.Description,
	}

	err = database.UpdateLocation(&location)
	if err != nil{
		if isUniqueViolation(err) {
			c.AbortWithError(http.StatusConflict, errors.New("location name already exists"))
			fmt.Fprint(c.Writer, "Location name already exists")
			return
		}
		c.AbortWithError(http.StatusInternalServerError, errors.New("error in updating locations: " + err.Error()))
		fmt.Fprint(c.Writer, "Error in updating locations")
		return
	}

	c.Status(http.StatusOK)
}