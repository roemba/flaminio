package handlers

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	"errors"
	"flaminio/database"
	"flaminio/models"
	"flaminio/utility"
	"github.com/lib/pq"
	"strings"
)

const (
	STATUS_SUCCESS   = "success"
	STATUS_FAIL      = "failed"
)

type Envelope struct {
	Status string `json:"status"`
	Data interface{} `json:"data"`
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

func isUniqueViolation(err error) (bool){
	const UNIQUE_VIOLATION_CODE = "23505"
	if err, ok := err.(*pq.Error); ok && string(err.Code) == UNIQUE_VIOLATION_CODE {
		return true
	}
	return false
}

func isForeignKeyViolation(err error) (bool){
	const FOREIGN_KEY_VIOLATION_CODE = "23503"
	if err, ok := err.(*pq.Error); ok && string(err.Code) == FOREIGN_KEY_VIOLATION_CODE {
		return true
	}
	return false
}

func removeSlash(s string) (string) {
	return strings.Replace(s, "/", "", -1)
}

