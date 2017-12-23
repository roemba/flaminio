package handlers

import (
	"fmt"
	"net/http"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"errors"
	"github.com/satori/go.uuid"
	"flaminio/database"
	"flaminio/models"
	"flaminio/utility"
)

const (
	STATUS_SUCCESS   = "success"
	STATUS_FAIL      = "failed"
)

type Response struct {
	Status string `json:"status"`
	Data interface{} `json:"data"`
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

