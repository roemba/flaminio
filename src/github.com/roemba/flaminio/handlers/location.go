package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
	"errors"
	"encoding/json"
	"github.com/roemba/flaminio/models"
	"github.com/roemba/flaminio/database"
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

	database.CreateLocation(&location)

	jsonResponse(Response{STATUS_SUCCESS, location}, c.Writer)
}
