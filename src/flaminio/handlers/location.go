package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
	"errors"
	"encoding/json"
	"flaminio/models"
	"flaminio/database"
	"github.com/satori/go.uuid"
)

func PUTLocationsHandler(c *gin.Context) {
	user := getUserFromContext(c)

	if !checkPermission(user, "canEditLocations") {
		c.AbortWithError(http.StatusForbidden, errors.New("user has invalid permissions to access this resource"))
		fmt.Fprint(c.Writer, "Invalid permissions to access this resource")
		return
	}

	type putLocationsBody struct {
		Name string `json:"name"`
		Description string `json:"description"`
	}

	var userInput putLocationsBody
	err := json.NewDecoder(c.Request.Body).Decode(&userInput)

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
			c.AbortWithError(http.StatusBadRequest, errors.New("location name already exists"))
			fmt.Fprint(c.Writer, "Location name already exists")
			return
		}
		c.AbortWithError(http.StatusInternalServerError, errors.New("error in getting reservations: " + err.Error()))
		fmt.Fprint(c.Writer, "Error in getting reservations")
		return
	}

	response := struct {
		UUID uuid.UUID `json:"uuid"`
	}{
		UUID: locationUUID,
	}
	jsonResponse(Response{STATUS_SUCCESS, response}, c.Writer)
}
