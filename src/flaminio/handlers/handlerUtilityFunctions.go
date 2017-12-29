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
	"github.com/dgrijalva/jwt-go"
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
		c.AbortWithError(http.StatusInternalServerError, errors.New("error extracting user from key"))
		fmt.Fprintln(c.Writer, "Error extracting the key")
		return
	}
	return value.(models.User)
}

func getTokenFromContext(c *gin.Context) (token *jwt.Token) {
	value, exists := c.Get("token")
	if !exists {
		c.AbortWithError(http.StatusInternalServerError, errors.New("error extracting ttoken from key"))
		fmt.Fprintln(c.Writer, "Error extracting the key")
		return
	}
	return value.(*jwt.Token)
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
	pqErr, ok := err.(*pq.Error)
	return ok && string(pqErr.Code) == UNIQUE_VIOLATION_CODE
}

func isExcludeViolation(err error) (bool){
	const EXCLUDE_VIOLATION_CODE = "23P01"
	pqErr, ok := err.(*pq.Error)
	return ok && string(pqErr.Code) == EXCLUDE_VIOLATION_CODE
}

func isForeignKeyViolation(err error) (bool){
	const FOREIGN_KEY_VIOLATION_CODE = "23503"
	pqErr, ok := err.(*pq.Error)
	return ok && string(pqErr.Code) == FOREIGN_KEY_VIOLATION_CODE
}

func removeSlash(s string) (string) {
	return strings.Replace(s, "/", "", -1)
}

