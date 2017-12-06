package handlers

import (
	"fmt"
	"time"
	"net/http"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"errors"
	"github.com/satori/go.uuid"
	"github.com/roemba/flaminio/database"
	"github.com/roemba/flaminio/models"
	"github.com/roemba/flaminio/utility"
)

const (
	STATUS_SUCCESS = "success"
	STATUS_FAIL = "failed"
	defaultDataFormat        = "02-01-2006"
	defaultDateAndTimeFormat = "02-01-2006 15:04:05"
)

type Response struct {
	Status string `json:"status"`
	Data interface{} `json:"data"`
}

type customDateAndTimeFormat struct {
	time.Time
}

func (c *customDateAndTimeFormat) UnmarshalJSON(j []byte) (err error) {
	s := string(j)
	s = s[1:len(s)-1]
	t, err := time.Parse(defaultDateAndTimeFormat, s)
	if err != nil {
		return err
	}
	c.Time = t
	return
}

func jsonResponse(response interface{}, w http.ResponseWriter) {
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

func toNullUUID(u uuid.UUID) uuid.NullUUID {
	if u == uuid.Nil {
		return uuid.NullUUID{
			UUID: u,
			Valid: false,
		}
	}
	return uuid.NullUUID{
		UUID: u,
		Valid: true,
	}
}

func getUserFromContext(c *gin.Context) (user models.User) {
	value, exists := c.Get("user")
	if !exists {
		c.AbortWithError(http.StatusInternalServerError, errors.New("error extracting the key"))
		fmt.Fprintln(c.Writer, "Error extracting the key")
		utility.Fatal(errors.New("could not load user from context"))
		return
	}
	return value.(models.User)
}

func checkPermission(user models.User, permissionKey string) (hasPermission bool) {
	permissionUUID := utility.GetUUIDFromMapSafely(permissionKey, database.PermissionsMap)
	for _, e := range user.Permissions {
		if e.UUID == permissionUUID {
			return true
		}
	}
	return false
}

